package pipeline

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"sync"
	"sync/atomic"
	"time"
)

type LogEntry struct {
	Time    string `json:"time"`
	Message string `json:"message"`
	Level   string `json:"level"`
}

type LogBuffer struct {
	mu      sync.Mutex
	entries []LogEntry
	maxSize int
	out     io.Writer
	version atomic.Int64
}

func NewLogBuffer(out io.Writer, maxSize int) *LogBuffer {
	return &LogBuffer{out: out, maxSize: maxSize}
}

func (lb *LogBuffer) Write(p []byte) (n int, err error) {
	n, err = lb.out.Write(p)
	msg := string(p)
	if len(msg) > 0 && msg[len(msg)-1] == '\n' {
		msg = msg[:len(msg)-1]
	}
	level := "info"
	if containsStr(msg, "ERROR") {
		level = "error"
	} else if containsStr(msg, "VERBOSE") {
		level = "verbose"
	}
	lb.mu.Lock()
	lb.entries = append(lb.entries, LogEntry{
		Time:    time.Now().Format("15:04:05"),
		Message: msg,
		Level:   level,
	})
	if len(lb.entries) > lb.maxSize {
		lb.entries = lb.entries[len(lb.entries)-lb.maxSize:]
	}
	lb.mu.Unlock()
	lb.version.Add(1)
	return
}

func (lb *LogBuffer) Entries(verbose bool) []LogEntry {
	lb.mu.Lock()
	defer lb.mu.Unlock()
	if verbose {
		out := make([]LogEntry, len(lb.entries))
		copy(out, lb.entries)
		return out
	}
	var out []LogEntry
	for _, e := range lb.entries {
		if e.Level != "verbose" {
			out = append(out, e)
		}
	}
	return out
}

func (lb *LogBuffer) Version() int64 {
	return lb.version.Load()
}

func containsStr(s, sub string) bool {
	return len(s) >= len(sub) && searchStr(s, sub)
}

func searchStr(s, sub string) bool {
	for i := 0; i <= len(s)-len(sub); i++ {
		if s[i:i+len(sub)] == sub {
			return true
		}
	}
	return false
}

type Dashboard struct {
	Runner        *Runner
	Port          int
	CycleInterval int
	ConfigPath    string
	stepCh        chan struct{}
	intervalCh    chan int
	mu            sync.Mutex
	lastLog       []string
	nextCycleAt   time.Time
	cycleRunning  bool
	cycleStarted  time.Time
	retryMessage  string
	LogBuf        *LogBuffer
	paused        bool
}

func NewDashboard(runner *Runner, port, cycleInterval int, configPath string, logBuf *LogBuffer) *Dashboard {
	return &Dashboard{
		Runner:        runner,
		Port:          port,
		CycleInterval: cycleInterval,
		ConfigPath:    configPath,
		stepCh:        make(chan struct{}, 1),
		intervalCh:    make(chan int, 1),
		LogBuf:        logBuf,
		paused:        true,
	}
}

func (d *Dashboard) SetNextCycleAt(t time.Time) {
	d.mu.Lock()
	defer d.mu.Unlock()
	d.nextCycleAt = t
}

func (d *Dashboard) SetCycleRunning() {
	d.mu.Lock()
	defer d.mu.Unlock()
	d.cycleRunning = true
	d.cycleStarted = time.Now()
}

func (d *Dashboard) SetCycleFinished() {
	d.mu.Lock()
	defer d.mu.Unlock()
	d.cycleRunning = false
}

func (d *Dashboard) SetRetryMessage(msg string) {
	d.mu.Lock()
	defer d.mu.Unlock()
	d.retryMessage = msg
}

func (d *Dashboard) StepChannel() <-chan struct{} {
	return d.stepCh
}

func (d *Dashboard) IntervalChannel() <-chan int {
	return d.intervalCh
}

func (d *Dashboard) IsPaused() bool {
	d.mu.Lock()
	defer d.mu.Unlock()
	return d.paused
}

func (d *Dashboard) SetLastLog(actions []string) {
	d.mu.Lock()
	defer d.mu.Unlock()
	d.lastLog = actions
}

func (d *Dashboard) Start() error {
	mux := http.NewServeMux()
	mux.HandleFunc("/", d.handleIndex)
	mux.HandleFunc("/api/status", d.handleStatus)
	mux.HandleFunc("/api/step", d.handleStep)
	mux.HandleFunc("/api/essays", d.handleEssays)
	mux.HandleFunc("/api/open", d.handleOpen)

	mux.HandleFunc("/api/open-docx", d.handleOpenDocx)
	mux.HandleFunc("/api/revert", d.handleRevert)
	mux.HandleFunc("/api/revert-all", d.handleRevertAll)
	mux.HandleFunc("/api/disk-stats", d.handleDiskStats)
	mux.HandleFunc("/api/logs", d.handleLogs)
	mux.HandleFunc("/api/pause", d.handlePause)
	mux.HandleFunc("/api/settings", d.handleSettings)
	mux.HandleFunc("/accounting", d.handleAccounting)

	addr := fmt.Sprintf("127.0.0.1:%d", d.Port)
	d.Runner.Log.Printf("Dashboard: http://%s", addr)
	return http.ListenAndServe(addr, mux)
}

func (d *Dashboard) handleIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	tmpl, err := template.New("dashboard").Parse(dashboardHTML)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	tmpl.Execute(w, nil)
}

type statusResponse struct {
	DryRun            bool                    `json:"dry_run"`
	Verbose           bool                    `json:"verbose"`
	Paused            bool                    `json:"paused"`
	Cycle             int                     `json:"cycle"`
	TotalCost         float64                 `json:"total_cost"`
	SessionDone       int                     `json:"session_done"`
	CycleInterval     int                     `json:"cycle_interval"`
	SecondsLeft       int                     `json:"seconds_left"`
	CycleRunning      bool                    `json:"cycle_running"`
	CycleElapsed      int                     `json:"cycle_elapsed"`
	RetryMessage      string                  `json:"retry_message"`
	ReadMean          float64                 `json:"read_mean"`
	ReadSpread        float64                 `json:"read_spread"`
	SkipRevertConfirm bool                    `json:"skip_revert_confirm"`
	Projects          []projectStatusResponse `json:"projects"`
	Summary           map[string]int          `json:"summary"`
	LastLog           []string                `json:"last_log"`
}

type projectStatusResponse struct {
	Name           string         `json:"name"`
	Cycle          int            `json:"cycle"`
	Cost           float64        `json:"cost"`
	Summary        map[string]int `json:"summary"`
	HasBlurb       bool           `json:"has_blurb"`
	HasCoverPrompt bool           `json:"has_cover_prompt"`
	HasCoverImage  bool           `json:"has_cover_image"`
}

func (d *Dashboard) handleStatus(w http.ResponseWriter, r *http.Request) {
	d.mu.Lock()
	lastLog := d.lastLog
	nextAt := d.nextCycleAt
	running := d.cycleRunning
	started := d.cycleStarted
	retryMsg := d.retryMessage
	paused := d.paused
	d.mu.Unlock()

	secsLeft := int(time.Until(nextAt).Seconds())
	if secsLeft < 0 {
		secsLeft = 0
	}

	var elapsed int
	if running {
		elapsed = int(time.Since(started).Seconds())
	}

	totalSummary := map[string]int{
		"pending": 0, "research": 0, "outline": 0,
		"draft": 0, "factcheck": 0, "draft2": 0, "illustrate": 0, "done": 0, "error": 0,
	}

	var projectStatuses []projectStatusResponse
	for _, ps := range d.Runner.Projects {
		s := ps.Summary()
		bookDir := filepath.Join(ps.BaseDir, "book")
		projectStatuses = append(projectStatuses, projectStatusResponse{
			Name:           ps.Project,
			Cycle:          ps.CycleCount,
			Cost:           ps.TotalCost,
			Summary:        s,
			HasBlurb:       bookDirHasFile(bookDir, "back-cover-blurb.md"),
			HasCoverPrompt: bookDirHasFile(bookDir, "front-cover-prompt.md"),
			HasCoverImage:  bookDirHasFile(bookDir, "front-cover.png"),
		})
		for k, v := range s {
			totalSummary[k] += v
		}
	}

	resp := statusResponse{
		DryRun:            d.Runner.Config.Pipeline.DryRun,
		Verbose:           d.Runner.Config.Pipeline.Verbose,
		Paused:            paused,
		Cycle:             d.Runner.TotalCycles(),
		TotalCost:         d.Runner.TotalCost(),
		SessionDone:       d.Runner.SessionDone(),
		CycleInterval:     d.CycleInterval,
		SecondsLeft:       secsLeft,
		CycleRunning:      running,
		CycleElapsed:      elapsed,
		RetryMessage:      retryMsg,
		ReadMean:          d.Runner.Config.Pipeline.ReadMean,
		ReadSpread:        d.Runner.Config.Pipeline.ReadSpread,
		SkipRevertConfirm: d.Runner.Config.Pipeline.SkipRevertConfirm,
		Projects:          projectStatuses,
		Summary:           totalSummary,
		LastLog:           lastLog,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func (d *Dashboard) handleStep(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "POST only", http.StatusMethodNotAllowed)
		return
	}

	select {
	case d.stepCh <- struct{}{}:
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, `{"ok":true}`)
	default:
		w.WriteHeader(http.StatusTooManyRequests)
		fmt.Fprintf(w, `{"ok":false,"error":"step already pending"}`)
	}
}

func (d *Dashboard) handleOpen(w http.ResponseWriter, r *http.Request) {
	project := r.URL.Query().Get("project")
	slug := r.URL.Query().Get("slug")
	if project == "" || slug == "" {
		http.Error(w, "project and slug required", http.StatusBadRequest)
		return
	}

	var ps *PipelineState
	for _, p := range d.Runner.Projects {
		if p.Project == project {
			ps = p
			break
		}
	}
	if ps == nil {
		http.Error(w, "project not found", http.StatusNotFound)
		return
	}

	essay := ps.Essays[slug]
	if essay == nil {
		http.Error(w, "essay not found", http.StatusNotFound)
		return
	}

	stage := essay.CurrentStage
	mdFile := filepath.Join(ps.BaseDir, stage.Dir(), slug+".md")

	if err := exec.Command("open", mdFile).Run(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, `{"ok":true,"path":"%s"}`, mdFile)
}

func (d *Dashboard) handleOpenDocx(w http.ResponseWriter, r *http.Request) {
	project := r.URL.Query().Get("project")
	slug := r.URL.Query().Get("slug")
	if project == "" || slug == "" {
		http.Error(w, "project and slug required", http.StatusBadRequest)
		return
	}

	var ps *PipelineState
	for _, p := range d.Runner.Projects {
		if p.Project == project {
			ps = p
			break
		}
	}
	if ps == nil {
		http.Error(w, "project not found", http.StatusNotFound)
		return
	}

	essay := ps.Essays[slug]
	if essay == nil {
		http.Error(w, "essay not found", http.StatusNotFound)
		return
	}

	exportName := exportFilename(essay)
	docxFile := filepath.Join(ps.BaseDir, "export", exportName)
	if err := exec.Command("open", docxFile).Run(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, `{"ok":true,"path":"%s"}`, docxFile)
}

var (
	reHTMLTag    = regexp.MustCompile(`<[^>]+>`)
	reFencedCode = regexp.MustCompile("(?s)```.*?```")
)

func sanitizeForDocx(md string) string {
	md = reFencedCode.ReplaceAllString(md, "[markdown removed]")
	md = reHTMLTag.ReplaceAllString(md, "")

	var lines []string
	for _, line := range strings.Split(md, "\n") {
		trimmed := strings.TrimSpace(line)
		if strings.HasPrefix(trimmed, "|") && strings.HasSuffix(trimmed, "|") {
			lines = append(lines, "[markdown removed]")
			continue
		}
		if strings.HasPrefix(trimmed, "![") {
			lines = append(lines, "[markdown removed]")
			continue
		}
		lines = append(lines, line)
	}
	return strings.Join(lines, "\n")
}

func (d *Dashboard) handleLogs(w http.ResponseWriter, r *http.Request) {
	verbose := r.URL.Query().Get("verbose") == "true"
	ver := d.LogBuf.Version()
	entries := d.LogBuf.Entries(verbose)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(struct {
		Version int64      `json:"version"`
		Entries []LogEntry `json:"entries"`
	}{Version: ver, Entries: entries})
}

func (d *Dashboard) handlePause(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "POST only", http.StatusMethodNotAllowed)
		return
	}
	d.mu.Lock()
	d.paused = !d.paused
	nowPaused := d.paused
	d.mu.Unlock()
	if nowPaused {
		d.Runner.Log.Println("Pipeline PAUSED")
	} else {
		d.Runner.Log.Println("Pipeline RESUMED")
	}
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, `{"ok":true,"paused":%v}`, nowPaused)
}

func (d *Dashboard) handleSettings(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "POST only", http.StatusMethodNotAllowed)
		return
	}

	var req struct {
		CycleInterval     *int     `json:"cycle_interval"`
		Verbose           *bool    `json:"verbose"`
		ReadMean          *float64 `json:"read_mean"`
		ReadSpread        *float64 `json:"read_spread"`
		SkipRevertConfirm *bool    `json:"skip_revert_confirm"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if req.CycleInterval != nil && *req.CycleInterval >= 5 {
		d.CycleInterval = *req.CycleInterval
		d.Runner.Config.Pipeline.CycleInterval = *req.CycleInterval
		select {
		case d.intervalCh <- *req.CycleInterval:
		default:
		}
	}
	if req.Verbose != nil {
		d.Runner.Config.Pipeline.Verbose = *req.Verbose
	}
	if req.ReadMean != nil && *req.ReadMean >= 1 && *req.ReadMean <= 15 {
		d.Runner.Config.Pipeline.ReadMean = *req.ReadMean
	}
	if req.ReadSpread != nil && *req.ReadSpread >= 0.5 && *req.ReadSpread <= 3 {
		d.Runner.Config.Pipeline.ReadSpread = *req.ReadSpread
	}
	if req.SkipRevertConfirm != nil {
		d.Runner.Config.Pipeline.SkipRevertConfirm = *req.SkipRevertConfirm
	}

	if err := SaveConfig(d.ConfigPath, d.Runner.Config); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, `{"ok":true}`)
}

func (d *Dashboard) handleAccounting(w http.ResponseWriter, r *http.Request) {
	var allEntries []AccountingEntry
	for _, ps := range d.Runner.Projects {
		af, err := loadAccounting(ps.accountingPath())
		if err != nil {
			continue
		}
		allEntries = append(allEntries, af.Reverted...)

		ps.mu.Lock()
		for _, es := range ps.Essays {
			for stage, meta := range es.Meta {
				if meta.Cost == 0 {
					continue
				}
				allEntries = append(allEntries, AccountingEntry{
					Slug:      meta.Slug,
					Stage:     stage.String(),
					Book:      meta.Book,
					TokensIn:  meta.Tokens,
					TokensOut: meta.TokensOut,
					Cost:      meta.Cost,
					Model:     meta.Model,
					Count:     1,
				})
			}
		}
		ps.mu.Unlock()
	}

	type row struct {
		Label               string
		Runs                int
		TokensIn, TokensOut int
		Cost                float64
		AvgCost             float64
	}

	aggregate := func(entries []AccountingEntry, keyFn func(AccountingEntry) string, order []string) []row {
		m := make(map[string]*row)
		for _, e := range entries {
			k := keyFn(e)
			r, ok := m[k]
			if !ok {
				r = &row{Label: k}
				m[k] = r
			}
			cnt := e.Count
			if cnt == 0 {
				cnt = 1
			}
			r.Runs += cnt
			r.TokensIn += e.TokensIn
			r.TokensOut += e.TokensOut
			r.Cost += e.Cost
		}
		var rows []row
		seen := make(map[string]bool)
		for _, k := range order {
			if r, ok := m[k]; ok {
				if r.Runs > 0 {
					r.AvgCost = r.Cost / float64(r.Runs)
				}
				rows = append(rows, *r)
				seen[k] = true
			}
		}
		for k, r := range m {
			if !seen[k] {
				if r.Runs > 0 {
					r.AvgCost = r.Cost / float64(r.Runs)
				}
				rows = append(rows, *r)
			}
		}
		var total row
		total.Label = "Total"
		for _, r := range rows {
			total.Runs += r.Runs
			total.TokensIn += r.TokensIn
			total.TokensOut += r.TokensOut
			total.Cost += r.Cost
		}
		if total.Runs > 0 {
			total.AvgCost = total.Cost / float64(total.Runs)
		}
		rows = append(rows, total)
		return rows
	}

	stageOrder := []string{"seed", "ideas", "research", "outline", "draft", "factcheck", "illustrate", "draft2", "export"}
	bookOrder := []string{"I", "II", "III"}

	byStage := aggregate(allEntries, func(e AccountingEntry) string { return e.Stage }, stageOrder)
	byBook := aggregate(allEntries, func(e AccountingEntry) string { return e.Book }, bookOrder)

	type bookStageGroup struct {
		Book string
		Rows []row
	}
	var byBookStage []bookStageGroup
	bookEntries := make(map[string][]AccountingEntry)
	for _, e := range allEntries {
		bookEntries[e.Book] = append(bookEntries[e.Book], e)
	}
	var grandTotal row
	grandTotal.Label = "Grand Total"
	for _, b := range bookOrder {
		entries := bookEntries[b]
		if len(entries) == 0 {
			continue
		}
		rows := aggregate(entries, func(e AccountingEntry) string { return e.Stage }, stageOrder)
		for _, r := range rows {
			if r.Label == "Total" {
				grandTotal.Runs += r.Runs
				grandTotal.TokensIn += r.TokensIn
				grandTotal.TokensOut += r.TokensOut
				grandTotal.Cost += r.Cost
			}
		}
		byBookStage = append(byBookStage, bookStageGroup{Book: b, Rows: rows})
	}
	if grandTotal.Runs > 0 {
		grandTotal.AvgCost = grandTotal.Cost / float64(grandTotal.Runs)
	}

	data := struct {
		ByStage     []row
		ByBook      []row
		ByBookStage []bookStageGroup
		GrandTotal  row
		EntryCount  int
	}{byStage, byBook, byBookStage, grandTotal, len(allEntries)}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	tmpl, err := template.New("accounting").Parse(accountingHTML)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	tmpl.Execute(w, data)
}

func (d *Dashboard) handleRevert(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "POST only", http.StatusMethodNotAllowed)
		return
	}

	project := r.URL.Query().Get("project")
	slug := r.URL.Query().Get("slug")
	stage := r.URL.Query().Get("stage")
	if project == "" || slug == "" || stage == "" {
		http.Error(w, "project, slug, and stage required", http.StatusBadRequest)
		return
	}

	target := StageFromString(stage)

	var ps *PipelineState
	for _, p := range d.Runner.Projects {
		if p.Project == project {
			ps = p
			break
		}
	}
	if ps == nil {
		http.Error(w, "project not found", http.StatusNotFound)
		return
	}

	removed, err := ps.RevertToStage(slug, target)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if target < StageIllustrate {
		invalidateBookArtifacts(ps.BaseDir)
	}

	d.Runner.Log.Printf("[%s] REVERT %s to %s (removed: %v)", project, slug, stage, removed)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(struct {
		Ok      bool     `json:"ok"`
		Removed []string `json:"removed"`
	}{Ok: true, Removed: removed})
}

func (d *Dashboard) handleRevertAll(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "POST only", http.StatusMethodNotAllowed)
		return
	}

	project := r.URL.Query().Get("project")
	stage := r.URL.Query().Get("stage")
	if project == "" || stage == "" {
		http.Error(w, "project and stage required", http.StatusBadRequest)
		return
	}

	target := StageFromString(stage)

	var ps *PipelineState
	for _, p := range d.Runner.Projects {
		if p.Project == project {
			ps = p
			break
		}
	}
	if ps == nil {
		http.Error(w, "project not found", http.StatusNotFound)
		return
	}

	var slugs []string
	ps.mu.Lock()
	for slug := range ps.Essays {
		slugs = append(slugs, slug)
	}
	ps.mu.Unlock()

	reverted := 0
	var errors []string
	for _, slug := range slugs {
		_, err := ps.RevertToStage(slug, target)
		if err != nil {
			errors = append(errors, slug+": "+err.Error())
			continue
		}
		reverted++
	}

	if target < StageIllustrate {
		invalidateBookArtifacts(ps.BaseDir)
	}

	d.Runner.Log.Printf("[%s] REVERT-ALL to %s (%d reverted, %d errors)", project, stage, reverted, len(errors))

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(struct {
		Ok       bool     `json:"ok"`
		Reverted int      `json:"reverted"`
		Errors   []string `json:"errors,omitempty"`
	}{Ok: len(errors) == 0, Reverted: reverted, Errors: errors})
}

func (d *Dashboard) handleDiskStats(w http.ResponseWriter, r *http.Request) {
	stages := []string{"ideas", "research", "outline", "draft", "factcheck", "illustrate", "draft2", "export"}
	counts := make(map[string]int, len(stages))

	for _, ps := range d.Runner.Projects {
		for _, stage := range stages {
			dir := filepath.Join(ps.BaseDir, stage)
			entries, err := os.ReadDir(dir)
			if err != nil {
				continue
			}
			for _, e := range entries {
				name := e.Name()
				if strings.HasPrefix(name, "~$") {
					continue
				}
				ext := filepath.Ext(name)
				if ext == ".md" || ext == ".docx" {
					counts[stage]++
				}
			}
		}
	}

	totalCost := 0.0
	for _, ps := range d.Runner.Projects {
		ps.mu.Lock()
		totalCost += ps.TotalCost + ps.RevertedCost()
		ps.mu.Unlock()
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(struct {
		Counts    map[string]int `json:"counts"`
		TotalCost float64        `json:"total_cost"`
	}{Counts: counts, TotalCost: totalCost})
}

type essayJSON struct {
	Project        string `json:"project"`
	Slug           string `json:"slug"`
	Title          string `json:"title"`
	Type           string `json:"type"`
	Series         string `json:"series,omitempty"`
	Book           string `json:"book"`
	Part           int    `json:"part"`
	PartTitle      string `json:"part_title"`
	Order          int    `json:"order"`
	Stage          string `json:"stage"`
	Status         string `json:"status"`
	Arc            string `json:"arc,omitempty"`
	Ending         string `json:"ending,omitempty"`
	Structure      string `json:"structure,omitempty"`
	Entry          string `json:"entry,omitempty"`
	Register       string `json:"register,omitempty"`
	Setting        string `json:"setting,omitempty"`
	MathVisibility string `json:"math_visibility,omitempty"`
	Error          string `json:"error,omitempty"`
	HasDocx        bool   `json:"has_docx"`
	WordCount      int    `json:"word_count,omitempty"`
	ReadMins       int    `json:"read_mins,omitempty"`
	Stale          bool   `json:"stale,omitempty"`
}

func (d *Dashboard) handleEssays(w http.ResponseWriter, r *http.Request) {
	projectFilter := r.URL.Query().Get("project")

	debugBook, debugPart, debugOrder, debugActive := ParseDebug(d.Runner.Config.Pipeline.Debug)

	var essays []essayJSON
	for _, ps := range d.Runner.Projects {
		if projectFilter != "" && ps.Project != projectFilter {
			continue
		}
		for _, e := range ps.SnapshotEssays() {
			if debugActive && !(e.Book == debugBook && e.Part == debugPart && e.Order == debugOrder) {
				continue
			}
			stage := e.CurrentStage.String()
			status := e.Status
			if e.IsDone() {
				stage = "done"
				status = "done"
			}
			errMsg := ""
			if meta, ok := e.Meta[e.CurrentStage]; ok && meta.Error != "" {
				errMsg = meta.Error
			}
			hasDocx := false
			var wordCount, readMins int
			if e.IsDone() {
				exportName := exportFilename(e)
				exportPath := filepath.Join(ps.BaseDir, "export", exportName)
				if _, err := os.Stat(exportPath); err == nil {
					hasDocx = true
				}
			}
			for _, s := range []string{"draft2", "illustrate", "draft", "outline"} {
				path := filepath.Join(ps.BaseDir, s, e.Slug+".md")
				if raw, err := os.ReadFile(path); err == nil {
					words := strings.Fields(string(raw))
					wordCount = len(words)
					readMins = (wordCount + 264) / 265
					break
				}
			}
			stale := false
			if e.CurrentStage > StageIdeas {
				metaPath := filepath.Join(ps.BaseDir, "ideas", e.Slug+".meta.yaml")
				if metaInfo, err := os.Stat(metaPath); err == nil {
					metaMod := metaInfo.ModTime()
					for _, s := range []string{"research", "outline", "draft", "factcheck", "illustrate", "draft2"} {
						cp := filepath.Join(ps.BaseDir, s, e.Slug+".md")
						if ci, err := os.Stat(cp); err == nil {
							if metaMod.After(ci.ModTime()) {
								stale = true
								break
							}
						}
					}
				}
			}
			essays = append(essays, essayJSON{
				Project:        ps.Project,
				Slug:           e.Slug,
				Title:          e.Title,
				Type:           e.Type,
				Series:         e.Series,
				Book:           e.Book,
				Part:           e.Part,
				PartTitle:      e.PartTitle,
				Order:          e.Order,
				Stage:          stage,
				Status:         status,
				Arc:            e.Arc,
				Ending:         e.Ending,
				Structure:      e.Structure,
				Entry:          e.Entry,
				Register:       e.Register,
				Setting:        e.Setting,
				MathVisibility: e.MathVisibility,
				Error:          errMsg,
				HasDocx:        hasDocx,
				WordCount:      wordCount,
				ReadMins:       readMins,
				Stale:          stale,
			})
		}
	}

	sort.Slice(essays, func(i, j int) bool {
		ri, rj := stageRank(essays[i].Stage), stageRank(essays[j].Stage)
		if ri != rj {
			return ri < rj
		}
		if essays[i].Stale != essays[j].Stale {
			return essays[i].Stale
		}
		if essays[i].Book != essays[j].Book {
			return essays[i].Book < essays[j].Book
		}
		if essays[i].Part != essays[j].Part {
			return essays[i].Part < essays[j].Part
		}
		return essays[i].Order < essays[j].Order
	})

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(essays)
}

func stageRank(stage string) int {
	switch stage {
	case "research":
		return 0
	case "outline":
		return 1
	case "draft":
		return 2
	case "factcheck":
		return 3
	case "illustrate":
		return 4
	case "draft2":
		return 5
	case "export":
		return 6
	case "ideas":
		return 7
	case "done":
		return 8
	default:
		return 9
	}
}

const dashboardHTML = `<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="utf-8">
<meta name="viewport" content="width=device-width, initial-scale=1">
<title>Pipeline Dashboard</title>
<style>
  * { box-sizing: border-box; margin: 0; padding: 0; }
  body { font-family: -apple-system, BlinkMacSystemFont, "Segoe UI", Roboto, sans-serif;
         background: #1a1a2e; color: #e0e0e0; }
  .layout { display: flex; height: 100vh; }
  .main-panel { flex: 1; min-width: 0; padding: 20px; overflow-y: auto; }
  .log-panel { width: 380px; min-width: 300px; background: #0f0f23; border-left: 1px solid #2a2a4e;
               display: flex; flex-direction: column; }
  .log-header { padding: 12px 16px; border-bottom: 1px solid #2a2a4e; display: flex;
                align-items: center; gap: 12px; flex-wrap: wrap; }
  .log-header h2 { color: #e94560; font-size: 1rem; margin: 0; white-space: nowrap; }
  .log-settings { display: flex; align-items: center; gap: 10px; flex-wrap: wrap; }
  .log-settings label { font-size: 0.75rem; color: #8899aa; display: flex; align-items: center; gap: 4px; cursor: pointer; }
  .log-settings input[type="checkbox"] { accent-color: #e94560; }
  .interval-ctrl { display: flex; align-items: center; gap: 4px; }
  .interval-ctrl input[type="number"] { width: 60px; background: #16213e; border: 1px solid #2a2a4e;
    color: #e0e0e0; border-radius: 4px; padding: 3px 6px; font-size: 0.8rem; text-align: center; }
  .interval-ctrl .save-btn { background: #e94560; color: white; border: none; padding: 3px 10px;
    border-radius: 4px; cursor: pointer; font-size: 0.75rem; }
  .interval-ctrl .save-btn:hover { background: #c73650; }
  .interval-ctrl .save-btn.saved { background: #4caf50; }
  .log-body { flex: 1; overflow-y: auto; overflow-x: auto; padding: 8px 12px;
              font-family: "SF Mono", Monaco, monospace; font-size: 0.75rem; line-height: 1.5; }
  .log-line { padding: 1px 0; white-space: pre; }
  .log-line.error { color: #f44336; }
  .log-line.verbose { color: #546e7a; }
  .log-line.info { color: #6a9955; }
  .log-count { font-size: 0.7rem; color: #546e7a; padding: 4px 12px; border-top: 1px solid #2a2a4e;
               display: flex; justify-content: space-between; align-items: center; }
  .copy-btn { background: #16213e; color: #8899aa; border: 1px solid #2a2a4e; padding: 2px 10px;
              border-radius: 4px; cursor: pointer; font-size: 0.7rem; }
  .copy-btn:hover { border-color: #e94560; color: #e0e0e0; }
  .copy-btn.copied { background: #4caf50; color: white; border-color: #4caf50; }
  h1 { color: #e94560; margin-bottom: 4px; font-size: 1.5rem; }
  .subtitle { color: #8899aa; font-size: 0.85rem; margin-bottom: 16px; }
  .grid { display: grid; grid-template-columns: repeat(auto-fit, minmax(100px, 1fr));
          gap: 10px; margin-bottom: 20px; }
  .card { background: #16213e; border-radius: 8px; padding: 12px; text-align: center; }
  .card .number { font-size: 1.6rem; font-weight: bold; color: #e94560; }
  .card .label { font-size: 0.7rem; color: #8899aa; text-transform: uppercase; letter-spacing: 1px; }
  .controls { margin-bottom: 20px; display: flex; align-items: center; gap: 16px; }
  .btn { background: #e94560; color: white; border: none; padding: 10px 24px;
         border-radius: 6px; cursor: pointer; font-size: 1rem; }
  .btn:hover { background: #c73650; }
  .btn:disabled { background: #444; cursor: not-allowed; }
  .mode { display: inline-block; padding: 4px 12px; border-radius: 4px;
          font-size: 0.85rem; font-weight: bold; }
  .mode.dry { background: #ff9800; color: #000; }
  .mode.live { background: #4caf50; color: #fff; }
  .countdown { font-family: "SF Mono", Monaco, monospace; font-size: 1.4rem;
               color: #e94560; min-width: 60px; }
  .countdown-label { color: #8899aa; font-size: 0.75rem; text-transform: uppercase; }
  .progress-ring { width: 48px; height: 48px; }
  .progress-ring circle { fill: none; stroke-width: 3; }
  .progress-ring .bg { stroke: #16213e; }
  .progress-ring .fg { stroke: #e94560; stroke-linecap: round;
                       transform: rotate(-90deg); transform-origin: center;
                       transition: stroke-dashoffset 0.5s ease; }
  .project-tabs { display: flex; gap: 8px; flex-wrap: wrap; margin-left: auto; }
  .tab { padding: 6px 16px; border-radius: 4px; cursor: pointer; background: #16213e;
         color: #8899aa; font-size: 0.85rem; border: 1px solid #2a2a4e; }
  .tab.active { background: #e94560; color: white; border-color: #e94560; }
  .tab:hover { border-color: #e94560; }
  .book-dots { display: inline-flex; gap: 3px; margin-left: 6px; vertical-align: middle; }
  .book-dot { width: 8px; height: 8px; border-radius: 50%; display: inline-block; }
  .book-dot.done { background: #4ade80; }
  .book-dot.missing { background: #555; }
  table { width: 100%; border-collapse: collapse; margin-top: 12px; }
  th { text-align: left; padding: 8px 12px; background: #16213e; color: #8899aa;
       font-size: 0.75rem; text-transform: uppercase; letter-spacing: 1px; }
  td { padding: 6px 12px; border-bottom: 1px solid #1a1a3e; font-size: 0.85rem; }
  tr:hover td { background: #16213e; }
  .stage-badge { display: inline-block; padding: 2px 8px; border-radius: 3px;
                 font-size: 0.75rem; font-weight: bold; }
  .stage-ideas { background: #333; }
  .stage-research { background: #1565c0; }
  .stage-outline { background: #6a1b9a; }
  .stage-draft { background: #2e7d32; }
  .stage-factcheck { background: #e65100; }
  .stage-draft2 { background: #00838f; }
  .stage-export { background: #4527a0; }
  .stage-done { background: #2e7d32; }
  .status-pending { color: #888; }
  .status-done { color: #4caf50; }
  .status-final { color: #4caf50; }
  .status-in-progress { color: #ff9800; }
  .status-error { color: #f44336; }
  .cost { color: #ff9800; font-weight: bold; }
  .retry-banner { background: #e65100; color: white; padding: 10px 16px; border-radius: 6px;
                  margin-bottom: 16px; font-weight: bold; font-size: 0.95rem; }
  .open-btn { background: #2e7d32; color: white; border: none; padding: 3px 10px; border-radius: 3px;
              font-size: 0.75rem; cursor: pointer; font-weight: bold; }
  .open-btn:hover { background: #388e3c; }
  .modal-overlay { display:none; position:fixed; top:0; left:0; width:100%; height:100%;
                   background:rgba(0,0,0,0.6); z-index:1000; align-items:center; justify-content:center; }
  .modal-overlay.active { display:flex; }
  .modal-box { background:#1a1a2e; border:1px solid #2a2a4e; border-radius:8px; padding:24px;
               max-width:420px; width:90%; color:#e0e0e0; }
  .modal-box h3 { margin:0 0 12px; color:#e94560; font-size:1rem; }
  .modal-box p { margin:0 0 16px; font-size:0.85rem; line-height:1.4; }
  .modal-box label { font-size:0.8rem; color:#8899aa; cursor:pointer; display:flex; align-items:center; gap:6px; }
  .modal-btns { display:flex; gap:10px; margin-top:16px; justify-content:flex-end; }
  .modal-btns button { padding:6px 16px; border:none; border-radius:4px; cursor:pointer; font-size:0.8rem; font-weight:bold; }
  .modal-btn-cancel { background:#2a2a4e; color:#8899aa; }
  .modal-btn-cancel:hover { background:#3a3a5e; }
  .modal-btn-confirm { background:#e94560; color:white; }
  .modal-btn-confirm:hover { background:#d63050; }
  .error-popover { position:fixed; z-index:2000; background:#2a1a1a; border:1px solid #e94560;
                   border-radius:6px; padding:12px 16px; max-width:500px; color:#e0e0e0;
                   font-size:0.8rem; line-height:1.4; box-shadow:0 4px 12px rgba(0,0,0,0.5); }
  .error-popover .close-pop { float:right; cursor:pointer; color:#8899aa; margin-left:8px; }
  .error-popover .close-pop:hover { color:#e94560; }
</style>
</head>
<body>
<div id="revertModal" class="modal-overlay">
  <div class="modal-box">
    <h3>Confirm Revert</h3>
    <p id="revertModalMsg"></p>
    <label><input type="checkbox" id="revertSkipCheck"> Do not show this again</label>
    <div class="modal-btns">
      <button class="modal-btn-cancel" onclick="revertModalResolve(false)">Cancel</button>
      <button class="modal-btn-confirm" onclick="revertModalResolve(true)">Revert</button>
    </div>
  </div>
</div>
<div class="layout">
<div class="main-panel">
  <h1>Pipeline Dashboard <span id="mode" class="mode dry">DRY RUN</span></h1>
  <div style="margin-bottom:8px"><a href="/accounting" style="color:#e94560;font-size:0.85rem;text-decoration:none;border:1px solid #2a2a4e;padding:4px 12px;border-radius:4px;background:#16213e">Accounting &rarr;</a></div>
  <div class="subtitle" id="subtitle">Loading...</div>
  <div id="retryBanner" class="retry-banner" style="display:none"></div>

  <div class="controls">
    <svg class="progress-ring" viewBox="0 0 52 52">
      <circle class="bg" cx="26" cy="26" r="22"/>
      <circle class="fg" id="ring" cx="26" cy="26" r="22" stroke-dasharray="138.2" stroke-dashoffset="0"/>
    </svg>
    <div>
      <div class="countdown" id="countdown">--</div>
      <div class="countdown-label" id="countdownLabel">Next cycle</div>
    </div>
    <button class="btn" id="stepBtn" onclick="doStep()">Run Now</button>
    <button class="btn" id="pauseBtn" onclick="togglePause()" style="background:#ff9800">Resume</button>
    <div class="project-tabs" id="projectTabs"></div>
  </div>

  <div class="grid">
    <div class="card"><div class="number" id="pending">-</div><div class="label">Pending</div></div>
    <div class="card"><div class="number" id="research">-</div><div class="label">Research</div></div>
    <div class="card"><div class="number" id="outline">-</div><div class="label">Outline</div></div>
    <div class="card"><div class="number" id="draft">-</div><div class="label">Draft</div></div>
    <div class="card"><div class="number" id="factcheck">-</div><div class="label">Factcheck</div></div>
    <div class="card"><div class="number" id="draft2">-</div><div class="label">Draft 2</div></div>
    <div class="card"><div class="number" id="done">-</div><div class="label">Done</div></div>
    <div class="card"><div class="number cost" id="cost">-</div><div class="label">Session Cost</div></div>
    <div class="card"><div class="number" id="cycle">-</div><div class="label">Cycles</div></div>
  </div>

  <div style="font-size:0.7rem;color:#546e7a;text-transform:uppercase;letter-spacing:1px;margin-bottom:6px;">On Disk (all time)</div>
  <div class="grid" style="margin-bottom:20px;">
    <div class="card"><div class="number" id="dk-ideas">-</div><div class="label">Ideas</div></div>
    <div class="card"><div class="number" id="dk-research">-</div><div class="label">Research</div></div>
    <div class="card"><div class="number" id="dk-outline">-</div><div class="label">Outline</div></div>
    <div class="card"><div class="number" id="dk-draft">-</div><div class="label">Draft</div></div>
    <div class="card"><div class="number" id="dk-factcheck">-</div><div class="label">Factcheck</div></div>
    <div class="card"><div class="number" id="dk-draft2">-</div><div class="label">Draft 2</div></div>
    <div class="card"><div class="number" id="dk-export">-</div><div class="label">Export</div></div>
    <div class="card"><div class="number cost" id="dk-cost">-</div><div class="label">Total Cost</div></div>
    <div class="card"><div class="number">&nbsp;</div><div class="label">&nbsp;</div></div>
  </div>

  <table>
    <thead>
      <tr><th>Project</th><th>#</th><th>Type</th><th>Revert</th><th>Title</th><th>Stage</th><th>Attributes</th><th>Length</th><th>Status</th><th></th></tr>
    </thead>
    <tbody id="essayTable"></tbody>
  </table>
  <div style="height:48px"></div>

  <div id="globalRevertFooter" style="position:sticky;bottom:0;background:#0f0f23;border-top:2px solid #2a2a4e;padding:10px 16px;display:flex;align-items:center;gap:12px;z-index:100;margin-top:12px;box-shadow:0 -4px 12px rgba(0,0,0,0.5)">
    <span style="color:#8899aa;font-size:0.8rem;text-transform:uppercase;letter-spacing:1px">Global Revert:</span>
    <select id="globalRevertSelect" style="background:#16213e;color:#e0e0e0;border:1px solid #2a2a4e;border-radius:4px;font-size:0.8rem;padding:4px 8px;cursor:pointer">
      <option value="">Select stage...</option>
      <option value="ideas">ideas</option>
      <option value="research">research</option>
      <option value="outline">outline</option>
      <option value="draft">draft</option>
      <option value="factcheck">factcheck</option>
      <option value="illustrate">illustrate</option>
      <option value="draft2">draft2</option>
      <option value="export">export</option>
    </select>
    <button class="btn" style="padding:6px 16px;font-size:0.8rem" onclick="doGlobalRevert()">Revert All</button>
    <span id="globalRevertStatus" style="color:#8899aa;font-size:0.8rem"></span>
  </div>
</div>

<div class="log-panel">
  <div class="log-header">
    <h2>Logs</h2>
    <div class="log-settings">
      <label><input type="checkbox" id="verboseToggle" onchange="toggleVerbose()"> Verbose</label>
      <div class="interval-ctrl">
        <label style="margin:0">Interval:</label>
        <input type="number" id="intervalInput" min="5" step="5" style="width:50px">
      </div>
      <div class="interval-ctrl">
        <label style="margin:0">Read:</label>
        <input type="number" id="readMeanInput" min="1" max="15" step="0.5" style="width:50px" title="Target reading time (minutes)">
        <label style="margin:0">&plusmn;</label>
        <input type="number" id="readSpreadInput" min="0.5" max="3" step="0.5" style="width:45px" title="Std deviation (minutes)">
        <button class="save-btn" id="saveIntervalBtn" onclick="saveSettings()">Save</button>
      </div>
    </div>
  </div>
  <div class="log-body" id="logBody"></div>
  <div class="log-count"><span id="logCount"></span> <span id="idleIndicator" style="color:#ff9800;display:none"></span> <button class="copy-btn" onclick="copyLogs()">Copy</button></div>
</div>
</div>

<script>
let cycleInterval = 15, secondsLeft = 15, activeProject = localStorage.getItem('activeProject') || '', isRunning = false, isPaused = true;
let verbose = false;
var lastEssays = [];
const circ = 2 * Math.PI * 22;

function updateRing() {
  const label = document.getElementById('countdownLabel');
  const ring = document.getElementById('ring');
  const cd = document.getElementById('countdown');
  const pct = cycleInterval > 0 ? secondsLeft / cycleInterval : 0;
  ring.setAttribute('stroke-dashoffset', String(circ * (1 - pct)));
  cd.textContent = secondsLeft + 's';
  if (isRunning) {
    ring.style.stroke = '#ff9800';
    label.textContent = 'Running';
  } else if (isPaused) {
    ring.style.stroke = '#888';
    label.textContent = 'Paused';
  } else {
    ring.style.stroke = '#e94560';
    label.textContent = 'Next cycle';
  }
}

setInterval(function() { if (secondsLeft > 0) secondsLeft--; updateRing(); }, 1000);

var lastLogText = '';
var lastLogVersion = -1;
var lastNewLogTime = Date.now();
async function refreshLogs() {
  try {
    var data = await fetch('/api/logs?verbose=' + verbose).then(function(r) { return r.json(); });
    if (data.version === lastLogVersion) {
      updateIdleIndicator();
      return;
    }
    lastLogVersion = data.version;
    lastNewLogTime = Date.now();
    var reversed = (data.entries || []).slice().reverse();
    var box = document.getElementById('logBody');
    box.innerHTML = reversed.map(function(e) {
      return '<div class="log-line ' + e.level + '"><span style="color:#546e7a">' + esc(e.time) + '</span> ' + esc(e.message) + '</div>';
    }).join('');
    lastLogText = reversed.map(function(e) { return e.time + ' ' + e.message; }).join('\n');
    document.getElementById('logCount').textContent = reversed.length + ' entries';
    updateIdleIndicator();
  } catch(e) {}
}
function updateIdleIndicator() {
  var el = document.getElementById('idleIndicator');
  var ago = Math.floor((Date.now() - lastNewLogTime) / 1000);
  if (ago >= 5) {
    el.textContent = 'idle ' + ago + 's';
    el.style.display = 'inline';
  } else {
    el.style.display = 'none';
  }
}

async function refresh() {
  try {
    var s = await fetch('/api/status').then(function(r) { return r.json(); });
  } catch(e) { return; }
  cycleInterval = s.cycle_interval || 15;
  secondsLeft = s.seconds_left;
  isRunning = s.cycle_running || false;
  isPaused = s.paused || false;
  var pauseBtn = document.getElementById('pauseBtn');
  if (isPaused) {
    pauseBtn.textContent = 'Resume';
    pauseBtn.style.background = '#ff9800';
  } else {
    pauseBtn.textContent = 'Pause';
    pauseBtn.style.background = '#e94560';
  }
  if (!document.getElementById('intervalInput').matches(':focus')) {
    document.getElementById('intervalInput').value = cycleInterval;
  }
  if (!document.getElementById('readMeanInput').matches(':focus')) {
    document.getElementById('readMeanInput').value = s.read_mean || 5;
  }
  if (!document.getElementById('readSpreadInput').matches(':focus')) {
    document.getElementById('readSpreadInput').value = s.read_spread || 1;
  }
  if (s.verbose !== undefined) {
    verbose = s.verbose;
    document.getElementById('verboseToggle').checked = verbose;
  }
  if (s.skip_revert_confirm !== undefined) {
    skipRevertConfirm = s.skip_revert_confirm;
  }
  document.getElementById('pending').textContent = s.summary.pending || 0;
  document.getElementById('research').textContent = s.summary.research || 0;
  document.getElementById('outline').textContent = s.summary.outline || 0;
  document.getElementById('draft').textContent = s.summary.draft || 0;
  document.getElementById('factcheck').textContent = s.summary.factcheck || 0;
  document.getElementById('draft2').textContent = s.summary.draft2 || 0;
  document.getElementById('done').textContent = s.session_done || 0;
  document.getElementById('cost').textContent = '$' + s.total_cost.toFixed(2);
  document.getElementById('cycle').textContent = s.cycle;
  var mode = document.getElementById('mode');
  if (s.dry_run) { mode.textContent = 'DRY RUN'; mode.className = 'mode dry'; }
  else { mode.textContent = 'LIVE'; mode.className = 'mode live'; }
  var banner = document.getElementById('retryBanner');
  if (s.retry_message) { banner.textContent = s.retry_message; banner.style.display = 'block'; }
  else { banner.style.display = 'none'; }
  var projects = s.projects || [];
  document.getElementById('subtitle').textContent =
    projects.map(function(p) { return p.name; }).join(', ') + ' \u2014 ' + (s.summary.done||0) + ' done of ' +
    Object.values(s.summary).reduce(function(a,b) { return a+b; },0) + ' total';

  var tabs = document.getElementById('projectTabs');
  var tabsHtml = '<div class="tab' + (activeProject===''?' active':'') + '" onclick="filterProject(\x27\x27)">All</div>';
  projects.forEach(function(p) {
    var dots = '<span class="book-dots" title="blurb / cover-prompt / cover-image">' +
      '<span class="book-dot ' + (p.has_blurb?'done':'missing') + '" title="blurb"></span>' +
      '<span class="book-dot ' + (p.has_cover_prompt?'done':'missing') + '" title="cover prompt"></span>' +
      '<span class="book-dot ' + (p.has_cover_image?'done':'missing') + '" title="cover image"></span>' +
      '</span>';
    tabsHtml += '<div class="tab' + (activeProject===p.name?' active':'') +
      '" onclick="filterProject(\x27' + esc(p.name) + '\x27)">' + esc(p.name) + dots + '</div>';
  });
  tabs.innerHTML = tabsHtml;
  updateRing();

  var url = activeProject ? '/api/essays?project=' + encodeURIComponent(activeProject) : '/api/essays';
  var essays = await fetch(url).then(function(r) { return r.json(); });
  lastEssays = essays || [];
  var tbody = document.getElementById('essayTable');
  tbody.innerHTML = (essays||[]).map(function(e) {
    var idStr = e.book + '.' + e.part + '.' + e.order;
    var lenStr = e.word_count ? e.word_count.toLocaleString() + '-' + e.read_mins + 'min' : '';
    return '<tr>' +
    '<td>' + esc(e.project) + '</td>' +
    '<td>' + idStr + '</td>' +
    '<td>' + esc(e.type || 'essay') + '</td>' +
    '<td>' + revertDropdown(e) + '</td>' +
    '<td><a href="#" onclick="' + (e.has_docx ? 'openDocx' : 'openFolder') + '(\x27' + encodeURIComponent(e.project) + '\x27,\x27' + encodeURIComponent(e.slug) + '\x27);return false" style="color:#e94560;text-decoration:none">' + (e.stale ? '<span title="Attributes newer than content" style="color:#ff6b6b;font-size:0.5em;vertical-align:middle;margin-right:3px">\u25cf</span>' : '') + esc(e.title) + '</a></td>' +
    '<td><span class="stage-badge stage-' + e.stage + '">' + e.stage + '</span></td>' +
    '<td title="' + attrTooltip(e) + '">' + attrCompact(e) + '</td>' +
    '<td>' + lenStr + '</td>' +
    '<td class="status-' + e.status + '">' + e.status + (e.error ? ' <span onclick="showError(this,\x27' + esc(e.error).replace(/'/g,'\x27') + '\x27)" style="cursor:pointer;color:#ff9800">\u26a0</span>' : '') + '</td>' +
    '<td>' + (e.has_docx ? '<button class="open-btn" onclick="openDocx(\x27' + encodeURIComponent(e.project) + '\x27,\x27' + encodeURIComponent(e.slug) + '\x27)">Open</button>' : '') + '</td>' +
    '</tr>';
  }).join('');
}
function esc(s) { var d = document.createElement('div'); d.textContent = s; return d.innerHTML; }
function attrCompact(e) {
  var parts = [e.structure, e.register].filter(Boolean);
  return parts.length ? esc(parts.join(' \u00b7 ')) : '-';
}
function attrTooltip(e) {
  var lines = [];
  if (e.arc) lines.push('Arc: ' + e.arc);
  if (e.ending) lines.push('Ending: ' + e.ending);
  if (e.structure) lines.push('Structure: ' + e.structure);
  if (e.entry) lines.push('Entry: ' + e.entry);
  if (e.register) lines.push('Register: ' + e.register);
  if (e.setting) lines.push('Setting: ' + e.setting);
  if (e.math_visibility) lines.push('MathVis: ' + e.math_visibility);
  return esc(lines.join('\n'));
}
function revertDropdown(e) {
  var stages = ['ideas','research','outline','draft','factcheck','illustrate','draft2','export'];
  var id = 'rev-' + e.slug;
  var html = '<select id="' + id + '" style="background:#16213e;color:#8899aa;border:1px solid #2a2a4e;border-radius:3px;font-size:0.7rem;padding:1px 2px;cursor:pointer" onfocus="refreshPaused=true" onblur="refreshPaused=false" onchange="doRevert(\x27' + encodeURIComponent(e.project) + '\x27,\x27' + encodeURIComponent(e.slug) + '\x27,this)">';
  html += '<option value="">\u21a9</option>';
  for (var i = 0; i < stages.length; i++) {
    html += '<option value="' + stages[i] + '">' + stages[i] + '</option>';
  }
  html += '</select>';
  return html;
}
var skipRevertConfirm = false;
var revertModalResolve = function() {};
var activePopover = null;
function showError(el, msg) {
  if (activePopover) { activePopover.remove(); activePopover = null; }
  var pop = document.createElement('div');
  pop.className = 'error-popover';
  pop.innerHTML = '<span class="close-pop" onclick="this.parentNode.remove();activePopover=null">&times;</span>' + esc(msg);
  document.body.appendChild(pop);
  var rect = el.getBoundingClientRect();
  pop.style.top = (rect.bottom + 4) + 'px';
  pop.style.left = Math.min(rect.left, window.innerWidth - 520) + 'px';
  activePopover = pop;
}
function showRevertModal(msg) {
  return new Promise(function(resolve) {
    document.getElementById('revertModalMsg').textContent = msg;
    document.getElementById('revertSkipCheck').checked = false;
    var modal = document.getElementById('revertModal');
    modal.classList.add('active');
    revertModalResolve = async function(ok) {
      modal.classList.remove('active');
      if (ok && document.getElementById('revertSkipCheck').checked) {
        skipRevertConfirm = true;
        await fetch('/api/settings', { method:'POST', headers:{'Content-Type':'application/json'}, body:JSON.stringify({skip_revert_confirm:true}) });
      }
      resolve(ok);
    };
  });
}
async function doGlobalRevert() {
  var sel = document.getElementById('globalRevertSelect');
  var stage = sel.value;
  if (!stage) { alert('Select a stage first'); return; }
  var project = activeProject;
  if (!project && lastEssays.length > 0) project = lastEssays[0].project;
  if (!project) { alert('No project found'); return; }
  var ok = await showRevertModal('Revert ALL works in ' + project + ' to ' + stage + '? Files from ' + stage + ' onward will be deleted for every essay.');
  if (!ok) return;
  var status = document.getElementById('globalRevertStatus');
  status.textContent = 'Reverting...';
  status.style.color = '#ff9800';
  var res = await fetch('/api/revert-all?project=' + encodeURIComponent(project) + '&stage=' + encodeURIComponent(stage), { method: 'POST' });
  var data = await res.json();
  if (data.ok) {
    status.textContent = data.reverted + ' reverted';
    status.style.color = '#4caf50';
  } else {
    status.textContent = data.reverted + ' reverted, ' + (data.errors ? data.errors.length : 0) + ' errors';
    status.style.color = '#f44336';
  }
  sel.value = '';
  refresh();
  refreshDiskStats();
  setTimeout(function() { status.textContent = ''; }, 5000);
}
async function doRevert(project, slug, sel) {
  var stage = sel.value;
  if (!stage) { refreshPaused = false; return; }
  if (!skipRevertConfirm) {
    var ok = await showRevertModal('Revert ' + decodeURIComponent(slug) + ' to ' + stage + '? Files from ' + stage + ' onward will be deleted.');
    if (!ok) { sel.value = ''; refreshPaused = false; return; }
  }
  var url = '/api/revert?project=' + project + '&slug=' + slug + '&stage=' + encodeURIComponent(stage);
  var res = await fetch(url, { method: 'POST' });
  var data = await res.json();
  sel.value = '';
  refreshPaused = false;
  if (data.ok) { refresh(); }
  else { alert('Revert failed'); }
}
function filterProject(p) { activeProject = p; localStorage.setItem('activeProject', p); refresh(); }
async function openFolder(project, slug) {
  await fetch('/api/open?project=' + project + '&slug=' + slug);
}
function openDocx(project, slug) {
  fetch('/api/open-docx?project=' + project + '&slug=' + slug);
}
async function doStep() {
  var btn = document.getElementById('stepBtn');
  btn.disabled = true; btn.textContent = 'Running...';
  await fetch('/api/step', { method: 'POST' });
  setTimeout(async function() { await refresh(); btn.disabled = false; btn.textContent = 'Run Now'; }, 2000);
}
async function togglePause() {
  await fetch('/api/pause', { method: 'POST' });
  await refresh();
}
function toggleVerbose() {
  verbose = document.getElementById('verboseToggle').checked;
  saveSettings();
  refreshLogs();
}
function copyLogs() {
  var btn = document.querySelector('.copy-btn');
  navigator.clipboard.writeText(lastLogText).then(function() {
    btn.textContent = 'Copied'; btn.classList.add('copied');
    setTimeout(function() { btn.textContent = 'Copy'; btn.classList.remove('copied'); }, 1500);
  });
}
async function saveSettings() {
  var btn = document.getElementById('saveIntervalBtn');
  var intVal = parseInt(document.getElementById('intervalInput').value, 10);
  var meanVal = parseFloat(document.getElementById('readMeanInput').value);
  var spreadVal = parseFloat(document.getElementById('readSpreadInput').value);
  var body = { verbose: verbose };
  if (intVal >= 5) body.cycle_interval = intVal;
  if (meanVal >= 1 && meanVal <= 15) body.read_mean = meanVal;
  if (spreadVal >= 0.5 && spreadVal <= 3) body.read_spread = spreadVal;
  await fetch('/api/settings', { method: 'POST', headers: {'Content-Type':'application/json'}, body: JSON.stringify(body) });
  btn.textContent = 'Saved';
  btn.classList.add('saved');
  setTimeout(function() { btn.textContent = 'Save'; btn.classList.remove('saved'); }, 1500);
}
async function refreshDiskStats() {
  var res = await fetch('/api/disk-stats');
  var d = await res.json();
  var c = d.counts || {};
  document.getElementById('dk-ideas').textContent = c.ideas || 0;
  document.getElementById('dk-research').textContent = c.research || 0;
  document.getElementById('dk-outline').textContent = c.outline || 0;
  document.getElementById('dk-draft').textContent = c.draft || 0;
  document.getElementById('dk-factcheck').textContent = c.factcheck || 0;
  document.getElementById('dk-draft2').textContent = c.draft2 || 0;
  document.getElementById('dk-export').textContent = c.export || 0;
  document.getElementById('dk-cost').textContent = '$' + d.total_cost.toFixed(2);
}
var refreshPaused = false;
refresh();
refreshLogs();
refreshDiskStats();
setInterval(function() { if (!refreshPaused) refresh(); }, 1000);
setInterval(refreshLogs, 500);
setInterval(refreshDiskStats, 10000);
</script>
</body>
</html>`

const accountingHTML = `<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="utf-8">
<meta name="viewport" content="width=device-width, initial-scale=1">
<title>Pipeline Accounting</title>
<style>
  * { box-sizing: border-box; margin: 0; padding: 0; }
  body { font-family: -apple-system, BlinkMacSystemFont, "Segoe UI", Roboto, sans-serif;
         background: #1a1a2e; color: #e0e0e0; padding: 24px; max-width: 1200px; margin: 0 auto; }
  h1 { color: #e94560; margin-bottom: 4px; font-size: 1.5rem; }
  .subtitle { color: #8899aa; font-size: 0.85rem; margin-bottom: 20px; }
  .nav-link { color: #e94560; font-size: 0.85rem; text-decoration: none;
              border: 1px solid #2a2a4e; padding: 4px 12px; border-radius: 4px;
              background: #16213e; display: inline-block; margin-bottom: 16px; }
  .nav-link:hover { border-color: #e94560; }
  h2 { color: #e94560; font-size: 1.1rem; margin: 24px 0 8px; }
  h3 { color: #8899aa; font-size: 0.9rem; margin: 16px 0 6px; }
  table { width: 100%; border-collapse: collapse; margin-bottom: 16px; }
  th { text-align: left; padding: 8px 12px; background: #16213e; color: #8899aa;
       font-size: 0.75rem; text-transform: uppercase; letter-spacing: 1px; }
  th.num { text-align: right; }
  td { padding: 6px 12px; border-bottom: 1px solid #1a1a3e; font-size: 0.85rem; }
  td.num { text-align: right; font-family: "SF Mono", Monaco, monospace; font-size: 0.8rem; }
  tr:hover td { background: #16213e; }
  tr.total-row td { border-top: 2px solid #2a2a4e; font-weight: bold; color: #e94560; }
  tr.grand-total td { border-top: 3px solid #e94560; font-weight: bold; color: #e94560; font-size: 0.9rem; }
  .cost { color: #ff9800; }
</style>
</head>
<body>
<a href="/" class="nav-link">&larr; Dashboard</a>
<h1>Pipeline Accounting</h1>
<div class="subtitle">{{.EntryCount}} consolidated entries across all projects</div>

<h2>By Book</h2>
<table>
<thead><tr><th>Book</th><th class="num">Runs</th><th class="num">Tokens In</th><th class="num">Tokens Out</th><th class="num">Cost</th><th class="num">Avg Cost</th></tr></thead>
<tbody>
{{range .ByBook}}<tr{{if eq .Label "Total"}} class="total-row"{{end}}>
<td>{{.Label}}</td><td class="num">{{.Runs}}</td><td class="num">{{.TokensIn}}</td><td class="num">{{.TokensOut}}</td><td class="num cost">${{printf "%.2f" .Cost}}</td><td class="num cost">${{printf "%.4f" .AvgCost}}</td>
</tr>{{end}}
</tbody>
</table>

<h2>By Stage</h2>
<table>
<thead><tr><th>Stage</th><th class="num">Runs</th><th class="num">Tokens In</th><th class="num">Tokens Out</th><th class="num">Cost</th><th class="num">Avg Cost</th></tr></thead>
<tbody>
{{range .ByStage}}<tr{{if eq .Label "Total"}} class="total-row"{{end}}>
<td>{{.Label}}</td><td class="num">{{.Runs}}</td><td class="num">{{.TokensIn}}</td><td class="num">{{.TokensOut}}</td><td class="num cost">${{printf "%.2f" .Cost}}</td><td class="num cost">${{printf "%.4f" .AvgCost}}</td>
</tr>{{end}}
</tbody>
</table>

<h2>By Book &times; Stage</h2>
{{range .ByBookStage}}
<h3>Book {{.Book}}</h3>
<table>
<thead><tr><th>Stage</th><th class="num">Runs</th><th class="num">Tokens In</th><th class="num">Tokens Out</th><th class="num">Cost</th><th class="num">Avg Cost</th></tr></thead>
<tbody>
{{range .Rows}}<tr{{if eq .Label "Total"}} class="total-row"{{end}}>
<td>{{.Label}}</td><td class="num">{{.Runs}}</td><td class="num">{{.TokensIn}}</td><td class="num">{{.TokensOut}}</td><td class="num cost">${{printf "%.2f" .Cost}}</td><td class="num cost">${{printf "%.4f" .AvgCost}}</td>
</tr>{{end}}
</tbody>
</table>
{{end}}

<table>
<thead><tr><th>Grand Total</th><th class="num">Runs</th><th class="num">Tokens In</th><th class="num">Tokens Out</th><th class="num">Cost</th><th class="num">Avg Cost</th></tr></thead>
<tbody>
<tr class="grand-total">
<td>All Books</td><td class="num">{{.GrandTotal.Runs}}</td><td class="num">{{.GrandTotal.TokensIn}}</td><td class="num">{{.GrandTotal.TokensOut}}</td><td class="num cost">${{printf "%.2f" .GrandTotal.Cost}}</td><td class="num cost">${{printf "%.4f" .GrandTotal.AvgCost}}</td>
</tr>
</tbody>
</table>
</body>
</html>`
