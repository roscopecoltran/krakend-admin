package search

type Infobox struct {
	Attributes []Attribute `json:"attributes,omitempty" yaml:"attributes,omitempty" toml:"attributes,omitempty"`
	Content    string      `json:"content,omitempty" yaml:"content,omitempty" toml:"content,omitempty"`
	Engine     string      `json:"engine,omitempty" yaml:"engine,omitempty" toml:"engine,omitempty"`
	ID         string      `json:"id,omitempty" yaml:"id,omitempty" toml:"id,omitempty"`
	ImgSrc     interface{} `json:"img_src,omitempty" yaml:"img_src,omitempty" toml:"img_src,omitempty"`
	Infobox    string      `json:"infobox,omitempty" yaml:"infobox,omitempty" toml:"infobox,omitempty"`
	Urls       []URL       `json:"urls,omitempty" yaml:"urls,omitempty" toml:"urls,omitempty"`
}
