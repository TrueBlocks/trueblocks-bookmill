package pipeline

import "testing"

func TestShouldSkipStage(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name      string
		itemType  string
		stage     Stage
		want      bool
	}{
		{"section skips research", "section", StageResearch, true},
		{"section skips outline", "section", StageOutline, true},
		{"section keeps draft2", "section", StageDraft2, false},
		{"section keeps export", "section", StageExport, false},
		{"intro skips research", "introduction", StageResearch, true},
		{"intro skips factcheck", "introduction", StageFactcheck, true},
		{"intro skips illustrate", "introduction", StageIllustrate, true},
		{"intro keeps outline", "introduction", StageOutline, false},
		{"intro keeps draft", "introduction", StageDraft, false},
		{"chapter skips nothing", "chapter", StageResearch, false},
		{"chapter skips nothing 2", "chapter", StageFactcheck, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := shouldSkipStage(tt.itemType, tt.stage); got != tt.want {
				t.Errorf("shouldSkipStage(%q, %s) = %v, want %v", tt.itemType, tt.stage, got, tt.want)
			}
		})
	}
}
