package pipeline

import (
	"context"
	"fmt"
	"math"
	randv2 "math/rand/v2"
	"os"
	"path/filepath"
	"time"

	"github.com/TrueBlocks/trueblocks-art/packages/ai"
)

func exportFilename(essay *EssayState) string {
	return exportFilenameFromParts(essay)
}

func shouldSkipStage(itemType string, stage Stage) bool {
	switch itemType {
	case typeSection:
		return stage != StageDraft2 && stage != StageExport
	case typeIntroduction:
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

	result := &ai.Result{}
	r.markComplete(ps, essay, targetStage, "skip", result)
	r.Log.Printf("    [skip] %s: %s (auto-skipped for %s)", essay.Slug, targetStage, essay.Type)
	return nil
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

	var result *ai.Result
	if r.Config.Pipeline.DryRun {
		result = DryRunResult(targetStage, essay.Title)
		r.Log.Printf("    [dry-run] %s", targetStage)
	} else {
		timeout := time.Duration(r.Config.Pipeline.APITimeout) * time.Second
		if timeout <= 0 {
			timeout = 300 * time.Second
		}
		timeout *= time.Duration(essay.ErrorRetries + 1)
		result, err = r.Client.Call(ctx, model, prompt, ai.CallOptions{
			MaxTokens: r.Config.API.MaxTokens,
			Timeout:   timeout,
			Effort:    r.Effort,
		})
		if err != nil {
			return fmt.Errorf("%s/%s: %w", essay.Slug, targetStage, err)
		}
		r.Log.Printf("    tokens: %d in / %d out, cost: $%.4f",
			result.InputTokens, result.OutputTokens, result.Cost)
		if lerr := ai.RecordCall("bookmill", result, 0); lerr != nil {
			r.Log.Printf("    WARNING: could not record cost: %v", lerr)
		}
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

	var prompt string
	model := r.Model
	targetWords := r.sampleWordTarget(essay)
	switch targetStage {
	case StageResearch:
		ideaContent := readContent(StageIdeas)
		hook := displayTitle
		hiddenMath := ""
		if len(ideaContent) > 0 {
			hook = ideaContent
		}
		prompt = r.researchPrompt(series, displayTitle, hook, hiddenMath, essay.Setting)

	case StageOutline:
		if essay.Type == typeIntroduction {
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
		if essay.Type == typeIntroduction {
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
		draft := readContent(StageDraft)
		research := readContent(StageResearch)
		prompt = r.factcheckPrompt(series, displayTitle, draft, research)

	case StageDraft2:
		switch essay.Type {
		case typeSection:
			ideaContent := readContent(StageIdeas)
			prompt = r.sectionDraft2Prompt(series, displayTitle, ideaContent, ideaMeta.PartTitle)
		case typeIntroduction:
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
		draft := readContent(StageDraft)
		factcheck := readContent(StageFactcheck)
		mathVis, _ := MathVisByName(essay.MathVisibility)
		prompt = r.illustratePrompt(series, displayTitle, draft, factcheck, essay.Slug, essay.Setting, mathVis)

	case StageContinuity:
		draft := readContent(StageDraft)
		outline := readContent(StageOutline)
		prompt = r.continuityPrompt(series, displayTitle, draft, outline)

	case StageRevision:
		draft := readContent(StageDraft)
		continuityNotes := readContent(StageContinuity)
		prompt = r.revisionPrompt(series, displayTitle, draft, continuityNotes, targetWords)

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
		Status:    statusInProgress,
		Model:     model,
		Arc:       essay.Arc,
		Ending:    essay.Ending,
		Created:   essay.Meta[StageIdeas].Created,
		Started:   nowString(),
	}
	_ = ps.WriteMeta(stage, meta)

	ps.mu.Lock()
	defer ps.mu.Unlock()
	essay.CurrentStage = stage
	essay.Status = statusInProgress
	essay.Meta[stage] = meta
}

func (r *Runner) markComplete(ps *PipelineState, essay *EssayState, stage Stage, model string, result *ai.Result) {
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
		Status:    statusFinal,
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
	_ = ps.WriteMeta(stage, meta)

	ps.mu.Lock()
	defer ps.mu.Unlock()
	essay.CurrentStage = stage
	essay.Status = statusFinal
	essay.Meta[stage] = meta
	if stage == StageExport {
		ps.SessionDone++
	}
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
		Status:    statusError,
		Model:     "",
		Arc:       essay.Arc,
		Ending:    essay.Ending,
		Created:   essay.Meta[StageIdeas].Created,
		Started:   nowString(),
		Error:     err.Error(),
		Retries:   essay.ErrorRetries + 1,
	}
	_ = ps.WriteMeta(stage, meta)

	ps.mu.Lock()
	defer ps.mu.Unlock()
	essay.Status = statusError
	essay.ErrorRetries++
	essay.Meta[stage] = meta
	if essay.ErrorRetries >= 3 {
		r.Log.Printf("  %s: giving up after %d retries (revert to retry)", essay.Slug, essay.ErrorRetries)
	}
}
