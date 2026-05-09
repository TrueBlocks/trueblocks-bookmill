package main

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	appkit "github.com/TrueBlocks/trueblocks-art/packages/appkit/v2"
	"github.com/TrueBlocks/trueblocks-art/packages/cli"
	"github.com/TrueBlocks/trueblocks-math/internal/dalle"
	"github.com/TrueBlocks/trueblocks-math/internal/pipeline"
	"gopkg.in/yaml.v3"
)

var version = "dev"

type imageMeta struct {
	Filename    string `yaml:"filename"`
	Method      string `yaml:"method"`
	Description string `yaml:"description"`
}

func main() {
	app := cli.App{
		Name:        "imagerender",
		Description: "Render images (mermaid / R / DALL-E) referenced by .meta.yaml files in a project's images directory.",
		Version:     version,
		Flags: []cli.FlagDef{
			{Name: "config", Help: "path to config.yaml", Default: pipeline.DefaultConfigPath()},
			{Name: "base-dir", Help: "project base directory (default: <cwd>/projects/math-books)"},
			{Name: "data", Help: "path to data directory containing common.R and mermaid-theme.json"},
			{Name: "slug", Help: "render images for a specific essay slug only"},
			{Name: "force", Help: "re-render even if PNG exists and is newer than source", Default: false},
		},
		Run: run,
	}
	cli.Exit(app.Main())
}

func run(c *cli.Context) error {
	configPath := c.String("config")
	dataDir := c.String("data")
	slug := c.String("slug")
	force := c.Bool("force")
	baseDir := c.String("base-dir")

	cfg, err := pipeline.LoadConfig(configPath)
	if err != nil {
		return fmt.Errorf("loading config: %w", err)
	}

	if baseDir == "" {
		cwd, err := os.Getwd()
		if err != nil {
			return fmt.Errorf("getting cwd: %w", err)
		}
		baseDir = filepath.Join(cwd, "projects", "math-books")
	}

	imagesDir := filepath.Join(baseDir, "images")
	if _, err := os.Stat(imagesDir); os.IsNotExist(err) {
		return fmt.Errorf("images directory not found: %s", imagesDir)
	}

	entries, err := os.ReadDir(imagesDir)
	if err != nil {
		return fmt.Errorf("reading images dir: %w", err)
	}

	var themeFile, commonR string
	if dataDir != "" {
		themeFile = filepath.Join(dataDir, "mermaid-theme.json")
		commonR = filepath.Join(dataDir, "common.R")
	}

	rendered := 0
	skipped := 0
	failed := 0

	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}
		if slug != "" && entry.Name() != slug {
			continue
		}

		slugDir := filepath.Join(imagesDir, entry.Name())
		metas, err := filepath.Glob(filepath.Join(slugDir, "*.meta.yaml"))
		if err != nil {
			log.Printf("glob error in %s: %v", slugDir, err)
			continue
		}

		for _, metaPath := range metas {
			data, err := os.ReadFile(metaPath)
			if err != nil {
				log.Printf("reading %s: %v", metaPath, err)
				failed++
				continue
			}

			var meta imageMeta
			if err := yaml.Unmarshal(data, &meta); err != nil {
				log.Printf("parsing %s: %v", metaPath, err)
				failed++
				continue
			}

			outPath := filepath.Join(slugDir, meta.Filename)

			srcPath := sourcePathForMeta(slugDir, meta)
			if srcPath == "" {
				log.Printf("unknown method %q for %s", meta.Method, meta.Filename)
				failed++
				continue
			}

			if _, err := os.Stat(srcPath); os.IsNotExist(err) {
				log.Printf("source not found: %s", srcPath)
				failed++
				continue
			}

			if !force && isNewer(outPath, srcPath) {
				skipped++
				continue
			}

			log.Printf("rendering %s (%s)", meta.Filename, meta.Method)

			const maxRepairs = 3
			var renderErr error
			var renderOutput string

			for attempt := 0; attempt <= maxRepairs; attempt++ {
				switch meta.Method {
				case "mermaid":
					renderOutput, renderErr = renderMermaid(srcPath, outPath, themeFile)
				case "r":
					renderOutput, renderErr = renderR(srcPath, outPath, commonR)
				case "ai":
					renderErr = renderAI(srcPath, outPath, cfg.API.OpenAIKey)
					renderOutput = ""
				}

				if renderErr == nil {
					break
				}

				if meta.Method == "ai" {
					if isSafetyViolation(renderErr) && attempt == 0 {
						log.Printf("  safety violation, sanitizing prompt...")
						if repairErr := sanitizeAIPrompt(srcPath, cfg.API.AnthropicKey); repairErr != nil {
							log.Printf("  sanitize failed: %v", repairErr)
							break
						}
						log.Printf("  sanitized %s, retrying...", filepath.Base(srcPath))
						continue
					}
					break
				}

				if attempt == maxRepairs {
					break
				}

				log.Printf("  attempt %d failed: %v", attempt+1, renderErr)
				if repairErr := repairSource(srcPath, meta.Method, renderOutput, cfg.API.AnthropicKey); repairErr != nil {
					log.Printf("  repair failed: %v", repairErr)
					break
				}
				log.Printf("  repaired %s, retrying...", filepath.Base(srcPath))
			}

			if renderErr != nil {
				log.Printf("  FAILED: %v", renderErr)
				promptData, _ := os.ReadFile(srcPath)
				if fbErr := renderTextFallback(outPath, string(promptData)); fbErr != nil {
					log.Printf("  fallback image also failed: %v", fbErr)
					failed++
				} else {
					log.Printf("  OK (text fallback): %s", outPath)
					rendered++
				}
			} else {
				log.Printf("  OK: %s", outPath)
				rendered++
			}
		}
	}

	log.Printf("done: %d rendered, %d skipped, %d failed", rendered, skipped, failed)
	return nil
}

func sourcePathForMeta(dir string, meta imageMeta) string {
	base := strings.TrimSuffix(meta.Filename, ".png")
	switch meta.Method {
	case "mermaid":
		return filepath.Join(dir, base+".mermaid")
	case "r":
		return filepath.Join(dir, base+".R")
	case "ai":
		return filepath.Join(dir, base+".ai-prompt.txt")
	default:
		return ""
	}
}

func isNewer(target, source string) bool {
	tInfo, err := os.Stat(target)
	if err != nil {
		return false
	}
	sInfo, err := os.Stat(source)
	if err != nil {
		return false
	}
	return tInfo.ModTime().After(sInfo.ModTime())
}

func renderMermaid(srcPath, outPath, themeFile string) (string, error) {
	args := []string{"-i", srcPath, "-o", outPath, "-s", "4"}
	if _, err := os.Stat(themeFile); err == nil {
		args = append(args, "--configFile", themeFile)
	}
	cmd := exec.Command("mmdc", args...)
	var buf bytes.Buffer
	cmd.Stdout = &buf
	cmd.Stderr = &buf
	err := cmd.Run()
	return buf.String(), err
}

func renderR(srcPath, outPath, commonR string) (string, error) {
	dir := filepath.Dir(srcPath)

	cmd := exec.Command("Rscript", srcPath)
	cmd.Dir = dir
	cmd.Env = append(os.Environ(), "IMAGERENDER_OUTPUT="+outPath)
	if commonR != "" {
		cmd.Env = append(cmd.Env, "COMMON_R="+commonR)
	}
	var buf bytes.Buffer
	cmd.Stdout = &buf
	cmd.Stderr = &buf
	err := cmd.Run()
	return buf.String(), err
}

func renderAI(srcPath, outPath, apiKey string) error {
	if apiKey == "" {
		return fmt.Errorf("no OpenAI API key configured (api.openai_key)")
	}

	promptData, err := os.ReadFile(srcPath)
	if err != nil {
		return fmt.Errorf("reading prompt: %w", err)
	}
	prompt := strings.TrimSpace(string(promptData))

	imgData, err := dalle.GenerateImage(apiKey, prompt, "dall-e-3", "1792x1024", "")
	if err != nil {
		return err
	}

	return os.WriteFile(outPath, imgData, appkit.FilePermissions)
}

func repairSource(srcPath, method, errorOutput, apiKey string) error {
	if apiKey == "" {
		return fmt.Errorf("no Anthropic API key for repair")
	}

	src, err := os.ReadFile(srcPath)
	if err != nil {
		return fmt.Errorf("reading source: %w", err)
	}

	var lang string
	switch method {
	case "r":
		lang = "R"
	case "mermaid":
		lang = "Mermaid"
	default:
		return fmt.Errorf("cannot repair method %q", method)
	}

	prompt := fmt.Sprintf(`Fix this %s source code. It failed with the error shown below.

SOURCE FILE (%s):
%s

ERROR OUTPUT:
%s

RULES:
- Output ONLY the fixed %s code, nothing else — no explanation, no markdown fences, no commentary.
- The output must be directly executable as a .%s file.
- For R scripts: only use packages from base R, ggplot2, dplyr, and scales. Do NOT use packages like: png, plotly, rgl, gridExtra, patchwork, cowplot, magick, grid (except grid::unit).
- MUST call save_chart(p) to save output — NEVER use ggsave() directly.
- Use linewidth= NOT size= for ALL line-drawing geoms: geom_line, geom_path, geom_step, geom_segment, geom_curve, geom_vline, geom_hline, geom_abline, and element_line.
- NEVER use annotate("arrow", ...). For arrows, use annotate("segment", ..., arrow = arrow(length = unit(0.1, "inches"))).
- In element_text(), use face= NOT fontface= (fontface= is only valid in annotate("text")).
- Source common.R using the COMMON_R environment variable: source(Sys.getenv("COMMON_R")).
- For Mermaid: output valid Mermaid syntax only.
- Preserve the original intent and visual design.
- Fix the specific error, do not rewrite from scratch unless necessary.`,
		lang, filepath.Base(srcPath), string(src), errorOutput, lang, method)

	client := &pipeline.AnthropicClient{APIKey: apiKey}
	result, err := client.Call(context.Background(), "claude-sonnet-4-20250514", prompt, 60*time.Second)
	if err != nil {
		return fmt.Errorf("API call: %w", err)
	}

	fixed := strings.TrimSpace(result.Content)
	fixed = strings.TrimPrefix(fixed, "```r")
	fixed = strings.TrimPrefix(fixed, "```R")
	fixed = strings.TrimPrefix(fixed, "```mermaid")
	fixed = strings.TrimPrefix(fixed, "```")
	fixed = strings.TrimSuffix(fixed, "```")
	fixed = strings.TrimSpace(fixed)

	return os.WriteFile(srcPath, []byte(fixed+"\n"), appkit.FilePermissions)
}

func isSafetyViolation(err error) bool {
	return err != nil && strings.Contains(err.Error(), "content_policy_violation")
}

func sanitizeAIPrompt(srcPath, apiKey string) error {
	if apiKey == "" {
		return fmt.Errorf("no Anthropic API key for sanitization")
	}

	src, err := os.ReadFile(srcPath)
	if err != nil {
		return fmt.Errorf("reading prompt: %w", err)
	}

	prompt := fmt.Sprintf(`Rewrite this DALL-E image generation prompt so it will pass OpenAI's safety filter.
The original was rejected with a content_policy_violation.

ORIGINAL PROMPT:
%s

RULES:
- Output ONLY the rewritten prompt text, nothing else — no explanation, no commentary.
- Do NOT mention medical or psychiatric conditions, diagnoses, or symptoms by name.
- Do NOT use character names that reference disabilities or mental health (e.g. "Chronically Ill", "Manic", "Dumpy").
- Focus on visual scene descriptions: settings, objects, lighting, expressions, poses, clothing.
- Describe characters by their appearance and actions, not by labels or conditions.
- Keep the same visual intent and composition as the original.
- Keep it concise — under 400 words.`, string(src))

	client := &pipeline.AnthropicClient{APIKey: apiKey}
	result, err := client.Call(context.Background(), "claude-sonnet-4-20250514", prompt, 60*time.Second)
	if err != nil {
		return fmt.Errorf("API call: %w", err)
	}

	sanitized := strings.TrimSpace(result.Content)
	return os.WriteFile(srcPath, []byte(sanitized+"\n"), appkit.FilePermissions)
}

func renderTextFallback(outPath, promptText string) error {
	promptText = strings.TrimSpace(promptText)
	if len(promptText) > 800 {
		promptText = promptText[:800] + "..."
	}

	cmd := exec.Command("magick",
		"-size", "1792x1024",
		"xc:#f0f0f0",
		"-font", "Times-Italic",
		"-pointsize", "28",
		"-fill", "#444444",
		"-gravity", "Center",
		"-annotate", "0",
		promptText,
		outPath,
	)
	var buf bytes.Buffer
	cmd.Stdout = &buf
	cmd.Stderr = &buf
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("magick: %s: %w", buf.String(), err)
	}
	return nil
}
