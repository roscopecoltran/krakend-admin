package local

import (
	"github.com/jinzhu/gorm"
)

type LocalLanguages struct {
	gorm.Model `json:"-" yaml:"-" toml:"-"`

	Debug    bool `gorm:"column:debug" default:"false" storm:"debug" json:"debug" yaml:"debug" toml:"debug"`
	Disabled bool `gorm:"column:disabled" default:"false" storm:"disabled" json:"disabled" yaml:"disabled" toml:"disabled"`

	Languages []LocalLanguage
}

type Language struct {
	gorm.Model `json:"-" yaml:"-" toml:"-"`

	Debug    bool `gorm:"column:debug" default:"false" storm:"debug" json:"debug" yaml:"debug" toml:"debug"`
	Disabled bool `gorm:"column:disabled" default:"false" storm:"disabled" json:"disabled" yaml:"disabled" toml:"disabled"`

	Extensions   []Extension   `gorm:"many2many:local_fs_extensions;" yaml:"extensions,omitempty" json:"extensions,omitempty" toml:"extensions,omitempty"`
	Filenames    []File        `gorm:"many2many:local_fs_filenames;" yaml:"filenames,omitempty" json:"filenames,omitempty" toml:"filenames,omitempty"`
	Interpreters []Interpreter `gorm:"many2many:local_fs_interpreters;" yaml:"interpreters,omitempty" json:"interpreters,omitempty" toml:"interpreters,omitempty"`
	Color        string        `yaml:"color,omitempty" json:"color,omitempty" toml:"color,omitempty"`
}

type Extension struct {
	gorm.Model `json:"-" yaml:"-" toml:"-"`

	Debug    bool `gorm:"column:debug" default:"false" storm:"debug" json:"debug" yaml:"debug" toml:"debug"`
	Disabled bool `gorm:"column:disabled" default:"false" storm:"disabled" json:"disabled" yaml:"disabled" toml:"disabled"`

	Value    string `gorm:"column:value" mapstructure:"value" storm:"value" json:"value" yaml:"value" toml:"value"`
	MimeType string `gorm:"column:mime_type" mapstructure:"mime_type" storm:"mime_type" json:"mime_type" yaml:"mime_type" toml:"mime_type"`
}

type Interpreter struct {
	gorm.Model `json:"-" yaml:"-" toml:"-"`

	Debug    bool `gorm:"column:debug" default:"false" storm:"debug" json:"debug" yaml:"debug" toml:"debug"`
	Disabled bool `gorm:"column:disabled" default:"false" storm:"disabled" json:"disabled" yaml:"disabled" toml:"disabled"`

	Name string `gorm:"column:name" mapstructure:"name" storm:"name" json:"name" yaml:"name" toml:"name"`
	Slug string `gorm:"column:slug" mapstructure:"slug" storm:"slug" json:"slug" yaml:"slug" toml:"slug"`
}

//	languages := map[string]*language{}
