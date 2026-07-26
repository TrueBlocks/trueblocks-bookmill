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
	"sort"
	"strings"
	"time"
	"unicode"

	"github.com/TrueBlocks/trueblocks-art/packages/cli"
	"gopkg.in/yaml.v3"
)

var version = "dev"

type BookEntry struct {
	File         string `yaml:"file" json:"file"`
	Title        string `yaml:"title" json:"title"`
	Author       string `yaml:"author" json:"author"`
	Year         string `yaml:"year" json:"year"`
	Pages        int    `yaml:"pages" json:"pages"`
	EmbeddedText string `yaml:"embedded_text" json:"embedded_text"` // "high", "medium", "low", "none"
	GutenbergURL string `yaml:"gutenberg_url,omitempty" json:"gutenberg_url,omitempty"`
	ArchiveURL   string `yaml:"archive_url,omitempty" json:"archive_url,omitempty"`
	Difficulty   string `yaml:"difficulty" json:"difficulty"` // "easy", "medium", "hard"
	NameValid    bool   `yaml:"name_valid" json:"name_valid"`
	NameIssue    string `yaml:"name_issue,omitempty" json:"name_issue,omitempty"`
}

type Inventory struct {
	GeneratedAt string      `yaml:"generated_at" json:"generated_at"`
	SourceDir   string      `yaml:"source_dir" json:"source_dir"`
	TotalBooks  int         `yaml:"total_books" json:"total_books"`
	Books       []BookEntry `yaml:"books" json:"books"`
}

// reNamePattern matches YYYY - Author - Title.pdf
var reNamePattern = regexp.MustCompile(`^(\d{4})\s*-\s*(.+?)\s*-\s*(.+)\.(pdf|PDF)$`)

func main() {
	app := cli.App{
		Name:        "inventory",
		Description: "Scan a directory of historical PDFs, validate naming, check for embedded text and online sources, and produce a sorted inventory.",
		Version:     version,
		Flags: []cli.FlagDef{
			{Name: "dir", Help: "path to the PDF directory to scan", Default: ""},
			{Name: "output", Help: "output file path (default: stdout)", Default: ""},
			{Name: "format", Help: "output format: yaml or json", Default: "yaml"},
			{Name: "skip-online", Help: "skip online source lookups (faster)", Default: false},
		},
		Run: run,
	}
	cli.Exit(app.Main())
}

func run(c *cli.Context) error {
	dir := c.String("dir")
	if dir == "" {
		dir = filepath.Join("..", "gutenberg", "Historical Books")
	}

	outputPath := c.String("output")
	format := c.String("format")
	skipOnline := c.Bool("skip-online")

	absDir, err := filepath.Abs(dir)
	if err != nil {
		return fmt.Errorf("resolving path: %w", err)
	}

	entries, err := os.ReadDir(absDir)
	if err != nil {
		return fmt.Errorf("reading directory %s: %w", absDir, err)
	}

	var pdfs []string
	for _, e := range entries {
		if e.IsDir() {
			continue
		}
		name := e.Name()
		lower := strings.ToLower(name)
		if !strings.HasSuffix(lower, ".pdf") {
			continue
		}
		pdfs = append(pdfs, name)
	}

	var books []BookEntry
	for i, name := range pdfs {
		pdfPath := filepath.Join(absDir, name)
		fmt.Fprintf(os.Stderr, "[%d/%d] %s\n", i+1, len(pdfs), name)
		entry := processBook(pdfPath, name, skipOnline)
		books = append(books, entry)
	}

	sort.Slice(books, func(i, j int) bool {
		di := difficultyRank(books[i].Difficulty)
		dj := difficultyRank(books[j].Difficulty)
		if di != dj {
			return di < dj
		}
		return books[i].File < books[j].File
	})

	inv := Inventory{
		GeneratedAt: time.Now().Format(time.RFC3339),
		SourceDir:   absDir,
		TotalBooks:  len(books),
		Books:       books,
	}

	var out []byte
	if format == "json" {
		out, err = json.MarshalIndent(inv, "", "  ")
	} else {
		out, err = yaml.Marshal(inv)
	}
	if err != nil {
		return fmt.Errorf("marshaling output: %w", err)
	}

	if outputPath != "" {
		if err := os.WriteFile(outputPath, out, 0644); err != nil {
			return fmt.Errorf("writing output: %w", err)
		}
		fmt.Fprintf(os.Stderr, "Wrote %d books to %s\n", len(books), outputPath)
	} else {
		fmt.Println(string(out))
	}

	return nil
}

func processBook(pdfPath, filename string, skipOnline bool) BookEntry {
	entry := BookEntry{File: filename}

	// Validate naming
	matches := reNamePattern.FindStringSubmatch(filename)
	if matches == nil {
		entry.NameValid = false
		entry.NameIssue = "does not match YYYY - Author - Title.pdf"
		parts := strings.SplitN(strings.TrimSuffix(filename, filepath.Ext(filename)), " - ", 3)
		if len(parts) >= 1 {
			entry.Year = strings.TrimSpace(parts[0])
		}
		if len(parts) >= 2 {
			entry.Title = strings.TrimSpace(parts[1])
		}
		if len(parts) >= 3 {
			entry.Author = strings.TrimSpace(parts[1])
			entry.Title = strings.TrimSpace(parts[2])
		}
	} else {
		entry.NameValid = true
		entry.Year = matches[1]
		entry.Author = matches[2]
		entry.Title = strings.TrimSuffix(matches[3], "."+matches[4])
	}

	// Page count
	entry.Pages = countPages(pdfPath)

	// Embedded text check
	entry.EmbeddedText = checkEmbeddedText(pdfPath, entry.Pages)

	// Online source lookup
	if !skipOnline {
		fmt.Fprintf(os.Stderr, "  searching Internet Archive...\n")
		entry.ArchiveURL = searchInternetArchive(entry.Title, entry.Author)
		if entry.ArchiveURL != "" {
			fmt.Fprintf(os.Stderr, "  found: %s\n", entry.ArchiveURL)
		} else {
			fmt.Fprintf(os.Stderr, "  searching Gutenberg...\n")
			entry.GutenbergURL = searchGutenberg(entry.Title, entry.Author)
			if entry.GutenbergURL != "" {
				fmt.Fprintf(os.Stderr, "  found: %s\n", entry.GutenbergURL)
			}
		}
	}

	// Difficulty scoring
	entry.Difficulty = scoreDifficulty(entry)

	return entry
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
			_, _ = fmt.Sscanf(strings.TrimPrefix(line, "Pages:"), "%d", &pages)
			return pages
		}
	}
	return 0
}

func checkEmbeddedText(pdfPath string, totalPages int) string {
	if totalPages == 0 {
		return "none"
	}

	// Sample pages at various points through the document
	samplePages := []int{1}
	if totalPages > 10 {
		samplePages = append(samplePages, totalPages/10)
	}
	if totalPages > 4 {
		samplePages = append(samplePages, totalPages/4)
	}
	if totalPages > 2 {
		samplePages = append(samplePages, totalPages/2)
	}
	if totalPages > 4 {
		samplePages = append(samplePages, totalPages*3/4)
	}
	if totalPages > 10 {
		samplePages = append(samplePages, totalPages*9/10)
	}

	goodPages := 0
	for _, p := range samplePages {
		if p < 1 {
			p = 1
		}
		if p > totalPages {
			p = totalPages
		}
		text := extractPage(pdfPath, p)
		if isCoherentText(text) {
			goodPages++
		}
	}

	ratio := float64(goodPages) / float64(len(samplePages))
	if ratio >= 0.7 {
		return "high"
	}
	if ratio >= 0.3 {
		return "medium"
	}
	if goodPages > 0 {
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
	// Strip whitespace and check if there's meaningful content
	cleaned := strings.TrimSpace(text)
	if len(cleaned) < 50 {
		return false
	}

	// Count alphabetic characters vs total
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

	// Text should be mostly letters (at least 60%)
	ratio := float64(alphaCount) / float64(totalCount)
	if ratio < 0.6 {
		return false
	}

	// Should contain real words (spaces between groups of letters)
	words := strings.Fields(cleaned)
	return len(words) >= 10
}

func searchGutenberg(title, author string) string {
	_ = author // Currently not using author for search, but could be added to improve accuracy
	query := sanitizeSearchTerm(title)
	searchURL := fmt.Sprintf("https://gutendex.com/books/?search=%s", url.QueryEscape(query))

	resp, err := httpGetJSON(searchURL)
	if err != nil {
		return ""
	}

	results, ok := resp["results"].([]interface{})
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

func searchInternetArchive(title, author string) string {
	_ = author // Currently not using author for search, but could be added to improve accuracy
	query := sanitizeSearchTerm(title)
	searchURL := fmt.Sprintf(
		"https://archive.org/advancedsearch.php?q=title%%3A(%s)%%20AND%%20mediatype%%3Atexts&output=json&rows=5&fl[]=identifier,title",
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
			identifier, _ := doc["identifier"].(string)
			if identifier != "" {
				return fmt.Sprintf("https://archive.org/details/%s", identifier)
			}
		}
	}

	// If no fuzzy match, return the first result as a likely candidate
	if len(docs) > 0 {
		doc, ok := docs[0].(map[string]interface{})
		if ok {
			identifier, _ := doc["identifier"].(string)
			if identifier != "" {
				return fmt.Sprintf("https://archive.org/details/%s", identifier)
			}
		}
	}

	return ""
}

func httpGetJSON(rawURL string) (map[string]interface{}, error) {
	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Get(rawURL)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("HTTP %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result map[string]interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func sanitizeSearchTerm(s string) string {
	// Remove common articles and noise words for better search
	s = strings.TrimSpace(s)
	for _, prefix := range []string{"The ", "A ", "An "} {
		if strings.HasPrefix(s, prefix) {
			s = strings.TrimPrefix(s, prefix)
			break
		}
	}
	return s
}

func fuzzyMatch(a, b string) bool {
	a = strings.ToLower(strings.TrimSpace(a))
	b = strings.ToLower(strings.TrimSpace(b))
	if a == b {
		return true
	}
	// Check if one contains the significant part of the other
	aSig := significantWords(a)
	bSig := significantWords(b)

	matchCount := 0
	for _, w := range aSig {
		for _, w2 := range bSig {
			if w == w2 {
				matchCount++
				break
			}
		}
	}

	// At least 50% of significant words match
	minLen := len(aSig)
	if len(bSig) < minLen {
		minLen = len(bSig)
	}
	if minLen == 0 {
		return false
	}
	return float64(matchCount)/float64(minLen) >= 0.5
}

func significantWords(s string) []string {
	noise := map[string]bool{
		"the": true, "a": true, "an": true, "of": true, "in": true,
		"and": true, "or": true, "to": true, "for": true, "with": true,
		"by": true, "from": true, "its": true, "being": true,
	}
	var words []string
	for _, w := range strings.Fields(s) {
		w = strings.Trim(w, ".,;:!?\"'()-")
		if len(w) > 2 && !noise[w] {
			words = append(words, w)
		}
	}
	return words
}

func scoreDifficulty(e BookEntry) string {
	if e.ArchiveURL != "" || e.GutenbergURL != "" {
		return "easy"
	}
	if e.EmbeddedText == "high" {
		return "easy"
	}
	if e.EmbeddedText == "medium" {
		return "medium"
	}
	return "hard"
}

func difficultyRank(d string) int {
	switch d {
	case "easy":
		return 0
	case "medium":
		return 1
	case "hard":
		return 2
	}
	return 3
}
