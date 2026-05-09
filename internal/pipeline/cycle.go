package pipeline

import (
	"context"
	"fmt"
	"sync"
)

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
		if e.NextActionForGenre(ps.Genre) == StageIllustrate {
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
		nextStage := essay.NextActionForGenre(ps.Genre)
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

		nextStage := essay.NextActionForGenre(ps.Genre)
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
