package search

type Attribute struct {
	Label string `json:"label,omitempty" yaml:"label,omitempty" toml:"label,omitempty"`
	Value string `json:"value,omitempty" yaml:"value,omitempty" toml:"value,omitempty"`
}
