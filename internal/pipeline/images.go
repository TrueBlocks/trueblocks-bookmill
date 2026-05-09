package pipeline

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
)

type imageSource struct {
	filename string
	method   string
	source   string
}

func parseImageSources(content string) []imageSource {
	var sources []imageSource
	sepIdx := strings.Index(content, "---IMAGE-SOURCES---")
	if sepIdx < 0 {
		return nil
	}
	remainder := content[sepIdx:]

	re := regexp.MustCompile(`---IMAGE:([^|]+)\|method:(\w+)---\n([\s\S]*?)---END-IMAGE---`)
	matches := re.FindAllStringSubmatch(remainder, -1)
	for _, m := range matches {
		sources = append(sources, imageSource{
			filename: strings.TrimSpace(m[1]),
			method:   strings.TrimSpace(m[2]),
			source:   strings.TrimSpace(m[3]),
		})
	}
	return sources
}

func (r *Runner) writeImageSources(ps *PipelineState, slug, content string) error {
	sources := parseImageSources(content)
	if len(sources) == 0 {
		r.Log.Printf("    no image sources found in illustrate output")
		return nil
	}

	imgDir := filepath.Join(ps.BaseDir, "images", slug)
	if err := os.MkdirAll(imgDir, 0755); err != nil {
		return fmt.Errorf("creating images dir: %w", err)
	}

	for _, src := range sources {
		baseName := strings.TrimSuffix(src.filename, filepath.Ext(src.filename))

		var ext string
		switch src.method {
		case "mermaid":
			ext = ".mermaid"
		case "r":
			ext = ".R"
		case "ai":
			ext = ".ai-prompt.txt"
		default:
			r.Log.Printf("    unknown image method: %s for %s", src.method, src.filename)
			continue
		}

		srcPath := filepath.Join(imgDir, baseName+ext)
		if err := os.WriteFile(srcPath, []byte(src.source), 0644); err != nil {
			return fmt.Errorf("writing %s: %w", srcPath, err)
		}

		metaContent := fmt.Sprintf("filename: %s\nmethod: %s\ndescription: %s\n",
			src.filename, src.method, baseName)
		metaPath := filepath.Join(imgDir, baseName+".meta.yaml")
		if err := os.WriteFile(metaPath, []byte(metaContent), 0644); err != nil {
			return fmt.Errorf("writing %s: %w", metaPath, err)
		}

		r.Log.Printf("    image source: %s (%s)", baseName+ext, src.method)
	}

	essayText := content
	if sepIdx := strings.Index(content, "---IMAGE-SOURCES---"); sepIdx > 0 {
		essayText = strings.TrimSpace(content[:sepIdx])
	}
	if err := ps.WriteContent(StageIllustrate, slug, essayText); err != nil {
		return fmt.Errorf("rewriting illustrate content: %w", err)
	}

	return nil
}

func (r *Runner) renderAllImages(ps *PipelineState) {
	r.Log.Printf("[%s] imagerender: starting", ps.Project)
	dataDir := filepath.Join(r.BaseDir, "data")
	cmd := exec.Command("imagerender", "--data", dataDir, ps.BaseDir)
	output, err := cmd.CombinedOutput()
	if err != nil {
		r.Log.Printf("[%s] WARNING: imagerender: %v\n%s", ps.Project, err, string(output))
		return
	}
	lines := strings.TrimSpace(string(output))
	for _, line := range strings.Split(lines, "\n") {
		line = strings.TrimSpace(line)
		if line != "" {
			r.Log.Printf("[%s] imagerender: %s", ps.Project, line)
		}
	}
}

func (r *Runner) renderImages(ps *PipelineState, slug string) {
	dataDir := filepath.Join(r.BaseDir, "data")
	cmd := exec.Command("imagerender", "--data", dataDir, "--slug", slug, ps.BaseDir)
	output, err := cmd.CombinedOutput()
	if err != nil {
		r.Log.Printf("    WARNING: imagerender %s: %v\n%s", slug, err, string(output))
		return
	}
	lines := strings.TrimSpace(string(output))
	if lines != "" {
		for _, line := range strings.Split(lines, "\n") {
			r.Log.Printf("    imagerender: %s", line)
		}
	}
}
