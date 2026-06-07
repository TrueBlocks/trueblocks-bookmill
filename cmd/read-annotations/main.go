package main

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/TrueBlocks/trueblocks-art/packages/cli"
	"github.com/pdfcpu/pdfcpu/pkg/api"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
)

var version = "dev"

type PageAnnotation struct {
	Page int
	Type string
	Text string
}

func main() {
	app := cli.App{
		Name:        "read-annotations",
		Description: "Read text annotations from a PDF and report rotate, no_sky, and chapter title markers by page.",
		Version:     version,
		Flags: []cli.FlagDef{
			{Name: "pdf", Help: "path to the annotated PDF file", Default: ""},
			{Name: "format", Help: "output format: text or yaml (default: text)", Default: "text"},
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

	f, err := os.Open(absPDF)
	if err != nil {
		return fmt.Errorf("opening pdf: %w", err)
	}
	defer f.Close()

	conf := model.NewDefaultConfiguration()
	annots, err := api.Annotations(f, nil, conf)
	if err != nil {
		return fmt.Errorf("reading annotations: %w", err)
	}

	var results []PageAnnotation
	for page, pgAnnots := range annots {
		for _, annot := range pgAnnots {
			for _, renderer := range annot.Map {
				text := strings.TrimSpace(renderer.ContentString())
				if text == "" {
					continue
				}
				kind := classify(text)
				results = append(results, PageAnnotation{
					Page: page,
					Type: kind,
					Text: text,
				})
			}
		}
	}

	sort.Slice(results, func(i, j int) bool {
		if results[i].Page != results[j].Page {
			return results[i].Page < results[j].Page
		}
		return results[i].Type < results[j].Type
	})

	format := c.String("format")
	switch format {
	case "yaml":
		printYAML(results)
	default:
		printText(results)
	}

	return nil
}

func classify(text string) string {
	lower := strings.ToLower(strings.TrimSpace(text))
	if lower == "rotate" {
		return "rotate"
	}
	if strings.HasPrefix(lower, "no sky") || strings.HasPrefix(lower, "no_sky") {
		return "no_sky"
	}
	return "chapter"
}

func printText(results []PageAnnotation) {
	for _, r := range results {
		fmt.Printf("page %3d  %-8s  %s\n", r.Page, r.Type, r.Text)
	}
	fmt.Fprintf(os.Stderr, "\n%d annotation(s) found\n", len(results))
}

func printYAML(results []PageAnnotation) {
	var rotates []PageAnnotation
	var noSkys []PageAnnotation
	var chapters []PageAnnotation
	for _, r := range results {
		switch r.Type {
		case "rotate":
			rotates = append(rotates, r)
		case "no_sky":
			noSkys = append(noSkys, r)
		case "chapter":
			chapters = append(chapters, r)
		}
	}

	if len(rotates) > 0 {
		fmt.Println("rotations:")
		for _, r := range rotates {
			fmt.Printf("  - page: %d\n", r.Page)
		}
	}
	if len(noSkys) > 0 {
		fmt.Println("no_sky:")
		for _, r := range noSkys {
			fmt.Printf("  - page: %d\n    text: %q\n", r.Page, r.Text)
		}
	}
	if len(chapters) > 0 {
		fmt.Println("chapters:")
		for _, r := range chapters {
			fmt.Printf("  - page: %d\n    title: %q\n", r.Page, r.Text)
		}
	}

	fmt.Fprintf(os.Stderr, "\n%d annotation(s) found\n", len(results))
}
