package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/TrueBlocks/trueblocks-art/packages/cli"
)

var version = "dev"

func main() {
	app := cli.App{
		Name:        "extract-text",
		Description: "Extract text from a historical PDF into markdown with page markers. Uses pdftotext for PDFs with embedded text, or downloads from Internet Archive as fallback.",
		Version:     version,
		Flags: []cli.FlagDef{
			{Name: "pdf", Help: "path to the source PDF file", Default: ""},
			{Name: "output", Help: "output markdown file path (default: stdout)", Default: ""},
			{Name: "source", Help: "text source: pdf, archive, or auto (default: auto)", Default: "auto"},
			{Name: "archive-id", Help: "Internet Archive identifier (e.g. asylvancityorqu00campgoog)", Default: ""},
			{Name: "min-chars", Help: "minimum characters for a page to be kept (default: 40)", Default: 40},
		},
		Run: run,
	}
	cli.Exit(app.Main())
}

func run(c *cli.Context) error {
	pdfPath := c.String("pdf")
	outputPath := c.String("output")
	source := c.String("source")
	archiveID := c.String("archive-id")
	minChars := c.Int("min-chars")

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

	var rawText string

	switch source {
	case "pdf":
		rawText, err = extractFromPDF(absPDF)
	case "archive":
		if archiveID == "" {
			return fmt.Errorf("--archive-id is required when --source=archive")
		}
		rawText, err = extractFromArchive(archiveID)
	case "auto":
		rawText, err = extractAuto(absPDF, archiveID)
	default:
		return fmt.Errorf("unknown source: %s (use pdf, archive, or auto)", source)
	}
	if err != nil {
		return err
	}

	markdown := convertToMarkdown(rawText, source != "pdf" && source != "auto", minChars)

	if outputPath != "" {
		dir := filepath.Dir(outputPath)
		if err := os.MkdirAll(dir, 0755); err != nil {
			return fmt.Errorf("creating output directory: %w", err)
		}
		if err := os.WriteFile(outputPath, []byte(markdown), 0644); err != nil {
			return fmt.Errorf("writing output: %w", err)
		}
		pages := strings.Count(markdown, "<!-- page ")
		fmt.Fprintf(os.Stderr, "Wrote %d pages to %s\n", pages, outputPath)
	} else {
		fmt.Print(markdown)
	}

	return nil
}

func extractFromPDF(pdfPath string) (string, error) {
	fmt.Fprintf(os.Stderr, "Extracting text from PDF: %s\n", filepath.Base(pdfPath))
	cmd := exec.Command("pdftotext", "-layout", pdfPath, "-")
	out, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("pdftotext failed: %w", err)
	}
	return string(out), nil
}

func extractFromArchive(archiveID string) (string, error) {
	textURL := fmt.Sprintf("https://archive.org/download/%s/%s_djvu.txt", archiveID, archiveID)
	fmt.Fprintf(os.Stderr, "Downloading text from Internet Archive: %s\n", textURL)

	resp, err := http.Get(textURL)
	if err != nil {
		return "", fmt.Errorf("downloading from archive: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("archive returned status %d for %s", resp.StatusCode, textURL)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("reading archive response: %w", err)
	}

	return string(body), nil
}

func extractAuto(pdfPath, archiveID string) (string, error) {
	quality := checkTextQuality(pdfPath)
	fmt.Fprintf(os.Stderr, "PDF text quality: %s\n", quality)

	if quality == "high" || quality == "medium" {
		return extractFromPDF(pdfPath)
	}

	if archiveID != "" {
		fmt.Fprintf(os.Stderr, "PDF text quality is %s, trying Internet Archive...\n", quality)
		text, err := extractFromArchive(archiveID)
		if err == nil && len(strings.TrimSpace(text)) > 100 {
			return text, nil
		}
		fmt.Fprintf(os.Stderr, "Archive text not usable, falling back to PDF\n")
	}

	return extractFromPDF(pdfPath)
}

func checkTextQuality(pdfPath string) string {
	pages := countPages(pdfPath)
	if pages == 0 {
		return "none"
	}

	samplePoints := []int{1}
	if pages > 10 {
		samplePoints = append(samplePoints, pages/10)
	}
	if pages > 4 {
		samplePoints = append(samplePoints, pages/4)
	}
	if pages > 2 {
		samplePoints = append(samplePoints, pages/2)
	}
	if pages > 4 {
		samplePoints = append(samplePoints, pages*3/4)
	}

	good := 0
	for _, p := range samplePoints {
		cmd := exec.Command("pdftotext", "-f", fmt.Sprintf("%d", p), "-l", fmt.Sprintf("%d", p), pdfPath, "-")
		out, err := cmd.Output()
		if err != nil {
			continue
		}
		text := strings.TrimSpace(string(out))
		if len(text) > 50 && letterRatio(text) > 0.6 {
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

func letterRatio(text string) float64 {
	letters := 0
	nonSpace := 0
	for _, r := range text {
		if r != ' ' && r != '\t' && r != '\n' && r != '\r' {
			nonSpace++
			if (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') {
				letters++
			}
		}
	}
	if nonSpace == 0 {
		return 0
	}
	return float64(letters) / float64(nonSpace)
}

// reExtraSpaces matches 2+ spaces within a line
var reExtraSpaces = regexp.MustCompile(`  +`)

// reBlankRuns matches 3+ consecutive blank lines
// var reBlankRuns = regexp.MustCompile(`\n{4,}`)

func convertToMarkdown(rawText string, isArchiveText bool, minChars int) string {
	if isArchiveText {
		return convertArchiveToMarkdown(rawText)
	}
	return convertPDFToMarkdown(rawText, minChars)
}

func convertPDFToMarkdown(rawText string, minChars int) string {
	pages := strings.Split(rawText, "\f")

	type cleanedPage struct {
		pageNum int
		text    string
	}

	var cleaned []cleanedPage
	for i, page := range pages {
		pageNum := i + 1
		text := cleanPage(page)
		if len(strings.TrimSpace(text)) < minChars {
			continue
		}
		cleaned = append(cleaned, cleanedPage{pageNum: pageNum, text: text})
	}

	// Join cross-page paragraphs: if page N ends mid-sentence and page N+1
	// starts with continuation text (not a chapter heading), merge them
	for i := 0; i < len(cleaned)-1; i++ {
		cur := cleaned[i].text
		next := cleaned[i+1].text

		// Don't merge into a chapter heading
		if strings.HasPrefix(next, "## ") {
			continue
		}

		if endsIncompletely(cur) && startsContinuation(next) {
			// Remove trailing newline from current, merge first paragraph of next
			nextLines := strings.SplitN(next, "\n\n", 2)
			firstPara := nextLines[0]

			// Join with space (or rejoin hyphen)
			if strings.HasSuffix(cur, "-") {
				cleaned[i].text = cur[:len(cur)-1] + firstPara
			} else {
				cleaned[i].text = cur + " " + firstPara
			}

			// Keep the rest of next page
			if len(nextLines) > 1 {
				cleaned[i+1].text = nextLines[1]
			} else {
				cleaned[i+1].text = ""
			}
		}
	}

	var sb strings.Builder
	for _, p := range cleaned {
		text := strings.TrimSpace(p.text)
		if len(text) == 0 {
			continue
		}
		if sb.Len() > 0 {
			sb.WriteString("\n\n")
		}
		sb.WriteString(fmt.Sprintf("<!-- page %d -->\n\n", p.pageNum))
		sb.WriteString(text)
	}

	sb.WriteString("\n")
	return sb.String()
}

func endsIncompletely(text string) bool {
	text = strings.TrimRight(text, " \t\n")
	if len(text) == 0 {
		return false
	}
	last := text[len(text)-1]
	// Ends with sentence-ending punctuation → complete
	if last == '.' || last == '!' || last == '?' || last == ':' || last == '"' {
		return false
	}
	return true
}

func startsContinuation(text string) bool {
	text = strings.TrimLeft(text, " \t\n")
	if len(text) == 0 {
		return false
	}
	// Starts with lowercase letter → continuation
	first := rune(text[0])
	if first >= 'a' && first <= 'z' {
		return true
	}
	// Starts with certain punctuation that continues a sentence
	if first == ',' || first == ';' || first == '-' {
		return true
	}
	return false
}

func convertArchiveToMarkdown(rawText string) string {
	lines := strings.Split(rawText, "\n")

	skipPrefixes := []string{
		"Google",
		"This  is  a  digital  copy",
		"It  has  survived",
		"Marks,  notations",
		"Usage  guidelines",
		"Google  is  proud",
		"+  Make  non-commercial",
		"+  Refrain",
		"+  Maintain  attribution",
		"+  Keep  it  legal",
		"About  Google  Book  Search",
		"Google's  mission",
	}

	var sb strings.Builder
	skippingHeader := true
	pageNum := 1
	sb.WriteString(fmt.Sprintf("<!-- page %d -->\n\n", pageNum))

	blankCount := 0
	for _, line := range lines {
		trimmed := strings.TrimSpace(line)

		if skippingHeader {
			isBoilerplate := false
			for _, prefix := range skipPrefixes {
				if strings.HasPrefix(trimmed, prefix) {
					isBoilerplate = true
					break
				}
			}
			if isBoilerplate || len(trimmed) == 0 {
				continue
			}
			if strings.Contains(trimmed, "http") && strings.Contains(trimmed, "google") {
				continue
			}
			skippingHeader = false
		}

		cleaned := reExtraSpaces.ReplaceAllString(line, " ")
		cleaned = strings.TrimRight(cleaned, " \t")

		if len(strings.TrimSpace(cleaned)) == 0 {
			blankCount++
			if blankCount >= 3 {
				pageNum++
				sb.WriteString(fmt.Sprintf("\n\n<!-- page %d -->\n\n", pageNum))
				blankCount = 0
			} else {
				sb.WriteString("\n")
			}
		} else {
			blankCount = 0
			sb.WriteString(cleaned)
			sb.WriteString("\n")
		}
	}

	sb.WriteString("\n")
	return sb.String()
}

// rePageNumber matches a line that is just a page number (with possible OCR noise)
var rePageNumber = regexp.MustCompile(`^\d[\d\s\w]{0,4}$`)

// reAnnotationText matches lines that are PDF annotation text bleeding through
var reAnnotationText = regexp.MustCompile(`(?i)^(no sky\s*\d*|rotate)$`)

func cleanPage(page string) string {
	lines := strings.Split(page, "\n")

	// First pass: collapse extra spaces, detect indented lines as paragraph starts
	var trimmed []string
	for _, line := range lines {
		raw := line
		line = reExtraSpaces.ReplaceAllString(line, " ")
		line = strings.TrimRight(line, " \t")

		indent := 0
		for _, ch := range raw {
			if ch == ' ' {
				indent++
			} else if ch == '\t' {
				indent += 4
			} else {
				break
			}
		}
		cleaned := strings.TrimLeft(line, " \t")

		if indent >= 3 && cleaned != "" && len(trimmed) > 0 && trimmed[len(trimmed)-1] != "" {
			trimmed = append(trimmed, "")
		}
		trimmed = append(trimmed, cleaned)
	}

	// Strip leading blank lines
	for len(trimmed) > 0 && trimmed[0] == "" {
		trimmed = trimmed[1:]
	}
	// Strip trailing blank lines
	for len(trimmed) > 0 && trimmed[len(trimmed)-1] == "" {
		trimmed = trimmed[:len(trimmed)-1]
	}

	if len(trimmed) == 0 {
		return ""
	}

	// Detect chapter start: first line is short all-caps title,
	// NOT preceded by a page number (running headers have page number first)
	chapterTitle := ""
	if isChapterStart(trimmed) {
		chapterTitle = toTitleCase(trimmed[0])
		trimmed = trimmed[1:]
		// Strip any blank lines after the chapter title
		for len(trimmed) > 0 && trimmed[0] == "" {
			trimmed = trimmed[1:]
		}
	} else {
		// Regular page: strip running headers
		trimmed = stripHeaders(trimmed)
	}

	// Strip footers
	trimmed = stripFooters(trimmed)

	// Group lines into paragraphs (blank line = paragraph break)
	// and rejoin hyphenated words at line boundaries
	var paragraphs []string
	var current []string

	for _, line := range trimmed {
		if line == "" {
			if len(current) > 0 {
				paragraphs = append(paragraphs, joinLines(current))
				current = nil
			}
		} else if reAnnotationText.MatchString(line) {
			continue
		} else {
			current = append(current, line)
		}
	}
	if len(current) > 0 {
		paragraphs = append(paragraphs, joinLines(current))
	}

	var sb strings.Builder
	if chapterTitle != "" {
		sb.WriteString("## ")
		sb.WriteString(chapterTitle)
		sb.WriteString("\n\n")
	}
	sb.WriteString(strings.Join(paragraphs, "\n\n"))

	result := sb.String()
	result = strings.TrimSpace(result)
	return result
}

func isChapterStart(lines []string) bool {
	if len(lines) < 3 {
		return false
	}
	first := lines[0]
	// Chapter start: first line is short, all-caps, NOT a page number
	if rePageNumber.MatchString(first) {
		return false
	}
	if len(first) < 5 || len(first) > 60 {
		return false
	}
	if !isUpperCase(first) {
		return false
	}
	// Must not have a trailing page number (running headers do: "A QUAKER SOLDIER. 31")
	if hasTrailingNumber(first) {
		return false
	}
	// Running headers with leading page numbers (like "40 A SYLVAN CITY") start with digits
	if len(first) > 0 && first[0] >= '0' && first[0] <= '9' {
		return false
	}
	// Filter out book-title running headers ("A SYLVAN CITY" and OCR variants)
	lower := strings.ToLower(first)
	if strings.Contains(lower, "sylvan") {
		return false
	}
	// Filter out OCR-garbled lines with stray digits mixed in
	digitCount := 0
	for _, c := range first {
		if c >= '0' && c <= '9' {
			digitCount++
		}
	}
	if digitCount > 0 {
		return false
	}
	// The next non-blank line should be body text (lowercase-dominant, longish)
	for i := 1; i < len(lines); i++ {
		if lines[i] == "" {
			continue
		}
		// If next non-blank line is a page number, this is a running header pair
		if rePageNumber.MatchString(lines[i]) {
			return false
		}
		// Body text should have lowercase letters and be reasonably long
		if !isUpperCase(lines[i]) && len(lines[i]) > 30 {
			return true
		}
		return false
	}
	return false
}

func toTitleCase(s string) string {
	words := strings.Fields(strings.ToLower(s))
	skip := map[string]bool{"a": true, "an": true, "the": true, "of": true, "in": true, "and": true, "or": true, "to": true, "for": true}
	for i, w := range words {
		w = strings.Trim(w, ".,;:!?\"'")
		if i == 0 || !skip[w] {
			runes := []rune(words[i])
			if len(runes) > 0 {
				runes[0] = []rune(strings.ToUpper(string(runes[0])))[0]
				words[i] = string(runes)
			}
		}
	}
	return strings.Join(words, " ")
}

func isHeaderLine(line string) bool {
	if line == "" {
		return true
	}
	// Page numbers: "40", "91", "90o" (OCR garble)
	if rePageNumber.MatchString(line) {
		return true
	}
	// Short all-caps lines like "A SYLVAN CITY." or "A QUAKER SOLDIER."
	if len(line) < 50 && isUpperCase(line) {
		return true
	}
	// Short lines with page numbers embedded: "THE BETTERING-HO USE. - 413"
	if len(line) < 60 && hasTrailingNumber(line) {
		return true
	}
	return false
}

func isUpperCase(s string) bool {
	letters := 0
	upper := 0
	for _, r := range s {
		if r >= 'a' && r <= 'z' {
			letters++
		} else if r >= 'A' && r <= 'Z' {
			letters++
			upper++
		}
	}
	if letters < 3 {
		return false
	}
	return float64(upper)/float64(letters) > 0.7
}

func hasTrailingNumber(s string) bool {
	s = strings.TrimSpace(s)
	if len(s) == 0 {
		return false
	}
	// Check if the line ends with digits (possibly after punctuation/spaces)
	for i := len(s) - 1; i >= 0; i-- {
		c := s[i]
		if c >= '0' && c <= '9' {
			return true
		}
		if c == ' ' || c == '.' || c == '-' {
			continue
		}
		break
	}
	return false
}

func stripHeaders(lines []string) []string {
	stripped := 0
	for stripped < len(lines) && stripped < 4 {
		if !isHeaderLine(lines[stripped]) {
			break
		}
		stripped++
	}
	// Don't strip everything — keep at least some content
	if stripped >= len(lines) {
		return lines
	}
	return lines[stripped:]
}

func stripFooters(lines []string) []string {
	stripped := 0
	for stripped < len(lines) && stripped < 2 {
		last := lines[len(lines)-1-stripped]
		if last == "" || rePageNumber.MatchString(last) {
			stripped++
		} else {
			break
		}
	}
	if stripped == 0 {
		return lines
	}
	return lines[:len(lines)-stripped]
}

func joinLines(lines []string) string {
	if len(lines) == 0 {
		return ""
	}
	result := lines[0]
	for i := 1; i < len(lines); i++ {
		// Rejoin hyphenated words: "de-" + "spair" → "despair"
		if strings.HasSuffix(result, "-") {
			result = result[:len(result)-1] + lines[i]
		} else {
			result = result + " " + lines[i]
		}
	}
	return result
}
