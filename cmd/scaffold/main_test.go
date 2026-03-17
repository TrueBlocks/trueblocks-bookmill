package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestSplitTableRow(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		line string
		want []string
	}{
		{
			name: "standard row",
			line: "| 1 | essay | my-slug | My Title | A hook | Hidden math |",
			want: []string{"1", "essay", "my-slug", "My Title", "A hook", "Hidden math"},
		},
		{
			name: "extra whitespace",
			line: "|  2  |  section  |  sec-slug  |  Sec Title  |  hook  |  math  |",
			want: []string{"2", "section", "sec-slug", "Sec Title", "hook", "math"},
		},
		{
			name: "with arc and ending columns",
			line: "| 3 | essay | slug | Title | Hook | Math | rise | happy |",
			want: []string{"3", "essay", "slug", "Title", "Hook", "Math", "rise", "happy"},
		},
		{
			name: "dash values",
			line: "| — | essay | slug | Title | Hook | Math | - | - |",
			want: []string{"—", "essay", "slug", "Title", "Hook", "Math", "-", "-"},
		},
		{
			name: "empty columns",
			line: "| 1 | essay | slug | Title | | |",
			want: []string{"1", "essay", "slug", "Title", "", ""},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := splitTableRow(tt.line)
			if len(got) != len(tt.want) {
				t.Fatalf("splitTableRow(%q) returned %d columns, want %d\ngot:  %v\nwant: %v",
					tt.line, len(got), len(tt.want), got, tt.want)
			}
			for i := range got {
				if got[i] != tt.want[i] {
					t.Errorf("column %d: got %q, want %q", i, got[i], tt.want[i])
				}
			}
		})
	}
}

func TestToSlug(t *testing.T) {
	t.Parallel()

	tests := []struct {
		input string
		want  string
	}{
		{"My Title", "my-title"},
		{"  Leading Spaces  ", "leading-spaces"},
		{"UPPER CASE", "upper-case"},
		{"already-slug", "already-slug"},
		{"Single", "single"},
		{"Multiple   Spaces", "multiple---spaces"},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			t.Parallel()
			got := toSlug(tt.input)
			if got != tt.want {
				t.Errorf("toSlug(%q) = %q, want %q", tt.input, got, tt.want)
			}
		})
	}
}

func TestParsePlanFile(t *testing.T) {
	t.Parallel()

	content := `# Plan for Test Book

## Book Introduction

| # | Type | Slug | Title | Hook | Hidden Theme |
|---|------|------|-------|------|--------------|
| — | introduction | test-intro | Test Intro | Intro hook | Intro math |

## Part 1: The Beginning

| # | Type | Slug | Title | Hook | Hidden Theme |
|---|------|------|-------|------|--------------|
| 1 | essay | first-essay | First Essay | First hook | First math |
| 2 | essay | second-essay | Second Essay | Second hook | Second math |

## Part 2: The Middle

| # | Type | Slug | Title | Hook | Hidden Theme | Arc | Ending |
|---|------|------|-------|------|--------------|-----|--------|
| 1 | section | mid-section | Mid Section | Section hook | Section math | rise | happy |
| 2 | essay | third-essay | Third Essay | Third hook | Third math | - | - |
`

	dir := t.TempDir()
	planPath := filepath.Join(dir, "Plan for Test.md")
	if err := os.WriteFile(planPath, []byte(content), 0644); err != nil {
		t.Fatalf("writing test plan: %v", err)
	}

	items, err := parsePlanFile(planPath, "I")
	if err != nil {
		t.Fatalf("parsePlanFile: %v", err)
	}

	if len(items) != 5 {
		t.Fatalf("expected 5 items, got %d", len(items))
	}

	// Check introduction
	intro := items[0]
	if intro.Type != "introduction" {
		t.Errorf("item 0 type: got %q, want %q", intro.Type, "introduction")
	}
	if intro.Slug != "test-intro" {
		t.Errorf("item 0 slug: got %q, want %q", intro.Slug, "test-intro")
	}
	if intro.Part != 0 {
		t.Errorf("item 0 part: got %d, want 0", intro.Part)
	}
	if intro.PartTitle != "Book Introduction" {
		t.Errorf("item 0 part title: got %q, want %q", intro.PartTitle, "Book Introduction")
	}
	if intro.Book != "I" {
		t.Errorf("item 0 book: got %q, want %q", intro.Book, "I")
	}

	// Check first essay in Part 1
	first := items[1]
	if first.Type != "essay" {
		t.Errorf("item 1 type: got %q, want %q", first.Type, "essay")
	}
	if first.Slug != "first-essay" {
		t.Errorf("item 1 slug: got %q, want %q", first.Slug, "first-essay")
	}
	if first.Part != 1 {
		t.Errorf("item 1 part: got %d, want 1", first.Part)
	}
	if first.PartTitle != "The Beginning" {
		t.Errorf("item 1 part title: got %q, want %q", first.PartTitle, "The Beginning")
	}
	if first.Order != 1 {
		t.Errorf("item 1 order: got %d, want 1", first.Order)
	}
	if first.Hook != "First hook" {
		t.Errorf("item 1 hook: got %q, want %q", first.Hook, "First hook")
	}
	if first.HiddenMath != "First math" {
		t.Errorf("item 1 hidden math: got %q, want %q", first.HiddenMath, "First math")
	}

	// Check second essay in Part 1
	second := items[2]
	if second.Order != 2 {
		t.Errorf("item 2 order: got %d, want 2", second.Order)
	}

	// Check section in Part 2 with arc and ending
	section := items[3]
	if section.Type != "section" {
		t.Errorf("item 3 type: got %q, want %q", section.Type, "section")
	}
	if section.Part != 2 {
		t.Errorf("item 3 part: got %d, want 2", section.Part)
	}
	if section.Arc != "rise" {
		t.Errorf("item 3 arc: got %q, want %q", section.Arc, "rise")
	}
	if section.Ending != "happy" {
		t.Errorf("item 3 ending: got %q, want %q", section.Ending, "happy")
	}

	// Check dash arc/ending normalized to empty
	thirdEssay := items[4]
	if thirdEssay.Arc != "" {
		t.Errorf("item 4 arc: got %q, want empty (dash should be normalized)", thirdEssay.Arc)
	}
	if thirdEssay.Ending != "" {
		t.Errorf("item 4 ending: got %q, want empty (dash should be normalized)", thirdEssay.Ending)
	}
}

func TestParsePlanFileSkipsEmptySlug(t *testing.T) {
	t.Parallel()

	content := `## Part 1: Test

| # | Type | Slug | Title | Hook | Hidden Theme |
|---|------|------|-------|------|--------------|
| 1 | essay |  | No Slug | hook | math |
| 2 |  | no-type | No Type | hook | math |
| 3 | essay | valid | Valid | hook | math |
`

	dir := t.TempDir()
	planPath := filepath.Join(dir, "plan.md")
	if err := os.WriteFile(planPath, []byte(content), 0644); err != nil {
		t.Fatalf("writing test plan: %v", err)
	}

	items, err := parsePlanFile(planPath, "I")
	if err != nil {
		t.Fatalf("parsePlanFile: %v", err)
	}

	if len(items) != 1 {
		t.Fatalf("expected 1 item (skipping empty slug/type), got %d", len(items))
	}
	if items[0].Slug != "valid" {
		t.Errorf("expected slug %q, got %q", "valid", items[0].Slug)
	}
}

func TestParsePlanFileDashOrder(t *testing.T) {
	t.Parallel()

	content := `## Part 1: Test

| # | Type | Slug | Title | Hook | Hidden Theme |
|---|------|------|-------|------|--------------|
| — | essay | dash-order | Dash Order | hook | math |
`

	dir := t.TempDir()
	planPath := filepath.Join(dir, "plan.md")
	if err := os.WriteFile(planPath, []byte(content), 0644); err != nil {
		t.Fatalf("writing test plan: %v", err)
	}

	items, err := parsePlanFile(planPath, "II")
	if err != nil {
		t.Fatalf("parsePlanFile: %v", err)
	}

	if len(items) != 1 {
		t.Fatalf("expected 1 item, got %d", len(items))
	}
	if items[0].Order != 0 {
		t.Errorf("dash order should be 0, got %d", items[0].Order)
	}
}

func TestLoadAttributes(t *testing.T) {
	t.Parallel()

	yamlContent := `essays:
  - slug: test-essay
    arc: fall
    ending: ambiguous
    structure: three-act
    entry: medias-res
    register: formal
    setting: "urban"
    math_visibility: hidden
  - slug: another-essay
    arc: rise
    ending: happy
    structure: linear
    entry: opening
    register: casual
    setting: "rural"
    math_visibility: overt
`

	dir := t.TempDir()
	if err := os.WriteFile(filepath.Join(dir, "attributes.yaml"), []byte(yamlContent), 0644); err != nil {
		t.Fatalf("writing attributes: %v", err)
	}

	attrs, err := loadAttributes(dir)
	if err != nil {
		t.Fatalf("loadAttributes: %v", err)
	}

	if len(attrs) != 2 {
		t.Fatalf("expected 2 attributes, got %d", len(attrs))
	}

	te := attrs["test-essay"]
	if te.Arc != "fall" {
		t.Errorf("test-essay arc: got %q, want %q", te.Arc, "fall")
	}
	if te.Ending != "ambiguous" {
		t.Errorf("test-essay ending: got %q, want %q", te.Ending, "ambiguous")
	}
	if te.Structure != "three-act" {
		t.Errorf("test-essay structure: got %q, want %q", te.Structure, "three-act")
	}
	if te.Setting != "urban" {
		t.Errorf("test-essay setting: got %q, want %q", te.Setting, "urban")
	}

	ae := attrs["another-essay"]
	if ae.Arc != "rise" {
		t.Errorf("another-essay arc: got %q, want %q", ae.Arc, "rise")
	}
}

func TestLoadAttributesEmptyDir(t *testing.T) {
	t.Parallel()

	dir := t.TempDir()
	attrs, err := loadAttributes(dir)
	if err != nil {
		t.Fatalf("loadAttributes on empty dir: %v", err)
	}
	if len(attrs) != 0 {
		t.Errorf("expected 0 attributes, got %d", len(attrs))
	}
}

func TestDiscoverSeries(t *testing.T) {
	t.Parallel()

	dir := t.TempDir()

	// Create plan files for a multi-book series
	for _, name := range []string{
		"Plan for Test Books I.md",
		"Plan for Test Books II.md",
	} {
		if err := os.WriteFile(filepath.Join(dir, name), []byte("# plan\n"), 0644); err != nil {
			t.Fatalf("writing %s: %v", name, err)
		}
	}

	// Create a single plan file (no numeral)
	if err := os.WriteFile(filepath.Join(dir, "Plan for Solo.md"), []byte("# plan\n"), 0644); err != nil {
		t.Fatalf("writing solo plan: %v", err)
	}

	result, err := discoverSeries(dir)
	if err != nil {
		t.Fatalf("discoverSeries: %v", err)
	}

	if len(result) != 2 {
		t.Fatalf("expected 2 series, got %d", len(result))
	}

	// Results are sorted by slug
	solo := result[0]
	if solo.slug != "solo" {
		t.Errorf("series 0 slug: got %q, want %q", solo.slug, "solo")
	}
	if len(solo.plans) != 1 {
		t.Errorf("solo plans: got %d, want 1", len(solo.plans))
	}

	books := result[1]
	if books.slug != "test-books" {
		t.Errorf("series 1 slug: got %q, want %q", books.slug, "test-books")
	}
	if len(books.plans) != 2 {
		t.Errorf("test-books plans: got %d, want 2", len(books.plans))
	}
}
