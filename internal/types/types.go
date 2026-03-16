package types

type EssayMeta struct {
	Slug      string `yaml:"slug"`
	Title     string `yaml:"title"`
	Type      string `yaml:"type"`
	Order     int    `yaml:"order"`
	PartTitle string `yaml:"part_title"`
}

type EssayContent struct {
	Slug    string
	Title   string
	Order   int
	Typ     string
	Content string
}
