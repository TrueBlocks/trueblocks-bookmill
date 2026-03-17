package bookutil

import (
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/TrueBlocks/trueblocks-math/internal/types"
	"gopkg.in/yaml.v3"
)

func FindPlan(projectDir string) string {
	project := filepath.Base(projectDir)
	base := filepath.Dir(filepath.Dir(projectDir))
	designDir := filepath.Join(base, "design", project)

	entries, err := os.ReadDir(designDir)
	if err != nil {
		return ""
	}
	for _, e := range entries {
		lower := strings.ToLower(e.Name())
		if strings.Contains(lower, "plan") && strings.HasSuffix(lower, ".md") {
			data, err := os.ReadFile(filepath.Join(designDir, e.Name()))
			if err != nil {
				continue
			}
			return string(data)
		}
	}
	return ""
}

func ReadDraft2(projectDir string, maxContentLen int) []types.EssayContent {
	draft2Dir := filepath.Join(projectDir, "draft2")
	entries, err := os.ReadDir(draft2Dir)
	if err != nil {
		return nil
	}

	var essays []types.EssayContent
	for _, e := range entries {
		if !strings.HasSuffix(e.Name(), ".md") || strings.HasSuffix(e.Name(), ".meta.yaml") {
			continue
		}
		slug := strings.TrimSuffix(e.Name(), ".md")

		content, err := os.ReadFile(filepath.Join(draft2Dir, e.Name()))
		if err != nil {
			continue
		}

		var meta types.EssayMeta
		metaPath := filepath.Join(draft2Dir, slug+".meta.yaml")
		if data, err := os.ReadFile(metaPath); err == nil {
			yaml.Unmarshal(data, &meta)
		}
		if meta.Title == "" {
			meta.Title = slug
		}

		text := string(content)
		if maxContentLen > 0 && len(text) > maxContentLen {
			text = text[:maxContentLen] + "..."
		}

		essays = append(essays, types.EssayContent{
			Slug:    slug,
			Title:   meta.Title,
			Order:   meta.Order,
			Typ:     meta.Type,
			Content: text,
		})
	}

	sort.Slice(essays, func(i, j int) bool {
		return essays[i].Order < essays[j].Order
	})
	return essays
}
