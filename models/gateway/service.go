package gateway

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Service struct {
	gorm.Model `json:"-" yaml:"-" toml:"-"`

	Disabled  bool          `gorm:"column:disabled" default:"false" storm:"disabled" json:"disabled" yaml:"disabled" toml:"disabled"`
	Debug     bool          `gorm:"column:debug" default:"false" example:"false" storm:"debug" json:"debug" yaml:"debug" toml:"debug"`
	Timeout   time.Duration `mapstructure:"timeout" gorm:"column:timeout" json:"timeout,omitempty" yaml:"timeout,omitempty" toml:"timeout,omitempty"`
	CacheTTL  time.Duration `mapstructure:"cache_ttl" gorm:"column:cache_ttl" json:"cache_ttl,omitempty" yaml:"cache_ttl,omitempty" toml:"cache_ttl,omitempty"`
	HostStr   string        `mapstructure:"host" gorm:"column:host" json:"host,omitempty" yaml:"host,omitempty" toml:"host,omitempty"`
	Port      int           `mapstructure:"port" gorm:"column:port" json:"port,omitempty" yaml:"port,omitempty" toml:"port,omitempty"`
	Version   int           `mapstructure:"version" gorm:"column:version" json:"version,omitempty" yaml:"version,omitempty" toml:"version,omitempty"`
	Endpoints []Endpoint    `mapstructure:"endpoints" gorm:"many2many:gateway_service_endpoints;" json:"endpoints,omitempty" yaml:"endpoints,omitempty" toml:"endpoints,omitempty"`
}
