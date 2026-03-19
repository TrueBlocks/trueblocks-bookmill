package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/TrueBlocks/trueblocks-art/packages/ai"
	"github.com/TrueBlocks/trueblocks-art/packages/bookgen"
	"github.com/TrueBlocks/trueblocks-math/internal/bookutil"
	"github.com/TrueBlocks/trueblocks-math/internal/pipeline"
	"github.com/TrueBlocks/trueblocks-math/internal/types"
)

func main() {
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}

	cmd := os.Args[1]
	os.Args = append(os.Args[:1], os.Args[2:]...)

	switch cmd {
	case "blurb":
		runBlurb()
	case "cover":
		runCover()
	default:
		fmt.Fprintf(os.Stderr, "Unknown command: %s\n", cmd)
		printUsage()
		os.Exit(1)
	}
}

func printUsage() {
	fmt.Fprintln(os.Stderr, "Usage: bookgen <command> [flags] <project-dir>")
	fmt.Fprintln(os.Stderr, "Commands:")
	fmt.Fprintln(os.Stderr, "  blurb   Generate back-cover blurb")
	fmt.Fprintln(os.Stderr, "  cover   Generate front-cover image")
}

func runBlurb() {
	fs := flag.NewFlagSet("blurb", flag.ExitOnError)
	configPath := fs.String("config", pipeline.DefaultConfigPath(), "path to config.yaml")
	model := fs.String("model", "claude-sonnet-4-20250514", "Anthropic model")
	dryRun := fs.Bool("dry-run", false, "print prompt without calling API")
	force := fs.Bool("force", false, "regenerate even if blurb exists")
	fs.Parse(os.Args[1:])

	if fs.NArg() < 1 {
		fmt.Fprintln(os.Stderr, "Usage: bookgen blurb [flags] <project-dir>")
		os.Exit(1)
	}
	projectDir := fs.Arg(0)

	outPath := filepath.Join(projectDir, "book", "back-cover-blurb.md")
	if !*force {
		if _, err := os.Stat(outPath); err == nil {
			fmt.Fprintln(os.Stderr, "blurb already exists:", outPath)
			fmt.Fprintln(os.Stderr, "use --force to regenerate")
			os.Exit(0)
		}
	}

	provider, err := loadProvider(*configPath, *dryRun)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	planText := bookutil.FindPlan(projectDir)
	rawEssays := bookutil.ReadDraft2(projectDir, 0)
	if len(rawEssays) == 0 {
		fmt.Fprintln(os.Stderr, "no draft2 essays found in", projectDir)
		os.Exit(1)
	}

	essays := convertEssays(rawEssays)

	ctx := context.Background()
	fmt.Fprintln(os.Stderr, "Generating back-cover blurb...")
	result, err := bookgen.GenerateBlurb(ctx, provider, bookgen.BlurbInput{
		Plan:   planText,
		Essays: essays,
		Model:  *model,
		DryRun: *dryRun,
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	if *dryRun {
		fmt.Println(result.Content)
		return
	}

	fmt.Fprintf(os.Stderr, "Tokens: %d in, %d out (cost: $%.4f)\n",
		result.InputTokens, result.OutputTokens, result.Cost)

	if err := os.MkdirAll(filepath.Dir(outPath), 0755); err != nil {
		fmt.Fprintf(os.Stderr, "Error creating output directory: %v\n", err)
		os.Exit(1)
	}
	if err := os.WriteFile(outPath, []byte(result.Content+"\n"), 0644); err != nil {
		fmt.Fprintf(os.Stderr, "Error writing output: %v\n", err)
		os.Exit(1)
	}
	fmt.Fprintf(os.Stderr, "Written to %s\n", outPath)
}

func runCover() {
	fs := flag.NewFlagSet("cover", flag.ExitOnError)
	configPath := fs.String("config", pipeline.DefaultConfigPath(), "path to config.yaml")
	model := fs.String("model", "claude-sonnet-4-20250514", "Anthropic model")
	dryRun := fs.Bool("dry-run", false, "print prompt without calling API")
	promptOnly := fs.Bool("prompt-only", false, "generate cover prompt, skip image")
	force := fs.Bool("force", false, "regenerate even if artifacts exist")
	title := fs.String("title", "", "book title (extracted from Plan if not specified)")
	author := fs.String("author", "Claude Jay Rush", "author name")
	fs.Parse(os.Args[1:])

	if fs.NArg() < 1 {
		fmt.Fprintln(os.Stderr, "Usage: bookgen cover [flags] <project-dir>")
		os.Exit(1)
	}
	projectDir := fs.Arg(0)

	bookDir := filepath.Join(projectDir, "book")
	promptPath := filepath.Join(bookDir, "front-cover-prompt.md")
	imagePath := filepath.Join(bookDir, "front-cover.png")

	needPrompt := *force || !fileExists(promptPath)
	needImage := *force || !fileExists(imagePath)

	if !needPrompt && !needImage {
		fmt.Fprintln(os.Stderr, "cover artifacts already exist:", bookDir)
		fmt.Fprintln(os.Stderr, "use --force to regenerate")
		os.Exit(0)
	}

	cfg, err := pipeline.LoadConfig(*configPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading config: %v\n", err)
		os.Exit(1)
	}

	planText := bookutil.FindPlan(projectDir)
	rawEssays := bookutil.ReadDraft2(projectDir, 500)

	bookTitle := *title
	if bookTitle == "" {
		bookTitle = extractTitleFromPlan(planText)
	}
	if bookTitle == "" {
		bookTitle = filepath.Base(projectDir)
	}

	blurbText := readBlurb(projectDir)
	essays := convertEssays(rawEssays)

	provider := &ai.Anthropic{APIKey: cfg.API.AnthropicKey}

	var imgProvider ai.ImageProvider
	if !*promptOnly && !*dryRun {
		imgProvider = &ai.DallE{APIKey: cfg.API.OpenAIKey}
	}

	fmt.Fprintln(os.Stderr, "Generating front cover...")
	ctx := context.Background()
	result, err := bookgen.GenerateCover(ctx, provider, imgProvider, bookgen.CoverInput{
		Title:  bookTitle,
		Author: *author,
		Plan:   planText,
		Blurb:  blurbText,
		Essays: essays,
		Model:  *model,
		DryRun: *dryRun,
	})
	if err != nil && result == nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	if *dryRun {
		fmt.Println(result.DesignDoc)
		return
	}

	if needPrompt && result.DesignDoc != "" {
		if mkErr := os.MkdirAll(bookDir, 0755); mkErr != nil {
			fmt.Fprintf(os.Stderr, "Error creating directory: %v\n", mkErr)
			os.Exit(1)
		}
		if wErr := os.WriteFile(promptPath, []byte(result.DesignDoc+"\n"), 0644); wErr != nil {
			fmt.Fprintf(os.Stderr, "Error writing prompt: %v\n", wErr)
			os.Exit(1)
		}
		fmt.Fprintf(os.Stderr, "Cover prompt written to %s\n", promptPath)
	}

	if result.ImageData != nil {
		if wErr := os.WriteFile(imagePath, result.ImageData, 0644); wErr != nil {
			fmt.Fprintf(os.Stderr, "Error writing image: %v\n", wErr)
			os.Exit(1)
		}
		fmt.Fprintf(os.Stderr, "Cover image written to %s\n", imagePath)
	}

	if err != nil {
		fmt.Fprintf(os.Stderr, "Warning: %v\n", err)
	}
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
