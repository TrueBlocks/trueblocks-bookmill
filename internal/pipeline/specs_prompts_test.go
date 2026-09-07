package pipeline

import (
	"io"
	"os"
	"path/filepath"
	"testing"
	"text/template"
)

func TestEveryPromptTemplateParsesAndRenders(t *testing.T) {
	root := filepath.Join("..", "..", "specs", "prompts")
	dirs := []string{filepath.Join(root, "generic")}
	seriesEntries, err := os.ReadDir(filepath.Join(root, "series"))
	if err != nil {
		t.Fatalf("reading series prompt dirs: %v", err)
	}
	for _, e := range seriesEntries {
		if e.IsDir() {
			dirs = append(dirs, filepath.Join(root, "series", e.Name()))
		}
	}

	raw := map[string]bool{"voice-summary": true, "essay-rules": true}
	sample := map[string]any{}

	found := 0
	for _, dir := range dirs {
		entries, err := os.ReadDir(dir)
		if err != nil {
			t.Fatalf("reading %s: %v", dir, err)
		}
		for _, e := range entries {
			if e.IsDir() || filepath.Ext(e.Name()) != ".md" {
				continue
			}
			name := e.Name()[:len(e.Name())-len(".md")]
			if raw[name] {
				continue
			}
			path := filepath.Join(dir, e.Name())
			content, err := os.ReadFile(path)
			if err != nil {
				t.Fatalf("reading %s: %v", path, err)
			}
			tmpl, err := template.New(name).Parse(string(content))
			if err != nil {
				t.Errorf("%s does not parse: %v", path, err)
				continue
			}
			if err := tmpl.Execute(io.Discard, sample); err != nil {
				t.Errorf("%s does not render: %v", path, err)
			}
			found++
		}
	}
	if found == 0 {
		t.Fatal("no prompt templates found to check")
	}
}
