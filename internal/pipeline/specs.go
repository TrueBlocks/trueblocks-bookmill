package pipeline

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

func (r *Runner) loadSpecs() {
	seriesNames := make(map[string]bool)
	for _, ps := range r.Projects {
		seriesNames[ps.Project] = true
	}

	for series := range seriesNames {
		specs := r.loadSeriesSpecs(series)
		r.SeriesSpecs[series] = specs
	}
}

func (r *Runner) resolveSpecFile(series, filename string) ([]byte, error) {
	seriesPath := filepath.Join(r.BaseDir, "specs", "prompts", "series", series, filename)
	if data, err := os.ReadFile(seriesPath); err == nil {
		return data, nil
	}
	genericPath := filepath.Join(r.BaseDir, "specs", "prompts", "generic", filename)
	return os.ReadFile(genericPath)
}

func (r *Runner) loadSeriesSpecs(series string) *seriesSpecs {
	specs := &seriesSpecs{
		PromptTemplates: make(map[string]*template.Template),
		Examples:        make(map[string]string),
	}

	data, err := r.resolveSpecFile(series, "voice-summary.md")
	if err != nil {
		r.Log.Printf("[%s] Voice summary load failed: %v", series, err)
	} else {
		full := string(data)
		specs.VoiceProfile = full

		const antiHeader = "## What This Voice Does NOT Do"
		idx := strings.Index(full, antiHeader)
		if idx >= 0 {
			section := full[idx:]
			if end := strings.Index(section[len(antiHeader):], "\n---"); end >= 0 {
				section = section[:len(antiHeader)+end]
			}
			specs.VoiceAntiPatterns = strings.TrimSpace(section)
		}
	}

	data, err = r.resolveSpecFile(series, "essay-rules.md")
	if err != nil {
		r.Log.Printf("[%s] Essay rules load failed: %v", series, err)
	} else {
		full := string(data)
		specs.DraftRules = extractSection(full, "## Draft Guidelines")
		specs.RevisionRules = extractSection(full, "## Revision Rules")
	}

	seriesPromptDir := filepath.Join(r.BaseDir, "specs", "prompts", "series", series)
	genericPromptDir := filepath.Join(r.BaseDir, "specs", "prompts", "generic")

	seen := make(map[string]bool)
	for _, dir := range []string{seriesPromptDir, genericPromptDir} {
		entries, err := os.ReadDir(dir)
		if err != nil {
			continue
		}
		for _, e := range entries {
			if e.IsDir() || filepath.Ext(e.Name()) != ".md" {
				continue
			}
			name := strings.TrimSuffix(e.Name(), ".md")
			if name == "voice-summary" || name == "essay-rules" {
				continue
			}
			if seen[name] {
				continue
			}
			content, err := os.ReadFile(filepath.Join(dir, e.Name()))
			if err != nil {
				r.Log.Printf("[%s] WARNING: could not read prompt template '%s': %v", series, name, err)
				continue
			}
			tmpl, err := template.New(name).Parse(string(content))
			if err != nil {
				r.Log.Printf("[%s] WARNING: could not parse prompt template '%s': %v", series, name, err)
				continue
			}
			specs.PromptTemplates[name] = tmpl
			seen[name] = true
		}
	}

	exampleDir := filepath.Join(r.BaseDir, "specs", "examples")
	exFiles, err := os.ReadDir(exampleDir)
	if err != nil {
		r.Log.Printf("[%s] Examples dir load failed: %v", series, err)
	} else {
		for _, f := range exFiles {
			if f.IsDir() || filepath.Ext(f.Name()) != ".md" {
				continue
			}
			category := strings.TrimSuffix(f.Name(), ".md")
			content, err := os.ReadFile(filepath.Join(exampleDir, f.Name()))
			if err != nil {
				r.Log.Printf("[%s] WARNING: could not read examples '%s': %v", series, category, err)
				continue
			}
			for k, v := range splitExamples(category, string(content)) {
				specs.Examples[k] = v
			}
		}
	}

	r.Log.Printf("[%s] Loaded specs (prompts: %d, examples: %d)", series, len(specs.PromptTemplates), len(specs.Examples))
	return specs
}

func (r *Runner) specsFor(series string) *seriesSpecs {
	if s, ok := r.SeriesSpecs[series]; ok {
		return s
	}
	r.Log.Printf("WARNING: no specs loaded for series '%s', loading on demand", series)
	s := r.loadSeriesSpecs(series)
	r.SeriesSpecs[series] = s
	return s
}

func (r *Runner) executePrompt(series, name string, data map[string]any) string {
	specs := r.specsFor(series)
	tmpl, ok := specs.PromptTemplates[name]
	if !ok {
		r.Log.Printf("WARNING: prompt template '%s' not found for series '%s'", name, series)
		return ""
	}
	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		r.Log.Printf("ERROR executing prompt template '%s': %v", name, err)
		return ""
	}
	return buf.String()
}

func (r *Runner) buildExamples(series string, keys ...string) string {
	specs := r.specsFor(series)
	var parts []string
	for _, key := range keys {
		if ex, ok := specs.Examples[key]; ok && ex != "" {
			parts = append(parts, ex)
		}
	}
	if len(parts) == 0 {
		return ""
	}
	return "\n## REFERENCE EXAMPLES (for this essay's attributes)\n\nStudy these examples. They show the level of craft expected for this specific combination of attributes. Adapt the techniques — do not copy.\n\n" + strings.Join(parts, "\n\n---\n\n") + "\n"
}

func extractSection(content, header string) string {
	idx := strings.Index(content, header)
	if idx < 0 {
		return ""
	}
	section := content[idx+len(header):]
	if end := strings.Index(section, "\n---"); end >= 0 {
		section = section[:end]
	}
	return strings.TrimSpace(section)
}

func splitExamples(category, content string) map[string]string {
	result := make(map[string]string)
	sections := strings.Split(content, "\n## ")
	for i, sec := range sections {
		if i == 0 {
			continue
		}
		newline := strings.IndexByte(sec, '\n')
		if newline < 0 {
			continue
		}
		name := strings.TrimSpace(sec[:newline])
		body := strings.TrimSpace(sec[newline+1:])
		if body != "" {
			result[category+"/"+name] = body
		}
	}
	return result
}
