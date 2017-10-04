package local

import (
	"github.com/jinzhu/gorm"
	"github.com/qor/media/oss"
)

var ValidMimeExtensions = []string{"yaml", "yml", "json", "ini", "conf", "cnf", "toml"} // ,".env-example"}

type LocalFile struct {
	gorm.Model `json:"-" yaml:"-" toml:"-"`

	Disabled   bool   `gorm:"column:disabled" default:"false" storm:"disabled" json:"disabled" yaml:"disabled" toml:"disabled"`
	Debug      bool   `gorm:"column:debug" default:"false" example:"false" storm:"debug" json:"debug" yaml:"debug" toml:"debug"`
	Name       string `gorm:"column:filename" storm:"filename" json:"filename" yaml:"filename" toml:"filename"`
	Path       string `gorm:"column:filepath" storm:"filepath" json:"filepath" yaml:"filepath" toml:"filepath"`
	PrefixPath string `gorm:"column:prefix_path" storm:"prefix_path" json:"prefix_path" yaml:"prefix_path" toml:"prefix_path"`
	File       oss.OSS
}
