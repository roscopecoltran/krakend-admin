package openapi

import (
	// 	"time"
	// "github.com/qor/media/oss"
	"github.com/jinzhu/gorm"
)

type DataAPI struct {
	gorm.Model `json:"-" yaml:"-" toml:"-"`
	// ID int `gorm:"column:id" json:"id" yaml:"id" toml:"id"`

	FullName  string `gorm:"column:api_name" json:"api_name" yaml:"api_name" toml:"api_name"`
	Company   string `gorm:"column:company" json:"company" yaml:"company" toml:"company"`
	Product   string `gorm:"column:product" json:"product" yaml:"product" toml:"product"`
	System    string `gorm:"column:system" json:"system" yaml:"system" toml:"system"`
	Interface string `gorm:"column:interface" json:"interface" yaml:"interface" toml:"interface"`
	Version   string `gorm:"column:version" json:"version" yaml:"version" toml:"version"`

	Method string `gorm:"column:method" json:"method" yaml:"method" toml:"method"`

	ProxyMode string `gorm:"column:proxy_mode" json:"proxy_mode" yaml:"proxy_mode" toml:"proxy_mode"`

	UpstreamMode  string `gorm:"column:upstream_mode" json:"upstream_mode" yaml:"upstream_mode" toml:"upstream_mode"`
	UpstreamValue string `gorm:"column:upstream_value" json:"upstream_value" yaml:"upstream_value" toml:"upstream_value"`
}

type DataInApi struct {
	gorm.Model `json:"-" yaml:"-" toml:"-"`

	ApiName   string `gorm:"column:api_name" json:"api_name" yaml:"api_name" toml:"api_name"`
	Company   string `gorm:"column:company" json:"company" yaml:"company" toml:"company"`
	Product   string `gorm:"column:product" json:"product" yaml:"product" toml:"product"`
	System    string `gorm:"column:system" json:"system" yaml:"system" toml:"system"`
	Interface string `gorm:"column:interface" json:"interface" yaml:"interface" toml:"interface"`

	Method   string `gorm:"column:method" json:"method" yaml:"method" toml:"method"`
	Register string `gorm:"column:register" json:"register" yaml:"register" toml:"register"`

	ApiDesc    string `gorm:"column:api_desc" json:"api_desc" yaml:"api_desc" toml:"api_desc"`
	ParamDesc  string `gorm:"column:param_desc" json:"param_desc" yaml:"param_desc" toml:"param_desc"`
	ReturnDesc string `gorm:"column:return_desc" json:"return_desc" yaml:"return_desc" toml:"return_desc"`
	InputDate  string `gorm:"column:input_date" json:"input_date" yaml:"input_date" toml:"input_date"`
	InputStaff string `gorm:"column:input_staff" json:"input_staff" yaml:"input_staff" toml:"input_staff"`
}
