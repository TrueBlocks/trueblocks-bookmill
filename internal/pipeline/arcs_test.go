package pipeline

import "testing"

func TestArcByName_Known(t *testing.T) {
	arc, ok := ArcByName("slow-build")
	if !ok {
		t.Fatal("expected slow-build to be found")
	}
	if arc.Name != "slow-build" {
		t.Errorf("Name = %q, want slow-build", arc.Name)
	}
	if arc.Label == "" || arc.OutlineHint == "" || arc.DraftHint == "" {
		t.Error("expected non-empty Label/OutlineHint/DraftHint")
	}
}

func TestArcByName_Unknown(t *testing.T) {
	if _, ok := ArcByName("does-not-exist"); ok {
		t.Error("expected lookup of unknown arc to fail")
	}
}

func TestArcByName_AllRegisteredArcsResolvable(t *testing.T) {
	if len(narrativeArcs) == 0 {
		t.Fatal("narrativeArcs registry is empty")
	}
	seen := make(map[string]bool, len(narrativeArcs))
	for _, arc := range narrativeArcs {
		if arc.Name == "" {
			t.Errorf("arc with empty Name: %+v", arc)
			continue
		}
		if seen[arc.Name] {
			t.Errorf("duplicate arc name: %q", arc.Name)
		}
		seen[arc.Name] = true

		got, ok := ArcByName(arc.Name)
		if !ok {
			t.Errorf("ArcByName(%q) returned ok=false", arc.Name)
			continue
		}
		if got.Name != arc.Name {
			t.Errorf("ArcByName(%q).Name = %q", arc.Name, got.Name)
		}
		if got.Label == "" {
			t.Errorf("arc %q missing Label", arc.Name)
		}
		if got.OutlineHint == "" {
			t.Errorf("arc %q missing OutlineHint", arc.Name)
		}
		if got.DraftHint == "" {
			t.Errorf("arc %q missing DraftHint", arc.Name)
		}
	}
}
