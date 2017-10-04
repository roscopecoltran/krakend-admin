package main

import (
	"github.com/jinzhu/gorm"
)

type CommonCurrency struct {
	gorm.Model `json:"-" yaml:"-" toml:"-"`

	Name      string          `gorm:"column:name" xml:"CcyNm" json:"currency_name" yaml:"currency_name" toml:"currency_name"`
	Code      string          `gorm:"column:iso_code" xml:"Ccy" json:"iso_code" yaml:"iso_code" toml:"iso_code"`
	Num       string          `gorm:"column:iso_num" xml:"CcyNbr" json:"iso_num" yaml:"iso_num" toml:"iso_num"`
	MinorUnit string          `gorm:"column:minor_unit" xml:"CcyMnrUnts" json:"minor_unit,omitempty" yaml:"minor_unit,omitempty" toml:"minor_unit,omitempty"`
	CodeIso   string          `gorm:"column:code_iso" xml:"CcyCodeIso" json:"code_iso,omitempty" yaml:"code_iso,omitempty" toml:"code_iso,omitempty"`
	Countries []CommonCountry `gorm:"column:countries" xml:"CcyTbl>CcyNtry" json:"countries" yaml:"countries" toml:"countries"`
}

type CommonCountry struct {
	gorm.Model `json:"-" yaml:"-" toml:"-"`

	Name    string `gorm:"column:country_name" xml:"CtryNm" json:"country_name,omitempty" yaml:"country_name,omitempty" toml:"country_name,omitempty"`
	Capital string `gorm:"column:capital_name" xml:"CtryCapitalNm" json:"capital_name,omitempty" yaml:"capital_name,omitempty" toml:"capital_name,omitempty"`

	Iso2 string `gorm:"column:country_iso2" xml:"CtryIso2" json:"country_iso2,omitempty" yaml:"country_iso2,omitempty" toml:"country_iso2,omitempty"`
	Iso3 string `gorm:"column:country_iso3" xml:"CtryIso3" json:"country_iso3,omitempty" yaml:"country_iso3,omitempty" toml:"country_iso3,omitempty"`
}

/*
type CurrencyIso4217 struct {
	gorm.Model `json:"-" yaml:"-" toml:"-"`
	Published  string     `gorm:"column:published" xml:"Pblshd,attr" json:"published" yaml:"published" toml:"published"`
	Currencies []Currency `gorm:"column:currencies" xml:"CcyTbl>CcyNtry" json:"currencies" yaml:"currencies" toml:"currencies"`
}
*/
