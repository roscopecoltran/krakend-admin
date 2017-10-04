package patterns

import (
	"github.com/jinzhu/gorm"
)

type QueryString struct {
	gorm.Model `json:"-" yaml:"-" toml:"-"`

	Disabled bool `gorm:"column:disabled" default:"false" storm:"disabled" json:"disabled" yaml:"disabled" toml:"disabled"`
	Debug    bool `gorm:"column:debug" default:"false" example:"false" storm:"debug" json:"debug" yaml:"debug" toml:"debug"`

	LimitString  string `gorm:"column:limit_str" storm:"limit_str" json:"limit_str" yaml:"limit_str" toml:"limit_str"`
	PageString   string `gorm:"column:page_string" storm:"page_string" json:"page_string" yaml:"page_string" toml:"page_string"`
	ExpandString string `gorm:"column:expand_string" storm:"expand_string" json:"expand_string" yaml:"expand_string" toml:"expand_string"`
	FieldsString string `gorm:"column:fields_string" storm:"fields_string" json:"fields_string" yaml:"fields_string" toml:"fields_string"`
	OrderString  string `gorm:"column:order_string" storm:"order_string" json:"order_string" yaml:"order_string" toml:"order_string"`
	AscString    string `gorm:"column:asc_string" storm:"asc_string" json:"asc_string" yaml:"asc_string" toml:"asc_string"`
	DescString   string `gorm:"column:desc_string" storm:"desc_string" json:"desc_string" yaml:"desc_string" toml:"desc_string"`
	QueryString  string `gorm:"column:query_string" storm:"query_string" json:"query_string" yaml:"query_string" toml:"query_string"`
	ParamString  string `gorm:"column:param_string" storm:"param_string" json:"param_string" yaml:"param_string" toml:"param_string"`
	LimitValue   int    `gorm:"column:limit_value" storm:"limit_value" json:"limit_value" yaml:"limit_value" toml:"limit_value"`
	PageValue    int    `gorm:"column:page_value" storm:"page_value" json:"page_value" yaml:"page_value" toml:"page_value"`
	LeftBracket  rune   `gorm:"column:left_bracket" storm:"left_bracket" json:"left_bracket" yaml:"left_bracket" toml:"left_bracket"`
	RightBracket rune   `gorm:"column:right_bracket" storm:"right_bracket" json:"right_bracket" yaml:"right_bracket" toml:"right_bracket"`
	Separator    rune   `gorm:"column:separator" storm:"separator" json:"separator" yaml:"separator" toml:"separator"`
	KVSeparator  rune   `gorm:"column:kv_sperator" storm:"kv_sperator" json:"kv_sperator" yaml:"kv_sperator" toml:"kv_sperator"`
}
