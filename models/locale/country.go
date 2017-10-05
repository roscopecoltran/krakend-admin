package locales

import (
	"github.com/jinzhu/gorm"
)

type Countries struct {
	gorm.Model `json:"-" yaml:"-" toml:"-"`

	Disabled bool `gorm:"column:disabled,omitempty" default:"false" storm:"disabled,omitempty" json:"disabled,omitempty" yaml:"disabled,omitempty" toml:"disabled,omitempty"`
	Debug    bool `gorm:"column:debug,omitempty" default:"false" example:"false" storm:"debug,omitempty" json:"debug,omitempty" yaml:"debug,omitempty" toml:"debug,omitempty"`

	List []Country
}

type Country struct {
	gorm.Model `json:"-" yaml:"-" toml:"-"`

	Disabled bool `gorm:"column:disabled,omitempty" default:"false" storm:"disabled,omitempty" json:"disabled,omitempty" yaml:"disabled,omitempty" toml:"disabled,omitempty"`
	Debug    bool `gorm:"column:debug,omitempty" default:"false" example:"false" storm:"debug,omitempty" json:"debug,omitempty" yaml:"debug,omitempty" toml:"debug,omitempty"`

	Name    string `gorm:"column:country_name" xml:"CtryNm" json:"country_name,omitempty" yaml:"country_name,omitempty" toml:"country_name,omitempty"`
	Capital string `gorm:"column:capital_name" xml:"CtryCapitalNm" json:"capital_name,omitempty" yaml:"capital_name,omitempty" toml:"capital_name,omitempty"`
	Iso2    string `gorm:"column:country_iso2" xml:"CtryIso2" json:"country_iso2,omitempty" yaml:"country_iso2,omitempty" toml:"country_iso2,omitempty"`
	Iso3    string `gorm:"column:country_iso3" xml:"CtryIso3" json:"country_iso3,omitempty" yaml:"country_iso3,omitempty" toml:"country_iso3,omitempty"`
}
