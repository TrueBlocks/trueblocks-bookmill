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
	NoSky   bool   `yaml:"no_sky,omitempty"`
	Rotated bool   `yaml:"rotated,omitempty"`
}

type ChapterEntry struct {
	Page  int    `yaml:"page"`
	Title string `yaml:"title"`
}

type Manifest struct {
	Imprint       string         `yaml:"imprint,omitempty"`
	Project       string         `yaml:"project,omitempty"`
	Book          string         `yaml:"book"`
	OutputDir     string         `yaml:"output_dir,omitempty"`
	SupportingDir string         `yaml:"supporting_dir,omitempty"`
	Images        []ImageEntry   `yaml:"images"`
	Chapters      []ChapterEntry `yaml:"chapters,omitempty"`
}

var rePageMarker = regexp.MustCompile(`<!-- page (\d+) -->`)
var reIMGTag = regexp.MustCompile(`\[\[IMG:([^\]|]+?)(?:\|[^\]]*?)?\]\]`)

func main() {
	app := cli.App{
		Name:        "compose",
		Description: "Compose final markdown by merging proofed text with image markers from the manifest. Splits into chapters using page boundaries from manifest.",
		Version:     version,
		Flags: []cli.FlagDef{
			{Name: "text", Help: "path to proofed markdown file", Default: ""},
			{Name: "manifest", Help: "path to manifest.yaml (required)", Default: ""},
			{Name: "output", Help: "output directory for chapter files (with --split) or single file", Default: ""},
			{Name: "image-dir", Help: "subdirectory name for image references (default: images)", Default: "images"},
			{Name: "source-images", Help: "directory containing extracted images to copy into output", Default: ""},
			{Name: "keep-page-markers", Help: "keep <!-- page N --> markers in output", Default: false},
			{Name: "split", Help: "split output into one .md per chapter using manifest chapters", Default: true},
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
	sourceImages := c.String("source-images")
	keepMarkers := c.Bool("keep-page-markers")
	split := c.Bool("split")

	if textPath == "" {
		return fmt.Errorf("--text is required")
	}
	if manifestPath == "" {
		return fmt.Errorf("--manifest is required")
	}

	absText, err := filepath.Abs(textPath)
	if err != nil {
		return fmt.Errorf("resolving text path: %w", err)
	}
	content, err := os.ReadFile(absText)
	if err != nil {
		return fmt.Errorf("reading text: %w", err)
	}

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

	imagesByPage := make(map[int][]ImageEntry)
	for _, img := range manifest.Images {
		imagesByPage[img.Page] = append(imagesByPage[img.Page], img)
	}
	for page := range imagesByPage {
		sort.Slice(imagesByPage[page], func(i, j int) bool {
			return imagesByPage[page][i].Seq < imagesByPage[page][j].Seq
		})
	}
	fmt.Fprintf(os.Stderr, "Loaded %d images across %d pages\n", len(manifest.Images), len(imagesByPage))

	text := string(content)
	pages := splitByPageMarkers(text)

	textPageSet := make(map[int]bool)
	for _, page := range pages {
		textPageSet[page.pageNum] = true
	}

	orphanAfter := make(map[int][]ImageEntry)
	for imgPage, imgs := range imagesByPage {
		if textPageSet[imgPage] {
			continue
		}
		bestPage := 0
		for _, page := range pages {
			if page.pageNum < imgPage && page.pageNum > bestPage {
				bestPage = page.pageNum
			}
		}
		if bestPage == 0 && len(pages) > 0 {
			bestPage = pages[0].pageNum
		}
		orphanAfter[bestPage] = append(orphanAfter[bestPage], imgs...)
	}

	var composed []pageBlock
	for _, page := range pages {
		var sb strings.Builder
		sb.WriteString(strings.TrimSpace(page.text))
		sb.WriteString("\n")

		if imgs, ok := imagesByPage[page.pageNum]; ok {
			for _, img := range imgs {
				imgPath := filepath.Join(imageDir, img.File)
				caption := fmt.Sprintf("Page %d", img.Page)
				sb.WriteString(fmt.Sprintf("\n[[IMG:%s|%s]]\n", imgPath, caption))
			}
		}

		if imgs, ok := orphanAfter[page.pageNum]; ok {
			for _, img := range imgs {
				imgPath := filepath.Join(imageDir, img.File)
				caption := fmt.Sprintf("Page %d", img.Page)
				sb.WriteString(fmt.Sprintf("\n[[IMG:%s|%s]]\n", imgPath, caption))
			}
		}

		composed = append(composed, pageBlock{pageNum: page.pageNum, text: sb.String()})
	}

	if split && outputPath != "" {
		return splitByChapters(composed, outputPath, manifest.Chapters, keepMarkers, sourceImages)
	}

	var sb strings.Builder
	for _, page := range composed {
		if keepMarkers {
			sb.WriteString(fmt.Sprintf("<!-- page %d -->\n\n", page.pageNum))
		}
		sb.WriteString(page.text)
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
		_, _ = fmt.Sscanf(marker, "<!-- page %d -->", &pageNum)

		var end int
		if i+1 < len(matches) {
			end = matches[i+1][0]
		} else {
			end = len(content)
		}

		text := content[loc[1]:end]
		text = strings.TrimPrefix(text, "\n")

		pages = append(pages, pageBlock{pageNum: pageNum, text: text})
	}

	return pages
}

func splitByChapters(pages []pageBlock, outputDir string, chapters []ChapterEntry, keepMarkers bool, sourceImages string) error {
	if len(chapters) == 0 {
		return fmt.Errorf("manifest has no chapters defined")
	}

	if err := os.MkdirAll(outputDir, 0755); err != nil {
		return fmt.Errorf("creating output dir: %w", err)
	}

	sort.Slice(chapters, func(i, j int) bool {
		return chapters[i].Page < chapters[j].Page
	})

	chapterStartPages := make(map[int]int)
	for i, ch := range chapters {
		chapterStartPages[ch.Page] = i
	}

	type chapterContent struct {
		num   int
		title string
		pages []pageBlock
	}

	var result []chapterContent
	currentIdx := -1

	for _, page := range pages {
		if idx, ok := chapterStartPages[page.pageNum]; ok {
			currentIdx = idx
			result = append(result, chapterContent{
				num:   idx,
				title: chapters[idx].Title,
			})
		}

		if currentIdx < 0 {
			result = append(result, chapterContent{
				num:   0,
				title: "Front Matter",
			})
			currentIdx = 0
		}

		result[len(result)-1].pages = append(result[len(result)-1].pages, page)
	}

	for _, ch := range result {
		var sb strings.Builder

		sb.WriteString(fmt.Sprintf("# %s\n\n", ch.title))

		isFirstPage := true
		for _, page := range ch.pages {
			if keepMarkers {
				sb.WriteString(fmt.Sprintf("<!-- page %d -->\n\n", page.pageNum))
			}
			text := stripOldChapterHeading(page.text)
			if isFirstPage {
				text = stripChapterPreamble(text, ch.title)
				isFirstPage = false
			}
			sb.WriteString(text)
			sb.WriteString("\n")
		}

		slug := fmt.Sprintf("ch%02d - %s", ch.num, sanitizeFilename(ch.title))
		mdPath := filepath.Join(outputDir, slug+".md")

		text := sb.String()
		if err := os.WriteFile(mdPath, []byte(text), 0644); err != nil {
			return fmt.Errorf("writing %s: %w", mdPath, err)
		}

		imgTags := reIMGTag.FindAllStringSubmatch(text, -1)
		if len(imgTags) > 0 {
			imgDir := filepath.Join(outputDir, "images")
			if err := os.MkdirAll(imgDir, 0755); err != nil {
				return fmt.Errorf("creating image dir %s: %w", imgDir, err)
			}
			for _, tag := range imgTags {
				imgFile := filepath.Base(tag[1])
				dst := filepath.Join(imgDir, imgFile)
				if _, err := os.Stat(dst); os.IsNotExist(err) {
					if sourceImages != "" {
						src := filepath.Join(sourceImages, imgFile)
						if data, readErr := os.ReadFile(src); readErr == nil {
							_ = os.WriteFile(dst, data, 0644)
						}
					}
				}
			}
			fmt.Fprintf(os.Stderr, "  %s: %d images\n", slug, len(imgTags))
		}

		fmt.Fprintf(os.Stderr, "Wrote %s\n", mdPath)
	}

	fmt.Fprintf(os.Stderr, "Split into %d chapters in %s\n", len(result), outputDir)
	return nil
}

func sanitizeFilename(name string) string {
	replacer := strings.NewReplacer(
		"/", "-",
		"\\", "-",
		":", "-",
		"\"", "",
		"?", "",
		"*", "",
		"<", "",
		">", "",
		"|", "",
	)
	result := replacer.Replace(name)
	result = strings.TrimRight(result, ". ")
	return result
}

var reOldHeading = regexp.MustCompile(`(?m)^## .+\n+`)

func stripOldChapterHeading(text string) string {
	return reOldHeading.ReplaceAllString(text, "")
}

var reAllCapsLine = regexp.MustCompile(`^[A-Z][A-Z\s,.\-:;'"!?]+$`)

func stripChapterPreamble(text, title string) string {
	lines := strings.Split(text, "\n")
	titleLower := strings.ToLower(strings.TrimSpace(title))

	stripped := 0
	for stripped < len(lines) && stripped < 10 {
		line := strings.TrimSpace(lines[stripped])
		if line == "" {
			stripped++
			continue
		}
		lineLower := strings.ToLower(line)
		if lineLower == titleLower {
			stripped++
			continue
		}
		if reAllCapsLine.MatchString(line) && len(line) < 60 {
			stripped++
			continue
		}
		break
	}
	if stripped > 0 {
		return strings.Join(lines[stripped:], "\n")
	}
	return text
}
