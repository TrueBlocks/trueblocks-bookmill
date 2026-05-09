package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/TrueBlocks/trueblocks-art/packages/ai"
	appkit "github.com/TrueBlocks/trueblocks-art/packages/appkit/v2"
	"github.com/TrueBlocks/trueblocks-art/packages/bookgen"
	"github.com/TrueBlocks/trueblocks-art/packages/cli"
	"github.com/TrueBlocks/trueblocks-math/internal/bookutil"
	"github.com/TrueBlocks/trueblocks-math/internal/pipeline"
	"github.com/TrueBlocks/trueblocks-math/internal/types"
)

var version = "dev"

func main() {
	app := cli.App{
		Name:        "bookgen",
		Description: "Generate book artifacts (back-cover blurb, front-cover image) from a project's draft2 essays.",
		Version:     version,
		Subcommands: []*cli.App{
			{
				Name:        "blurb",
				Description: "Generate back-cover blurb",
				ArgsUsage:   "<project-dir>",
				MinArgs:     1,
				Flags: []cli.FlagDef{
					{Name: "config", Help: "path to config.yaml", Default: pipeline.DefaultConfigPath()},
					{Name: "model", Help: "Anthropic model", Default: "claude-sonnet-4-20250514"},
					{Name: "dry-run", Help: "print prompt without calling API", Default: false},
					{Name: "force", Help: "regenerate even if blurb exists", Default: false},
				},
				Run: runBlurb,
			},
			{
				Name:        "cover",
				Description: "Generate front-cover prompt + image",
				ArgsUsage:   "<project-dir>",
				MinArgs:     1,
				Flags: []cli.FlagDef{
					{Name: "config", Help: "path to config.yaml", Default: pipeline.DefaultConfigPath()},
					{Name: "model", Help: "Anthropic model", Default: "claude-sonnet-4-20250514"},
					{Name: "dry-run", Help: "print prompt without calling API", Default: false},
					{Name: "prompt-only", Help: "generate cover prompt, skip image", Default: false},
					{Name: "force", Help: "regenerate even if artifacts exist", Default: false},
					{Name: "title", Help: "book title (extracted from Plan if not specified)"},
					{Name: "author", Help: "author name", Default: "Claude Jay Rush"},
				},
				Run: runCover,
			},
		},
	}
	cli.Exit(app.Main())
}

func runBlurb(c *cli.Context) error {
	configPath := c.String("config")
	model := c.String("model")
	dryRun := c.Bool("dry-run")
	force := c.Bool("force")
	projectDir := c.Args[0]

	outPath := filepath.Join(projectDir, "book", "back-cover-blurb.md")
	if !force {
		if _, err := os.Stat(outPath); err == nil {
			c.Logger.Info("blurb already exists, use --force to regenerate", "path", outPath)
			return nil
		}
	}

	provider, err := loadProvider(configPath, dryRun)
	if err != nil {
		return err
	}

	planText := bookutil.FindPlan(projectDir)
	rawEssays := bookutil.ReadDraft2(projectDir, 0)
	if len(rawEssays) == 0 {
		return fmt.Errorf("no draft2 essays found in %s", projectDir)
	}

	essays := convertEssays(rawEssays)

	c.Logger.Info("generating back-cover blurb")
	result, err := bookgen.GenerateBlurb(c.Context, provider, bookgen.BlurbInput{
		Plan:   planText,
		Essays: essays,
		Model:  model,
		DryRun: dryRun,
	})
	if err != nil {
		return err
	}

	if dryRun {
		fmt.Println(result.Content)
		return nil
	}

	c.Logger.Info("api result",
		"input_tokens", result.InputTokens,
		"output_tokens", result.OutputTokens,
		"cost_usd", result.Cost,
	)

	if err := os.MkdirAll(filepath.Dir(outPath), appkit.DirPermissions); err != nil {
		return fmt.Errorf("creating output directory: %w", err)
	}
	if err := os.WriteFile(outPath, []byte(result.Content+"\n"), appkit.FilePermissions); err != nil {
		return fmt.Errorf("writing output: %w", err)
	}
	c.Logger.Info("wrote blurb", "path", outPath)
	return nil
}

func runCover(c *cli.Context) error {
	configPath := c.String("config")
	model := c.String("model")
	dryRun := c.Bool("dry-run")
	promptOnly := c.Bool("prompt-only")
	force := c.Bool("force")
	title := c.String("title")
	author := c.String("author")
	projectDir := c.Args[0]

	bookDir := filepath.Join(projectDir, "book")
	promptPath := filepath.Join(bookDir, "front-cover-prompt.md")
	imagePath := filepath.Join(bookDir, "front-cover.png")

	needPrompt := force || !fileExists(promptPath)
	needImage := force || !fileExists(imagePath)

	if !needPrompt && !needImage {
		c.Logger.Info("cover artifacts already exist, use --force to regenerate", "dir", bookDir)
		return nil
	}

	cfg, err := pipeline.LoadConfig(configPath)
	if err != nil {
		return fmt.Errorf("loading config: %w", err)
	}

	planText := bookutil.FindPlan(projectDir)
	rawEssays := bookutil.ReadDraft2(projectDir, 500)

	bookTitle := title
	if bookTitle == "" {
		bookTitle = extractTitleFromPlan(planText)
	}
	if bookTitle == "" {
		bookTitle = filepath.Base(projectDir)
	}

	blurbText := bookgen.ExtractBookBlurb(readBlurb(projectDir))
	essays := convertEssays(rawEssays)

	provider := &ai.Anthropic{APIKey: cfg.API.AnthropicKey}

	var imgProvider ai.ImageProvider
	if !promptOnly && !dryRun {
		imgProvider = &ai.DallE{APIKey: cfg.API.OpenAIKey}
	}

	c.Logger.Info("generating front cover")
	result, err := bookgen.GenerateCover(c.Context, provider, imgProvider, bookgen.CoverInput{
		Title:  bookTitle,
		Author: author,
		Plan:   planText,
		Blurb:  blurbText,
		Essays: essays,
		Model:  model,
		DryRun: dryRun,
	})
	if err != nil && result == nil {
		return err
	}

	if dryRun {
		fmt.Println(result.DesignDoc)
		return nil
	}

	if needPrompt && result.DesignDoc != "" {
		if mkErr := os.MkdirAll(bookDir, appkit.DirPermissions); mkErr != nil {
			return fmt.Errorf("creating directory: %w", mkErr)
		}
		if wErr := os.WriteFile(promptPath, []byte(result.DesignDoc+"\n"), appkit.FilePermissions); wErr != nil {
			return fmt.Errorf("writing prompt: %w", wErr)
		}
		c.Logger.Info("wrote cover prompt", "path", promptPath)
	}

	if result.ImageData != nil {
		if wErr := os.WriteFile(imagePath, result.ImageData, appkit.FilePermissions); wErr != nil {
			return fmt.Errorf("writing image: %w", wErr)
		}
		c.Logger.Info("wrote cover image", "path", imagePath)
	}

	if err != nil {
		c.Logger.Warn("partial cover result", "error", err)
	}
	return nil
}

func loadProvider(configPath string, dryRun bool) (ai.Provider, error) {
	if dryRun {
		return &ai.Anthropic{}, nil
	}
	cfg, err := pipeline.LoadConfig(configPath)
	if err != nil {
		return nil, fmt.Errorf("loading config: %w", err)
	}
	if cfg.API.AnthropicKey == "" {
		return nil, fmt.Errorf("no anthropic_key in config")
	}
	return &ai.Anthropic{APIKey: cfg.API.AnthropicKey}, nil
}

func convertEssays(raw []types.EssayContent) []bookgen.Essay {
	essays := make([]bookgen.Essay, len(raw))
	for i, e := range raw {
		essays[i] = bookgen.Essay{Title: e.Title, Type: e.Typ, Content: e.Content}
	}
	return essays
}

func readBlurb(projectDir string) string {
	data, err := os.ReadFile(filepath.Join(projectDir, "book", "back-cover-blurb.md"))
	if err != nil {
		return ""
	}
	return string(data)
}

func extractTitleFromPlan(plan string) string {
	for _, line := range strings.Split(plan, "\n") {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "# ") {
			t := strings.TrimPrefix(line, "# ")
			t = strings.TrimPrefix(t, "Plan for ")
			return strings.TrimSpace(t)
		}
	}
	return ""
}

func fileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}
