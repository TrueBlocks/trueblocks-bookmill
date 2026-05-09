package pipeline

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

// exportYear is the year used in export filenames. Set by NewRunner from config.
var exportYear = time.Now().Format("2006")

func (r *Runner) docxWorker() {
	defer r.docxWg.Done()
	for job := range r.docxCh {
		r.Log.Printf("  [docx] %s: export starting", job.essay.Slug)
		if err := r.exportEssay(job.ps, job.essay); err != nil {
			r.Log.Printf("  [docx] ERROR export %s: %v", job.essay.Slug, err)
			r.markError(job.ps, job.essay, StageExport, err)
			continue
		}

		dstFile := filepath.Join(job.ps.BaseDir, "export", exportFilename(job.essay))

		r.Log.Printf("  [docx] %s: Word upgrade starting", job.essay.Slug)
		if err := r.upgradeDocx(dstFile); err != nil {
			r.Log.Printf("  [docx] WARNING: Word upgrade %s: %v", job.essay.Slug, err)
		}

		if err := r.swapImages(job.ps, job.essay, dstFile); err != nil {
			r.Log.Printf("  [docx] ERROR imageswap %s: %v", job.essay.Slug, err)
		}

		r.Log.Printf("  [docx] %s: complete", job.essay.Slug)
	}
}

func (r *Runner) processDocxSync(ps *PipelineState, essay *EssayState) error {
	if err := r.exportEssay(ps, essay); err != nil {
		return err
	}
	dstFile := filepath.Join(ps.BaseDir, "export", exportFilename(essay))
	if err := r.upgradeDocx(dstFile); err != nil {
		r.Log.Printf("    WARNING: Word upgrade %s: %v", essay.Slug, err)
	}
	if err := r.swapImages(ps, essay, dstFile); err != nil {
		r.Log.Printf("    WARNING: imageswap %s: %v", essay.Slug, err)
	}
	return nil
}

func (r *Runner) findExportSource(ps *PipelineState, slug string) string {
	// Novel pipeline: revision → draft
	// Essay pipeline: draft2 → illustrate
	var candidates []string
	if ps.Genre != nil && ps.Genre.IsNovel() {
		candidates = []string{"revision", "draft"}
	} else {
		candidates = []string{"draft2", "illustrate"}
	}
	for _, stage := range candidates {
		path := filepath.Join(ps.BaseDir, stage, slug+".md")
		if _, err := os.Stat(path); err == nil {
			return path
		}
	}
	return ""
}

func (r *Runner) exportEssay(ps *PipelineState, essay *EssayState) error {
	r.markInProgress(ps, essay, StageExport, "local")

	mdFile := r.findExportSource(ps, essay.Slug)
	if mdFile == "" {
		return fmt.Errorf("no exportable draft found for %s", essay.Slug)
	}

	raw, err := os.ReadFile(mdFile)
	if err != nil {
		return fmt.Errorf("reading markdown: %w", err)
	}
	cleaned := sanitizeForDocx(string(raw))

	tmpFile, err := os.CreateTemp("", "export-*.md")
	if err != nil {
		return fmt.Errorf("creating temp file: %w", err)
	}
	defer os.Remove(tmpFile.Name())
	if _, err := tmpFile.WriteString(cleaned); err != nil {
		tmpFile.Close()
		return fmt.Errorf("writing temp file: %w", err)
	}
	tmpFile.Close()

	exportDir := filepath.Join(ps.BaseDir, "export")
	if err := os.MkdirAll(exportDir, 0755); err != nil {
		return fmt.Errorf("creating export dir: %w", err)
	}

	home, _ := os.UserHomeDir()
	templatePath := filepath.Join(home, ".local", "share", "trueblocks", "works", "works", "templates", "book-template.dotm")
	exportName := exportFilename(essay)
	outFile := filepath.Join(exportDir, exportName)

	cmd := exec.Command("md2docx", templatePath, tmpFile.Name(), outFile)
	if output, err := cmd.CombinedOutput(); err != nil {
		return fmt.Errorf("md2docx: %w %s", err, string(output))
	}

	if strings.Contains(cleaned, "[[IMG:") {
		dataDir := filepath.Join(r.BaseDir, "data")
		renderCmd := exec.Command("imagerender", "--data", dataDir, "--slug", essay.Slug, ps.BaseDir)
		if output, err := renderCmd.CombinedOutput(); err != nil {
			return fmt.Errorf("imagerender %s: %w %s", essay.Slug, err, string(output))
		} else {
			r.Log.Printf("    imagerender: %s", strings.TrimSpace(string(output)))
		}
	}

	r.Log.Printf("    exported → %s", exportName)

	result := &APIResult{}
	r.markComplete(ps, essay, StageExport, "local", result)
	return nil
}

func (r *Runner) swapImages(ps *PipelineState, essay *EssayState, docxPath string) error {
	mdFile := filepath.Join(ps.BaseDir, "draft2", essay.Slug+".md")
	if _, err := os.Stat(mdFile); os.IsNotExist(err) {
		mdFile = filepath.Join(ps.BaseDir, "illustrate", essay.Slug+".md")
	}
	raw, err := os.ReadFile(mdFile)
	if err != nil {
		return nil
	}
	if !strings.Contains(string(raw), "[[IMG:") {
		return nil
	}
	r.Log.Printf("  [docx] %s: imageswap starting", essay.Slug)
	swapCmd := exec.Command("imageswap", "--slug", essay.Slug, "--images", filepath.Join(ps.BaseDir, "images"), docxPath)
	if output, err := swapCmd.CombinedOutput(); err != nil {
		return fmt.Errorf("imageswap %s: %w %s", essay.Slug, err, string(output))
	}
	return nil
}

func (r *Runner) upgradeDocx(docxPath string) error {
	absPath, err := filepath.Abs(docxPath)
	if err != nil {
		return fmt.Errorf("abs path: %w", err)
	}

	home, _ := os.UserHomeDir()
	templatePath := filepath.Join(home, ".local", "share", "trueblocks", "works", "works", "templates", "book-template.dotm")

	escaped := func(s string) string {
		var out strings.Builder
		for _, c := range s {
			switch c {
			case '"', '\u201C', '\u201D':
				out.WriteString(`\"`)
			case '\\':
				out.WriteString(`\\`)
			default:
				out.WriteRune(c)
			}
		}
		return out.String()
	}

	baseName := filepath.Base(docxPath)
	title := baseName[:len(baseName)-len(filepath.Ext(baseName))]

	script := `
tell application "Microsoft Word"
	launch
end tell
tell application "System Events"
	set visible of process "Microsoft Word" to false
end tell
tell application "Microsoft Word"
	open (POSIX file "` + escaped(templatePath) + `" as text)
	set templateDoc to active document
	set templatePageSetup to page setup of section 1 of templateDoc
	set templatePageWidth to page width of templatePageSetup
	set templatePageHeight to page height of templatePageSetup
	set templateTopMargin to top margin of templatePageSetup
	set templateBottomMargin to bottom margin of templatePageSetup
	set templateLeftMargin to left margin of templatePageSetup
	set templateRightMargin to right margin of templatePageSetup
	set templateMirrorMargins to mirror margins of templatePageSetup
	set templateGutter to gutter of templatePageSetup
	close templateDoc saving no
	open (POSIX file "` + escaped(absPath) + `" as text)
	set theDoc to active document
	set docPageSetup to page setup of section 1 of theDoc
	set docPageWidth to page width of docPageSetup
	set docPageHeight to page height of docPageSetup
	set docIsLandscape to (docPageWidth > docPageHeight)
	if docIsLandscape then
		set tempWidth to templatePageWidth
		set templatePageWidth to templatePageHeight
		set templatePageHeight to tempWidth
		set tempTopMargin to templateTopMargin
		set tempBottomMargin to templateBottomMargin
		set templateTopMargin to templateLeftMargin
		set templateBottomMargin to templateRightMargin
		set templateLeftMargin to tempTopMargin
		set templateRightMargin to tempBottomMargin
	end if
	set widthRatio to templatePageWidth / docPageWidth
	set heightRatio to templatePageHeight / docPageHeight
	if widthRatio < heightRatio then
		set scaleFactor to widthRatio
	else
		set scaleFactor to heightRatio
	end if
	set inlineShapeCount to count of inline shapes of theDoc
	repeat with i from 1 to inlineShapeCount
		try
			set inlineShape to inline shape i of theDoc
			set width of inlineShape to (width of inlineShape) * scaleFactor
			set height of inlineShape to (height of inlineShape) * scaleFactor
		end try
	end repeat
	set floatShapeCount to count of shapes of theDoc
	repeat with i from 1 to floatShapeCount
		try
			set floatShape to shape i of theDoc
			set width of floatShape to (width of floatShape) * scaleFactor
			set height of floatShape to (height of floatShape) * scaleFactor
		end try
	end repeat
	tell theDoc to copy styles from template template (POSIX file "` + escaped(templatePath) + `" as text)
	set page width of docPageSetup to templatePageWidth
	set page height of docPageSetup to templatePageHeight
	set mirror margins of docPageSetup to templateMirrorMargins
	set gutter of docPageSetup to templateGutter
	set top margin of docPageSetup to templateTopMargin
	set bottom margin of docPageSetup to templateBottomMargin
	set left margin of docPageSetup to templateLeftMargin
	set right margin of docPageSetup to templateRightMargin
	set title of properties of theDoc to "` + escaped(title) + `"
	save theDoc
	close theDoc
	quit
end tell
`

	tmpFile, err := os.CreateTemp("", "upgrade-*.scpt")
	if err != nil {
		return fmt.Errorf("creating temp script: %w", err)
	}
	scriptPath := tmpFile.Name()
	if _, err := tmpFile.WriteString(script); err != nil {
		tmpFile.Close()
		os.Remove(scriptPath)
		return fmt.Errorf("writing script: %w", err)
	}
	tmpFile.Close()
	defer os.Remove(scriptPath)

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	cmd := exec.CommandContext(ctx, "osascript", scriptPath)
	output, err := cmd.CombinedOutput()
	if ctx.Err() == context.DeadlineExceeded {
		return fmt.Errorf("Word upgrade timed out")
	}
	if err != nil {
		return fmt.Errorf("osascript: %w: %s", err, string(output))
	}
	return nil
}
