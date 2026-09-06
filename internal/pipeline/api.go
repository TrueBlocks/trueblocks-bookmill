package pipeline

import (
	"fmt"

	"github.com/TrueBlocks/trueblocks-art/packages/ai"
)

func DryRunResult(stage Stage, title string) *ai.Result {
	content := fmt.Sprintf("# DRY RUN: %s — %s\n\n"+
		"This is placeholder content generated in dry-run mode.\n\n"+
		"In live mode, this would contain the AI-generated %s output for \"%s\".\n\n"+
		"The prompt template and model configuration are ready.\n"+
		"Set `dry_run: false` in config.yaml and provide an API key to generate real content.\n",
		stage.String(), title, stage.String(), title)

	return &ai.Result{
		Content: content,
	}
}
