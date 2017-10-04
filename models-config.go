package main

import (
	"github.com/jinzhu/gorm"
	// "github.com/roscopecoltran/krakend/config"
)

const (
	DefaultLimitValue   int    = 25
	DefaultPageValue    int    = 1
	DefaultLimitString  string = "limit"
	DefaultPageString   string = "page"
	DefaultExpandString string = "expand"
	DefaultFieldsString string = "fields"
	DefaultOrderString  string = "order"
	DefaultAscString    string = "asc"
	DefaultDescString   string = "desc"
	DefaultQueryString  string = "q"
	DefaultParamString  string = "p"
	DefaultLeftBracket  rune   = '('
	DefaultRightBracket rune   = ')'
	DefaultSeparator    rune   = ','
	DefaultKVSeparator  rune   = ':'
)

//type KrakendService *config.ServiceConfig
// var cfg config.ServiceConfig

type Config struct {
	gorm.Model `json:"-" yaml:"-" toml:"-"`

	Debug    bool   `gorm:"column:debug" default:"false" storm:"debug" json:"debug" yaml:"debug" toml:"debug"`
	Instance string `gorm:"column:instance" storm:"instance" json:"instance" yaml:"instance" toml:"instance"`

	//Env        []EnvConfig      `gorm:"many2many:gateway_local_envs;" storm:"env" json:"env" yaml:"env" file:"environment" toml:"env"`
	Service   []GatewayService     `gorm:"many2many:gateway_services;" storm:"service" json:"service" yaml:"service" toml:"service"`
	Backends  []BackendExport      `gorm:"-" storm:"backends" json:"backends" yaml:"backends" file:"backends" toml:"backends"`
	Providers []Provider           `gorm:"-" storm:"providers" json:"providers" yaml:"providers" file:"providers" toml:"providers"`
	Krakend   KrakendServiceConfig `gorm:"-" storm:"krakend" json:"krakend" yaml:"krakend" file:"krakend" toml:"krakend"`
	Specs     SpecsConfig          `gorm:"-" storm:"specs" json:"specs" yaml:"specs" file:"specs" toml:"specs"`
}
