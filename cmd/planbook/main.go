package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/TrueBlocks/trueblocks-math/internal/pipeline"
)

func main() {
	originsDir := flag.String("origins", "", "path to the origins folder (required)")
	examplePlan := flag.String("example", "", "path to an existing Plan file to use as format reference")
	outputPath := flag.String("output", "", "output file path (default: stdout)")
	bookTitle := flag.String("title", "", "working title for the book")
	dryRun := flag.Bool("dry-run", false, "print the prompt without calling the API")
	model := flag.String("model", "claude-sonnet-4-20250514", "Anthropic model to use")
	configPath := flag.String("config", pipeline.DefaultConfigPath(), "path to config.yaml for API key")
	flag.Parse()

	if *originsDir == "" {
		fmt.Fprintln(os.Stderr, "Error: --origins is required")
		flag.Usage()
		os.Exit(1)
	}

	designDir := filepath.Dir(filepath.Clean(*originsDir))
	genre, err := pipeline.LoadGenre(designDir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading genre: %v\n", err)
		os.Exit(1)
	}
	fmt.Fprintf(os.Stderr, "Genre: %s/%s\n", genre.Form, genre.Flavor)

	origins, err := readOrigins(*originsDir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading origins: %v\n", err)
		os.Exit(1)
	}
	if len(origins) == 0 {
		fmt.Fprintf(os.Stderr, "No files found in %s\n", *originsDir)
		os.Exit(1)
	}

	var example string
	if *examplePlan != "" {
		data, err := os.ReadFile(*examplePlan)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading example plan: %v\n", err)
			os.Exit(1)
		}
		example = string(data)
	}

	if *outputPath != "" {
		if _, err := os.Stat(*outputPath); err == nil {
			fmt.Fprintf(os.Stderr, "Warning: overwriting existing %s\n", *outputPath)
		}
	}

	prompt := buildPrompt(origins, example, *bookTitle, genre)

	if *dryRun {
		fmt.Println(prompt)
		return
	}

	cfg, err := pipeline.LoadConfig(*configPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading config: %v\n", err)
		os.Exit(1)
	}
	if cfg.API.AnthropicKey == "" {
		fmt.Fprintln(os.Stderr, "Error: no anthropic_key in config")
		os.Exit(1)
	}

	client := &pipeline.AnthropicClient{APIKey: cfg.API.AnthropicKey}

	fmt.Fprintln(os.Stderr, "Calling API...")
	ctx := context.Background()
	result, err := client.Call(ctx, *model, prompt, 5*time.Minute)
	if err != nil {
		fmt.Fprintf(os.Stderr, "API error: %v\n", err)
		os.Exit(1)
	}

	fmt.Fprintf(os.Stderr, "Tokens: %d in, %d out (cost: $%.4f)\n",
		result.InputTokens, result.OutputTokens, result.Cost)

	output := result.Content

	if *outputPath != "" {
		dir := filepath.Dir(*outputPath)
		if err := os.MkdirAll(dir, 0755); err != nil {
			fmt.Fprintf(os.Stderr, "Error creating output directory: %v\n", err)
			os.Exit(1)
		}
		if err := os.WriteFile(*outputPath, []byte(output), 0644); err != nil {
			fmt.Fprintf(os.Stderr, "Error writing output: %v\n", err)
			os.Exit(1)
		}
		fmt.Fprintf(os.Stderr, "Written to %s\n", *outputPath)
	} else {
		fmt.Println(output)
	}
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
