package pipeline

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"math"
	randv2 "math/rand/v2"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
	"sync"
	"text/template"
	"time"
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
		Config:      cfg,
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

// exportYear is the year used in export filenames. Set by NewRunner from config.
var exportYear = time.Now().Format("2006")

func exportFilename(essay *EssayState) string {
	return exportFilenameFromParts(essay)
}

func shouldSkipStage(itemType string, stage Stage) bool {
	switch itemType {
	case "section":
		return stage != StageDraft2 && stage != StageExport
	case "introduction":
		return stage == StageResearch || stage == StageFactcheck || stage == StageIllustrate
	}
	return false
}

func (r *Runner) autoSkipStage(ps *PipelineState, essay *EssayState, targetStage Stage) error {
	prevStage := targetStage - 1
	if prevStage < StageIdeas {
		prevStage = StageIdeas
	}

	srcPath := filepath.Join(ps.BaseDir, prevStage.Dir(), essay.Slug+".md")
	content, err := os.ReadFile(srcPath)
	if err != nil {
		return fmt.Errorf("auto-skip reading %s: %w", srcPath, err)
	}

	if err := ps.WriteContent(targetStage, essay.Slug, string(content)); err != nil {
		return fmt.Errorf("auto-skip writing content: %w", err)
	}

	result := &APIResult{}
	r.markComplete(ps, essay, targetStage, "skip", result)
	r.Log.Printf("    [skip] %s: %s (auto-skipped for %s)", essay.Slug, targetStage, essay.Type)
	return nil
}

func (r *Runner) docxWorker() {
	defer r.docxWg.Done()
	for job := range r.docxCh {
		r.Log.Printf("  [docx] %s: export starting", job.essay.Slug)
		if err := r.exportEssay(job.ps, job.essay); err != nil {
			r.Log.Printf("  [docx] ERROR export %s: %v", job.essay.Slug, err)
			r.markError(job.ps, job.essay, StageExport, err)
			continue
		}

		dstFile := filepath.Join(job.ps.BaseDir, "export", exportFilename(job.essay))

		r.Log.Printf("  [docx] %s: Word upgrade starting", job.essay.Slug)
		if err := r.upgradeDocx(dstFile); err != nil {
			r.Log.Printf("  [docx] WARNING: Word upgrade %s: %v", job.essay.Slug, err)
		}

		if err := r.swapImages(job.ps, job.essay, dstFile); err != nil {
			r.Log.Printf("  [docx] ERROR imageswap %s: %v", job.essay.Slug, err)
		}

		r.Log.Printf("  [docx] %s: complete", job.essay.Slug)
	}
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

func (r *Runner) loadSpecs() {
	seriesNames := make(map[string]bool)
	for _, ps := range r.Projects {
		seriesNames[ps.Project] = true
	}

	for series := range seriesNames {
		specs := r.loadSeriesSpecs(series)
		r.SeriesSpecs[series] = specs
	}
}

func (r *Runner) resolveSpecFile(series, filename string) ([]byte, error) {
	seriesPath := filepath.Join(r.BaseDir, "specs", "prompts", "series", series, filename)
	if data, err := os.ReadFile(seriesPath); err == nil {
		return data, nil
	}
	genericPath := filepath.Join(r.BaseDir, "specs", "prompts", "generic", filename)
	return os.ReadFile(genericPath)
}

func (r *Runner) loadSeriesSpecs(series string) *seriesSpecs {
	specs := &seriesSpecs{
		PromptTemplates: make(map[string]*template.Template),
		Examples:        make(map[string]string),
	}

	data, err := r.resolveSpecFile(series, "voice-summary.md")
	if err != nil {
		r.Log.Printf("[%s] Voice summary load failed: %v", series, err)
	} else {
		full := string(data)
		specs.VoiceProfile = full

		const antiHeader = "## What This Voice Does NOT Do"
		idx := strings.Index(full, antiHeader)
		if idx >= 0 {
			section := full[idx:]
			if end := strings.Index(section[len(antiHeader):], "\n---"); end >= 0 {
				section = section[:len(antiHeader)+end]
			}
			specs.VoiceAntiPatterns = strings.TrimSpace(section)
		}
	}

	data, err = r.resolveSpecFile(series, "essay-rules.md")
	if err != nil {
		r.Log.Printf("[%s] Essay rules load failed: %v", series, err)
	} else {
		full := string(data)
		specs.DraftRules = extractSection(full, "## Draft Guidelines")
		specs.RevisionRules = extractSection(full, "## Revision Rules")
	}

	seriesPromptDir := filepath.Join(r.BaseDir, "specs", "prompts", "series", series)
	genericPromptDir := filepath.Join(r.BaseDir, "specs", "prompts", "generic")

	seen := make(map[string]bool)
	for _, dir := range []string{seriesPromptDir, genericPromptDir} {
		entries, err := os.ReadDir(dir)
		if err != nil {
			continue
		}
		for _, e := range entries {
			if e.IsDir() || filepath.Ext(e.Name()) != ".md" {
				continue
			}
			name := strings.TrimSuffix(e.Name(), ".md")
			if name == "voice-summary" || name == "essay-rules" {
				continue
			}
			if seen[name] {
				continue
			}
			content, err := os.ReadFile(filepath.Join(dir, e.Name()))
			if err != nil {
				r.Log.Printf("[%s] WARNING: could not read prompt template '%s': %v", series, name, err)
				continue
			}
			tmpl, err := template.New(name).Parse(string(content))
			if err != nil {
				r.Log.Printf("[%s] WARNING: could not parse prompt template '%s': %v", series, name, err)
				continue
			}
			specs.PromptTemplates[name] = tmpl
			seen[name] = true
		}
	}

	exampleDir := filepath.Join(r.BaseDir, "specs", "examples")
	exFiles, err := os.ReadDir(exampleDir)
	if err != nil {
		r.Log.Printf("[%s] Examples dir load failed: %v", series, err)
	} else {
		for _, f := range exFiles {
			if f.IsDir() || filepath.Ext(f.Name()) != ".md" {
				continue
			}
			category := strings.TrimSuffix(f.Name(), ".md")
			content, err := os.ReadFile(filepath.Join(exampleDir, f.Name()))
			if err != nil {
				r.Log.Printf("[%s] WARNING: could not read examples '%s': %v", series, category, err)
				continue
			}
			for k, v := range splitExamples(category, string(content)) {
				specs.Examples[k] = v
			}
		}
	}

	r.Log.Printf("[%s] Loaded specs (prompts: %d, examples: %d)", series, len(specs.PromptTemplates), len(specs.Examples))
	return specs
}

func (r *Runner) specsFor(series string) *seriesSpecs {
	if s, ok := r.SeriesSpecs[series]; ok {
		return s
	}
	r.Log.Printf("WARNING: no specs loaded for series '%s', loading on demand", series)
	s := r.loadSeriesSpecs(series)
	r.SeriesSpecs[series] = s
	return s
}

func (r *Runner) executePrompt(series, name string, data map[string]any) string {
	specs := r.specsFor(series)
	tmpl, ok := specs.PromptTemplates[name]
	if !ok {
		r.Log.Printf("WARNING: prompt template '%s' not found for series '%s'", name, series)
		return ""
	}
	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		r.Log.Printf("ERROR executing prompt template '%s': %v", name, err)
		return ""
	}
	return buf.String()
}

func (r *Runner) buildExamples(series string, keys ...string) string {
	specs := r.specsFor(series)
	var parts []string
	for _, key := range keys {
		if ex, ok := specs.Examples[key]; ok && ex != "" {
			parts = append(parts, ex)
		}
	}
	if len(parts) == 0 {
		return ""
	}
	return "\n## REFERENCE EXAMPLES (for this essay's attributes)\n\nStudy these examples. They show the level of craft expected for this specific combination of attributes. Adapt the techniques — do not copy.\n\n" + strings.Join(parts, "\n\n---\n\n") + "\n"
}

func extractSection(content, header string) string {
	idx := strings.Index(content, header)
	if idx < 0 {
		return ""
	}
	section := content[idx+len(header):]
	if end := strings.Index(section, "\n---"); end >= 0 {
		section = section[:end]
	}
	return strings.TrimSpace(section)
}

func splitExamples(category, content string) map[string]string {
	result := make(map[string]string)
	sections := strings.Split(content, "\n## ")
	for i, sec := range sections {
		if i == 0 {
			continue
		}
		newline := strings.IndexByte(sec, '\n')
		if newline < 0 {
			continue
		}
		name := strings.TrimSpace(sec[:newline])
		body := strings.TrimSpace(sec[newline+1:])
		if body != "" {
			result[category+"/"+name] = body
		}
	}
	return result
}

func (r *Runner) RunCycle(ctx context.Context) ([]string, error) {
	r.ReloadConfig()
	r.loadSpecs()
	var allActions []string
	for _, ps := range r.Projects {
		if ctx.Err() != nil {
			return allActions, ctx.Err()
		}
		actions, err := r.runProjectCycle(ctx, ps)
		if err != nil {
			return allActions, fmt.Errorf("project %s: %w", ps.Project, err)
		}
		allActions = append(allActions, actions...)
	}
	return allActions, nil
}

func (r *Runner) runProjectCycle(ctx context.Context, ps *PipelineState) ([]string, error) {
	if err := ps.LoadFromDisk(); err != nil {
		return nil, fmt.Errorf("loading state: %w", err)
	}

	r.generateBookArtifacts(ps)

	selected := ps.SelectForCycle(&r.Config.Pipeline)
	if len(selected) == 0 {
		r.Log.Printf("[%s] Cycle: nothing to do", ps.Project)
		return nil, nil
	}

	for _, e := range selected {
		if e.NextAction() == StageIllustrate {
			r.renderAllImages(ps)
			break
		}
	}

	if r.Config.Pipeline.Debug != "" {
		return r.runDebugCycle(ctx, ps, selected[0])
	}

	ps.CycleCount++
	r.Log.Printf("[%s] Cycle %d: processing %d essays", ps.Project, ps.CycleCount, len(selected))

	concurrency := r.Config.Pipeline.Concurrency
	if concurrency <= 0 {
		concurrency = 3
	}

	type result struct {
		action string
		err    error
	}

	results := make([]result, len(selected))
	sem := make(chan struct{}, concurrency)
	var wg sync.WaitGroup

	for i, essay := range selected {
		nextStage := essay.NextAction()
		if nextStage == StageDone {
			continue
		}

		action := fmt.Sprintf("[%s] %s → %s", ps.Project, essay.Slug, nextStage)
		r.Log.Printf("  %s → %s", essay.Slug, nextStage)

		wg.Add(1)
		go func(idx int, e *EssayState, ns Stage, act string) {
			defer wg.Done()
			select {
			case sem <- struct{}{}:
			case <-ctx.Done():
				results[idx] = result{action: fmt.Sprintf("%s (cancelled)", act)}
				return
			}
			defer func() { <-sem }()

			if err := r.processEssay(ctx, ps, e, ns); err != nil {
				r.Log.Printf("  ERROR: %s: %v", e.Slug, err)
				r.markError(ps, e, ns, err)
				results[idx] = result{action: fmt.Sprintf("%s (ERROR: %v)", act, err), err: err}
				return
			}
			results[idx] = result{action: act}
		}(i, essay, nextStage, action)
	}

	wg.Wait()

	r.generateBookArtifacts(ps)

	var actions []string
	for _, res := range results {
		if res.action != "" {
			actions = append(actions, res.action)
		}
	}

	return actions, nil
}

func (r *Runner) runDebugCycle(ctx context.Context, ps *PipelineState, essay *EssayState) ([]string, error) {
	r.Log.Printf("[%s] DEBUG MODE: %s (%s)", ps.Project, essay.Slug, r.Config.Pipeline.Debug)

	var actions []string
	for {
		if ctx.Err() != nil {
			return actions, ctx.Err()
		}

		nextStage := essay.NextAction()
		if nextStage == StageDone {
			r.Log.Printf("  %s: all stages complete", essay.Slug)
			break
		}

		action := fmt.Sprintf("[%s] %s → %s", ps.Project, essay.Slug, nextStage)
		r.Log.Printf("  %s → %s", essay.Slug, nextStage)

		if err := r.processEssay(ctx, ps, essay, nextStage); err != nil {
			r.Log.Printf("  ERROR: %s: %v", essay.Slug, err)
			r.markError(ps, essay, nextStage, err)
			actions = append(actions, fmt.Sprintf("%s (ERROR: %v)", action, err))
			return actions, err
		}
		actions = append(actions, action)
	}

	return actions, nil
}

func (r *Runner) processEssay(ctx context.Context, ps *PipelineState, essay *EssayState, targetStage Stage) error {
	if shouldSkipStage(essay.Type, targetStage) {
		return r.autoSkipStage(ps, essay, targetStage)
	}

	if targetStage == StageExport {
		if r.Config.Pipeline.Debug != "" {
			return r.processDocxSync(ps, essay)
		}
		r.markInProgress(ps, essay, StageExport, "local")
		r.docxCh <- docxJob{ps: ps, essay: essay}
		return nil
	}

	prompt, model, err := r.buildPrompt(ps, essay, targetStage)
	if err != nil {
		return fmt.Errorf("building prompt: %w", err)
	}

	r.markInProgress(ps, essay, targetStage, model)

	var result *APIResult
	if r.Config.Pipeline.DryRun {
		result = DryRunResult(targetStage, essay.Title)
		r.Log.Printf("    [dry-run] %s", targetStage)
	} else {
		timeout := time.Duration(r.Config.Pipeline.APITimeout) * time.Second
		if timeout <= 0 {
			timeout = 300 * time.Second
		}
		timeout *= time.Duration(essay.ErrorRetries + 1)
		result, err = r.Client.Call(ctx, model, prompt, timeout)
		if err != nil {
			return fmt.Errorf("%s/%s: %w", essay.Slug, targetStage, err)
		}
		r.Log.Printf("    tokens: %d in / %d out, cost: $%.4f",
			result.InputTokens, result.OutputTokens, result.Cost)
	}

	if err := ps.WriteContent(targetStage, essay.Slug, result.Content); err != nil {
		return fmt.Errorf("writing content: %w", err)
	}

	if targetStage == StageIllustrate {
		if err := r.writeImageSources(ps, essay.Slug, result.Content); err != nil {
			r.Log.Printf("    WARNING: writing image sources: %v", err)
		} else if !r.Config.Pipeline.DryRun {
			r.renderImages(ps, essay.Slug)
		}
	}

	r.markComplete(ps, essay, targetStage, model, result)
	ps.TotalCost += result.Cost
	ps.SessionCost += result.Cost

	return nil
}

func (r *Runner) processDocxSync(ps *PipelineState, essay *EssayState) error {
	if err := r.exportEssay(ps, essay); err != nil {
		return err
	}
	dstFile := filepath.Join(ps.BaseDir, "export", exportFilename(essay))
	if err := r.upgradeDocx(dstFile); err != nil {
		r.Log.Printf("    WARNING: Word upgrade %s: %v", essay.Slug, err)
	}
	if err := r.swapImages(ps, essay, dstFile); err != nil {
		r.Log.Printf("    WARNING: imageswap %s: %v", essay.Slug, err)
	}
	return nil
}

func (r *Runner) exportEssay(ps *PipelineState, essay *EssayState) error {
	r.markInProgress(ps, essay, StageExport, "local")

	mdFile := filepath.Join(ps.BaseDir, "draft2", essay.Slug+".md")
	if _, err := os.Stat(mdFile); os.IsNotExist(err) {
		mdFile = filepath.Join(ps.BaseDir, "illustrate", essay.Slug+".md")
		if _, err := os.Stat(mdFile); os.IsNotExist(err) {
			return fmt.Errorf("neither draft2 nor illustrate file found for %s", essay.Slug)
		}
	}

	raw, err := os.ReadFile(mdFile)
	if err != nil {
		return fmt.Errorf("reading markdown: %w", err)
	}
	cleaned := sanitizeForDocx(string(raw))

	tmpFile, err := os.CreateTemp("", "export-*.md")
	if err != nil {
		return fmt.Errorf("creating temp file: %w", err)
	}
	defer os.Remove(tmpFile.Name())
	if _, err := tmpFile.WriteString(cleaned); err != nil {
		tmpFile.Close()
		return fmt.Errorf("writing temp file: %w", err)
	}
	tmpFile.Close()

	exportDir := filepath.Join(ps.BaseDir, "export")
	if err := os.MkdirAll(exportDir, 0755); err != nil {
		return fmt.Errorf("creating export dir: %w", err)
	}

	home, _ := os.UserHomeDir()
	templatePath := filepath.Join(home, ".works", "templates", "book-template.dotm")
	exportName := exportFilename(essay)
	outFile := filepath.Join(exportDir, exportName)

	cmd := exec.Command("md2docx", templatePath, tmpFile.Name(), outFile)
	if output, err := cmd.CombinedOutput(); err != nil {
		return fmt.Errorf("md2docx: %w %s", err, string(output))
	}

	if strings.Contains(cleaned, "[[IMG:") {
		dataDir := filepath.Join(r.BaseDir, "data")
		renderCmd := exec.Command("imagerender", "--data", dataDir, "--slug", essay.Slug, ps.BaseDir)
		if output, err := renderCmd.CombinedOutput(); err != nil {
			return fmt.Errorf("imagerender %s: %w %s", essay.Slug, err, string(output))
		} else {
			r.Log.Printf("    imagerender: %s", strings.TrimSpace(string(output)))
		}
	}

	r.Log.Printf("    exported → %s", exportName)

	result := &APIResult{}
	r.markComplete(ps, essay, StageExport, "local", result)
	return nil
}

func (r *Runner) swapImages(ps *PipelineState, essay *EssayState, docxPath string) error {
	mdFile := filepath.Join(ps.BaseDir, "draft2", essay.Slug+".md")
	if _, err := os.Stat(mdFile); os.IsNotExist(err) {
		mdFile = filepath.Join(ps.BaseDir, "illustrate", essay.Slug+".md")
	}
	raw, err := os.ReadFile(mdFile)
	if err != nil {
		return nil
	}
	if !strings.Contains(string(raw), "[[IMG:") {
		return nil
	}
	r.Log.Printf("  [docx] %s: imageswap starting", essay.Slug)
	swapCmd := exec.Command("imageswap", "--slug", essay.Slug, "--images", filepath.Join(ps.BaseDir, "images"), docxPath)
	if output, err := swapCmd.CombinedOutput(); err != nil {
		return fmt.Errorf("imageswap %s: %w %s", essay.Slug, err, string(output))
	}
	return nil
}

func (r *Runner) upgradeDocx(docxPath string) error {
	absPath, err := filepath.Abs(docxPath)
	if err != nil {
		return fmt.Errorf("abs path: %w", err)
	}

	home, _ := os.UserHomeDir()
	templatePath := filepath.Join(home, ".works", "templates", "book-template.dotm")

	escaped := func(s string) string {
		var out strings.Builder
		for _, c := range s {
			switch c {
			case '"', '\u201C', '\u201D':
				out.WriteString(`\"`)
			case '\\':
				out.WriteString(`\\`)
			default:
				out.WriteRune(c)
			}
		}
		return out.String()
	}

	baseName := filepath.Base(docxPath)
	title := baseName[:len(baseName)-len(filepath.Ext(baseName))]

	script := `
tell application "Microsoft Word"
	launch
end tell
tell application "System Events"
	set visible of process "Microsoft Word" to false
end tell
tell application "Microsoft Word"
	open (POSIX file "` + escaped(templatePath) + `" as text)
	set templateDoc to active document
	set templatePageSetup to page setup of section 1 of templateDoc
	set templatePageWidth to page width of templatePageSetup
	set templatePageHeight to page height of templatePageSetup
	set templateTopMargin to top margin of templatePageSetup
	set templateBottomMargin to bottom margin of templatePageSetup
	set templateLeftMargin to left margin of templatePageSetup
	set templateRightMargin to right margin of templatePageSetup
	set templateMirrorMargins to mirror margins of templatePageSetup
	set templateGutter to gutter of templatePageSetup
	close templateDoc saving no
	open (POSIX file "` + escaped(absPath) + `" as text)
	set theDoc to active document
	set docPageSetup to page setup of section 1 of theDoc
	set docPageWidth to page width of docPageSetup
	set docPageHeight to page height of docPageSetup
	set docIsLandscape to (docPageWidth > docPageHeight)
	if docIsLandscape then
		set tempWidth to templatePageWidth
		set templatePageWidth to templatePageHeight
		set templatePageHeight to tempWidth
		set tempTopMargin to templateTopMargin
		set tempBottomMargin to templateBottomMargin
		set templateTopMargin to templateLeftMargin
		set templateBottomMargin to templateRightMargin
		set templateLeftMargin to tempTopMargin
		set templateRightMargin to tempBottomMargin
	end if
	set widthRatio to templatePageWidth / docPageWidth
	set heightRatio to templatePageHeight / docPageHeight
	if widthRatio < heightRatio then
		set scaleFactor to widthRatio
	else
		set scaleFactor to heightRatio
	end if
	set inlineShapeCount to count of inline shapes of theDoc
	repeat with i from 1 to inlineShapeCount
		try
			set inlineShape to inline shape i of theDoc
			set width of inlineShape to (width of inlineShape) * scaleFactor
			set height of inlineShape to (height of inlineShape) * scaleFactor
		end try
	end repeat
	set floatShapeCount to count of shapes of theDoc
	repeat with i from 1 to floatShapeCount
		try
			set floatShape to shape i of theDoc
			set width of floatShape to (width of floatShape) * scaleFactor
			set height of floatShape to (height of floatShape) * scaleFactor
		end try
	end repeat
	tell theDoc to copy styles from template template (POSIX file "` + escaped(templatePath) + `" as text)
	set page width of docPageSetup to templatePageWidth
	set page height of docPageSetup to templatePageHeight
	set mirror margins of docPageSetup to templateMirrorMargins
	set gutter of docPageSetup to templateGutter
	set top margin of docPageSetup to templateTopMargin
	set bottom margin of docPageSetup to templateBottomMargin
	set left margin of docPageSetup to templateLeftMargin
	set right margin of docPageSetup to templateRightMargin
	set title of properties of theDoc to "` + escaped(title) + `"
	save theDoc
	close theDoc
	quit
end tell
`

	tmpFile, err := os.CreateTemp("", "upgrade-*.scpt")
	if err != nil {
		return fmt.Errorf("creating temp script: %w", err)
	}
	scriptPath := tmpFile.Name()
	if _, err := tmpFile.WriteString(script); err != nil {
		tmpFile.Close()
		os.Remove(scriptPath)
		return fmt.Errorf("writing script: %w", err)
	}
	tmpFile.Close()
	defer os.Remove(scriptPath)

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	cmd := exec.CommandContext(ctx, "osascript", scriptPath)
	output, err := cmd.CombinedOutput()
	if ctx.Err() == context.DeadlineExceeded {
		return fmt.Errorf("Word upgrade timed out")
	}
	if err != nil {
		return fmt.Errorf("osascript: %w: %s", err, string(output))
	}
	return nil
}

func (r *Runner) buildPrompt(ps *PipelineState, essay *EssayState, targetStage Stage) (string, string, error) {
	ideaMeta := essay.Meta[StageIdeas]
	if ideaMeta == nil {
		return "", "", fmt.Errorf("no idea metadata for %s", essay.Slug)
	}

	series := ps.Project

	readContent := func(stage Stage) string {
		path := filepath.Join(ps.BaseDir, stage.Dir(), essay.Slug+".md")
		data, err := os.ReadFile(path)
		if err != nil {
			return ""
		}
		return string(data)
	}

	displayTitle := ideaMeta.Title
	if essay.Subtitle != "" {
		displayTitle = ideaMeta.Title + ": " + essay.Subtitle
	}

	var prompt, model string
	targetWords := r.sampleWordTarget(essay)
	switch targetStage {
	case StageResearch:
		model = r.Config.Models.Research
		ideaContent := readContent(StageIdeas)
		hook := displayTitle
		hiddenMath := ""
		if len(ideaContent) > 0 {
			hook = ideaContent
		}
		prompt = r.researchPrompt(series, displayTitle, hook, hiddenMath, essay.Setting)

	case StageOutline:
		model = r.Config.Models.Outline
		if essay.Type == "introduction" {
			ideaContent := readContent(StageIdeas)
			prompt = r.introOutlinePrompt(series, displayTitle, ideaContent)
		} else {
			research := readContent(StageResearch)
			arc, _ := ArcByName(essay.Arc)
			structure, _ := StructureByName(essay.Structure)
			entry, _ := EntryByName(essay.Entry)
			mathVis, _ := MathVisByName(essay.MathVisibility)
			prompt = r.outlinePrompt(series, displayTitle, research, targetWords, arc, structure, entry, mathVis)
		}

	case StageDraft:
		model = r.Config.Models.Draft
		if essay.Type == "introduction" {
			outline := readContent(StageOutline)
			ideaContent := readContent(StageIdeas)
			prompt = r.introDraftPrompt(series, displayTitle, outline, ideaContent)
		} else {
			outline := readContent(StageOutline)
			research := readContent(StageResearch)
			arc, _ := ArcByName(essay.Arc)
			structure, _ := StructureByName(essay.Structure)
			entry, _ := EntryByName(essay.Entry)
			register, _ := RegisterByName(essay.Register)
			mathVis, _ := MathVisByName(essay.MathVisibility)
			prompt = r.draftPrompt(series, displayTitle, outline, research, targetWords, arc, structure, entry, register, essay.Setting, mathVis)
		}

	case StageFactcheck:
		model = r.Config.Models.Factcheck
		draft := readContent(StageDraft)
		research := readContent(StageResearch)
		prompt = r.factcheckPrompt(series, displayTitle, draft, research)

	case StageDraft2:
		model = r.Config.Models.Draft2
		switch essay.Type {
		case "section":
			ideaContent := readContent(StageIdeas)
			prompt = r.sectionDraft2Prompt(series, displayTitle, ideaContent, ideaMeta.PartTitle)
		case "introduction":
			draft := readContent(StageDraft)
			prompt = r.introDraft2Prompt(series, displayTitle, draft)
		default:
			draft := readContent(StageDraft)
			factcheck := readContent(StageFactcheck)
			illustrate := readContent(StageIllustrate)
			arc, _ := ArcByName(essay.Arc)
			register, _ := RegisterByName(essay.Register)
			prompt = r.draft2Prompt(series, displayTitle, draft, factcheck, illustrate, targetWords, arc, register)
		}

	case StageIllustrate:
		model = r.Config.Models.Illustrate
		draft := readContent(StageDraft)
		factcheck := readContent(StageFactcheck)
		mathVis, _ := MathVisByName(essay.MathVisibility)
		prompt = r.illustratePrompt(series, displayTitle, draft, factcheck, essay.Slug, essay.Setting, mathVis)

	default:
		return "", "", fmt.Errorf("unknown target stage: %s", targetStage)
	}

	return prompt, model, nil
}

func (r *Runner) sampleWordTarget(essay *EssayState) int {
	mean := r.Config.Pipeline.ReadMean
	spread := r.Config.Pipeline.ReadSpread
	// Per-essay overrides take precedence
	if essay.ReadMean > 0 {
		mean = essay.ReadMean
	}
	if essay.ReadSpread > 0 {
		spread = essay.ReadSpread
	}
	if mean <= 0 {
		mean = 5.0
	}
	if spread <= 0 {
		spread = 1.0
	}
	minutes := randv2.NormFloat64()*spread + mean
	minutes = math.Max(mean-2*spread, math.Min(mean+2*spread, minutes))
	if minutes < 1 {
		minutes = 1
	}
	return int(minutes * 265)
}

func (r *Runner) markInProgress(ps *PipelineState, essay *EssayState, stage Stage, model string) {
	meta := &EssayMeta{
		Slug:      essay.Slug,
		Title:     essay.Title,
		Subtitle:  essay.Subtitle,
		Type:      essay.Type,
		Series:    essay.Series,
		Book:      essay.Book,
		Part:      essay.Part,
		PartTitle: essay.PartTitle,
		Order:     essay.Order,
		Status:    "in-progress",
		Model:     model,
		Arc:       essay.Arc,
		Ending:    essay.Ending,
		Created:   essay.Meta[StageIdeas].Created,
		Started:   nowString(),
	}
	ps.WriteMeta(stage, meta)

	ps.mu.Lock()
	essay.CurrentStage = stage
	essay.Status = "in-progress"
	essay.Meta[stage] = meta
	ps.mu.Unlock()
}

func (r *Runner) markComplete(ps *PipelineState, essay *EssayState, stage Stage, model string, result *APIResult) {
	meta := &EssayMeta{
		Slug:      essay.Slug,
		Title:     essay.Title,
		Subtitle:  essay.Subtitle,
		Type:      essay.Type,
		Series:    essay.Series,
		Book:      essay.Book,
		Part:      essay.Part,
		PartTitle: essay.PartTitle,
		Order:     essay.Order,
		Status:    "final",
		Model:     model,
		Arc:       essay.Arc,
		Ending:    essay.Ending,
		Created:   essay.Meta[StageIdeas].Created,
		Started:   nowString(),
		Completed: nowString(),
		Tokens:    result.InputTokens,
		TokensOut: result.OutputTokens,
		Cost:      result.Cost,
	}
	ps.WriteMeta(stage, meta)

	ps.mu.Lock()
	essay.CurrentStage = stage
	essay.Status = "final"
	essay.Meta[stage] = meta
	if stage == StageExport {
		ps.SessionDone++
	}
	ps.mu.Unlock()
}

func (r *Runner) markError(ps *PipelineState, essay *EssayState, stage Stage, err error) {
	meta := &EssayMeta{
		Slug:      essay.Slug,
		Title:     essay.Title,
		Subtitle:  essay.Subtitle,
		Type:      essay.Type,
		Series:    essay.Series,
		Book:      essay.Book,
		Part:      essay.Part,
		PartTitle: essay.PartTitle,
		Order:     essay.Order,
		Status:    "error",
		Model:     "",
		Arc:       essay.Arc,
		Ending:    essay.Ending,
		Created:   essay.Meta[StageIdeas].Created,
		Started:   nowString(),
		Error:     err.Error(),
		Retries:   essay.ErrorRetries + 1,
	}
	ps.WriteMeta(stage, meta)

	ps.mu.Lock()
	essay.Status = "error"
	essay.ErrorRetries++
	essay.Meta[stage] = meta
	ps.mu.Unlock()
	if essay.ErrorRetries >= 3 {
		r.Log.Printf("  %s: giving up after %d retries (revert to retry)", essay.Slug, essay.ErrorRetries)
	}
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

type imageSource struct {
	filename string
	method   string
	source   string
}

func parseImageSources(content string) []imageSource {
	var sources []imageSource
	sepIdx := strings.Index(content, "---IMAGE-SOURCES---")
	if sepIdx < 0 {
		return nil
	}
	remainder := content[sepIdx:]

	re := regexp.MustCompile(`---IMAGE:([^|]+)\|method:(\w+)---\n([\s\S]*?)---END-IMAGE---`)
	matches := re.FindAllStringSubmatch(remainder, -1)
	for _, m := range matches {
		sources = append(sources, imageSource{
			filename: strings.TrimSpace(m[1]),
			method:   strings.TrimSpace(m[2]),
			source:   strings.TrimSpace(m[3]),
		})
	}
	return sources
}

func (r *Runner) writeImageSources(ps *PipelineState, slug, content string) error {
	sources := parseImageSources(content)
	if len(sources) == 0 {
		r.Log.Printf("    no image sources found in illustrate output")
		return nil
	}

	imgDir := filepath.Join(ps.BaseDir, "images", slug)
	if err := os.MkdirAll(imgDir, 0755); err != nil {
		return fmt.Errorf("creating images dir: %w", err)
	}

	for _, src := range sources {
		baseName := strings.TrimSuffix(src.filename, filepath.Ext(src.filename))

		var ext string
		switch src.method {
		case "mermaid":
			ext = ".mermaid"
		case "r":
			ext = ".R"
		case "ai":
			ext = ".ai-prompt.txt"
		default:
			r.Log.Printf("    unknown image method: %s for %s", src.method, src.filename)
			continue
		}

		srcPath := filepath.Join(imgDir, baseName+ext)
		if err := os.WriteFile(srcPath, []byte(src.source), 0644); err != nil {
			return fmt.Errorf("writing %s: %w", srcPath, err)
		}

		metaContent := fmt.Sprintf("filename: %s\nmethod: %s\ndescription: %s\n",
			src.filename, src.method, baseName)
		metaPath := filepath.Join(imgDir, baseName+".meta.yaml")
		if err := os.WriteFile(metaPath, []byte(metaContent), 0644); err != nil {
			return fmt.Errorf("writing %s: %w", metaPath, err)
		}

		r.Log.Printf("    image source: %s (%s)", baseName+ext, src.method)
	}

	essayText := content
	if sepIdx := strings.Index(content, "---IMAGE-SOURCES---"); sepIdx > 0 {
		essayText = strings.TrimSpace(content[:sepIdx])
	}
	if err := ps.WriteContent(StageIllustrate, slug, essayText); err != nil {
		return fmt.Errorf("rewriting illustrate content: %w", err)
	}

	return nil
}

func (r *Runner) renderAllImages(ps *PipelineState) {
	r.Log.Printf("[%s] imagerender: starting", ps.Project)
	dataDir := filepath.Join(r.BaseDir, "data")
	cmd := exec.Command("imagerender", "--data", dataDir, ps.BaseDir)
	output, err := cmd.CombinedOutput()
	if err != nil {
		r.Log.Printf("[%s] WARNING: imagerender: %v\n%s", ps.Project, err, string(output))
		return
	}
	lines := strings.TrimSpace(string(output))
	for _, line := range strings.Split(lines, "\n") {
		line = strings.TrimSpace(line)
		if line != "" {
			r.Log.Printf("[%s] imagerender: %s", ps.Project, line)
		}
	}
}

func (r *Runner) renderImages(ps *PipelineState, slug string) {
	dataDir := filepath.Join(r.BaseDir, "data")
	cmd := exec.Command("imagerender", "--data", dataDir, "--slug", slug, ps.BaseDir)
	output, err := cmd.CombinedOutput()
	if err != nil {
		r.Log.Printf("    WARNING: imagerender %s: %v\n%s", slug, err, string(output))
		return
	}
	lines := strings.TrimSpace(string(output))
	if lines != "" {
		for _, line := range strings.Split(lines, "\n") {
			r.Log.Printf("    imagerender: %s", line)
		}
	}
}

func (r *Runner) allEssaysAtOrPast(ps *PipelineState, stage Stage) bool {
	ps.mu.Lock()
	defer ps.mu.Unlock()
	if len(ps.Essays) == 0 {
		return false
	}
	for _, essay := range ps.Essays {
		if essay.CurrentStage < stage {
			return false
		}
		if essay.CurrentStage == stage && essay.Status != "final" {
			return false
		}
	}
	return true
}

func (r *Runner) generateBookArtifacts(ps *PipelineState) {
	if !r.allEssaysAtOrPast(ps, StageDraft2) {
		return
	}

	bookDir := filepath.Join(ps.BaseDir, "book")

	needBlurb := !bookDirHasFile(bookDir, "back-cover-blurb.md")
	needCover := !bookDirHasFile(bookDir, "front-cover-prompt.md") || !bookDirHasFile(bookDir, "front-cover.png")

	if !needBlurb && !needCover {
		return
	}

	if r.Config.Pipeline.DryRun {
		r.Log.Printf("[%s] [book] dry-run: would generate book artifacts", ps.Project)
		return
	}

	go func() {
		if needBlurb {
			r.Log.Printf("[%s] [book] generating back-cover blurb...", ps.Project)
			cmd := exec.Command("bookblurb", ps.BaseDir)
			output, err := cmd.CombinedOutput()
			if err != nil {
				r.Log.Printf("[%s] [book] WARNING: bookblurb: %v\n%s", ps.Project, err, string(output))
				return
			}
			r.Log.Printf("[%s] [book] blurb generated", ps.Project)
		}

		if needCover {
			r.Log.Printf("[%s] [book] generating front-cover prompt and image...", ps.Project)
			cmd := exec.Command("bookcover", ps.BaseDir)
			output, err := cmd.CombinedOutput()
			if err != nil {
				r.Log.Printf("[%s] [book] WARNING: bookcover: %v\n%s", ps.Project, err, string(output))
				return
			}
			r.Log.Printf("[%s] [book] cover generated", ps.Project)
		}
	}()
}

func fileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

func invalidateBookArtifacts(baseDir string) {
	bookDir := filepath.Join(baseDir, "book")
	for _, name := range []string{"back-cover-blurb.md", "front-cover-prompt.md", "front-cover.png"} {
		p := filepath.Join(bookDir, name)
		if fileExists(p) {
			os.Remove(p)
		}
		entries, err := os.ReadDir(bookDir)
		if err != nil {
			continue
		}
		for _, e := range entries {
			if e.IsDir() {
				sp := filepath.Join(bookDir, e.Name(), name)
				if fileExists(sp) {
					os.Remove(sp)
				}
			}
		}
	}
}

func bookDirHasFile(bookDir, filename string) bool {
	if fileExists(filepath.Join(bookDir, filename)) {
		return true
	}
	entries, err := os.ReadDir(bookDir)
	if err != nil {
		return false
	}
	for _, e := range entries {
		if e.IsDir() && fileExists(filepath.Join(bookDir, e.Name(), filename)) {
			return true
		}
	}
	return false
}
