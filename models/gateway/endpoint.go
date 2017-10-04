package gateway

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Endpoint struct {
	gorm.Model `json:"-" yaml:"-" toml:"-"`

	Disabled bool `gorm:"column:disabled" default:"false" storm:"disabled" json:"disabled" yaml:"disabled" toml:"disabled"`
	Debug    bool `gorm:"column:debug" default:"false" example:"false" storm:"debug" json:"debug" yaml:"debug" toml:"debug"`

	Endpoint        string        `gorm:"column:endpoint" mapstructure:"endpoint" storm:"endpoint" json:"endpoint" yaml:"endpoint" toml:"endpoint"`
	Method          string        `gorm:"column:method" mapstructure:"method" storm:"method" json:"method" yaml:"method" toml:"method"`
	QueryStringStr  string        `gorm:"column:querystring_params" mapstructure:"querystring_params" storm:"querystring_params" json:"querystring_params" yaml:"querystring_params" toml:"querystring_params"`
	ExtraConfigStr  string        `gorm:"column:extra_config" mapstructure:"extra_config" storm:"extra_config" json:"extra_config" yaml:"extra_config" toml:"extra_config"`
	ConcurrentCalls int           `gorm:"column:concurrent_calls" mapstructure:"concurrent_calls" storm:"concurrent_calls" json:"concurrent_calls" yaml:"concurrent_calls" toml:"concurrent_calls"`
	Timeout         time.Duration `gorm:"column:timeout" mapstructure:"timeout" storm:"timeout" json:"timeout" yaml:"timeout" toml:"timeout"`
	CacheTTL        time.Duration `gorm:"column:cache_ttl" mapstructure:"cache_ttl" storm:"cache_ttl" json:"cache_ttl" yaml:"cache_ttl" toml:"cache_ttl"`
	Backend         []Backend     `gorm:"many2many:gateway_service_backends;" mapstructure:"backends" storm:"backends" json:"backends" yaml:"backends" toml:"backends"`
}
