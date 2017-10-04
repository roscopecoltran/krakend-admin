package local

import (
	"github.com/jinzhu/gorm"
	"github.com/qor/media/oss"
)

var DefaultEnvFiles = []string{".env"} // ,".env-example"}

type LocalEnvironment struct {
	gorm.Model `json:"-" yaml:"-" toml:"-"`

	Debug    bool   `gorm:"column:debug" default:"false" storm:"debug" json:"debug" yaml:"debug" toml:"debug"`
	Instance string `gorm:"column:instance" mapstructure:"instance" storm:"instance" json:"instance" yaml:"instance" toml:"instance"`

	File oss.OSS
	Envs []LocalEnvironmentVariable `gorm:"many2many:gateway_local_envvars;" json:"variables,omitempty" yaml:"variables,omitempty" toml:"variables,omitempty"`
}

type LocalEnvironmentVariable struct {
	gorm.Model `json:"-" yaml:"-" toml:"-"`

	Disabled bool `gorm:"column:disabled" default:"false" storm:"disabled" json:"disabled" yaml:"disabled" toml:"disabled"`
	Debug    bool `gorm:"column:debug" default:"false" example:"false" storm:"debug" json:"debug" yaml:"debug" toml:"debug"`

	Key string `gorm:"column:key" storm:"key" json:"key" yaml:"key" toml:"key"`
	Val string `gorm:"column:val" storm:"val" json:"val" yaml:"val" toml:"val"`
	// SourceFile string `gorm:"column:file" storm:"file" json:"file" yaml:"file" toml:"file"`
}
