package main

import (
	"strings"
	"testing"
)

func TestColorizePromptDefaultAndNoSky(t *testing.T) {
	base := colorizePrompt("", false)
	if !strings.Contains(base, "Colorize this black and white engraving") {
		t.Error("default colorize prompt lost its instruction")
	}
	if !strings.Contains(base, "Keep all lines, details, and textures exactly as they are.") {
		t.Error("default colorize prompt lost its fidelity clause")
	}
	if strings.Contains(base, "does not contain sky") {
		t.Error("without noSky the prompt should not mention sky absence")
	}

	sky := colorizePrompt("", true)
	if !strings.Contains(sky, "This image does not contain sky.") {
		t.Error("noSky should append the sky-absence clause")
	}
}

func TestColorizePromptOverrideWins(t *testing.T) {
	got := colorizePrompt("just make it green", true)
	if got != "just make it green" {
		t.Errorf("an override should replace the whole prompt, got %q", got)
	}
}
