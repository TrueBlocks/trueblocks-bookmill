package prompts

import "testing"

func TestApplyTemplate(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		tmpl string
		vars map[string]string
		want string
	}{
		{"simple substitution", "Hello {{.Name}}", map[string]string{"Name": "World"}, "Hello World"},
		{"multiple vars", "{{.A}} and {{.B}}", map[string]string{"A": "one", "B": "two"}, "one and two"},
		{"missing var becomes empty", "Hello {{.Name}}", map[string]string{}, "Hello "},
		{"no vars", "plain text", map[string]string{}, "plain text"},
		{"empty template", "", map[string]string{"A": "val"}, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := ApplyTemplate(tt.tmpl, tt.vars)
			if got != tt.want {
				t.Errorf("ApplyTemplate() = %q, want %q", got, tt.want)
			}
		})
	}
}
