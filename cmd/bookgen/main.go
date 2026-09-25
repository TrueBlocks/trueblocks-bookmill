package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/TrueBlocks/trueblocks-art/packages/ai"
	"github.com/TrueBlocks/trueblocks-art/packages/aiflags"
	appkit "github.com/TrueBlocks/trueblocks-art/packages/appkit/v2"
	"github.com/TrueBlocks/trueblocks-art/packages/bookgen"
	"github.com/TrueBlocks/trueblocks-art/packages/cli"
	"github.com/TrueBlocks/trueblocks-art/packages/creds"
	"github.com/TrueBlocks/trueblocks-bookmill/internal/bookutil"
	"github.com/TrueBlocks/trueblocks-bookmill/internal/pipeline"
	"github.com/TrueBlocks/trueblocks-bookmill/internal/types"
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
					aiflags.SpendFlag(),
					{Name: "model", Help: "model that writes (see the ai registry); overrides --spend", Default: ""},
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
					aiflags.SpendFlag(),
					{Name: "model", Help: "model that writes (see the ai registry); overrides --spend", Default: ""},
					aiflags.ImageModelFlag(""),
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

// resolveModel picks the writing model: --model when set, otherwise the
// registry's compose model at --spend, with that tier's compose effort.
func resolveModel(c *cli.Context) (string, string, error) {
	model, spec, effort, err := ai.TierWriter(c.String("spend"), c.String("model"))
	if err != nil {
		return "", "", cli.NewUsageError(err)
	}
	if spec.Provider != ai.ProviderAnthropic {
		return "", "", cli.NewUsageError(fmt.Errorf("bookgen calls Anthropic; %s is a %s model", model, spec.Provider))
	}
	return model, effort, nil
}

func runBlurb(c *cli.Context) error {
	configPath := c.String("config")
	model, effort, err := resolveModel(c)
	if err != nil {
		return err
	}
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
		Effort: effort,
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
	model, effort, err := resolveModel(c)
	if err != nil {
		return err
	}
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

	provider := &ai.Anthropic{APIKey: cfg.API.AnthropicKey, Pricing: ai.ProviderPricing(ai.ProviderAnthropic)}

	// The cover draws with the image model at --spend unless --image-model
	// names another drawer.
	imageModel, imageSpec, err := aiflags.ResolveTierImageModel(c)
	if err != nil {
		return err
	}
	var imgProvider ai.ImageProvider
	if !promptOnly && !dryRun {
		keyName, kErr := ai.KeyNameForProvider(imageSpec.Provider)
		if kErr != nil {
			return kErr
		}
		imageKey, kErr := creds.Get(keyName)
		if kErr != nil {
			return fmt.Errorf("reading %s: %w", keyName, kErr)
		}
		switch imageSpec.Provider {
		case ai.ProviderGemini:
			imgProvider = &ai.Gemini{APIKey: imageKey}
		case ai.ProviderOpenAI:
			imgProvider = &ai.DallE{APIKey: imageKey}
		default:
			return cli.NewUsageError(fmt.Errorf("bookgen draws with Gemini or OpenAI models, not %s (%s)", imageModel, imageSpec.Provider))
		}
	}

	c.Logger.Info("generating front cover")
	result, err := bookgen.GenerateCover(c.Context, provider, imgProvider, bookgen.CoverInput{
		Title:           bookTitle,
		Author:          author,
		Plan:            planText,
		Blurb:           blurbText,
		Essays:          essays,
		Model:           model,
		Effort:          effort,
		ImageModel:      imageModel,
		ImageQuality:    imageSpec.ImageQuality,
		ImageResolution: imageSpec.ImageResolution,
		DryRun:          dryRun,
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
		// The pipeline finds the cover as front-cover.png, so a JPEG from
		// Gemini is re-encoded rather than renamed.
		pngData, pErr := ai.EnsurePNG(result.ImageData)
		if pErr != nil {
			return pErr
		}
		if wErr := os.WriteFile(imagePath, pngData, appkit.FilePermissions); wErr != nil {
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
	return &ai.Anthropic{APIKey: cfg.API.AnthropicKey, Pricing: ai.ProviderPricing(ai.ProviderAnthropic)}, nil
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
