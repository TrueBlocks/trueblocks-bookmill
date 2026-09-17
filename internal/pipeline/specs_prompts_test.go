package pipeline

import (
	"io"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"text/template"

	"github.com/TrueBlocks/trueblocks-art/packages/writing"
)

func TestEveryPromptTemplateParsesAndRenders(t *testing.T) {
	raw := map[string]bool{"voice-summary": true, "essay-rules": true}
	sample := map[string]any{}
	found := 0

	check := func(name, source, content string) {
		if raw[name] {
			return
		}
		tmpl, err := template.New(name).Parse(content)
		if err != nil {
			t.Errorf("%s does not parse: %v", source, err)
			return
		}
		if err := tmpl.Execute(io.Discard, sample); err != nil {
			t.Errorf("%s does not render: %v", source, err)
		}
		found++
	}

	// Generic prompt chain now lives in the shared packages/writing home.
	for _, res := range writing.Names() {
		if !strings.HasPrefix(res, genericPromptPrefix) || filepath.Ext(res) != ".md" {
			continue
		}
		name := strings.TrimSuffix(strings.TrimPrefix(res, genericPromptPrefix), ".md")
		content, err := writing.Read(res)
		if err != nil {
			t.Fatalf("reading %s: %v", res, err)
		}
		check(name, res, string(content))
	}

	// Series prompts stay on disk.
	seriesRoot := filepath.Join("..", "..", "specs", "prompts", "series")
	seriesEntries, err := os.ReadDir(seriesRoot)
	if err != nil {
		t.Fatalf("reading series prompt dirs: %v", err)
	}
	for _, s := range seriesEntries {
		if !s.IsDir() {
			continue
		}
		dir := filepath.Join(seriesRoot, s.Name())
		entries, err := os.ReadDir(dir)
		if err != nil {
			t.Fatalf("reading %s: %v", dir, err)
		}
		for _, e := range entries {
			if e.IsDir() || filepath.Ext(e.Name()) != ".md" {
				continue
			}
			name := strings.TrimSuffix(e.Name(), ".md")
			path := filepath.Join(dir, e.Name())
			content, err := os.ReadFile(path)
			if err != nil {
				t.Fatalf("reading %s: %v", path, err)
			}
			check(name, path, string(content))
		}
	}

	if found == 0 {
		t.Fatal("no prompt templates found to check")
	}
}
