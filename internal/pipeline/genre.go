package pipeline

import (
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

type Genre struct {
	Form   string   `yaml:"form"`
	Flavor string   `yaml:"flavor"`
	Stages []string `yaml:"stages"`
}

var defaultEssayGenre = Genre{
	Form:   "essay",
	Flavor: "math",
	Stages: []string{"ideas", "research", "outline", "draft", "factcheck", "illustrate", "draft2", "export"},
}

func LoadGenre(designDir string) (*Genre, error) {
	path := filepath.Join(designDir, "genre.yaml")
	data, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			g := defaultEssayGenre
			return &g, nil
		}
		return nil, fmt.Errorf("reading genre.yaml: %w", err)
	}
	var g Genre
	if err := yaml.Unmarshal(data, &g); err != nil {
		return nil, fmt.Errorf("parsing genre.yaml: %w", err)
	}
	if g.Form == "" {
		g.Form = "essay"
	}
	if g.Flavor == "" {
		g.Flavor = "math"
	}
	if len(g.Stages) == 0 {
		g.Stages = defaultEssayGenre.Stages
	}
	return &g, nil
}

func LoadGenreFromProject(projectDir string) (*Genre, error) {
	path := filepath.Join(projectDir, "genre.yaml")
	data, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			g := defaultEssayGenre
			return &g, nil
		}
		return nil, fmt.Errorf("reading genre.yaml: %w", err)
	}
	var g Genre
	if err := yaml.Unmarshal(data, &g); err != nil {
		return nil, fmt.Errorf("parsing genre.yaml: %w", err)
	}
	if g.Form == "" {
		g.Form = "essay"
	}
	if g.Flavor == "" {
		g.Flavor = "math"
	}
	if len(g.Stages) == 0 {
		g.Stages = defaultEssayGenre.Stages
	}
	return &g, nil
}

func (g *Genre) HasStage(name string) bool {
	for _, s := range g.Stages {
		if s == name {
			return true
		}
	}
	return false
}

func (g *Genre) NextStageAfter(current string) string {
	for i, s := range g.Stages {
		if s == current && i+1 < len(g.Stages) {
			return g.Stages[i+1]
		}
	}
	return "done"
}

func (g *Genre) FirstContentStage() string {
	if len(g.Stages) > 1 {
		return g.Stages[1]
	}
	return "done"
}

func (g *Genre) IsNovel() bool {
	return g.Form == "novel"
}
