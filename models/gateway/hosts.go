package gateway

import (
	"github.com/jinzhu/gorm"
)

type Host struct {
	gorm.Model `json:"-" yaml:"-" toml:"-"`

	Disabled bool `gorm:"column:disabled" default:"false" storm:"disabled" json:"disabled" yaml:"disabled" toml:"disabled"`
	Debug    bool `gorm:"column:debug" default:"false" example:"false" storm:"debug" json:"debug" yaml:"debug" toml:"debug"`

	Address string `gorm:"column:address" mapstructure:"address" json:"address,omitempty" yaml:"address,omitempty" toml:"address,omitempty"`
	Port    int    `gorm:"column:port" mapstructure:"port" json:"port,omitempty" yaml:"port,omitempty" toml:"port,omitempty"`
}
