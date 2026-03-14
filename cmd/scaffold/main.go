package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"gopkg.in/yaml.v3"
)

type item struct {
	Type           string
	Slug           string
	Title          string
	Book           string
	Part           int
	PartTitle      string
	Order          int
	Hook           string
	HiddenMath     string
	Arc            string
	Ending         string
	Structure      string
	Entry          string
	Register       string
	Setting        string
	MathVisibility string
}

type attrEntry struct {
	Slug           string `yaml:"slug"`
	Arc            string `yaml:"arc"`
	Ending         string `yaml:"ending"`
	Structure      string `yaml:"structure"`
	Entry          string `yaml:"entry"`
	Register       string `yaml:"register"`
	Setting        string `yaml:"setting"`
	MathVisibility string `yaml:"math_visibility"`
}

type attrFile struct {
	Essays []attrEntry `yaml:"essays"`
}

func loadAttributes(designDir string) (map[string]attrEntry, error) {
	lookup := make(map[string]attrEntry)
	err := filepath.WalkDir(designDir, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() || !strings.HasSuffix(d.Name(), "attributes.yaml") {
			return nil
		}
		data, err := os.ReadFile(path)
		if err != nil {
			return fmt.Errorf("reading %s: %w", path, err)
		}
		var af attrFile
		if err := yaml.Unmarshal(data, &af); err != nil {
			return fmt.Errorf("parsing %s: %w", path, err)
		}
		for _, e := range af.Essays {
			lookup[e.Slug] = e
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return lookup, nil
}

func parsePlanFile(path string, bookNum string) ([]item, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("opening %s: %w", path, err)
	}
	defer f.Close()

	var items []item
	var currentPart int
	var currentPartTitle string
	partRe := regexp.MustCompile(`^## Part (\d+):\s*(.+)$`)
	introRe := regexp.MustCompile(`^## Book Introduction`)
	scanner := bufio.NewScanner(f)
	inTable := false
	headerSeen := false

	for scanner.Scan() {
		line := scanner.Text()

		if m := partRe.FindStringSubmatch(line); m != nil {
			p, _ := strconv.Atoi(m[1])
			currentPart = p
			currentPartTitle = strings.TrimSpace(m[2])
			inTable = false
			headerSeen = false
			continue
		}

		if introRe.MatchString(line) {
			currentPart = 0
			currentPartTitle = "Book Introduction"
			inTable = false
			headerSeen = false
			continue
		}

		if strings.HasPrefix(line, "| #") || strings.HasPrefix(line, "|-") || strings.HasPrefix(line, "| -") {
			if strings.HasPrefix(line, "| #") {
				inTable = true
				headerSeen = false
			} else if inTable {
				headerSeen = true
			}
			continue
		}

		if !inTable || !headerSeen {
			if !strings.HasPrefix(line, "|") {
				inTable = false
				headerSeen = false
			}
			continue
		}

		if !strings.HasPrefix(line, "|") {
			inTable = false
			headerSeen = false
			continue
		}

		cols := splitTableRow(line)
		if len(cols) < 6 {
			continue
		}

		num := strings.TrimSpace(cols[0])
		typ := strings.TrimSpace(cols[1])
		slug := strings.TrimSpace(cols[2])
		title := strings.TrimSpace(cols[3])
		hook := strings.TrimSpace(cols[4])
		hiddenMath := strings.TrimSpace(cols[5])
		arc := ""
		ending := ""
		if len(cols) > 6 {
			arc = strings.TrimSpace(cols[6])
		}
		if len(cols) > 7 {
			ending = strings.TrimSpace(cols[7])
		}

		if typ == "" || slug == "" {
			continue
		}

		order := 0
		if num != "—" && num != "-" {
			order, _ = strconv.Atoi(num)
		}

		if arc == "-" {
			arc = ""
		}
		if ending == "-" {
			ending = ""
		}

		it := item{
			Type:       typ,
			Slug:       slug,
			Title:      title,
			Book:       bookNum,
			Part:       currentPart,
			PartTitle:  currentPartTitle,
			Order:      order,
			Hook:       hook,
			HiddenMath: hiddenMath,
			Arc:        arc,
			Ending:     ending,
		}
		items = append(items, it)
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("scanning %s: %w", path, err)
	}

	return items, nil
}

func splitTableRow(line string) []string {
	line = strings.TrimSpace(line)
	line = strings.Trim(line, "|")
	parts := strings.Split(line, "|")
	for i := range parts {
		parts[i] = strings.TrimSpace(parts[i])
	}
	return parts
}

func toSlug(label string) string {
	return strings.ToLower(strings.ReplaceAll(strings.TrimSpace(label), " ", "-"))
}

type series struct {
	slug  string
	plans []struct {
		file string
		book string
	}
}

func discoverSeries(designDir string) ([]series, error) {
	romanNumerals := []string{"I", "II", "III", "IV", "V"}
	groups := map[string][]struct {
		file string
		book string
	}{}
	hasNumeral := map[string]bool{}

	err := filepath.WalkDir(designDir, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() || !strings.HasPrefix(d.Name(), "Plan for ") || !strings.HasSuffix(d.Name(), ".md") {
			return nil
		}
		relPath, err := filepath.Rel(designDir, path)
		if err != nil {
			return err
		}
		name := strings.TrimPrefix(d.Name(), "Plan for ")
		name = strings.TrimSuffix(name, ".md")
		name = strings.TrimSpace(name)

		base := name
		book := ""
		for _, r := range romanNumerals {
			suffix := " " + r
			if strings.HasSuffix(name, suffix) {
				base = strings.TrimSuffix(name, suffix)
				book = r
				hasNumeral[base] = true
				break
			}
		}

		groups[base] = append(groups[base], struct {
			file string
			book string
		}{file: relPath, book: book})
		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("walking %s: %w", designDir, err)
	}

	var result []series
	for base, plans := range groups {
		trimmed := base
		trimmed = strings.TrimSuffix(trimmed, " Books")
		trimmed = strings.TrimSuffix(trimmed, " books")
		trimmed = strings.TrimSuffix(trimmed, " Book")
		trimmed = strings.TrimSuffix(trimmed, " book")
		slug := toSlug(trimmed)
		if hasNumeral[base] {
			slug += "-books"
		}
		sort.Slice(plans, func(i, j int) bool {
			return plans[i].file < plans[j].file
		})
		result = append(result, series{slug: slug, plans: plans})
	}
	sort.Slice(result, func(i, j int) bool {
		return result[i].slug < result[j].slug
	})
	return result, nil
}

func loadAllItems(designDir string, plans []struct {
	file string
	book string
}) ([]item, error) {
	var all []item
	for _, p := range plans {
		path := filepath.Join(designDir, p.file)
		items, err := parsePlanFile(path, p.book)
		if err != nil {
			return nil, err
		}
		all = append(all, items...)
	}
	return all, nil
}

func main() {
	designDir := "design"
	if len(os.Args) > 1 {
		designDir = os.Args[1]
	}

	allSeries, err := discoverSeries(designDir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error discovering plans: %v\n", err)
		os.Exit(1)
	}
	if len(allSeries) == 0 {
		fmt.Fprintf(os.Stderr, "No 'Plan for ...' files found in %s/\n", designDir)
		os.Exit(1)
	}

	attrs, err := loadAttributes(designDir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading attributes: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Loaded %d attribute entries\n", len(attrs))

	for _, s := range allSeries {
		items, err := loadAllItems(designDir, s.plans)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error loading plans for %s: %v\n", s.slug, err)
			os.Exit(1)
		}

		baseDir := filepath.Join("projects", s.slug, "ideas")
		if err := os.MkdirAll(baseDir, 0755); err != nil {
			fmt.Fprintf(os.Stderr, "Error creating directory: %v\n", err)
			os.Exit(1)
		}

		fmt.Printf("Series: %s (%d plan files)\n", s.slug, len(s.plans))

		for i := range items {
			if a, ok := attrs[items[i].Slug]; ok {
				if items[i].Arc == "" {
					items[i].Arc = a.Arc
				}
				if items[i].Ending == "" {
					items[i].Ending = a.Ending
				}
				items[i].Structure = a.Structure
				items[i].Entry = a.Entry
				items[i].Register = a.Register
				items[i].Setting = a.Setting
				items[i].MathVisibility = a.MathVisibility
			}
		}

		for _, it := range items {
			var ideaContent string
			switch it.Type {
			case "section":
				ideaContent = fmt.Sprintf("# %s\n\n"+
					"**Slug:** %s\n"+
					"**Type:** section\n\n"+
					"## Part Title Page\n\n"+
					"This is a section divider for Part %d: %s.\n\n"+
					"## Placement\n\n"+
					"- **Book:** %s\n"+
					"- **Part %d:** %s\n",
					it.Title, it.Slug, it.Part, it.PartTitle, it.Book, it.Part, it.PartTitle)
			case "introduction":
				ideaContent = fmt.Sprintf("# %s\n\n"+
					"**Slug:** %s\n"+
					"**Type:** introduction\n\n"+
					"## Book Introduction\n\n"+
					"%s\n\n"+
					"## The Hidden Math\n\n"+
					"%s\n\n"+
					"## Placement\n\n"+
					"- **Book:** %s\n",
					it.Title, it.Slug, it.Hook, it.HiddenMath, it.Book)
			default:
				ideaContent = fmt.Sprintf("# %s\n\n"+
					"**Slug:** %s\n"+
					"**Type:** essay\n\n"+
					"## The Everyday Experience\n\n"+
					"%s\n\n"+
					"## The Hidden Math\n\n"+
					"%s\n\n"+
					"## Placement\n\n"+
					"- **Book:** %s\n"+
					"- **Part %d:** %s\n"+
					"- **Order:** %d\n",
					it.Title, it.Slug, it.Hook, it.HiddenMath, it.Book, it.Part, it.PartTitle, it.Order)
			}

			metaContent := fmt.Sprintf("slug: %s\n"+
				"title: \"%s\"\n"+
				"type: %s\n"+
				"series: %s\n"+
				"book: %s\n"+
				"part: %d\n"+
				"part_title: \"%s\"\n"+
				"order: %d\n"+
				"status: pending\n"+
				"model: claude-sonnet-4-20250514\n"+
				"arc: %s\n"+
				"ending: %s\n"+
				"structure: %s\n"+
				"entry: %s\n"+
				"register: %s\n"+
				"setting: \"%s\"\n"+
				"math_visibility: %s\n",
				it.Slug, it.Title, it.Type, s.slug, it.Book, it.Part, it.PartTitle, it.Order,
				it.Arc, it.Ending, it.Structure, it.Entry, it.Register, it.Setting, it.MathVisibility)

			mdPath := filepath.Join(baseDir, it.Slug+".md")
			metaPath := filepath.Join(baseDir, it.Slug+".meta.yaml")

			if err := os.WriteFile(mdPath, []byte(ideaContent), 0644); err != nil {
				fmt.Fprintf(os.Stderr, "Error writing %s: %v\n", mdPath, err)
				os.Exit(1)
			}

			if err := os.WriteFile(metaPath, []byte(metaContent), 0644); err != nil {
				fmt.Fprintf(os.Stderr, "Error writing %s: %v\n", metaPath, err)
				os.Exit(1)
			}

			fmt.Printf("Created: %s [%s] (.md + .meta.yaml)\n", it.Slug, it.Type)
		}

		fmt.Printf("\nTotal: %d items scaffolded in %s/\n", len(items), baseDir)

		typeCounts := map[string]map[string]int{}
		for _, it := range items {
			if typeCounts[it.Book] == nil {
				typeCounts[it.Book] = map[string]int{}
			}
			typeCounts[it.Book][it.Type]++
		}
		for _, b := range []string{"I", "II", "III"} {
			c := typeCounts[b]
			if c == nil {
				continue
			}
			fmt.Printf("  Book %s: %d essays, %d sections, %d introductions\n",
				b, c["essay"], c["section"], c["introduction"])
		}
	}
}
