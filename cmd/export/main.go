package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/TrueBlocks/trueblocks-art/packages/cli"
	"gopkg.in/yaml.v3"
)

var version = "dev"

var reNamePattern = regexp.MustCompile(`^(\d{4})\s*-\s*(.+?)\s*-\s*(.+)\.(pdf|PDF)$`)

func main() {
	app := cli.App{
		Name:        "export",
		Description: "Export composed markdown to .docx using md2docx and imageswap. Produces the final book file for import into the works system.",
		Version:     version,
		Flags: []cli.FlagDef{
			{Name: "input", Help: "path to composed markdown file", Default: ""},
			{Name: "manifest", Help: "path to manifest.yaml (output_dir read from here)", Default: ""},
			{Name: "output-dir", Help: "override output directory for .docx file", Default: ""},
			{Name: "template", Help: "path to .dotm template", Default: ""},
			{Name: "image-dir", Help: "directory containing colorized images for imageswap", Default: ""},
			{Name: "slug", Help: "override imageswap slug (default: derived from docx filename)", Default: ""},
			{Name: "book-file", Help: "original PDF filename (for naming the output)", Default: ""},
			{Name: "skip-imageswap", Help: "skip imageswap step", Default: false},
		},
		Run: run,
	}
	cli.Exit(app.Main())
}

func run(c *cli.Context) error {
	inputPath := c.String("input")
	if inputPath == "" {
		return fmt.Errorf("--input is required")
	}

	absInput, err := filepath.Abs(inputPath)
	if err != nil {
		return fmt.Errorf("resolving input path: %w", err)
	}
	if _, err := os.Stat(absInput); err != nil {
		return fmt.Errorf("input not found: %s", absInput)
	}

	outputDir := c.String("output-dir")
	if outputDir == "" {
		manifestPath := c.String("manifest")
		if manifestPath != "" {
			if data, err := os.ReadFile(manifestPath); err == nil {
				var m struct {
					OutputDir string `yaml:"output_dir"`
				}
				if yaml.Unmarshal(data, &m) == nil && m.OutputDir != "" {
					outputDir = expandHome(m.OutputDir)
				}
			}
		}
	}
	if outputDir == "" {
		outputDir = filepath.Join("works", "imports", "files")
	}
	absOutput, err := filepath.Abs(outputDir)
	if err != nil {
		return fmt.Errorf("resolving output dir: %w", err)
	}
	if err := os.MkdirAll(absOutput, 0755); err != nil {
		return fmt.Errorf("creating output dir: %w", err)
	}

	templatePath := c.String("template")
	if templatePath == "" {
		home, err := os.UserHomeDir()
		if err != nil {
			return fmt.Errorf("getting home dir: %w", err)
		}
		templatePath = filepath.Join(home, ".local", "share", "trueblocks", "works", "works", "templates", "book-template.dotm")
	}
	if _, err := os.Stat(templatePath); err != nil {
		return fmt.Errorf("template not found: %s", templatePath)
	}

	imageDir := c.String("image-dir")
	skipImageswap := c.Bool("skip-imageswap")
	bookFile := c.String("book-file")

	docxName := buildOutputName(bookFile, absInput)
	docxPath := filepath.Join(absOutput, docxName)

	fmt.Fprintf(os.Stderr, "Running md2docx...\n")
	md2docxPath := findTool("md2docx")
	if md2docxPath == "" {
		return fmt.Errorf("md2docx not found in PATH or ~/source/")
	}

	cmd := exec.Command(md2docxPath,
		templatePath,
		absInput,
		docxPath,
	)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("md2docx failed: %w", err)
	}
	fmt.Fprintf(os.Stderr, "Created %s\n", docxPath)

	if !skipImageswap && imageDir != "" {
		absImageDir, err := filepath.Abs(imageDir)
		if err != nil {
			return fmt.Errorf("resolving image dir: %w", err)
		}

		imageswapPath := findTool("imageswap")
		if imageswapPath == "" {
			fmt.Fprintf(os.Stderr, "Warning: imageswap not found, skipping image insertion\n")
		} else {
			fmt.Fprintf(os.Stderr, "Running imageswap...\n")
			slug := c.String("slug")
			args := []string{"--images", absImageDir, docxPath}
			if slug != "" {
				args = []string{"--images", absImageDir, "--slug", slug, docxPath}
			}
			cmd := exec.Command(imageswapPath, args...)
			cmd.Stderr = os.Stderr
			cmd.Stdout = os.Stdout
			if err := cmd.Run(); err != nil {
				fmt.Fprintf(os.Stderr, "Warning: imageswap failed: %v\n", err)
			} else {
				fmt.Fprintf(os.Stderr, "Images inserted into %s\n", docxPath)
			}
		}
	}

	fmt.Fprintf(os.Stderr, "Export complete: %s\n", docxPath)
	return nil
}

func buildOutputName(bookFile, inputPath string) string {
	if bookFile != "" {
		matches := reNamePattern.FindStringSubmatch(bookFile)
		if matches != nil {
			year := matches[1]
			author := matches[2]
			title := strings.TrimSuffix(matches[3], "."+matches[4])
			return fmt.Sprintf("cEssay - %s - %s - %s.docx", year, author, title)
		}
	}

	base := filepath.Base(inputPath)
	name := strings.TrimSuffix(base, filepath.Ext(base))
	return fmt.Sprintf("cEssay - %s.docx", name)
}

func findTool(name string) string {
	if path, err := exec.LookPath(name); err == nil {
		return path
	}

	home, err := os.UserHomeDir()
	if err != nil {
		return ""
	}
	candidate := filepath.Join(home, "source", name)
	if _, err := os.Stat(candidate); err == nil {
		return candidate
	}

	return ""
}

func expandHome(path string) string {
	if strings.HasPrefix(path, "~/") {
		home, err := os.UserHomeDir()
		if err != nil {
			return path
		}
		return filepath.Join(home, path[2:])
	}
	return path
}
