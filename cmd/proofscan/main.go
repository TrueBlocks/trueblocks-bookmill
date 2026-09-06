// Command proofscan is the no-spend, exploratory first pass of the research
// corpus text-quality audit (issue #727). It reads every extracted .txt file
// beside its source document (.pdf/.html) and flags the ones whose conversion
// obviously failed or looks garbled — empty files, placeholder error text,
// Unicode replacement characters, or a low ratio of readable characters. It
// calls no model and costs nothing; its job is to tell us where the paid,
// vision-based `proof` pass should later aim.
package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"unicode"
	"unicode/utf8"

	"github.com/TrueBlocks/trueblocks-art/packages/cli"
)

var version = "dev"

func main() {
	app := cli.App{
		Name:        "proofscan",
		Description: "No-spend audit of the research corpus: flag broken or garbled PDF/HTML-to-text conversions using cheap heuristics.",
		Version:     version,
		ArgsUsage:   "[corpus-dir]",
		Flags: []cli.FlagDef{
			{Name: "min-bytes", Help: "flag non-empty .txt smaller than this as suspect", Default: 500},
			{Name: "limit", Help: "max offenders to list per verdict (0 = all)", Default: 40},
			{Name: "only-bad", Help: "list only FAILED and SUSPECT files, skip the OK summary detail", Default: false},
		},
		Run: run,
	}
	cli.Exit(app.Main())
}

// verdict classifies one source/text pair.
type verdict int

const (
	vOK verdict = iota
	vSuspect
	vFailed
)

func (v verdict) String() string {
	switch v {
	case vFailed:
		return "FAILED"
	case vSuspect:
		return "SUSPECT"
	default:
		return "OK"
	}
}

// sourceExts are the document types whose extracted text we audit.
var sourceExts = map[string]bool{".pdf": true, ".html": true, ".htm": true}

// placeholderPhrases are exact-ish strings that mean the conversion produced
// no real content — a scanner or database placeholder rather than the document.
var placeholderPhrases = []string{
	"does not exist within the ldap database",
	"has not yet been digitized",
	"this file has not been digitized",
	"no text could be extracted",
	"page not found",
}

type finding struct {
	source  string // path to the .pdf/.html, relative to root
	shelf   string // top-level dir under root (the "shelf")
	txt     string // path to the .txt, relative to root ("" if missing)
	size    int64
	verdict verdict
	reasons []string
}

func run(c *cli.Context) error {
	root := "."
	if len(c.Args) > 0 {
		root = c.Args[0]
	}
	absRoot, err := filepath.Abs(root)
	if err != nil {
		return fmt.Errorf("resolving corpus dir: %w", err)
	}
	info, err := os.Stat(absRoot)
	if err != nil || !info.IsDir() {
		return fmt.Errorf("corpus dir not found: %s", absRoot)
	}

	minBytes := int64(c.Int("min-bytes"))
	limit := c.Int("limit")
	onlyBad := c.Bool("only-bad")

	var findings []finding
	err = filepath.WalkDir(absRoot, func(path string, d os.DirEntry, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}
		if d.IsDir() {
			return nil
		}
		ext := strings.ToLower(filepath.Ext(path))
		if !sourceExts[ext] {
			return nil
		}
		findings = append(findings, audit(absRoot, path, minBytes))
		return nil
	})
	if err != nil {
		return fmt.Errorf("walking corpus: %w", err)
	}

	if len(findings) == 0 {
		return fmt.Errorf("no source documents (.pdf/.html) found under %s", absRoot)
	}

	report(c.App.Stdout, absRoot, findings, limit, onlyBad)
	return nil
}

// audit computes a verdict for one source document by reading its sibling .txt.
func audit(root, sourceAbs string, minBytes int64) finding {
	rel, _ := filepath.Rel(root, sourceAbs)
	f := finding{
		source: rel,
		shelf:  shelfOf(rel),
	}

	txtAbs := strings.TrimSuffix(sourceAbs, filepath.Ext(sourceAbs)) + ".txt"
	txtInfo, err := os.Stat(txtAbs)
	if err != nil {
		f.verdict = vFailed
		f.reasons = append(f.reasons, "no .txt produced")
		return f
	}
	f.txt, _ = filepath.Rel(root, txtAbs)
	f.size = txtInfo.Size()

	if f.size == 0 {
		f.verdict = vFailed
		f.reasons = append(f.reasons, "empty .txt")
		return f
	}

	data, err := os.ReadFile(txtAbs)
	if err != nil {
		f.verdict = vSuspect
		f.reasons = append(f.reasons, "unreadable: "+err.Error())
		return f
	}

	f.verdict, f.reasons = classify(string(data), f.size, minBytes)
	return f
}

// classify runs the heuristics over the text and returns the worst verdict
// with the reasons that fired.
func classify(text string, size, minBytes int64) (verdict, []string) {
	var reasons []string
	worst := vOK
	bump := func(v verdict, reason string) {
		reasons = append(reasons, reason)
		if v > worst {
			worst = v
		}
	}

	lower := strings.ToLower(strings.TrimSpace(text))

	// Placeholder failure text (often short) — a certain failure.
	for _, p := range placeholderPhrases {
		if strings.Contains(lower, p) {
			bump(vFailed, "placeholder text: "+p)
			break
		}
	}

	// Character-quality pass.
	var letters, digits, spaces, other, replacement, total int
	for _, r := range text {
		total++
		switch {
		case r == unicode.ReplacementChar:
			replacement++
			other++
		case unicode.IsLetter(r):
			letters++
		case unicode.IsDigit(r):
			digits++
		case unicode.IsSpace(r):
			spaces++
		case unicode.IsPunct(r) || unicode.IsSymbol(r):
			// readable enough
		case !unicode.IsGraphic(r):
			other++
		}
	}
	invalidUTF8 := !utf8.ValidString(text)

	if replacement > 0 && total > 0 {
		ratio := float64(replacement) / float64(total)
		switch {
		case ratio > 0.02:
			bump(vFailed, fmt.Sprintf("%.1f%% replacement chars", ratio*100))
		default:
			bump(vSuspect, fmt.Sprintf("%d replacement chars", replacement))
		}
	}
	if invalidUTF8 && replacement == 0 {
		bump(vSuspect, "invalid UTF-8")
	}

	// Readable-character ratio (letters+digits+spaces vs everything).
	if total > 0 {
		readable := float64(letters+digits+spaces) / float64(total)
		switch {
		case readable < 0.50:
			bump(vFailed, fmt.Sprintf("%.0f%% readable chars", readable*100))
		case readable < 0.75:
			bump(vSuspect, fmt.Sprintf("%.0f%% readable chars", readable*100))
		}
	}

	// Suspiciously small — only if nothing worse already fired.
	if worst == vOK && size < minBytes {
		bump(vSuspect, fmt.Sprintf("only %d bytes", size))
	}

	return worst, reasons
}

// shelfOf returns the top-level directory under the corpus root (the "shelf"),
// or "(root)" for a document sitting directly in the corpus dir.
func shelfOf(rel string) string {
	parts := strings.SplitN(rel, string(filepath.Separator), 2)
	if len(parts) < 2 {
		return "(root)"
	}
	return parts[0]
}

func report(w io.Writer, root string, findings []finding, limit int, onlyBad bool) {
	// Per-shelf tallies.
	type tally struct {
		ok, suspect, failed, total int
	}
	shelves := map[string]*tally{}
	var order []string
	var totals tally
	for _, f := range findings {
		t, ok := shelves[f.shelf]
		if !ok {
			t = &tally{}
			shelves[f.shelf] = t
			order = append(order, f.shelf)
		}
		t.total++
		totals.total++
		switch f.verdict {
		case vFailed:
			t.failed++
			totals.failed++
		case vSuspect:
			t.suspect++
			totals.suspect++
		default:
			t.ok++
			totals.ok++
		}
	}
	sort.Strings(order)

	var b strings.Builder
	fmt.Fprintf(&b, "# Research corpus text-quality scan\n\n")
	fmt.Fprintf(&b, "Corpus: %s\n", root)
	fmt.Fprintf(&b, "Documents scanned: %d\n\n", totals.total)

	fmt.Fprintf(&b, "## Summary by shelf\n\n")
	fmt.Fprintf(&b, "%-40s  %6s  %7s  %6s  %5s\n", "Shelf", "OK", "SUSPECT", "FAILED", "Total")
	fmt.Fprintf(&b, "%s\n", strings.Repeat("-", 72))
	for _, s := range order {
		t := shelves[s]
		fmt.Fprintf(&b, "%-40s  %6d  %7d  %6d  %5d\n", truncate(s, 40), t.ok, t.suspect, t.failed, t.total)
	}
	fmt.Fprintf(&b, "%s\n", strings.Repeat("-", 72))
	fmt.Fprintf(&b, "%-40s  %6d  %7d  %6d  %5d\n\n", "TOTAL", totals.ok, totals.suspect, totals.failed, totals.total)

	listVerdict(&b, findings, vFailed, limit)
	listVerdict(&b, findings, vSuspect, limit)

	if !onlyBad && totals.failed == 0 && totals.suspect == 0 {
		fmt.Fprintf(&b, "No conversions flagged. The corpus looks clean by these heuristics.\n")
	}

	fmt.Fprintf(&b, "\nNote: heuristics only, no model was called. FAILED = almost certainly bad; "+
		"SUSPECT = worth a human or vision-`proof` look.\n")

	_, _ = io.WriteString(w, b.String())
}

func listVerdict(b *strings.Builder, findings []finding, v verdict, limit int) {
	var matched []finding
	for _, f := range findings {
		if f.verdict == v {
			matched = append(matched, f)
		}
	}
	if len(matched) == 0 {
		return
	}
	sort.Slice(matched, func(i, j int) bool {
		if matched[i].shelf != matched[j].shelf {
			return matched[i].shelf < matched[j].shelf
		}
		return matched[i].source < matched[j].source
	})

	fmt.Fprintf(b, "## %s (%d)\n\n", v, len(matched))
	shown := matched
	if limit > 0 && len(matched) > limit {
		shown = matched[:limit]
	}
	for _, f := range shown {
		fmt.Fprintf(b, "- %s\n  %s\n", f.source, strings.Join(f.reasons, "; "))
	}
	if len(shown) < len(matched) {
		fmt.Fprintf(b, "  ... and %d more (raise --limit to see all)\n", len(matched)-len(shown))
	}
	fmt.Fprintln(b)
}

func truncate(s string, n int) string {
	if len(s) <= n {
		return s
	}
	return s[:n-1] + "…"
}
