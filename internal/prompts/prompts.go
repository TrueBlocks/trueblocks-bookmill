package prompts

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

func LoadTemplate(projectDir, filename string) (string, error) {
	base := filepath.Dir(filepath.Dir(projectDir))
	path := filepath.Join(base, "specs", "prompts", "generic", filename)
	data, err := os.ReadFile(path)
	if err != nil {
		return "", fmt.Errorf("could not load prompt template %s: %w", path, err)
	}
	return string(data), nil
}

func ApplyTemplate(tmpl string, vars map[string]string) string {
	t, err := template.New("prompt").Option("missingkey=zero").Parse(tmpl)
	if err != nil {
		for k, v := range vars {
			tmpl = strings.ReplaceAll(tmpl, "{{."+k+"}}", v)
		}
		return tmpl
	}
	var buf strings.Builder
	if err := t.Execute(&buf, vars); err != nil {
		for k, v := range vars {
			tmpl = strings.ReplaceAll(tmpl, "{{."+k+"}}", v)
		}
		return tmpl
	}
	return buf.String()
}
