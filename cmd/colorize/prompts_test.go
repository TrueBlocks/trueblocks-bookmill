package main

import (
	"strings"
	"testing"
)

func TestColorizePromptDefaultAndNoSky(t *testing.T) {
	t.Setenv("TRUEBLOCKS_DATA_DIR", t.TempDir())
	base, err := colorizePrompt("", false)
	if err != nil {
		t.Fatalf("rendering the default colorize prompt: %v", err)
	}
	if !strings.Contains(base, "Colorize this black and white engraving") {
		t.Error("default colorize prompt lost its instruction")
	}
	if !strings.Contains(base, "Keep all lines, details, and textures exactly as they are.") {
		t.Error("default colorize prompt lost its fidelity clause")
	}
	if strings.Contains(base, "does not contain sky") {
		t.Error("without noSky the prompt should not mention sky absence")
	}

	sky, err := colorizePrompt("", true)
	if err != nil {
		t.Fatalf("rendering the noSky colorize prompt: %v", err)
	}
	if !strings.Contains(sky, "This image does not contain sky.") {
		t.Error("noSky should append the sky-absence clause")
	}
}

func TestColorizePromptOverrideWins(t *testing.T) {
	got, err := colorizePrompt("just make it green", true)
	if err != nil {
		t.Fatalf("override colorize prompt: %v", err)
	}
	if got != "just make it green" {
		t.Errorf("an override should replace the whole prompt, got %q", got)
	}
}
