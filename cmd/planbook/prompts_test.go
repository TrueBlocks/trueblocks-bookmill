package main

import (
	"strings"
	"testing"

	"github.com/TrueBlocks/trueblocks-bookmill/internal/pipeline"
)

func TestBuildPromptNovelBranchCarriesSubplotAndOrigins(t *testing.T) {
	origins := []originFile{{name: "seed.md", content: "A detective loses his memory."}}
	p := buildPrompt(origins, "", "The Forgetting", &pipeline.Genre{Form: "novel"})

	for _, want := range []string{
		"expert novel plotter",
		"Subplot",
		"BOOK TITLE: The Forgetting",
		"### File: seed.md",
		"A detective loses his memory.",
		"Now produce the Plan",
	} {
		if !strings.Contains(p, want) {
			t.Errorf("novel plan prompt is missing %q", want)
		}
	}
	if strings.Contains(p, "Hidden Theme") {
		t.Error("novel branch should not carry the book branch's Hidden Theme column")
	}
}

func TestBuildPromptBookBranchCarriesHiddenTheme(t *testing.T) {
	origins := []originFile{{name: "notes.md", content: "Essays about vanished trades."}}
	p := buildPrompt(origins, "", "", &pipeline.Genre{Form: "essay"})

	if !strings.Contains(p, "expert book editor") {
		t.Error("book plan prompt is missing the book-editor framing")
	}
	if !strings.Contains(p, "Hidden Theme") {
		t.Error("book plan prompt is missing the Hidden Theme column")
	}
	if strings.Contains(p, "Subplot") {
		t.Error("book branch should not carry the novel branch's Subplot column")
	}
	if strings.Contains(p, "BOOK TITLE:") {
		t.Error("an empty title should leave no BOOK TITLE line")
	}
}
