package locales

import (
	"github.com/jinzhu/gorm"
)

type Currencies struct {
	gorm.Model `json:"-" yaml:"-" toml:"-"`

	Disabled bool `gorm:"column:disabled,omitempty" default:"false" storm:"disabled,omitempty" json:"disabled,omitempty" yaml:"disabled,omitempty" toml:"disabled,omitempty"`
	Debug    bool `gorm:"column:debug,omitempty" default:"false" example:"false" storm:"debug,omitempty" json:"debug,omitempty" yaml:"debug,omitempty" toml:"debug,omitempty"`

	List []Currency
}

type Currency struct {
	gorm.Model `json:"-" yaml:"-" toml:"-"`

	Disabled bool `gorm:"column:disabled,omitempty" default:"false" storm:"disabled,omitempty" json:"disabled,omitempty" yaml:"disabled,omitempty" toml:"disabled,omitempty"`
	Debug    bool `gorm:"column:debug,omitempty" default:"false" example:"false" storm:"debug,omitempty" json:"debug,omitempty" yaml:"debug,omitempty" toml:"debug,omitempty"`

	Name      string    `gorm:"column:name" xml:"CcyNm" json:"currency_name" yaml:"currency_name" toml:"currency_name"`
	Code      string    `gorm:"column:iso_code" xml:"Ccy" json:"iso_code" yaml:"iso_code" toml:"iso_code"`
	Num       string    `gorm:"column:iso_num" xml:"CcyNbr" json:"iso_num" yaml:"iso_num" toml:"iso_num"`
	MinorUnit string    `gorm:"column:minor_unit" xml:"CcyMnrUnts" json:"minor_unit,omitempty" yaml:"minor_unit,omitempty" toml:"minor_unit,omitempty"`
	CodeIso   string    `gorm:"column:code_iso" xml:"CcyCodeIso" json:"code_iso,omitempty" yaml:"code_iso,omitempty" toml:"code_iso,omitempty"`
	Countries []Country `gorm:"column:countries" xml:"CcyTbl>CcyNtry" json:"countries" yaml:"countries" toml:"countries"`
}
