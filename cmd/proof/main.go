package main

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	"github.com/TrueBlocks/trueblocks-art/packages/ai"
	"github.com/TrueBlocks/trueblocks-art/packages/cli"
)

var version = "dev"

var rePageMarker = regexp.MustCompile(`<!-- page (\d+) -->`)

func main() {
	app := cli.App{
		Name:        "proof",
		Description: "Proofread extracted text by comparing against rendered PDF page images using a vision model. Fixes OCR errors while preserving original spelling and style.",
		Version:     version,
		Flags: []cli.FlagDef{
			{Name: "input", Help: "path to extracted markdown file (from extract-text)", Default: ""},
			{Name: "pdf", Help: "path to the source PDF file (for rendering page images)", Default: ""},
			{Name: "output", Help: "output proofed markdown file (default: stdout)", Default: ""},
			{Name: "text-model", Help: "vision model that proofs each page", Default: "gpt-4o"},
			{Name: "dpi", Help: "DPI for rendering PDF pages (default: 200)", Default: 200},
			{Name: "start-page", Help: "start proofreading from this page number (default: 1)", Default: 1},
			{Name: "end-page", Help: "stop proofreading at this page number (default: all)", Default: 0},
			{Name: "dry-run", Help: "show what would be proofed without calling the API", Default: false},
		},
		Run: run,
	}
	cli.Exit(app.Main())
}

func run(c *cli.Context) error {
	inputPath := c.String("input")
	pdfPath := c.String("pdf")
	outputPath := c.String("output")
	model := c.String("text-model")
	dpi := c.Int("dpi")
	startPage := c.Int("start-page")
	endPage := c.Int("end-page")
	dryRun := c.Bool("dry-run")

	if inputPath == "" {
		return fmt.Errorf("--input is required")
	}
	if pdfPath == "" {
		return fmt.Errorf("--pdf is required")
	}

	var provider *ai.OpenAI
	if !dryRun {
		cfg, err := ai.LoadSharedConfig()
		if err != nil {
			return fmt.Errorf("loading AI config: %w", err)
		}
		provider = cfg.NewOpenAI()
	}

	absInput, err := filepath.Abs(inputPath)
	if err != nil {
		return fmt.Errorf("resolving input path: %w", err)
	}
	absPDF, err := filepath.Abs(pdfPath)
	if err != nil {
		return fmt.Errorf("resolving pdf path: %w", err)
	}

	content, err := os.ReadFile(absInput)
	if err != nil {
		return fmt.Errorf("reading input: %w", err)
	}

	pages := splitByPageMarkers(string(content))
	fmt.Fprintf(os.Stderr, "Found %d pages to proof\n", len(pages))

	var proofed []string
	var totalCost float64
	for i, page := range pages {
		if page.pageNum < startPage {
			proofed = append(proofed, page.raw)
			continue
		}
		if endPage > 0 && page.pageNum > endPage {
			proofed = append(proofed, page.raw)
			continue
		}

		text := strings.TrimSpace(page.text)
		if len(text) < 40 {
			proofed = append(proofed, page.raw)
			continue
		}

		fmt.Fprintf(os.Stderr, "[%d/%d] Proofing page %d...\n", i+1, len(pages), page.pageNum)

		if dryRun {
			proofed = append(proofed, page.raw)
			continue
		}

		pageImage, err := renderPage(absPDF, page.pageNum, dpi)
		if err != nil {
			fmt.Fprintf(os.Stderr, "  warning: could not render page %d: %v\n", page.pageNum, err)
			proofed = append(proofed, page.raw)
			continue
		}

		corrected, cost, err := proofWithVision(provider, model, text, pageImage)
		if err != nil {
			fmt.Fprintf(os.Stderr, "  warning: vision API error on page %d: %v\n", page.pageNum, err)
			proofed = append(proofed, page.raw)
			continue
		}
		totalCost += cost

		proofed = append(proofed, fmt.Sprintf("<!-- page %d -->\n\n%s", page.pageNum, corrected))
	}

	if !dryRun {
		fmt.Fprintf(os.Stderr, "Total API cost: $%.4f\n", totalCost)
	}

	result := strings.Join(proofed, "\n\n") + "\n"

	if outputPath != "" {
		dir := filepath.Dir(outputPath)
		if err := os.MkdirAll(dir, 0755); err != nil {
			return fmt.Errorf("creating output dir: %w", err)
		}
		if err := os.WriteFile(outputPath, []byte(result), 0644); err != nil {
			return fmt.Errorf("writing output: %w", err)
		}
		fmt.Fprintf(os.Stderr, "Wrote proofed text to %s\n", outputPath)
	} else {
		fmt.Print(result)
	}

	return nil
}

type pageBlock struct {
	pageNum int
	text    string
	raw     string
}

func splitByPageMarkers(content string) []pageBlock {
	matches := rePageMarker.FindAllStringIndex(content, -1)
	if len(matches) == 0 {
		return []pageBlock{{pageNum: 1, text: content, raw: content}}
	}

	var pages []pageBlock
	for i, loc := range matches {
		marker := content[loc[0]:loc[1]]
		var pageNum int
		_, _ = fmt.Sscanf(marker, "<!-- page %d -->", &pageNum)

		var end int
		if i+1 < len(matches) {
			end = matches[i+1][0]
		} else {
			end = len(content)
		}

		raw := content[loc[0]:end]
		text := strings.TrimPrefix(content[loc[1]:end], "\n")

		pages = append(pages, pageBlock{
			pageNum: pageNum,
			text:    text,
			raw:     strings.TrimRight(raw, "\n"),
		})
	}

	return pages
}

func renderPage(pdfPath string, pageNum, dpi int) ([]byte, error) {
	tmpFile := filepath.Join(os.TempDir(), fmt.Sprintf("proof-page-%d.png", pageNum))
	defer func() { _ = os.Remove(tmpFile) }()

	cmd := exec.Command("pdftoppm",
		"-f", fmt.Sprintf("%d", pageNum),
		"-l", fmt.Sprintf("%d", pageNum),
		"-r", fmt.Sprintf("%d", dpi),
		"-png",
		"-singlefile",
		pdfPath,
		strings.TrimSuffix(tmpFile, ".png"),
	)
	if err := cmd.Run(); err != nil {
		return nil, fmt.Errorf("pdftoppm failed: %w", err)
	}

	data, err := os.ReadFile(tmpFile)
	if err != nil {
		return nil, fmt.Errorf("reading rendered page: %w", err)
	}
	return data, nil
}

func proofWithVision(provider *ai.OpenAI, model, text string, pageImage []byte) (string, float64, error) {
	prompt := `You are proofreading OCR-extracted text from a historical book (pre-1926).

Compare the extracted text against the page image. Fix ONLY:
- OCR errors (garbled characters, wrong letters, broken ligatures)
- Missing or extra characters from scanning artifacts
- Long-s (ſ) misread as f — correct to s
- Broken words at line endings that should be rejoined

Do NOT change:
- Original spelling (colour, connexion, etc.)
- Archaic word usage or grammar
- Original punctuation
- Capitalization style

If you are uncertain about a correction, wrap it in: <!-- uncertain: original_text → corrected_text -->

Return ONLY the corrected text. No explanations, no markdown fences.

Here is the extracted text from this page:

` + text

	result, err := provider.Call(context.Background(), model, prompt, ai.CallOptions{
		MaxTokens: 4096,
		Timeout:   120 * time.Second,
		Images:    []ai.ImageInput{{MediaType: "image/png", Data: pageImage}},
	})
	if err != nil {
		return "", 0, err
	}

	return strings.TrimSpace(result.Content), result.Cost, nil
}
