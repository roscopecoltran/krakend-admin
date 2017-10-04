package patterns

import (
	"github.com/jinzhu/gorm"
)

type PathMXJ struct {
	gorm.Model `json:"-" yaml:"-" toml:"-"`

	Disabled bool `gorm:"column:disabled" default:"false" storm:"disabled" json:"disabled" yaml:"disabled" toml:"disabled"`
	Debug    bool `gorm:"column:debug" default:"false" example:"false" storm:"debug" json:"debug" yaml:"debug" toml:"debug"`

	Active   int    `gorm:"column:active" default:"1" example:"2" storm:"active" json:"active" yaml:"active" toml:"active"`
	Encoding string `gorm:"column:encoding" storm:"encoding" json:"encoding" yaml:"encoding" toml:"encoding"`
	Path     string `gorm:"column:path" storm:"path" json:"path" yaml:"path" toml:"path"`
}
