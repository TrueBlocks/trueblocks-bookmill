package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"

	"github.com/TrueBlocks/trueblocks-art/packages/cli"
	"gopkg.in/yaml.v3"
)

var version = "dev"

type ImageEntry struct {
	File    string `yaml:"file"`
	Page    int    `yaml:"page"`
	Seq     int    `yaml:"seq"`
	Width   int    `yaml:"width"`
	Height  int    `yaml:"height"`
	HadBlue bool   `yaml:"had_blue,omitempty"`
}

type Manifest struct {
	Book   string       `yaml:"book"`
	Images []ImageEntry `yaml:"images"`
}

var rePageMarker = regexp.MustCompile(`<!-- page (\d+) -->`)

func main() {
	app := cli.App{
		Name:        "compose",
		Description: "Compose final markdown by merging proofed text with image markers from the manifest. Inserts [[IMG:...]] tags at corresponding page locations and removes page markers.",
		Version:     version,
		Flags: []cli.FlagDef{
			{Name: "text", Help: "path to proofed markdown file", Default: ""},
			{Name: "manifest", Help: "path to image manifest.yaml (from colorize or extract-images)", Default: ""},
			{Name: "output", Help: "output composed markdown file (default: stdout)", Default: ""},
			{Name: "image-dir", Help: "subdirectory name for image references (default: images)", Default: "images"},
			{Name: "keep-page-markers", Help: "keep <!-- page N --> markers in output", Default: false},
		},
		Run: run,
	}
	cli.Exit(app.Main())
}

func run(c *cli.Context) error {
	textPath := c.String("text")
	manifestPath := c.String("manifest")
	outputPath := c.String("output")
	imageDir := c.String("image-dir")
	keepMarkers := c.Bool("keep-page-markers")

	if textPath == "" {
		return fmt.Errorf("--text is required")
	}

	absText, err := filepath.Abs(textPath)
	if err != nil {
		return fmt.Errorf("resolving text path: %w", err)
	}

	content, err := os.ReadFile(absText)
	if err != nil {
		return fmt.Errorf("reading text: %w", err)
	}

	imagesByPage := make(map[int][]ImageEntry)
	if manifestPath != "" {
		absManifest, err := filepath.Abs(manifestPath)
		if err != nil {
			return fmt.Errorf("resolving manifest path: %w", err)
		}
		manifestData, err := os.ReadFile(absManifest)
		if err != nil {
			return fmt.Errorf("reading manifest: %w", err)
		}
		var manifest Manifest
		if err := yaml.Unmarshal(manifestData, &manifest); err != nil {
			return fmt.Errorf("parsing manifest: %w", err)
		}

		for _, img := range manifest.Images {
			imagesByPage[img.Page] = append(imagesByPage[img.Page], img)
		}

		for page := range imagesByPage {
			sort.Slice(imagesByPage[page], func(i, j int) bool {
				return imagesByPage[page][i].Seq < imagesByPage[page][j].Seq
			})
		}

		fmt.Fprintf(os.Stderr, "Loaded %d images across %d pages\n", len(manifest.Images), len(imagesByPage))
	}

	text := string(content)
	pages := splitByPageMarkers(text)

	var sb strings.Builder
	for _, page := range pages {
		if keepMarkers {
			sb.WriteString(fmt.Sprintf("<!-- page %d -->\n\n", page.pageNum))
		}

		sb.WriteString(strings.TrimSpace(page.text))
		sb.WriteString("\n")

		if imgs, ok := imagesByPage[page.pageNum]; ok {
			for _, img := range imgs {
				imgPath := filepath.Join(imageDir, img.File)
				caption := fmt.Sprintf("Page %d", img.Page)
				sb.WriteString(fmt.Sprintf("\n[[IMG:%s|%s]]\n", imgPath, caption))
			}
		}

		sb.WriteString("\n")
	}

	result := sb.String()

	if outputPath != "" {
		dir := filepath.Dir(outputPath)
		if err := os.MkdirAll(dir, 0755); err != nil {
			return fmt.Errorf("creating output dir: %w", err)
		}
		if err := os.WriteFile(outputPath, []byte(result), 0644); err != nil {
			return fmt.Errorf("writing output: %w", err)
		}
		fmt.Fprintf(os.Stderr, "Wrote composed markdown to %s\n", outputPath)
	} else {
		fmt.Print(result)
	}

	return nil
}

type pageBlock struct {
	pageNum int
	text    string
}

func splitByPageMarkers(content string) []pageBlock {
	matches := rePageMarker.FindAllStringIndex(content, -1)
	if len(matches) == 0 {
		return []pageBlock{{pageNum: 1, text: content}}
	}

	var pages []pageBlock
	for i, loc := range matches {
		marker := content[loc[0]:loc[1]]
		var pageNum int
		fmt.Sscanf(marker, "<!-- page %d -->", &pageNum)

		var end int
		if i+1 < len(matches) {
			end = matches[i+1][0]
		} else {
			end = len(content)
		}

		text := content[loc[1]:end]
		text = strings.TrimPrefix(text, "\n")

		pages = append(pages, pageBlock{
			pageNum: pageNum,
			text:    text,
		})
	}

	return pages
}
