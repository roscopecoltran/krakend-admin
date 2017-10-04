package config

/*
import (
	"github.com/jinzhu/gorm"
	"github.com/roscopecoltran/krakend-admin/models"
	"github.com/roscopecoltran/krakend-admin/models/provider"
)

type Config struct {
	gorm.Model `json:"-" yaml:"-" toml:"-"`

	Debug    bool   `gorm:"column:debug" default:"false" storm:"debug" json:"debug" yaml:"debug" toml:"debug"`
	Instance string `gorm:"column:instance" storm:"instance" json:"instance" yaml:"instance" toml:"instance"`

	//Env        []EnvConfig      `gorm:"many2many:gateway_local_envs;" storm:"env" json:"env" yaml:"env" file:"environment" toml:"env"`
	Service   []models.GatewayService     `gorm:"many2many:gateway_services;" storm:"service" json:"service" yaml:"service" toml:"service"`
	Backends  []models.BackendExport      `gorm:"-" storm:"backends" json:"backends" yaml:"backends" file:"backends" toml:"backends"`
	Providers []provider.Provider           `gorm:"-" storm:"providers" json:"providers" yaml:"providers" file:"providers" toml:"providers"`
	Krakend   models.KrakendServiceConfig `gorm:"-" storm:"krakend" json:"krakend" yaml:"krakend" file:"krakend" toml:"krakend"`
	Specs     models.SpecsConfig          `gorm:"-" storm:"specs" json:"specs" yaml:"specs" file:"specs" toml:"specs"`
}
*/
