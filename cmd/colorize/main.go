package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
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

func main() {
	app := cli.App{
		Name:        "colorize",
		Description: "Colorize B&W images extracted from historical books. Wraps an external colorization tool and copies results to the output directory.",
		Version:     version,
		Flags: []cli.FlagDef{
			{Name: "input-dir", Help: "directory containing extracted B&W images and manifest.yaml", Default: ""},
			{Name: "output-dir", Help: "directory for colorized images (default: ./colorize/)", Default: ""},
			{Name: "copy-to", Help: "additional directory to copy colorized images to (e.g. ~/Home/Classics/...)", Default: ""},
			{Name: "tool", Help: "colorization tool to use: deoldify, python-script, or copy (default: copy)", Default: "copy"},
			{Name: "script", Help: "path to custom Python colorization script (used with --tool=python-script)", Default: ""},
		},
		Run: run,
	}
	cli.Exit(app.Main())
}

func run(c *cli.Context) error {
	inputDir := c.String("input-dir")
	if inputDir == "" {
		return fmt.Errorf("--input-dir is required")
	}

	absInput, err := filepath.Abs(inputDir)
	if err != nil {
		return fmt.Errorf("resolving input dir: %w", err)
	}

	manifestPath := filepath.Join(absInput, "manifest.yaml")
	manifestData, err := os.ReadFile(manifestPath)
	if err != nil {
		return fmt.Errorf("reading manifest: %w", err)
	}

	var manifest Manifest
	if err := yaml.Unmarshal(manifestData, &manifest); err != nil {
		return fmt.Errorf("parsing manifest: %w", err)
	}

	outputDir := c.String("output-dir")
	if outputDir == "" {
		outputDir = "./colorize"
	}
	absOutput, err := filepath.Abs(outputDir)
	if err != nil {
		return fmt.Errorf("resolving output dir: %w", err)
	}
	if err := os.MkdirAll(absOutput, 0755); err != nil {
		return fmt.Errorf("creating output dir: %w", err)
	}

	copyTo := c.String("copy-to")
	if copyTo != "" {
		absCopy, err := filepath.Abs(copyTo)
		if err != nil {
			return fmt.Errorf("resolving copy-to dir: %w", err)
		}
		copyTo = absCopy
		if err := os.MkdirAll(copyTo, 0755); err != nil {
			return fmt.Errorf("creating copy-to dir: %w", err)
		}
	}

	tool := c.String("tool")
	scriptPath := c.String("script")

	fmt.Fprintf(os.Stderr, "Colorizing %d images using %s...\n", len(manifest.Images), tool)

	for i, entry := range manifest.Images {
		srcPath := filepath.Join(absInput, entry.File)
		dstPath := filepath.Join(absOutput, entry.File)

		fmt.Fprintf(os.Stderr, "[%d/%d] %s\n", i+1, len(manifest.Images), entry.File)

		var err error
		switch tool {
		case "copy":
			err = copyFile(srcPath, dstPath)
		case "deoldify":
			err = runDeOldify(srcPath, dstPath)
		case "python-script":
			if scriptPath == "" {
				return fmt.Errorf("--script is required when --tool=python-script")
			}
			err = runPythonScript(scriptPath, srcPath, dstPath)
		default:
			return fmt.Errorf("unknown tool: %s", tool)
		}

		if err != nil {
			fmt.Fprintf(os.Stderr, "  warning: colorize failed for %s: %v\n", entry.File, err)
			if copyErr := copyFile(srcPath, dstPath); copyErr != nil {
				fmt.Fprintf(os.Stderr, "  warning: fallback copy also failed: %v\n", copyErr)
			}
			continue
		}

		if copyTo != "" {
			copyDst := filepath.Join(copyTo, entry.File)
			if err := copyFile(dstPath, copyDst); err != nil {
				fmt.Fprintf(os.Stderr, "  warning: copy to %s failed: %v\n", copyTo, err)
			}
		}
	}

	outManifest := manifest
	manifestOut := filepath.Join(absOutput, "manifest.yaml")
	data, err := yaml.Marshal(outManifest)
	if err != nil {
		return fmt.Errorf("marshaling output manifest: %w", err)
	}
	if err := os.WriteFile(manifestOut, data, 0644); err != nil {
		return fmt.Errorf("writing output manifest: %w", err)
	}

	fmt.Fprintf(os.Stderr, "Colorized %d images to %s\n", len(manifest.Images), absOutput)
	return nil
}

func copyFile(src, dst string) error {
	data, err := os.ReadFile(src)
	if err != nil {
		return err
	}
	return os.WriteFile(dst, data, 0644)
}

func runDeOldify(src, dst string) error {
	cmd := exec.Command("python3", "-c", fmt.Sprintf(`
import sys
try:
    from deoldify import device
    from deoldify.device_id import DeviceId
    device.set(device=DeviceId.CPU)
    from deoldify.visualize import get_image_colorizer
    colorizer = get_image_colorizer(artistic=True)
    colorizer.plot_transformed_image(path="%s", render_factor=35, display_render_factor=True, figsize=(8,8))
    import shutil
    result = "%s" .replace(".png", "_result.png")
    shutil.move(result, "%s")
except ImportError:
    print("DeOldify not installed", file=sys.stderr)
    sys.exit(1)
`, src, src, dst))
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func runPythonScript(script, src, dst string) error {
	absScript, err := filepath.Abs(script)
	if err != nil {
		return err
	}
	cmd := exec.Command("python3", absScript, src, dst)
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func slugFromBook(bookFile string) string {
	name := strings.TrimSuffix(bookFile, filepath.Ext(bookFile))
	parts := strings.SplitN(name, " - ", 3)
	if len(parts) == 3 {
		return slugify(parts[2])
	}
	return slugify(name)
}

func slugify(s string) string {
	s = strings.ToLower(s)
	var result []byte
	for i := 0; i < len(s); i++ {
		c := s[i]
		if (c >= 'a' && c <= 'z') || (c >= '0' && c <= '9') {
			result = append(result, c)
		} else if c == ' ' || c == '-' || c == '_' {
			if len(result) > 0 && result[len(result)-1] != '-' {
				result = append(result, '-')
			}
		}
	}
	return strings.Trim(string(result), "-")
}
