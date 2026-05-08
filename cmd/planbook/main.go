package main

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	appkit "github.com/TrueBlocks/trueblocks-art/packages/appkit/v2"
	"github.com/TrueBlocks/trueblocks-art/packages/cli"
	"github.com/TrueBlocks/trueblocks-math/internal/pipeline"
)

// version is injected via -ldflags at build time.
var version = "dev"

func main() {
	app := cli.App{
		Name:        "planbook",
		Description: "Generate a structured book Plan from origin material via the Anthropic API.",
		Version:     version,
		Flags: []cli.FlagDef{
			{Name: "origins", Help: "path to the origins folder", Required: true},
			{Name: "example", Help: "path to an existing Plan file to use as format reference"},
			{Name: "output", Help: "output file path (default: stdout)"},
			{Name: "title", Help: "working title for the book"},
			{Name: "dry-run", Help: "print the prompt without calling the API", Default: false},
			{Name: "model", Help: "Anthropic model to use", Default: "claude-sonnet-4-20250514"},
			{Name: "config", Help: "path to config.yaml for API key", Default: pipeline.DefaultConfigPath()},
		},
		Run: run,
	}
	cli.Exit(app.Main())
}

func run(c *cli.Context) error {
	originsDir := c.MustString("origins")
	examplePlan := c.String("example")
	outputPath := c.String("output")
	bookTitle := c.String("title")
	dryRun := c.Bool("dry-run")
	model := c.String("model")
	configPath := c.String("config")

	designDir := filepath.Dir(filepath.Clean(originsDir))
	genre, err := pipeline.LoadGenre(designDir)
	if err != nil {
		return fmt.Errorf("loading genre: %w", err)
	}
	c.Logger.Info("loaded genre", "form", genre.Form, "flavor", genre.Flavor)

	origins, err := readOrigins(originsDir)
	if err != nil {
		return fmt.Errorf("reading origins: %w", err)
	}
	if len(origins) == 0 {
		return fmt.Errorf("no files found in %s", originsDir)
	}

	var example string
	if examplePlan != "" {
		data, err := os.ReadFile(examplePlan)
		if err != nil {
			return fmt.Errorf("reading example plan: %w", err)
		}
		example = string(data)
	}

	if outputPath != "" {
		if _, err := os.Stat(outputPath); err == nil {
			c.Logger.Warn("overwriting existing output file", "path", outputPath)
		}
	}

	prompt := buildPrompt(origins, example, bookTitle, genre)

	if dryRun {
		fmt.Println(prompt)
		return nil
	}

	cfg, err := pipeline.LoadConfig(configPath)
	if err != nil {
		return fmt.Errorf("loading config: %w", err)
	}
	if cfg.API.AnthropicKey == "" {
		return fmt.Errorf("no anthropic_key in config")
	}

	client := &pipeline.AnthropicClient{APIKey: cfg.API.AnthropicKey}

	c.Logger.Info("calling API", "model", model)
	callCtx, cancel := context.WithTimeout(c.Context, 5*time.Minute)
	defer cancel()
	result, err := client.Call(callCtx, model, prompt, 5*time.Minute)
	if err != nil {
		return fmt.Errorf("API: %w", err)
	}

	c.Logger.Info("api result",
		"input_tokens", result.InputTokens,
		"output_tokens", result.OutputTokens,
		"cost_usd", result.Cost,
	)

	output := result.Content

	if outputPath != "" {
		dir := filepath.Dir(outputPath)
		if err := os.MkdirAll(dir, appkit.DirPermissions); err != nil {
			return fmt.Errorf("creating output directory: %w", err)
		}
		if err := os.WriteFile(outputPath, []byte(output), appkit.FilePermissions); err != nil {
			return fmt.Errorf("writing output: %w", err)
		}
		c.Logger.Info("wrote output", "path", outputPath)
		return nil
	}

	fmt.Println(output)
	return nil
}

type originFile struct {
	name    string
	content string
}

func readOrigins(dir string) ([]originFile, error) {
	var origins []originFile
	err := filepath.WalkDir(dir, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			return nil
		}
		ext := strings.ToLower(filepath.Ext(d.Name()))
		if ext != ".md" && ext != ".txt" && ext != ".yaml" && ext != ".yml" {
			return nil
		}
		data, err := os.ReadFile(path)
		if err != nil {
			return fmt.Errorf("reading %s: %w", path, err)
		}
		relPath, _ := filepath.Rel(dir, path)
		origins = append(origins, originFile{name: relPath, content: string(data)})
		return nil
	})
	return origins, err
}

func buildPrompt(origins []originFile, example, title string, genre *pipeline.Genre) string {
	var sb strings.Builder

	if genre.IsNovel() {
		sb.WriteString("You are an expert novel plotter and story architect. Your task is to read the origin material below and produce a structured Plan for a novel.\n\n")
	} else {
		sb.WriteString("You are an expert book editor and planner. Your task is to read the origin material below and produce a structured Plan for a book.\n\n")
	}
	sb.WriteString("RULES:\n")
	sb.WriteString("1. The Plan must be a single markdown document.\n")
	sb.WriteString("2. Start with a top-level heading: \"# Plan for <Book Title>\"\n")
	if genre.IsNovel() {
		sb.WriteString("3. Include a \"## Thesis\" section (2-3 paragraphs describing the novel's premise, central conflict, and character arcs).\n")
		sb.WriteString("4. Include a \"## The Storyline in One Sentence\" section with a single italicized sentence in a blockquote.\n")
		sb.WriteString("5. Then list all chapters in a single markdown table with these columns:\n")
		sb.WriteString("   | # | Type | Slug | Title | Hook | Subplot |\n")
		sb.WriteString("   - \"#\" is the chapter number (integer, sequential starting at 1)\n")
		sb.WriteString("   - \"Type\" is one of: chapter, introduction, epilogue\n")
		sb.WriteString("   - \"Slug\" is a lowercase-hyphenated identifier derived from the title\n")
		sb.WriteString("   - \"Title\" is the chapter title\n")
		sb.WriteString("   - \"Hook\" is a 1-sentence description of the scene or event that drives the chapter\n")
		sb.WriteString("   - \"Subplot\" is the secondary thread, character arc, or thematic undercurrent in this chapter\n")
	} else {
		sb.WriteString("3. Include a \"## Thesis\" section (2-3 paragraphs describing the book's theme, character, and arc).\n")
		sb.WriteString("4. Include a \"## The Storyline in One Sentence\" section with a single italicized sentence in a blockquote.\n")
		sb.WriteString("5. Then list all chapters in a single markdown table with these columns:\n")
		sb.WriteString("   | # | Type | Slug | Title | Hook | Hidden Theme |\n")
		sb.WriteString("   - \"#\" is the chapter number (integer, sequential starting at 1)\n")
		sb.WriteString("   - \"Type\" is one of: chapter, introduction, epilogue\n")
		sb.WriteString("   - \"Slug\" is a lowercase-hyphenated identifier derived from the title\n")
		sb.WriteString("   - \"Title\" is the chapter title\n")
		sb.WriteString("   - \"Hook\" is a 1-sentence description of what draws the reader in\n")
		sb.WriteString("   - \"Hidden Theme\" is the deeper idea or pattern the chapter explores beneath the surface\n")
	}
	sb.WriteString("6. Do NOT use Parts or Sections -- this is a flat list of chapters.\n")
	sb.WriteString("7. End with a \"## Summary\" section showing total chapter count.\n")
	sb.WriteString("8. Derive ALL chapter ideas from the origin material. Expand, restructure, and fill gaps -- but stay rooted in what the origins provide.\n")
	if genre.IsNovel() {
		sb.WriteString("9. The Plan should feel like a coherent novel with narrative momentum, not a collection of unrelated scenes.\n")
		sb.WriteString("10. Chapters should build on each other -- each chapter should advance the plot, deepen a character, or raise the stakes.\n\n")
	} else {
		sb.WriteString("9. The Plan should feel like a coherent book, not a collection of unrelated pieces.\n\n")
	}

	if title != "" {
		sb.WriteString(fmt.Sprintf("BOOK TITLE: %s\n\n", title))
	}

	if example != "" {
		sb.WriteString("EXAMPLE PLAN (use this as a format reference -- match its style and structure):\n")
		sb.WriteString("---\n")
		sb.WriteString(example)
		if !strings.HasSuffix(example, "\n") {
			sb.WriteString("\n")
		}
		sb.WriteString("---\n\n")
	}

	sb.WriteString("ORIGIN MATERIAL:\n\n")
	for _, o := range origins {
		sb.WriteString(fmt.Sprintf("### File: %s\n", o.name))
		sb.WriteString("```\n")
		sb.WriteString(o.content)
		if !strings.HasSuffix(o.content, "\n") {
			sb.WriteString("\n")
		}
		sb.WriteString("```\n\n")
	}

	sb.WriteString("Now produce the Plan. Output ONLY the markdown document, no commentary.\n")

	return sb.String()
}
