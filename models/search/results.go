package search

type SearchResults struct {
	Answers             []interface{} `json:"answers,omitempty" yaml:"answers,omitempty" toml:"answers,omitempty"`
	Corrections         []interface{} `json:"corrections,omitempty" yaml:"corrections,omitempty" toml:"corrections,omitempty"`
	Infoboxes           []Infobox     `json:"infoboxes,omitempty" yaml:"infoboxes,omitempty" toml:"infoboxes,omitempty"`
	NumberOfResults     float64       `json:"number_of_results,omitempty" yaml:"number_of_results,omitempty" toml:"number_of_results,omitempty"`
	Query               string        `json:"query,omitempty" yaml:"query,omitempty" toml:"query,omitempty"`
	Results             []Result      `json:"results,omitempty" yaml:"results,omitempty" toml:"results,omitempty"`
	Suggestions         []string      `json:"suggestions,omitempty" yaml:"suggestions,omitempty" toml:"suggestions,omitempty"`
	UnresponsiveEngines []interface{} `json:"unresponsive_engines,omitempty" yaml:"unresponsive_engines,omitempty" toml:"unresponsive_engines,omitempty"`
}

type Result struct {
	Category      string   `json:"category,omitempty" yaml:"category,omitempty" toml:"category,omitempty"`
	Content       string   `json:"content,omitempty" yaml:"content,omitempty" toml:"content,omitempty"`
	Engine        string   `json:"engine,omitempty" yaml:"engine,omitempty" toml:"engine,omitempty"`
	Engines       []string `json:"engines,omitempty" yaml:"engines,omitempty" toml:"engines,omitempty"`
	ImgSrc        string   `json:"img_src,omitempty" yaml:"img_src,omitempty" toml:"img_src,omitempty"`
	ParsedURL     []string `json:"parsed_url,omitempty" yaml:"parsed_url,omitempty" toml:"parsed_url,omitempty"`
	Positions     []int    `json:"positions,omitempty" yaml:"positions,omitempty" toml:"positions,omitempty"`
	PrettyURL     string   `json:"pretty_url,omitempty" yaml:"pretty_url,omitempty" toml:"pretty_url,omitempty"`
	Pubdate       string   `json:"pubdate,omitempty" yaml:"pubdate,omitempty" toml:"pubdate,omitempty"`
	PublishedDate string   `json:"publishedDate,omitempty" yaml:"publishedDate,omitempty" toml:"publishedDate,omitempty"`
	Score         float64  `json:"score,omitempty" yaml:"score,omitempty" toml:"score,omitempty"`
	Template      string   `json:"template,omitempty" yaml:"template,omitempty" toml:"template,omitempty"`
	ThumbnailSrc  string   `json:"thumbnail_src,omitempty" yaml:"thumbnail_src,omitempty" toml:"thumbnail_src,omitempty"`
	Title         string   `json:"title,omitempty" yaml:"title,omitempty" toml:"title,omitempty"`
	URL           string   `json:"url,omitempty" yaml:"url,omitempty" toml:"url,omitempty"`
}
