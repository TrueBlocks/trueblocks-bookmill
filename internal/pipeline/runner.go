package pipeline

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sync"
	"text/template"
)

type docxJob struct {
	ps    *PipelineState
	essay *EssayState
}

type seriesSpecs struct {
	VoiceProfile      string
	VoiceAntiPatterns string
	DraftRules        string
	RevisionRules     string
	PromptTemplates   map[string]*template.Template
	Examples          map[string]string
}

type Runner struct {
	Config      *Config
	Projects    []*PipelineState
	Client      *AnthropicClient
	Log         *log.Logger
	BaseDir     string
	ConfigPath  string
	CLIDryRun   bool
	SeriesSpecs map[string]*seriesSpecs
	docxCh      chan docxJob
	docxWg      sync.WaitGroup
}

func NewRunner(cfg *Config, baseDir string) *Runner {
	r := &Runner{
		Config: cfg,
		Client: &AnthropicClient{
			APIKey:     cfg.API.AnthropicKey,
			APIVersion: cfg.API.Version,
			MaxTokens:  cfg.API.MaxTokens,
			Pricing:    cfg.Pricing,
		},
		Log:         log.New(os.Stdout, "", 0),
		BaseDir:     baseDir,
		SeriesSpecs: make(map[string]*seriesSpecs),
		docxCh:      make(chan docxJob, 200),
	}
	if cfg.ExportYear != "" {
		exportYear = cfg.ExportYear
	}
	r.docxWg.Add(1)
	go r.docxWorker()
	return r
}

func (r *Runner) Shutdown() {
	close(r.docxCh)
	r.docxWg.Wait()
}

func (r *Runner) DiscoverProjects() error {
	projectsDir := filepath.Join(r.BaseDir, "projects")
	entries, err := os.ReadDir(projectsDir)
	if err != nil {
		return fmt.Errorf("reading projects dir: %w", err)
	}

	r.Projects = nil
	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}
		ideasDir := filepath.Join(projectsDir, entry.Name(), "ideas")
		if _, err := os.Stat(ideasDir); os.IsNotExist(err) {
			continue
		}
		ps := NewPipelineState(entry.Name(), filepath.Join(projectsDir, entry.Name()))
		r.Projects = append(r.Projects, ps)
	}

	if len(r.Projects) == 0 {
		return fmt.Errorf("no projects found in %s", projectsDir)
	}

	return nil
}

func (r *Runner) LoadState() error {
	designDir := filepath.Join(r.BaseDir, "design")
	for _, ps := range r.Projects {
		repairs := ps.RepairOrphans()
		for _, msg := range repairs {
			r.Log.Printf("[%s] REPAIR: %s", ps.Project, msg)
		}
		if err := ps.LoadFromDisk(); err != nil {
			return fmt.Errorf("project %s: %w", ps.Project, err)
		}
		n, err := ps.ApplyAttributes(designDir)
		if err != nil {
			r.Log.Printf("[%s] WARNING: apply attributes: %v", ps.Project, err)
		}
		if n > 0 {
			r.Log.Printf("[%s] Applied variation attributes to %d essays", ps.Project, n)
		}
	}
	return nil
}

func (r *Runner) ReloadConfig() {
	if r.ConfigPath == "" {
		return
	}
	updated, err := LoadConfig(r.ConfigPath)
	if err != nil {
		r.Log.Printf("Config reload failed: %v (keeping previous settings)", err)
		return
	}
	port := r.Config.Dashboard.Port
	wasCliDry := r.CLIDryRun
	*r.Config = *updated
	r.Config.Dashboard.Port = port
	if wasCliDry {
		r.Config.Pipeline.DryRun = true
	}
	r.Client.APIKey = r.Config.API.AnthropicKey
}

func (r *Runner) TotalCycles() int {
	total := 0
	for _, ps := range r.Projects {
		total += ps.CycleCount
	}
	return total
}

func (r *Runner) TotalCost() float64 {
	total := 0.0
	for _, ps := range r.Projects {
		total += ps.SessionCost
	}
	return total
}

func (r *Runner) SessionDone() int {
	total := 0
	for _, ps := range r.Projects {
		total += ps.SessionDone
	}
	return total
}

func (r *Runner) RevertedCost() float64 {
	total := 0.0
	for _, ps := range r.Projects {
		total += ps.RevertedCost()
	}
	return total
}
