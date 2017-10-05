package krakend

type Path struct {
	Disabled bool   `default:"false" storm:"disabled" json:"disabled" yaml:"disabled" toml:"disabled"`
	Debug    bool   `default:"false" example:"false" storm:"debug" json:"debug" yaml:"debug" toml:"debug"`
	Name     string `json:"name" yaml:"name" toml:"name"`
	Path     string `json:"path" yaml:"path" toml:"path"`
}

type PathBlocks struct {
	Disabled bool `default:"false" storm:"disabled" json:"disabled" yaml:"disabled" toml:"disabled"`
	Debug    bool `default:"false" example:"false" storm:"debug" json:"debug" yaml:"debug" toml:"debug"`

	Article             string            `json:"article,omitempty" yaml:"article,omitempty" toml:"article,omitempty"`
	Author              string            `json:"author,omitempty" yaml:"author,omitempty" toml:"author,omitempty"`
	Content             string            `json:"content,omitempty" yaml:"content,omitempty" toml:"content,omitempty"`
	ContentMisc         string            `json:"content_misc,omitempty" yaml:"content_misc,omitempty" toml:"content_misc,omitempty"`
	RelatedLinks        string            `json:"related_links,omitempty" yaml:"related_links,omitempty" toml:"related_links,omitempty"`
	Results             string            `json:"results,omitempty" yaml:"results,omitempty" toml:"results,omitempty"`
	SpellingSuggestions string            `json:"spelling_suggestions,omitempty" yaml:"spelling_suggestions,omitempty" toml:"spelling_suggestions,omitempty"`
	Suggestions         string            `json:"suggestions,omitempty" yaml:"suggestions,omitempty" toml:"suggestions,omitempty"`
	Tags                string            `json:"tags,omitempty" yaml:"tags,omitempty" toml:"tags,omitempty"`
	Title               string            `json:"title,omitempty" yaml:"title,omitempty" toml:"title,omitempty"`
	Thumbnail           string            `json:"thumbnail,omitempty" yaml:"thumbnail,omitempty" toml:"thumbnail,omitempty"`
	URL                 string            `json:"url,omitempty" yaml:"url,omitempty" toml:"url,omitempty"`
	Extra               []PathExtraBlocks `json:"extra,omitempty" yaml:"extra,omitempty" toml:"extra,omitempty"`
}

type PathExtraBlocks struct {
	Disabled bool   `default:"false" storm:"disabled" json:"disabled" yaml:"disabled" toml:"disabled"`
	Debug    bool   `default:"false" example:"false" storm:"debug" json:"debug" yaml:"debug" toml:"debug"`
	Name     string `json:"name,omitempty" yaml:"name,omitempty" toml:"name,omitempty"`
	Path     string `json:"path,omitempty" yaml:"path,omitempty" toml:"path,omitempty"`
}
