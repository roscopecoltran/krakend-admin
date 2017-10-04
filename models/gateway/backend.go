package gateway

import (
	"time"

	"github.com/jinzhu/gorm"
	"github.com/roscopecoltran/krakend-admin/models/patterns"
	//"github.com/roscopecoltran/krakend-admin/models/local"
)

type Backend struct {
	gorm.Model `json:"-" yaml:"-" toml:"-"`

	Disabled bool `gorm:"column:disabled" default:"false" storm:"disabled" json:"disabled" yaml:"disabled" toml:"disabled"`
	Debug    bool `gorm:"column:debug" default:"false" example:"false" storm:"debug" json:"debug" yaml:"debug" toml:"debug"`

	IsCollection             bool               `gorm:"column:is_collection" mapstructure:"is_collection" storm:"is_collection" json:"is_collection" yaml:"is_collection" toml:"is_collection"`
	HostSanitizationDisabled bool               `gorm:"column:disable_host_sanitize" mapstructure:"disable_host_sanitize" storm:"disable_host_sanitize" json:"disable_host_sanitize" yaml:"disable_host_sanitize" toml:"disable_host_sanitize"`
	Slug                     string             `gorm:"column:slug" json:"slug,omitempty" yaml:"slug,omitempty" toml:"slug,omitempty"`
	Group                    string             `gorm:"column:group" mapstructure:"group" storm:"group" json:"group" yaml:"group" toml:"group"`
	Method                   string             `gorm:"column:method" mapstructure:"method" storm:"method" json:"method" yaml:"method" toml:"method"`
	URLPattern               string             `gorm:"column:url_pattern" mapstructure:"url_pattern" storm:"url_pattern" json:"url_pattern" yaml:"url_pattern" toml:"url_pattern"`
	Encoding                 string             `gorm:"column:encoding" mapstructure:"encoding" storm:"encoding" json:"encoding" yaml:"encoding" toml:"encoding"`
	Target                   string             `gorm:"column:target" mapstructure:"target" storm:"target" json:"target" yaml:"target" toml:"target"`
	HostStr                  string             `gorm:"host" mapstructure:"host_flat" storm:"host_flat" json:"host_flat" yaml:"host_flat" toml:"host_flat"`
	HeaderStr                string             `gorm:"column:header" mapstructure:"header_flat" storm:"header_flat" json:"header_flat" yaml:"header_flat" toml:"header_flat"`
	QueryStringsStr          string             `gorm:"column:query_strings" mapstructure:"query_strings_flat" storm:"query_strings_flat" json:"query_strings_flat" yaml:"query_strings_flat" toml:"query_strings_flat"`
	ParametersStr            string             `gorm:"parameters" mapstructure:"parameters_flat" storm:"parameters_flat" json:"parameters_flat" yaml:"parameters_flat" toml:"parameters_flat"`
	BlacklistStr             string             `gorm:"column:blacklist_flat" mapstructure:"blacklist_flat" storm:"blacklist_flat" json:"blacklist_flat" yaml:"blacklist_flat" toml:"blacklist_flat"`
	WhitelistStr             string             `gorm:"column:whitelist" mapstructure:"whitelist" storm:"whitelist" json:"whitelist" yaml:"whitelist" toml:"whitelist"`
	PathsStr                 string             `gorm:"paths" mapstructure:"paths_flat" storm:"paths_flat" json:"paths_flat" yaml:"paths_flat" toml:"paths_flat"`
	MappingStr               string             `gorm:"mapping" mapstructure:"mapping_flat" storm:"mapping_flat" json:"mapping_flat" yaml:"mapping_flat" toml:"mapping_flat"`
	HoldersStr               string             `gorm:"holders" mapstructure:"holders_flat" storm:"holders_flat" json:"holders_flat" yaml:"holders_flat" toml:"holders_flat"`
	URLKeysStr               string             `gorm:"urls_keys" mapstructure:"urls_keys_flat" storm:"urls_keys_flat" json:"urls_keys_flat" yaml:"urls_keys_flat" toml:"urls_keys_flat"`
	ConcurrentCalls          int                `gorm:"column:concurrent_calls" default:"1" storm:"concurrent_calls" json:"concurrent_calls" yaml:"concurrent_calls" toml:"concurrent_calls"`
	Timeout                  time.Duration      `gorm:"column:timeout" default:"3s" storm:"timeout" json:"timeout" yaml:"timeout" toml:"timeout"`
	Paths                    []patterns.PathMXJ `gorm:"many2many:gateway_service_backend_paths_mxj;" json:"paths,omitempty" yaml:"paths,omitempty" toml:"paths,omitempty"`

	// TaskQ        []local.TaskQueue       `storm:"taskq" json:"-" yaml:"-" toml:"-"`
}

type BackendKrakend struct {
	Disabled bool `default:"false" storm:"disabled" json:"disabled" yaml:"disabled" toml:"disabled"`
	Debug    bool `default:"false" example:"false" storm:"debug" json:"debug" yaml:"debug" toml:"debug"`

	IsCollection             bool          `mapstructure:"is_collection" storm:"is_collection" json:"is_collection" yaml:"is_collection" toml:"is_collection"`
	HostSanitizationDisabled bool          `mapstructure:"disable_host_sanitize" storm:"disable_host_sanitize" json:"disable_host_sanitize" yaml:"disable_host_sanitize" toml:"disable_host_sanitize"`
	Slug                     string        `json:"slug,omitempty" yaml:"slug,omitempty" toml:"slug,omitempty"`
	Group                    string        `mapstructure:"group" storm:"group" json:"group" yaml:"group" toml:"group"`
	Method                   string        `mapstructure:"method" storm:"method" json:"method" yaml:"method" toml:"method"`
	URLPattern               string        `mapstructure:"url_pattern" storm:"url_pattern" json:"url_pattern" yaml:"url_pattern" toml:"url_pattern"`
	Encoding                 string        `mapstructure:"encoding" storm:"encoding" json:"encoding" yaml:"encoding" toml:"encoding"`
	Target                   string        `mapstructure:"target" storm:"target" json:"target" yaml:"target" toml:"target"`
	ConcurrentCalls          int           `default:"1" storm:"concurrent_calls" json:"concurrent_calls" yaml:"concurrent_calls" toml:"concurrent_calls"`
	Timeout                  time.Duration `default:"3s" storm:"timeout" json:"timeout" yaml:"timeout" toml:"timeout"`

	Host         []string            `mapstructure:"host" storm:"host" json:"host" yaml:"host" toml:"host"`
	Header       []string            `mapstructure:"header" storm:"header" json:"header" yaml:"header" toml:"header"`
	QueryStrings map[string]string   `mapstructure:"query_strings" storm:"query_strings" json:"query_strings" yaml:"query_strings" toml:"query_strings"`
	Parameters   map[string]string   `mapstructure:"parameters" storm:"parameters" json:"parameters" yaml:"parameters" toml:"parameters"`
	Blacklist    []string            `mapstructure:"blacklist" storm:"blacklist" json:"blacklist" yaml:"blacklist" toml:"blacklist"`
	Whitelist    []string            `mapstructure:"whitelist" storm:"whitelist" json:"whitelist" yaml:"whitelist" toml:"whitelist"`
	Paths        []map[string]string `mapstructure:"paths" storm:"paths" json:"paths" yaml:"paths" toml:"paths"`
	Mapping      map[string]string   `mapstructure:"mapping" storm:"mapping" json:"mapping" yaml:"mapping" toml:"mapping"`
	Holders      []map[string]string `mapstructure:"holders" storm:"holders" json:"holders" yaml:"holders" toml:"holders"`
	URLKeys      []string            `mapstructure:"urls_keys_flat" storm:"urls_keys_flat" json:"urls_keys_flat" yaml:"urls_keys_flat" toml:"urls_keys_flat"`
}
