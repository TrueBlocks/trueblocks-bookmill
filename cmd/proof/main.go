package main

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
	"time"

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
			{Name: "model", Help: "vision model to use (default: gpt-4o)", Default: "gpt-4o"},
			{Name: "api-key", Help: "OpenAI API key (default: from OPENAI_API_KEY env)", Default: ""},
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
	model := c.String("model")
	apiKey := c.String("api-key")
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

	if apiKey == "" {
		apiKey = os.Getenv("OPENAI_API_KEY")
	}
	if apiKey == "" && !dryRun {
		return fmt.Errorf("--api-key or OPENAI_API_KEY environment variable is required")
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

		corrected, err := proofWithVision(apiKey, model, text, pageImage)
		if err != nil {
			fmt.Fprintf(os.Stderr, "  warning: vision API error on page %d: %v\n", page.pageNum, err)
			proofed = append(proofed, page.raw)
			continue
		}

		proofed = append(proofed, fmt.Sprintf("<!-- page %d -->\n\n%s", page.pageNum, corrected))
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

func proofWithVision(apiKey, model, text string, pageImage []byte) (string, error) {
	imageB64 := base64.StdEncoding.EncodeToString(pageImage)

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

Return ONLY the corrected text. No explanations, no markdown fences.`

	payload := map[string]interface{}{
		"model": model,
		"messages": []map[string]interface{}{
			{
				"role":    "system",
				"content": prompt,
			},
			{
				"role": "user",
				"content": []map[string]interface{}{
					{
						"type": "text",
						"text": fmt.Sprintf("Here is the extracted text from this page:\n\n%s", text),
					},
					{
						"type": "image_url",
						"image_url": map[string]string{
							"url":    fmt.Sprintf("data:image/png;base64,%s", imageB64),
							"detail": "high",
						},
					},
				},
			},
		},
		"max_tokens":  4096,
		"temperature": 0.1,
	}

	body, err := json.Marshal(payload)
	if err != nil {
		return "", fmt.Errorf("marshaling request: %w", err)
	}

	req, err := http.NewRequest("POST", "https://api.openai.com/v1/chat/completions", bytes.NewReader(body))
	if err != nil {
		return "", fmt.Errorf("creating request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+apiKey)

	client := &http.Client{Timeout: 120 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("API request failed: %w", err)
	}
	defer func() { _ = resp.Body.Close() }()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("reading response: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("API returned status %d: %s", resp.StatusCode, string(respBody))
	}

	var result struct {
		Choices []struct {
			Message struct {
				Content string `json:"content"`
			} `json:"message"`
		} `json:"choices"`
	}
	if err := json.Unmarshal(respBody, &result); err != nil {
		return "", fmt.Errorf("parsing response: %w", err)
	}

	if len(result.Choices) == 0 {
		return "", fmt.Errorf("no choices in response")
	}

	return strings.TrimSpace(result.Choices[0].Message.Content), nil
}
