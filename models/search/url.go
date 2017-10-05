package search

type URL struct {
	Title string `json:"title,omitempty" yaml:"title,omitempty" toml:"title,omitempty"`
	URL   string `json:"url,omitempty" yaml:"url,omitempty" toml:"url,omitempty"`
}
