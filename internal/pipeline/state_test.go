package pipeline

import "testing"

func TestStageString(t *testing.T) {
	t.Parallel()
	tests := []struct {
		stage Stage
		want  string
	}{
		{StageIdeas, "ideas"},
		{StageResearch, "research"},
		{StageOutline, "outline"},
		{StageDraft, "draft"},
		{StageFactcheck, "factcheck"},
		{StageIllustrate, "illustrate"},
		{StageDraft2, "draft2"},
		{StageExport, "export"},
		{StageDone, "done"},
		{Stage(99), "done"},
	}
	for _, tt := range tests {
		t.Run(tt.want, func(t *testing.T) {
			t.Parallel()
			if got := tt.stage.String(); got != tt.want {
				t.Errorf("Stage(%d).String() = %q, want %q", tt.stage, got, tt.want)
			}
		})
	}
}

func TestStageFromString(t *testing.T) {
	t.Parallel()
	for _, name := range stageNames {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			stage := StageFromString(name)
			if stage.String() != name {
				t.Errorf("StageFromString(%q).String() = %q", name, stage.String())
			}
		})
	}
	t.Run("unknown defaults to ideas", func(t *testing.T) {
		t.Parallel()
		if got := StageFromString("bogus"); got != StageIdeas {
			t.Errorf("StageFromString(bogus) = %d, want StageIdeas", got)
		}
	})
}

func TestNextStage(t *testing.T) {
	t.Parallel()
	tests := []struct {
		in   Stage
		want Stage
	}{
		{StageIdeas, StageResearch},
		{StageResearch, StageOutline},
		{StageDraft2, StageRevision},
		{StageRevision, StageExport},
		{StageExport, StageDone},
		{StageDone, StageDone},
	}
	for _, tt := range tests {
		t.Run(tt.in.String(), func(t *testing.T) {
			t.Parallel()
			if got := NextStage(tt.in); got != tt.want {
				t.Errorf("NextStage(%s) = %s, want %s", tt.in, got, tt.want)
			}
		})
	}
}

func TestEssayStateNextAction(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name  string
		essay EssayState
		want  Stage
	}{
		{"error retries current stage", EssayState{Status: "error", CurrentStage: StageDraft}, StageDraft},
		{"pending ideas goes to research", EssayState{Status: "pending", CurrentStage: StageIdeas}, StageResearch},
		{"final advances to next", EssayState{Status: "final", CurrentStage: StageOutline}, StageDraft},
		{"final at export is done", EssayState{Status: "final", CurrentStage: StageExport}, StageDone},
		{"in-progress is done", EssayState{Status: "in-progress", CurrentStage: StageDraft}, StageDone},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := tt.essay.NextAction(); got != tt.want {
				t.Errorf("NextAction() = %s, want %s", got, tt.want)
			}
		})
	}
}

func TestEssayStateIsDone(t *testing.T) {
	t.Parallel()
	done := EssayState{CurrentStage: StageExport, Status: "final"}
	if !done.IsDone() {
		t.Error("expected IsDone() = true for export+final")
	}
	notDone := EssayState{CurrentStage: StageDraft, Status: "final"}
	if notDone.IsDone() {
		t.Error("expected IsDone() = false for draft+final")
	}
}

func TestEssayStateIsAvailable(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name  string
		essay EssayState
		want  bool
	}{
		{"pending ideas is available", EssayState{Status: "pending", CurrentStage: StageIdeas, Meta: map[Stage]*EssayMeta{}}, true},
		{"in-progress not available", EssayState{Status: "in-progress", CurrentStage: StageDraft, Meta: map[Stage]*EssayMeta{}}, false},
		{"error under retry limit", EssayState{Status: "error", CurrentStage: StageDraft, ErrorRetries: 1, Meta: map[Stage]*EssayMeta{}}, true},
		{"error at retry limit", EssayState{Status: "error", CurrentStage: StageDraft, ErrorRetries: 3, Meta: map[Stage]*EssayMeta{}}, false},
		{"done not available", EssayState{Status: "final", CurrentStage: StageExport, Meta: map[Stage]*EssayMeta{}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := tt.essay.IsAvailable(); got != tt.want {
				t.Errorf("IsAvailable() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseDebug(t *testing.T) {
	t.Parallel()
	tests := []struct {
		input     string
		wantBook  string
		wantPart  int
		wantOrder int
		wantOk    bool
	}{
		{"mybook.2.5", "mybook", 2, 5, true},
		{"book.0.0", "book", 0, 0, true},
		{"too.few", "", 0, 0, false},
		{"a.b.c", "", 0, 0, false},
		{"book.1.notnum", "", 0, 0, false},
		{"", "", 0, 0, false},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			t.Parallel()
			book, part, order, ok := ParseDebug(tt.input)
			if ok != tt.wantOk {
				t.Fatalf("ParseDebug(%q) ok = %v, want %v", tt.input, ok, tt.wantOk)
			}
			if ok {
				if book != tt.wantBook || part != tt.wantPart || order != tt.wantOrder {
					t.Errorf("ParseDebug(%q) = (%q, %d, %d), want (%q, %d, %d)",
						tt.input, book, part, order, tt.wantBook, tt.wantPart, tt.wantOrder)
				}
			}
		})
	}
}
