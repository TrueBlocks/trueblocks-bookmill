package pipeline

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	"github.com/TrueBlocks/trueblocks-art/packages/ai"
	"gopkg.in/yaml.v3"
)

type Config struct {
	API        APIConfig                `yaml:"api"`
	Models     ModelsConfig             `yaml:"models"`
	Pricing    map[string]*ModelPricing `yaml:"pricing"`
	Pipeline   PipelineConfig           `yaml:"pipeline"`
	Dashboard  DashboardConfig          `yaml:"dashboard"`
	ExportYear string                   `yaml:"export_year"`
}

type APIConfig struct {
	AnthropicKey string `yaml:"anthropic_key"`
	OpenAIKey    string `yaml:"openai_key"`
	Version      string `yaml:"version"`
	MaxTokens    int    `yaml:"max_tokens"`
}

type ModelPricing struct {
	InputPer1M  float64 `yaml:"input_per_1m"`
	OutputPer1M float64 `yaml:"output_per_1m"`
}

type ModelsConfig struct {
	Research   string `yaml:"research"`
	Outline    string `yaml:"outline"`
	Draft      string `yaml:"draft"`
	Factcheck  string `yaml:"factcheck"`
	Draft2     string `yaml:"draft2"`
	Illustrate string `yaml:"illustrate"`
}

type PipelineConfig struct {
	MaxPerCycle       int     `yaml:"max_per_cycle"`
	NewPerCycle       int     `yaml:"new_per_cycle"`
	DryRun            bool    `yaml:"dry_run"`
	CycleInterval     int     `yaml:"cycle_interval"`
	APITimeout        int     `yaml:"api_timeout"`
	Concurrency       int     `yaml:"concurrency"`
	Verbose           bool    `yaml:"verbose"`
	ReadMean          float64 `yaml:"read_mean"`
	ReadSpread        float64 `yaml:"read_spread"`
	Debug             string  `yaml:"debug"`
	SkipRevertConfirm bool    `yaml:"skip_revert_confirm"`
}

type DashboardConfig struct {
	Port int `yaml:"port"`
}

func DefaultConfigPath() string {
	home, _ := os.UserHomeDir()
	return filepath.Join(home, ".local", "share", "trueblocks", "math", "config.yaml")
}

func LoadConfig(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("reading config %s: %w", path, err)
	}

	cfg := &Config{
		API: APIConfig{
			Version:   "2023-06-01",
			MaxTokens: 8192,
		},
		Models: ModelsConfig{
			Research:   "claude-sonnet-4-20250514",
			Outline:    "claude-sonnet-4-20250514",
			Draft:      "claude-sonnet-4-20250514",
			Factcheck:  "claude-sonnet-4-20250514",
			Draft2:     "claude-sonnet-4-20250514",
			Illustrate: "claude-sonnet-4-20250514",
		},
		Pricing: map[string]*ModelPricing{
			"opus":    {InputPer1M: 15.0, OutputPer1M: 75.0},
			"default": {InputPer1M: 3.0, OutputPer1M: 15.0},
		},
		Pipeline: PipelineConfig{
			MaxPerCycle:   6,
			NewPerCycle:   4,
			DryRun:        true,
			CycleInterval: 15,
			APITimeout:    300,
			Concurrency:   3,
			ReadMean:      5.0,
			ReadSpread:    1.0,
		},
		Dashboard: DashboardConfig{
			Port: 8787,
		},
		ExportYear: time.Now().Format("2006"),
	}

	if err := yaml.Unmarshal(data, cfg); err != nil {
		return nil, fmt.Errorf("parsing config %s: %w", path, err)
	}

	if cfg.API.AnthropicKey == "" || cfg.API.OpenAIKey == "" {
		shared, _ := ai.LoadSharedConfig()
		if cfg.API.AnthropicKey == "" {
			cfg.API.AnthropicKey = shared.AnthropicAPIKey
		}
		if cfg.API.OpenAIKey == "" {
			cfg.API.OpenAIKey = shared.OpenAIAPIKey
		}
	}

	return cfg, nil
}

func SaveConfig(path string, cfg *Config) error {
	raw, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("reading config for update: %w", err)
	}

	lines := strings.Split(string(raw), "\n")
	set := func(key, val string) {
		re := regexp.MustCompile(`(?i)^(\s*` + regexp.QuoteMeta(key) + `\s*:\s*)(\S.*?)(\s*#.*)?$`)
		found := false
		for i, line := range lines {
			if m := re.FindStringSubmatchIndex(line); m != nil {
				prefix := line[m[2]:m[3]]
				suffix := ""
				if m[6] >= 0 {
					suffix = line[m[6]:m[7]]
				}
				lines[i] = prefix + val + suffix
				found = true
			}
		}
		if !found {
			sectionRe := regexp.MustCompile(`(?i)^\s*pipeline\s*:`)
			for i, line := range lines {
				if sectionRe.MatchString(line) {
					newLine := "  " + key + ": " + val
					lines = append(lines[:i+1], append([]string{newLine}, lines[i+1:]...)...)
					break
				}
			}
		}
	}

	set("cycle_interval", fmt.Sprintf("%d", cfg.Pipeline.CycleInterval))
	set("verbose", fmt.Sprintf("%t", cfg.Pipeline.Verbose))
	set("read_mean", fmt.Sprintf("%.1f", cfg.Pipeline.ReadMean))
	set("read_spread", fmt.Sprintf("%.1f", cfg.Pipeline.ReadSpread))
	set("skip_revert_confirm", fmt.Sprintf("%t", cfg.Pipeline.SkipRevertConfirm))

	return os.WriteFile(path, []byte(strings.Join(lines, "\n")), 0644)
}
