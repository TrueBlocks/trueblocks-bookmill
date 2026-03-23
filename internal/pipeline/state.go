package pipeline

import (
	"encoding/json"
	"fmt"
	randv2 "math/rand/v2"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"sync"
	"syscall"
	"time"

	"gopkg.in/yaml.v3"
)

type Stage int

const (
	StageIdeas Stage = iota
	StageResearch
	StageOutline
	StageDraft
	StageFactcheck
	StageContinuity
	StageIllustrate
	StageDraft2
	StageRevision
	StageExport
	StageDone
)

var stageNames = []string{"ideas", "research", "outline", "draft", "factcheck", "continuity", "illustrate", "draft2", "revision", "export"}

func (s Stage) String() string {
	if int(s) < len(stageNames) {
		return stageNames[s]
	}
	return "done"
}

func (s Stage) Dir() string { return s.String() }

func NextStage(s Stage) Stage {
	if s < StageExport {
		return s + 1
	}
	return StageDone
}

func StageFromString(s string) Stage {
	for i, name := range stageNames {
		if name == s {
			return Stage(i)
		}
	}
	return StageIdeas
}

type EssayMeta struct {
	Slug           string  `yaml:"slug"`
	Title          string  `yaml:"title"`
	Subtitle       string  `yaml:"subtitle,omitempty"`
	Type           string  `yaml:"type"`
	Series         string  `yaml:"series,omitempty"`
	Book           string  `yaml:"book"`
	Part           int     `yaml:"part"`
	PartTitle      string  `yaml:"part_title"`
	Order          int     `yaml:"order"`
	Status         string  `yaml:"status"`
	Model          string  `yaml:"model"`
	Arc            string  `yaml:"arc,omitempty"`
	Ending         string  `yaml:"ending,omitempty"`
	Structure      string  `yaml:"structure,omitempty"`
	Entry          string  `yaml:"entry,omitempty"`
	Register       string  `yaml:"register,omitempty"`
	Setting        string  `yaml:"setting,omitempty"`
	MathVisibility string  `yaml:"math_visibility,omitempty"`
	ReadMean       float64 `yaml:"read_mean,omitempty"`
	ReadSpread     float64 `yaml:"read_spread,omitempty"`
	Created        string  `yaml:"created"`
	Started        string  `yaml:"started,omitempty"`
	Completed      string  `yaml:"completed,omitempty"`
	Tokens         int     `yaml:"tokens,omitempty"`
	TokensOut      int     `yaml:"tokens_out,omitempty"`
	Cost           float64 `yaml:"cost,omitempty"`
	Error          string  `yaml:"error,omitempty"`
	Retries        int     `yaml:"retries,omitempty"`
}

type EssayState struct {
	Slug           string
	Title          string
	Subtitle       string
	Type           string
	Series         string
	Book           string
	Part           int
	PartTitle      string
	Order          int
	Arc            string
	Ending         string
	Structure      string
	Entry          string
	Register       string
	Setting        string
	MathVisibility string
	ReadMean       float64
	ReadSpread     float64
	CurrentStage   Stage
	Status         string
	ErrorRetries   int
	Meta           map[Stage]*EssayMeta
}

func (e *EssayState) NextAction() Stage {
	return e.NextActionForGenre(nil)
}

func (e *EssayState) NextActionForGenre(g *Genre) Stage {
	if e.Status == "error" {
		return e.CurrentStage
	}
	if g != nil && len(g.Stages) > 0 {
		currentName := e.CurrentStage.String()
		if currentName == "ideas" && e.Status == "pending" {
			return StageFromString(g.FirstContentStage())
		}
		if e.Status == "final" {
			next := g.NextStageAfter(currentName)
			return StageFromString(next)
		}
		return StageDone
	}
	if e.CurrentStage == StageIdeas && e.Status == "pending" {
		return StageResearch
	}
	if e.Status == "final" && e.CurrentStage < StageExport {
		return NextStage(e.CurrentStage)
	}
	return StageDone
}

func (e *EssayState) IsDone() bool {
	return e.CurrentStage == StageExport && e.Status == "final"
}

func (e *EssayState) IsDoneForGenre(g *Genre) bool {
	if g == nil || len(g.Stages) == 0 {
		return e.IsDone()
	}
	lastStage := g.Stages[len(g.Stages)-1]
	return e.CurrentStage == StageFromString(lastStage) && e.Status == "final"
}

func (e *EssayState) IsAvailable() bool {
	next := e.NextAction()
	if next == StageDone || e.Status == "in-progress" {
		return false
	}
	if e.Status == "error" && e.ErrorRetries >= 3 {
		return false
	}
	return true
}

type PipelineState struct {
	Project     string
	BaseDir     string
	Genre       *Genre
	Essays      map[string]*EssayState
	CycleCount  int
	TotalCost   float64
	SessionCost float64
	SessionDone int
	mu          sync.Mutex
}

func NewPipelineState(project, baseDir string) *PipelineState {
	for _, name := range append(stageNames, "images") {
		os.MkdirAll(filepath.Join(baseDir, name), 0755)
	}
	genre, _ := LoadGenreFromProject(baseDir)
	return &PipelineState{
		Project: project,
		BaseDir: baseDir,
		Genre:   genre,
		Essays:  make(map[string]*EssayState),
	}
}

func (ps *PipelineState) LoadFromDisk() error {
	ps.mu.Lock()
	defer ps.mu.Unlock()
	ps.Essays = make(map[string]*EssayState)
	ps.TotalCost = 0

	for stageIdx, stageName := range stageNames {
		stage := Stage(stageIdx)
		dir := filepath.Join(ps.BaseDir, stageName)
		entries, err := os.ReadDir(dir)
		if err != nil {
			if os.IsNotExist(err) {
				continue
			}
			return fmt.Errorf("reading %s: %w", dir, err)
		}

		for _, entry := range entries {
			if filepath.Ext(entry.Name()) != ".yaml" {
				continue
			}
			slug := entry.Name()
			slug = slug[:len(slug)-len(".meta.yaml")]

			data, err := os.ReadFile(filepath.Join(dir, entry.Name()))
			if err != nil {
				return fmt.Errorf("reading %s/%s: %w", stageName, entry.Name(), err)
			}

			var meta EssayMeta
			if err := yaml.Unmarshal(data, &meta); err != nil {
				return fmt.Errorf("parsing %s/%s: %w", stageName, entry.Name(), err)
			}

			if meta.Slug == "" && meta.Title == "" {
				continue
			}

			essay, ok := ps.Essays[slug]
			if !ok {
				essay = &EssayState{
					Slug:           meta.Slug,
					Title:          meta.Title,
					Subtitle:       meta.Subtitle,
					Type:           meta.Type,
					Series:         meta.Series,
					Book:           meta.Book,
					Part:           meta.Part,
					PartTitle:      meta.PartTitle,
					Order:          meta.Order,
					Arc:            meta.Arc,
					Ending:         meta.Ending,
					Structure:      meta.Structure,
					Entry:          meta.Entry,
					Register:       meta.Register,
					Setting:        meta.Setting,
					MathVisibility: meta.MathVisibility,
					ReadMean:       meta.ReadMean,
					ReadSpread:     meta.ReadSpread,
					CurrentStage:   stage,
					Status:         meta.Status,
					Meta:           make(map[Stage]*EssayMeta),
				}
				ps.Essays[slug] = essay
			}

			essay.Meta[stage] = &meta
			if meta.Subtitle != "" {
				essay.Subtitle = meta.Subtitle
			}
			if meta.Type != "" {
				essay.Type = meta.Type
			}
			if meta.Series != "" {
				essay.Series = meta.Series
			}
			if meta.Arc != "" {
				essay.Arc = meta.Arc
			}
			if meta.Ending != "" {
				essay.Ending = meta.Ending
			}
			if meta.Structure != "" {
				essay.Structure = meta.Structure
			}
			if meta.Entry != "" {
				essay.Entry = meta.Entry
			}
			if meta.Register != "" {
				essay.Register = meta.Register
			}
			if meta.Setting != "" {
				essay.Setting = meta.Setting
			}
			if meta.MathVisibility != "" {
				essay.MathVisibility = meta.MathVisibility
			}
			if meta.ReadMean > 0 {
				essay.ReadMean = meta.ReadMean
			}
			if meta.ReadSpread > 0 {
				essay.ReadSpread = meta.ReadSpread
			}
			if stage > essay.CurrentStage {
				essay.CurrentStage = stage
				essay.Status = meta.Status
			}
			if meta.Status == "error" && stage == essay.CurrentStage {
				essay.ErrorRetries = meta.Retries
			}
			ps.TotalCost += meta.Cost
		}
	}

	exportDir := filepath.Join(ps.BaseDir, "export")
	for _, essay := range ps.Essays {
		if essay.CurrentStage == StageExport && essay.Status == "final" {
			continue
		}
		docxPath := filepath.Join(exportDir, exportFilename(essay))
		if _, err := os.Stat(docxPath); err == nil {
			essay.CurrentStage = StageExport
			essay.Status = "final"
		}
	}

	return nil
}

func (ps *PipelineState) WriteMeta(stage Stage, meta *EssayMeta) error {
	dir := filepath.Join(ps.BaseDir, stage.Dir())
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}

	data, err := yaml.Marshal(meta)
	if err != nil {
		return err
	}

	metaPath := filepath.Join(dir, meta.Slug+".meta.yaml")
	return os.WriteFile(metaPath, data, 0644)
}

func (ps *PipelineState) WriteContent(stage Stage, slug, content string) error {
	dir := filepath.Join(ps.BaseDir, stage.Dir())
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}

	mdPath := filepath.Join(dir, slug+".md")
	return os.WriteFile(mdPath, []byte(content), 0644)
}

func ParseDebug(debug string) (book string, part, order int, ok bool) {
	parts := strings.Split(debug, ".")
	if len(parts) != 3 {
		return "", 0, 0, false
	}
	book = parts[0]
	part, err := strconv.Atoi(parts[1])
	if err != nil {
		return "", 0, 0, false
	}
	order, err = strconv.Atoi(parts[2])
	if err != nil {
		return "", 0, 0, false
	}
	return book, part, order, true
}

type attributeEntry struct {
	Slug           string `yaml:"slug"`
	Arc            string `yaml:"arc"`
	Ending         string `yaml:"ending"`
	Structure      string `yaml:"structure"`
	Entry          string `yaml:"entry"`
	Register       string `yaml:"register"`
	Setting        string `yaml:"setting"`
	MathVisibility string `yaml:"math_visibility"`
}

type attributeFile struct {
	Essays []attributeEntry `yaml:"essays"`
}

func (ps *PipelineState) ApplyAttributes(designDir string) (int, error) {
	pattern := filepath.Join(designDir, "*attributes.yaml")
	matches, err := filepath.Glob(pattern)
	if err != nil {
		return 0, err
	}

	lookup := make(map[string]attributeEntry)
	for _, path := range matches {
		data, err := os.ReadFile(path)
		if err != nil {
			return 0, fmt.Errorf("reading %s: %w", path, err)
		}
		var af attributeFile
		if err := yaml.Unmarshal(data, &af); err != nil {
			return 0, fmt.Errorf("parsing %s: %w", path, err)
		}
		for _, e := range af.Essays {
			lookup[e.Slug] = e
		}
	}

	updated := 0
	for slug, essay := range ps.Essays {
		attr, ok := lookup[slug]
		if !ok {
			continue
		}
		changed := false
		if attr.Arc != "" && essay.Arc != attr.Arc {
			essay.Arc = attr.Arc
			changed = true
		}
		if attr.Ending != "" && essay.Ending != attr.Ending {
			essay.Ending = attr.Ending
			changed = true
		}
		if attr.Structure != "" && essay.Structure != attr.Structure {
			essay.Structure = attr.Structure
			changed = true
		}
		if attr.Entry != "" && essay.Entry != attr.Entry {
			essay.Entry = attr.Entry
			changed = true
		}
		if attr.Register != "" && essay.Register != attr.Register {
			essay.Register = attr.Register
			changed = true
		}
		if attr.Setting != "" && essay.Setting != attr.Setting {
			essay.Setting = attr.Setting
			changed = true
		}
		if attr.MathVisibility != "" && essay.MathVisibility != attr.MathVisibility {
			essay.MathVisibility = attr.MathVisibility
			changed = true
		}
		if changed {
			if meta, ok := essay.Meta[StageIdeas]; ok {
				meta.Arc = essay.Arc
				meta.Ending = essay.Ending
				meta.Structure = essay.Structure
				meta.Entry = essay.Entry
				meta.Register = essay.Register
				meta.Setting = essay.Setting
				meta.MathVisibility = essay.MathVisibility
				if err := ps.WriteMeta(StageIdeas, meta); err != nil {
					return updated, fmt.Errorf("writing meta for %s: %w", slug, err)
				}
			}
			updated++
		}
	}
	return updated, nil
}

type AccountingEntry struct {
	Slug       string  `json:"slug"`
	Stage      string  `json:"stage"`
	Series     string  `json:"series,omitempty"`
	Book       string  `json:"book,omitempty"`
	RevertedAt string  `json:"reverted_at,omitempty"`
	RevertedTo string  `json:"reverted_to,omitempty"`
	TokensIn   int     `json:"tokens_in,omitempty"`
	TokensOut  int     `json:"tokens_out,omitempty"`
	Cost       float64 `json:"cost,omitempty"`
	Model      string  `json:"model,omitempty"`
	Count      int     `json:"count,omitempty"`
}

type AccountingFile struct {
	Reverted []AccountingEntry `json:"reverted"`
}

func lockAccounting(path string) (*os.File, error) {
	lockPath := path + ".lock"
	f, err := os.OpenFile(lockPath, os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		return nil, fmt.Errorf("opening lock file: %w", err)
	}
	if err := syscall.Flock(int(f.Fd()), syscall.LOCK_EX); err != nil {
		f.Close()
		return nil, fmt.Errorf("acquiring lock: %w", err)
	}
	return f, nil
}

func unlockAccounting(f *os.File) {
	if f == nil {
		return
	}
	syscall.Flock(int(f.Fd()), syscall.LOCK_UN)
	f.Close()
}

func loadAccounting(path string) (*AccountingFile, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			return &AccountingFile{}, nil
		}
		return nil, err
	}
	var af AccountingFile
	if err := json.Unmarshal(data, &af); err != nil {
		return nil, err
	}
	return &af, nil
}

func saveAccounting(path string, af *AccountingFile) error {
	data, err := json.MarshalIndent(af, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(path, data, 0644)
}

func consolidateAccounting(af *AccountingFile) {
	type key struct{ slug, stage string }
	merged := make(map[key]*AccountingEntry)
	var order []key
	for _, e := range af.Reverted {
		k := key{e.Slug, e.Stage}
		if existing, ok := merged[k]; ok {
			existing.TokensIn += e.TokensIn
			existing.TokensOut += e.TokensOut
			existing.Cost += e.Cost
			if e.Count > 0 {
				existing.Count += e.Count
			} else {
				existing.Count++
			}
		} else {
			copy := e
			copy.RevertedAt = ""
			copy.RevertedTo = ""
			if copy.Count == 0 {
				copy.Count = 1
			}
			merged[k] = &copy
			order = append(order, k)
		}
	}
	af.Reverted = make([]AccountingEntry, 0, len(order))
	for _, k := range order {
		af.Reverted = append(af.Reverted, *merged[k])
	}
}

func (ps *PipelineState) accountingPath() string {
	return filepath.Join(ps.BaseDir, "accounting.json")
}

func (ps *PipelineState) RevertedCost() float64 {
	af, err := loadAccounting(ps.accountingPath())
	if err != nil {
		return 0
	}
	var total float64
	for _, e := range af.Reverted {
		total += e.Cost
	}
	return total
}

func (ps *PipelineState) RevertToStage(slug string, target Stage) ([]string, error) {
	ps.mu.Lock()
	defer ps.mu.Unlock()

	essay, ok := ps.Essays[slug]
	if !ok {
		return nil, fmt.Errorf("essay %q not found", slug)
	}

	if essay.CurrentStage < target {
		return nil, nil
	}

	var removed []string
	start := int(target)
	if target <= StageIdeas {
		start = int(StageResearch)
	}

	now := nowString()
	lockFile, lockErr := lockAccounting(ps.accountingPath())
	if lockErr != nil {
		return nil, fmt.Errorf("locking accounting: %w", lockErr)
	}
	defer unlockAccounting(lockFile)

	af, _ := loadAccounting(ps.accountingPath())
	if af == nil {
		af = &AccountingFile{}
	}

	for stageIdx := start; stageIdx <= int(StageExport); stageIdx++ {
		stage := Stage(stageIdx)
		dir := filepath.Join(ps.BaseDir, stage.Dir())

		metaPath := filepath.Join(dir, slug+".meta.yaml")
		if data, err := os.ReadFile(metaPath); err == nil {
			var meta EssayMeta
			if yaml.Unmarshal(data, &meta) == nil && meta.Cost > 0 {
				af.Reverted = append(af.Reverted, AccountingEntry{
					Slug:       slug,
					Stage:      stage.String(),
					Series:     ps.Project,
					Book:       essay.Book,
					RevertedAt: now,
					RevertedTo: target.String(),
					TokensIn:   meta.Tokens,
					TokensOut:  meta.TokensOut,
					Cost:       meta.Cost,
					Model:      meta.Model,
				})
			}
		}

		for _, ext := range []string{".md", ".meta.yaml"} {
			path := filepath.Join(dir, slug+ext)
			if _, err := os.Stat(path); err == nil {
				if err := os.Remove(path); err != nil {
					return removed, fmt.Errorf("removing %s: %w", path, err)
				}
				removed = append(removed, stage.String()+ext)
			}
		}
		delete(essay.Meta, stage)
	}

	consolidateAccounting(af)
	saveAccounting(ps.accountingPath(), af)

	exportDir := filepath.Join(ps.BaseDir, "export")
	docxPath := filepath.Join(exportDir, exportFilenameFromParts(essay))
	if _, err := os.Stat(docxPath); err == nil {
		if err := os.Remove(docxPath); err != nil {
			return removed, fmt.Errorf("removing docx: %w", err)
		}
		removed = append(removed, "export.docx")
	}

	if target <= StageIdeas {
		essay.CurrentStage = StageIdeas
		essay.Status = "pending"
		if meta, ok := essay.Meta[StageIdeas]; ok {
			meta.Status = "pending"
			if err := ps.WriteMeta(StageIdeas, meta); err != nil {
				return removed, fmt.Errorf("resetting ideas meta: %w", err)
			}
		}
	} else {
		prev := target - 1
		essay.CurrentStage = prev
		essay.Status = "final"
	}

	return removed, nil
}

func exportFilenameFromParts(essay *EssayState) string {
	switch essay.Type {
	case "section":
		return fmt.Sprintf("cSection - %s - %s.%02d.00 %s.docx",
			exportYear, essay.Book, essay.Part, essay.Title)
	case "introduction":
		return fmt.Sprintf("cChapter - %s - %s.00.00 %s.docx",
			exportYear, essay.Book, essay.Title)
	default:
		return fmt.Sprintf("cChapter - %s - %s.%02d.%02d %s.docx",
			exportYear, essay.Book, essay.Part, essay.Order, essay.Title)
	}
}

func (ps *PipelineState) SelectForCycle(cfg *PipelineConfig) []*EssayState {
	ps.mu.Lock()
	defer ps.mu.Unlock()

	if cfg.Debug != "" {
		book, part, order, ok := ParseDebug(cfg.Debug)
		if !ok {
			return nil
		}
		for _, essay := range ps.Essays {
			if essay.Book == book && essay.Part == part && essay.Order == order {
				return []*EssayState{essay}
			}
		}
		return nil
	}

	maxActions := cfg.MaxPerCycle

	var newEssays, continuingEssays []*EssayState
	for _, essay := range ps.Essays {
		if !essay.IsAvailable() {
			continue
		}
		if essay.CurrentStage == StageIdeas {
			newEssays = append(newEssays, essay)
		} else {
			continuingEssays = append(continuingEssays, essay)
		}
	}

	sortByBookOrder(newEssays)
	randv2.Shuffle(len(newEssays), func(i, j int) { newEssays[i], newEssays[j] = newEssays[j], newEssays[i] })
	sortByBookOrder(continuingEssays)

	var selected []*EssayState
	newCount := 0
	maxContinuing := maxActions
	if len(newEssays) > 0 && cfg.NewPerCycle > 0 {
		reserved := cfg.NewPerCycle
		if reserved > len(newEssays) {
			reserved = len(newEssays)
		}
		if maxContinuing > maxActions-reserved {
			maxContinuing = maxActions - reserved
		}
		if maxContinuing < 0 {
			maxContinuing = 0
		}
	}
	for _, e := range continuingEssays {
		if len(selected) >= maxContinuing {
			break
		}
		selected = append(selected, e)
	}
	for _, e := range newEssays {
		if len(selected) >= maxActions {
			break
		}
		if newCount >= cfg.NewPerCycle {
			break
		}
		selected = append(selected, e)
		newCount++
	}

	return selected
}

func sortByBookOrder(essays []*EssayState) {
	sort.Slice(essays, func(i, j int) bool {
		if essays[i].Book != essays[j].Book {
			return essays[i].Book < essays[j].Book
		}
		if essays[i].Part != essays[j].Part {
			return essays[i].Part < essays[j].Part
		}
		return essays[i].Order < essays[j].Order
	})
}

func (ps *PipelineState) SnapshotEssays() []*EssayState {
	ps.mu.Lock()
	defer ps.mu.Unlock()
	out := make([]*EssayState, 0, len(ps.Essays))
	for _, e := range ps.Essays {
		out = append(out, e)
	}
	return out
}

func (ps *PipelineState) Summary() map[string]int {
	ps.mu.Lock()
	defer ps.mu.Unlock()
	counts := map[string]int{
		"pending": 0, "research": 0, "outline": 0,
		"draft": 0, "factcheck": 0, "draft2": 0, "illustrate": 0, "done": 0, "error": 0,
	}
	for _, e := range ps.Essays {
		if e.Status == "error" {
			counts["error"]++
		} else if e.IsDone() {
			counts["done"]++
		} else if e.CurrentStage == StageIdeas {
			counts["pending"]++
		} else {
			counts[e.CurrentStage.String()]++
		}
	}
	return counts
}

func (ps *PipelineState) RepairOrphans() []string {
	var repairs []string
	for stageIdx, stageName := range stageNames {
		stage := Stage(stageIdx)
		if stage == StageIdeas {
			continue
		}
		dir := filepath.Join(ps.BaseDir, stageName)
		entries, err := os.ReadDir(dir)
		if err != nil {
			continue
		}
		for _, entry := range entries {
			if filepath.Ext(entry.Name()) != ".yaml" {
				continue
			}
			slug := entry.Name()[:len(entry.Name())-len(".meta.yaml")]
			yamlPath := filepath.Join(dir, entry.Name())
			mdPath := filepath.Join(dir, slug+".md")

			if _, err := os.Stat(mdPath); os.IsNotExist(err) {
				os.Remove(yamlPath)
				repairs = append(repairs, fmt.Sprintf("removed orphan %s/%s (no .md file)", stageName, entry.Name()))
				continue
			}

			data, err := os.ReadFile(yamlPath)
			if err != nil {
				continue
			}
			var meta EssayMeta
			if err := yaml.Unmarshal(data, &meta); err != nil {
				continue
			}
			if meta.Status == "in-progress" {
				meta.Status = "final"
				meta.Completed = nowString()
				if fixed, err := yaml.Marshal(&meta); err == nil {
					os.WriteFile(yamlPath, fixed, 0644)
					repairs = append(repairs, fmt.Sprintf("fixed %s/%s (in-progress → final)", stageName, entry.Name()))
				}
			}
		}
	}
	return repairs
}

func nowString() string {
	return time.Now().Format("2006-01-02T15:04:05")
}
