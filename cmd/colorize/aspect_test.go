package main

import (
	"image"
	"image/png"
	"os"
	"path/filepath"
	"testing"
)

func TestNearestGeminiAspect(t *testing.T) {
	for _, tc := range []struct {
		name          string
		width, height int
		want          string
	}{
		{"square", 100, 100, "1:1"},
		{"portrait 2:3", 100, 150, "2:3"},
		{"landscape 3:2", 150, 100, "3:2"},
		{"tall scan near 9:16", 116, 202, "9:16"},
		{"wide 16:9", 160, 90, "16:9"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			src := filepath.Join(t.TempDir(), "img.png")
			f, err := os.Create(src)
			if err != nil {
				t.Fatal(err)
			}
			if err := png.Encode(f, image.NewRGBA(image.Rect(0, 0, tc.width, tc.height))); err != nil {
				t.Fatal(err)
			}
			f.Close()
			if got := nearestGeminiAspect(src); got != tc.want {
				t.Errorf("nearestGeminiAspect(%dx%d) = %q, want %q", tc.width, tc.height, got, tc.want)
			}
		})
	}
	// A missing or unreadable source falls back to the default, never a panic.
	if got := nearestGeminiAspect(filepath.Join(t.TempDir(), "nope.png")); got != "2:3" {
		t.Errorf("missing source = %q, want the 2:3 fallback", got)
	}
}
