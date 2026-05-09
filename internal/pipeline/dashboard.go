package pipeline

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"sync"
	"time"
)

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

		func() {
			ps.mu.Lock()
			defer ps.mu.Unlock()
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
		}()
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
	func() {
		ps.mu.Lock()
		defer ps.mu.Unlock()
		for slug := range ps.Essays {
			slugs = append(slugs, slug)
		}
	}()

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
		func() {
			ps.mu.Lock()
			defer ps.mu.Unlock()
			totalCost += ps.TotalCost + ps.RevertedCost()
		}()
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

