package main

import (
	"errors"
	"os"
	"path/filepath"
	"testing"
	"time"

	"gopkg.in/yaml.v3"
)

func TestSourcePathForMeta(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		dir      string
		meta     imageMeta
		wantPath string
	}{
		{
			name:     "mermaid method",
			dir:      "/images/slug",
			meta:     imageMeta{Filename: "diagram.png", Method: "mermaid"},
			wantPath: "/images/slug/diagram.mermaid",
		},
		{
			name:     "r method",
			dir:      "/images/slug",
			meta:     imageMeta{Filename: "chart.png", Method: "r"},
			wantPath: "/images/slug/chart.R",
		},
		{
			name:     "ai method",
			dir:      "/images/slug",
			meta:     imageMeta{Filename: "art.png", Method: "ai"},
			wantPath: "/images/slug/art.ai-prompt.txt",
		},
		{
			name:     "unknown method returns empty",
			dir:      "/images/slug",
			meta:     imageMeta{Filename: "thing.png", Method: "unknown"},
			wantPath: "",
		},
		{
			name:     "empty method returns empty",
			dir:      "/images/slug",
			meta:     imageMeta{Filename: "thing.png", Method: ""},
			wantPath: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := sourcePathForMeta(tt.dir, tt.meta)
			if got != tt.wantPath {
				t.Errorf("sourcePathForMeta(%q, %+v) = %q, want %q",
					tt.dir, tt.meta, got, tt.wantPath)
			}
		})
	}
}

func TestIsNewer(t *testing.T) {
	t.Parallel()

	dir := t.TempDir()
	older := filepath.Join(dir, "older.txt")
	newer := filepath.Join(dir, "newer.txt")

	if err := os.WriteFile(older, []byte("old"), 0644); err != nil {
		t.Fatal(err)
	}
	// Ensure different modification times
	time.Sleep(50 * time.Millisecond)
	if err := os.WriteFile(newer, []byte("new"), 0644); err != nil {
		t.Fatal(err)
	}

	if !isNewer(newer, older) {
		t.Error("expected newer file to be newer than older file")
	}
	if isNewer(older, newer) {
		t.Error("expected older file NOT to be newer than newer file")
	}
}

func TestIsNewerMissingTarget(t *testing.T) {
	t.Parallel()

	dir := t.TempDir()
	source := filepath.Join(dir, "source.txt")
	if err := os.WriteFile(source, []byte("src"), 0644); err != nil {
		t.Fatal(err)
	}

	missing := filepath.Join(dir, "nonexistent.txt")
	if isNewer(missing, source) {
		t.Error("missing target should not be newer")
	}
}

func TestIsNewerMissingSource(t *testing.T) {
	t.Parallel()

	dir := t.TempDir()
	target := filepath.Join(dir, "target.txt")
	if err := os.WriteFile(target, []byte("tgt"), 0644); err != nil {
		t.Fatal(err)
	}

	missing := filepath.Join(dir, "nonexistent.txt")
	if isNewer(target, missing) {
		t.Error("missing source should cause isNewer to return false")
	}
}

func TestIsSafetyViolation(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		err  error
		want bool
	}{
		{
			name: "nil error",
			err:  nil,
			want: false,
		},
		{
			name: "content policy violation",
			err:  errors.New("request failed: content_policy_violation"),
			want: true,
		},
		{
			name: "other error",
			err:  errors.New("connection timeout"),
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := isSafetyViolation(tt.err)
			if got != tt.want {
				t.Errorf("isSafetyViolation(%v) = %v, want %v", tt.err, got, tt.want)
			}
		})
	}
}

func TestImageMetaYAMLParsing(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		yaml    string
		want    imageMeta
		wantErr bool
	}{
		{
			name: "mermaid meta",
			yaml: `filename: flow-diagram.png
method: mermaid
description: A flow diagram showing the process
`,
			want: imageMeta{
				Filename:    "flow-diagram.png",
				Method:      "mermaid",
				Description: "A flow diagram showing the process",
			},
		},
		{
			name: "r meta",
			yaml: `filename: scatter-plot.png
method: r
description: Scatter plot of data
`,
			want: imageMeta{
				Filename:    "scatter-plot.png",
				Method:      "r",
				Description: "Scatter plot of data",
			},
		},
		{
			name: "ai meta",
			yaml: `filename: hero-image.png
method: ai
description: Hero image for the essay
`,
			want: imageMeta{
				Filename:    "hero-image.png",
				Method:      "ai",
				Description: "Hero image for the essay",
			},
		},
		{
			name: "minimal meta with only filename",
			yaml: `filename: bare.png
`,
			want: imageMeta{
				Filename: "bare.png",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			var got imageMeta
			err := yaml.Unmarshal([]byte(tt.yaml), &got)
			if (err != nil) != tt.wantErr {
				t.Fatalf("yaml.Unmarshal error = %v, wantErr %v", err, tt.wantErr)
			}
			if err != nil {
				return
			}
			if got != tt.want {
				t.Errorf("parsed meta = %+v, want %+v", got, tt.want)
			}
		})
	}
}

func TestImageMetaYAMLRoundTrip(t *testing.T) {
	t.Parallel()

	original := imageMeta{
		Filename:    "test-image.png",
		Method:      "mermaid",
		Description: "A test diagram",
	}

	data, err := yaml.Marshal(&original)
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}

	var decoded imageMeta
	if err := yaml.Unmarshal(data, &decoded); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}

	if decoded != original {
		t.Errorf("round-trip mismatch: got %+v, want %+v", decoded, original)
	}
}

func TestImageMetaFromFile(t *testing.T) {
	t.Parallel()

	dir := t.TempDir()
	content := `filename: diagram.png
method: mermaid
description: "A diagram with special chars: colons, quotes"
`
	metaPath := filepath.Join(dir, "diagram.meta.yaml")
	if err := os.WriteFile(metaPath, []byte(content), 0644); err != nil {
		t.Fatal(err)
	}

	data, err := os.ReadFile(metaPath)
	if err != nil {
		t.Fatal(err)
	}

	var meta imageMeta
	if err := yaml.Unmarshal(data, &meta); err != nil {
		t.Fatalf("parsing meta file: %v", err)
	}

	if meta.Filename != "diagram.png" {
		t.Errorf("filename: got %q, want %q", meta.Filename, "diagram.png")
	}
	if meta.Method != "mermaid" {
		t.Errorf("method: got %q, want %q", meta.Method, "mermaid")
	}
}
