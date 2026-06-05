package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
	"time"
	"unicode"

	"github.com/TrueBlocks/trueblocks-art/packages/cli"
	"gopkg.in/yaml.v3"
)

var version = "dev"

type SourceMeta struct {
	File        string `yaml:"file"`
	Title       string `yaml:"title"`
	Author      string `yaml:"author"`
	Year        string `yaml:"year"`
	Pages       int    `yaml:"pages"`
	SourceType  string `yaml:"source_type"` // "embedded", "archive", "gutenberg", "ocr-needed"
	SourceURL   string `yaml:"source_url,omitempty"`
	TextQuality string `yaml:"text_quality"` // "high", "medium", "low", "none"
	Difficulty  string `yaml:"difficulty"`   // "easy", "medium", "hard"
	CheckedAt   string `yaml:"checked_at"`
}

var reNamePattern = regexp.MustCompile(`^(\d{4})\s*-\s*(.+?)\s*-\s*(.+)\.(pdf|PDF)$`)

func main() {
	app := cli.App{
		Name:        "check-source",
		Description: "Determine the best text source for a historical PDF and write a .meta.yaml file.",
		Version:     version,
		Flags: []cli.FlagDef{
			{Name: "pdf", Help: "path to the source PDF file", Default: ""},
			{Name: "project-dir", Help: "project directory for output (default: auto from bookmill/projects/)", Default: ""},
			{Name: "skip-online", Help: "skip online source lookups", Default: false},
		},
		Run: run,
	}
	cli.Exit(app.Main())
}

func run(c *cli.Context) error {
	pdfPath := c.String("pdf")
	if pdfPath == "" {
		return fmt.Errorf("--pdf is required")
	}

	absPDF, err := filepath.Abs(pdfPath)
	if err != nil {
		return fmt.Errorf("resolving pdf path: %w", err)
	}
	if _, err := os.Stat(absPDF); err != nil {
		return fmt.Errorf("pdf not found: %s", absPDF)
	}

	skipOnline := c.Bool("skip-online")
	projectDir := c.String("project-dir")

	filename := filepath.Base(absPDF)
	meta := SourceMeta{
		File:      filename,
		CheckedAt: time.Now().Format(time.RFC3339),
	}

	matches := reNamePattern.FindStringSubmatch(filename)
	if matches != nil {
		meta.Year = matches[1]
		meta.Author = matches[2]
		meta.Title = strings.TrimSuffix(matches[3], "."+matches[4])
	} else {
		parts := strings.SplitN(strings.TrimSuffix(filename, filepath.Ext(filename)), " - ", 3)
		if len(parts) >= 3 {
			meta.Year = strings.TrimSpace(parts[0])
			meta.Author = strings.TrimSpace(parts[1])
			meta.Title = strings.TrimSpace(parts[2])
		}
	}

	meta.Pages = countPages(absPDF)
	meta.TextQuality = checkTextQuality(absPDF, meta.Pages)
	fmt.Fprintf(os.Stderr, "PDF text quality: %s\n", meta.TextQuality)

	meta.SourceType = "ocr-needed"
	meta.Difficulty = "hard"

	if meta.TextQuality == "high" || meta.TextQuality == "medium" {
		meta.SourceType = "embedded"
		meta.Difficulty = "easy"
		if meta.TextQuality == "medium" {
			meta.Difficulty = "medium"
		}
	}

	if !skipOnline {
		archiveURL := searchInternetArchive(meta.Title, meta.Author)
		if archiveURL != "" {
			fmt.Fprintf(os.Stderr, "Found Internet Archive: %s\n", archiveURL)
			if meta.SourceType == "ocr-needed" {
				meta.SourceType = "archive"
				meta.Difficulty = "easy"
			}
			meta.SourceURL = archiveURL
		} else {
			gutenbergURL := searchGutenberg(meta.Title, meta.Author)
			if gutenbergURL != "" {
				fmt.Fprintf(os.Stderr, "Found Gutenberg: %s\n", gutenbergURL)
				if meta.SourceType == "ocr-needed" {
					meta.SourceType = "gutenberg"
					meta.Difficulty = "easy"
				}
				meta.SourceURL = gutenbergURL
			}
		}
	}

	out, err := yaml.Marshal(meta)
	if err != nil {
		return fmt.Errorf("marshaling meta: %w", err)
	}

	if projectDir != "" {
		stageDir := filepath.Join(projectDir, "check-source")
		if err := os.MkdirAll(stageDir, 0755); err != nil {
			return fmt.Errorf("creating stage dir: %w", err)
		}
		metaPath := filepath.Join(stageDir, ".meta.yaml")
		if err := os.WriteFile(metaPath, out, 0644); err != nil {
			return fmt.Errorf("writing meta: %w", err)
		}
		fmt.Fprintf(os.Stderr, "Wrote %s\n", metaPath)
	} else {
		fmt.Print(string(out))
	}

	return nil
}

func countPages(pdfPath string) int {
	cmd := exec.Command("pdfinfo", pdfPath)
	out, err := cmd.Output()
	if err != nil {
		return 0
	}
	for _, line := range strings.Split(string(out), "\n") {
		if strings.HasPrefix(line, "Pages:") {
			var pages int
			fmt.Sscanf(strings.TrimPrefix(line, "Pages:"), "%d", &pages)
			return pages
		}
	}
	return 0
}

func checkTextQuality(pdfPath string, totalPages int) string {
	if totalPages == 0 {
		return "none"
	}

	samplePoints := []int{1}
	if totalPages > 10 {
		samplePoints = append(samplePoints, totalPages/10)
	}
	if totalPages > 4 {
		samplePoints = append(samplePoints, totalPages/4)
	}
	if totalPages > 2 {
		samplePoints = append(samplePoints, totalPages/2)
	}
	if totalPages > 4 {
		samplePoints = append(samplePoints, totalPages*3/4)
	}

	good := 0
	for _, p := range samplePoints {
		text := extractPage(pdfPath, p)
		if isCoherentText(text) {
			good++
		}
	}

	ratio := float64(good) / float64(len(samplePoints))
	if ratio >= 0.7 {
		return "high"
	}
	if ratio >= 0.3 {
		return "medium"
	}
	if good > 0 {
		return "low"
	}
	return "none"
}

func extractPage(pdfPath string, page int) string {
	cmd := exec.Command("pdftotext", "-f", fmt.Sprintf("%d", page), "-l", fmt.Sprintf("%d", page), pdfPath, "-")
	out, err := cmd.Output()
	if err != nil {
		return ""
	}
	return string(out)
}

func isCoherentText(text string) bool {
	cleaned := strings.TrimSpace(text)
	if len(cleaned) < 50 {
		return false
	}
	alphaCount := 0
	totalCount := 0
	for _, r := range cleaned {
		if !unicode.IsSpace(r) {
			totalCount++
			if unicode.IsLetter(r) {
				alphaCount++
			}
		}
	}
	if totalCount == 0 {
		return false
	}
	if float64(alphaCount)/float64(totalCount) < 0.6 {
		return false
	}
	words := strings.Fields(cleaned)
	return len(words) >= 10
}

func searchInternetArchive(title, author string) string {
	_ = author
	query := sanitizeSearchTerm(title)
	searchURL := fmt.Sprintf(
		"https://archive.org/advancedsearch.php?q=title:(%s)+AND+mediatype:(texts)&fl[]=identifier,title&rows=5&output=json",
		url.QueryEscape(query),
	)

	resp, err := httpGetJSON(searchURL)
	if err != nil {
		return ""
	}

	response, ok := resp["response"].(map[string]interface{})
	if !ok {
		return ""
	}
	docs, ok := response["docs"].([]interface{})
	if !ok || len(docs) == 0 {
		return ""
	}

	for _, d := range docs {
		doc, ok := d.(map[string]interface{})
		if !ok {
			continue
		}
		docTitle, _ := doc["title"].(string)
		if fuzzyMatch(docTitle, title) {
			id, _ := doc["identifier"].(string)
			if id != "" {
				return fmt.Sprintf("https://archive.org/details/%s", id)
			}
		}
	}
	return ""
}

func searchGutenberg(title, author string) string {
	_ = author
	query := sanitizeSearchTerm(title)
	searchURL := fmt.Sprintf("https://gutendex.com/books/?search=%s", url.QueryEscape(query))

	client := &http.Client{Timeout: 10 * time.Second}
	req, err := http.NewRequest("GET", searchURL, nil)
	if err != nil {
		return ""
	}
	httpResp, err := client.Do(req)
	if err != nil {
		return ""
	}
	defer httpResp.Body.Close()
	if httpResp.StatusCode != http.StatusOK {
		return ""
	}
	body, err := io.ReadAll(httpResp.Body)
	if err != nil {
		return ""
	}

	var result map[string]interface{}
	if err := decodeJSON(body, &result); err != nil {
		return ""
	}

	results, ok := result["results"].([]interface{})
	if !ok || len(results) == 0 {
		return ""
	}

	for _, r := range results {
		book, ok := r.(map[string]interface{})
		if !ok {
			continue
		}
		bookTitle, _ := book["title"].(string)
		if fuzzyMatch(bookTitle, title) {
			id, _ := book["id"].(float64)
			if id > 0 {
				return fmt.Sprintf("https://www.gutenberg.org/ebooks/%d", int(id))
			}
		}
	}
	return ""
}

func httpGetJSON(rawURL string) (map[string]interface{}, error) {
	client := &http.Client{Timeout: 15 * time.Second}
	resp, err := client.Get(rawURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status %d", resp.StatusCode)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var result map[string]interface{}
	if err := decodeJSON(body, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func decodeJSON(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}

func sanitizeSearchTerm(s string) string {
	s = strings.TrimSpace(s)
	words := strings.Fields(s)
	if len(words) > 6 {
		words = words[:6]
	}
	return strings.Join(words, " ")
}

func fuzzyMatch(candidate, target string) bool {
	candidate = strings.ToLower(candidate)
	target = strings.ToLower(target)
	if strings.Contains(candidate, target) || strings.Contains(target, candidate) {
		return true
	}
	targetWords := strings.Fields(target)
	if len(targetWords) < 2 {
		return false
	}
	matchCount := 0
	for _, w := range targetWords {
		if len(w) > 2 && strings.Contains(candidate, w) {
			matchCount++
		}
	}
	return float64(matchCount)/float64(len(targetWords)) >= 0.5
}
