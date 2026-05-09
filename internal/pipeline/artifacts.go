package pipeline

import (
	"os"
	"os/exec"
	"path/filepath"
)

func (r *Runner) allEssaysAtOrPast(ps *PipelineState, stage Stage) bool {
	ps.mu.Lock()
	defer ps.mu.Unlock()
	if len(ps.Essays) == 0 {
		return false
	}
	for _, essay := range ps.Essays {
		if essay.CurrentStage < stage {
			return false
		}
		if essay.CurrentStage == stage && essay.Status != "final" {
			return false
		}
	}
	return true
}

func (r *Runner) generateBookArtifacts(ps *PipelineState) {
	if !r.allEssaysAtOrPast(ps, StageDraft2) {
		return
	}

	bookDir := filepath.Join(ps.BaseDir, "book")

	needBlurb := !bookDirHasFile(bookDir, "back-cover-blurb.md")
	needCover := !bookDirHasFile(bookDir, "front-cover-prompt.md") || !bookDirHasFile(bookDir, "front-cover.png")

	if !needBlurb && !needCover {
		return
	}

	if r.Config.Pipeline.DryRun {
		r.Log.Printf("[%s] [book] dry-run: would generate book artifacts", ps.Project)
		return
	}

	go func() {
		if needBlurb {
			r.Log.Printf("[%s] [book] generating back-cover blurb...", ps.Project)
			cmd := exec.Command("bookblurb", ps.BaseDir)
			output, err := cmd.CombinedOutput()
			if err != nil {
				r.Log.Printf("[%s] [book] WARNING: bookblurb: %v\n%s", ps.Project, err, string(output))
				return
			}
			r.Log.Printf("[%s] [book] blurb generated", ps.Project)
		}

		if needCover {
			r.Log.Printf("[%s] [book] generating front-cover prompt and image...", ps.Project)
			cmd := exec.Command("bookcover", ps.BaseDir)
			output, err := cmd.CombinedOutput()
			if err != nil {
				r.Log.Printf("[%s] [book] WARNING: bookcover: %v\n%s", ps.Project, err, string(output))
				return
			}
			r.Log.Printf("[%s] [book] cover generated", ps.Project)
		}
	}()
}

func fileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

func invalidateBookArtifacts(baseDir string) {
	bookDir := filepath.Join(baseDir, "book")
	for _, name := range []string{"back-cover-blurb.md", "front-cover-prompt.md", "front-cover.png"} {
		p := filepath.Join(bookDir, name)
		if fileExists(p) {
			os.Remove(p)
		}
		entries, err := os.ReadDir(bookDir)
		if err != nil {
			continue
		}
		for _, e := range entries {
			if e.IsDir() {
				sp := filepath.Join(bookDir, e.Name(), name)
				if fileExists(sp) {
					os.Remove(sp)
				}
			}
		}
	}
}

func bookDirHasFile(bookDir, filename string) bool {
	if fileExists(filepath.Join(bookDir, filename)) {
		return true
	}
	entries, err := os.ReadDir(bookDir)
	if err != nil {
		return false
	}
	for _, e := range entries {
		if e.IsDir() && fileExists(filepath.Join(bookDir, e.Name(), filename)) {
			return true
		}
	}
	return false
}
