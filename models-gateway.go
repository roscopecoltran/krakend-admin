package main

import (
	"time"

	"github.com/jinzhu/gorm"
)

type GatewayService struct {
	gorm.Model `json:"-" yaml:"-" toml:"-"`

	Disabled bool `gorm:"column:disabled,omitempty" default:"false" storm:"disabled,omitempty" json:"disabled,omitempty" yaml:"disabled,omitempty" toml:"disabled,omitempty"`
	Debug    bool `gorm:"column:debug,omitempty" default:"false" example:"false" storm:"debug,omitempty" json:"debug,omitempty" yaml:"debug,omitempty" toml:"debug,omitempty"`

	Name             string        `mapstructure:"name" gorm:"column:name" json:"name,omitempty" yaml:"name,omitempty" toml:"name,omitempty"`
	Hosts            []GatewayHost `mapstructure:"host" gorm:"many2many:gateway_service_hosts;" json:"host,omitempty" yaml:"host,omitempty" toml:"host,omitempty"`
	HostStr          string        `mapstructure:"hosts" gorm:"column:hosts" json:"hosts,omitempty" yaml:"hosts,omitempty" toml:"hosts,omitempty"`
	ThrottlingHeader string        `mapstructure:"throttling_header" gorm:"column:throttling_header" default:"" json:"throttling_header,omitempty" yaml:"throttling_header,omitempty" toml:"throttling_header,omitempty"`
	Port             int           `mapstructure:"port" gorm:"column:port" json:"port,omitempty" default:"8080" yaml:"port,omitempty" toml:"port,omitempty"`
	Version          int           `mapstructure:"version" gorm:"column:version" json:"version,omitempty" default:"0" yaml:"version,omitempty" toml:"version,omitempty"`
	DefaultMaxRate   int           `mapstructure:"default_max_rate" gorm:"column:default_max_rate" default:"0" json:"default_max_rate,omitempty" yaml:"default_max_rate,omitempty" toml:"default_max_rate,omitempty"`
	ClientMaxRate    int           `mapstructure:"client_max_rate" gorm:"column:client_max_rate" default:"0" json:"client_max_rate,omitempty" yaml:"client_max_rate,omitempty" toml:"client_max_rate,omitempty"`
	Timeout          time.Duration `mapstructure:"timeout" gorm:"column:timeout" default:"3s" json:"timeout,omitempty" yaml:"timeout,omitempty" toml:"timeout,omitempty"`
	CacheTTL         time.Duration `mapstructure:"cache_ttl" gorm:"column:cache_ttl" default:"3600s" json:"cache_ttl,omitempty" yaml:"cache_ttl,omitempty" toml:"cache_ttl,omitempty"`

	Endpoints []GatewayEndpoint `mapstructure:"endpoints" gorm:"many2many:gateway_service_endpoints;" json:"endpoints,omitempty" yaml:"endpoints,omitempty" toml:"endpoints,omitempty"`
}

type GatewayHost struct {
	gorm.Model `json:"-" yaml:"-" toml:"-"`

	Disabled bool `gorm:"column:disabled,omitempty" default:"false" storm:"disabled,omitempty" json:"disabled,omitempty" yaml:"disabled,omitempty" toml:"disabled,omitempty"`
	Debug    bool `gorm:"column:debug,omitempty" default:"false" example:"false" storm:"debug,omitempty" json:"debug,omitempty" yaml:"debug,omitempty" toml:"debug,omitempty"`

	Address  string `gorm:"column:address" mapstructure:"address" storm:"address" json:"address" yaml:"address" toml:"address"`
	Protocol string `gorm:"column:protocol" mapstructure:"protocol" storm:"protocol" json:"protocol" yaml:"protocol" toml:"protocol"`
	Port     int    `gorm:"column:port" mapstructure:"port" storm:"port" json:"port" yaml:"port" toml:"port"`
}

type GatewayKV struct {
	gorm.Model `json:"-" yaml:"-" toml:"-"`

	Disabled bool `gorm:"column:disabled,omitempty" default:"false" storm:"disabled,omitempty" json:"disabled,omitempty" yaml:"disabled,omitempty" toml:"disabled,omitempty"`
	Debug    bool `gorm:"column:debug,omitempty" default:"false" example:"false" storm:"debug,omitempty" json:"debug,omitempty" yaml:"debug,omitempty" toml:"debug,omitempty"`

	Type  string `gorm:"column:type" mapstructure:"type" storm:"type" json:"type" yaml:"type" toml:"type"`
	Group string `gorm:"column:group" mapstructure:"group" storm:"group" json:"group" yaml:"group" toml:"group"`
	Key   string `gorm:"column:key" mapstructure:"key" storm:"key" json:"key" yaml:"key" toml:"key"`
	Value string `gorm:"column:value" mapstructure:"value" storm:"value" json:"value" yaml:"value" toml:"value"`
}

type GatewayHeaders GatewayKV
type GatewayQueryStrings GatewayKV
type GatewayParameters GatewayKV
type GatewayBlacklist GatewayKV
type GatewayWhitelist GatewayKV
type GatewayUrlKeys GatewayKV
type GatewayHolders GatewayKV
type GatewayMappings GatewayKV

type GatewayEndpoint struct {
	gorm.Model `json:"-" yaml:"-" toml:"-"`

	Disabled bool `gorm:"column:disabled,omitempty" default:"false" storm:"disabled,omitempty" json:"disabled,omitempty" yaml:"disabled,omitempty" toml:"disabled,omitempty"`
	Debug    bool `gorm:"column:debug,omitempty" default:"false" example:"false" storm:"debug,omitempty" json:"debug,omitempty" yaml:"debug,omitempty" toml:"debug,omitempty"`

	Endpoint        string           `gorm:"column:endpoint" mapstructure:"endpoint" storm:"endpoint" json:"endpoint" yaml:"endpoint" toml:"endpoint"`
	Method          string           `gorm:"column:method" mapstructure:"method" storm:"method" json:"method" yaml:"method" toml:"method"`
	QueryStringStr  string           `gorm:"column:querystring_params" mapstructure:"querystring_params" storm:"querystring_params" json:"querystring_params" yaml:"querystring_params" toml:"querystring_params"`
	ExtraConfigStr  string           `gorm:"column:extra_config" mapstructure:"extra_config" storm:"extra_config" json:"extra_config" yaml:"extra_config" toml:"extra_config"`
	ConcurrentCalls int              `gorm:"column:concurrent_calls" mapstructure:"concurrent_calls" storm:"concurrent_calls" json:"concurrent_calls" yaml:"concurrent_calls" toml:"concurrent_calls"`
	Timeout         time.Duration    `gorm:"column:timeout" mapstructure:"timeout" storm:"timeout" json:"timeout" yaml:"timeout" toml:"timeout"`
	CacheTTL        time.Duration    `gorm:"column:cache_ttl" mapstructure:"cache_ttl" storm:"cache_ttl" json:"cache_ttl" yaml:"cache_ttl" toml:"cache_ttl"`
	Backend         []GatewayBackend `gorm:"many2many:gateway_service_backends;" mapstructure:"backends" storm:"backends" json:"backends" yaml:"backends" toml:"backends"`
}

type GatewayBackend2 struct {
	Disabled bool `default:"false" storm:"disabled,omitempty" json:"disabled,omitempty" yaml:"disabled,omitempty" toml:"disabled,omitempty"`
	Debug    bool `default:"false" example:"false" storm:"debug,omitempty" json:"debug,omitempty" yaml:"debug,omitempty" toml:"debug,omitempty"`

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

	Hosts  []GatewayHost    `mapstructure:"host" gorm:"many2many:gateway_backend_hosts;" json:"host,omitempty" yaml:"host,omitempty" toml:"host,omitempty"`
	Header []GatewayHeaders `mapstructure:"header" storm:"header" json:"header" yaml:"header" toml:"header"`

	QueryStrings []GatewayQueryStrings `mapstructure:"query_strings" storm:"query_strings" json:"query_strings" yaml:"query_strings" toml:"query_strings"`
	Parameters   []GatewayParameters   `mapstructure:"parameters" storm:"parameters" json:"parameters" yaml:"parameters" toml:"parameters"`
	Blacklist    []GatewayBlacklist    `mapstructure:"blacklist" storm:"blacklist" json:"blacklist" yaml:"blacklist" toml:"blacklist"`
	Whitelist    []GatewayWhitelist    `mapstructure:"whitelist" storm:"whitelist" json:"whitelist" yaml:"whitelist" toml:"whitelist"`
	Paths        []GatewayQueryStrings `mapstructure:"paths" storm:"paths" json:"paths" yaml:"paths" toml:"paths"`
	Mapping      []GatewayMappings     `mapstructure:"mapping" storm:"mapping" json:"mapping" yaml:"mapping" toml:"mapping"`
	Holders      []GatewayHolders      `mapstructure:"holders" storm:"holders" json:"holders" yaml:"holders" toml:"holders"`
	URLKeys      []GatewayUrlKeys      `mapstructure:"urls_keys_flat" storm:"urls_keys_flat" json:"urls_keys_flat" yaml:"urls_keys_flat" toml:"urls_keys_flat"`
	// TaskQ        []TaskQueue       `storm:"taskq" json:"-" yaml:"-" toml:"-"`
}

type BackendExport struct {
	Disabled bool `default:"false" storm:"disabled,omitempty" json:"disabled,omitempty" yaml:"disabled,omitempty" toml:"disabled,omitempty"`
	Debug    bool `default:"false" example:"false" storm:"debug,omitempty" json:"debug,omitempty" yaml:"debug,omitempty" toml:"debug,omitempty"`

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
	// TaskQ        []TaskQueue       `storm:"taskq" json:"-" yaml:"-" toml:"-"`
}

type GatewayBackend struct {
	gorm.Model `json:"-" yaml:"-" toml:"-"`

	Disabled bool `gorm:"column:disabled,omitempty" default:"false" storm:"disabled,omitempty" json:"disabled,omitempty" yaml:"disabled,omitempty" toml:"disabled,omitempty"`
	Debug    bool `gorm:"column:debug,omitempty" default:"false" example:"false" storm:"debug,omitempty" json:"debug,omitempty" yaml:"debug,omitempty" toml:"debug,omitempty"`

	IsCollection             bool           `gorm:"column:is_collection" mapstructure:"is_collection" storm:"is_collection" json:"is_collection" yaml:"is_collection" toml:"is_collection"`
	HostSanitizationDisabled bool           `gorm:"column:disable_host_sanitize" mapstructure:"disable_host_sanitize" storm:"disable_host_sanitize" json:"disable_host_sanitize" yaml:"disable_host_sanitize" toml:"disable_host_sanitize"`
	Slug                     string         `gorm:"column:slug" json:"slug,omitempty" yaml:"slug,omitempty" toml:"slug,omitempty"`
	Group                    string         `gorm:"column:group" mapstructure:"group" storm:"group" json:"group" yaml:"group" toml:"group"`
	Method                   string         `gorm:"column:method" mapstructure:"method" storm:"method" json:"method" yaml:"method" toml:"method"`
	URLPattern               string         `gorm:"column:url_pattern" mapstructure:"url_pattern" storm:"url_pattern" json:"url_pattern" yaml:"url_pattern" toml:"url_pattern"`
	Encoding                 string         `gorm:"column:encoding" mapstructure:"encoding" storm:"encoding" json:"encoding" yaml:"encoding" toml:"encoding"`
	Target                   string         `gorm:"column:target" mapstructure:"target" storm:"target" json:"target" yaml:"target" toml:"target"`
	HostStr                  string         `gorm:"host" mapstructure:"host_flat" storm:"host_flat" json:"host_flat" yaml:"host_flat" toml:"host_flat"`
	HeaderStr                string         `gorm:"column:header" mapstructure:"header_flat" storm:"header_flat" json:"header_flat" yaml:"header_flat" toml:"header_flat"`
	QueryStringsStr          string         `gorm:"column:query_strings" mapstructure:"query_strings_flat" storm:"query_strings_flat" json:"query_strings_flat" yaml:"query_strings_flat" toml:"query_strings_flat"`
	ParametersStr            string         `gorm:"parameters" mapstructure:"parameters_flat" storm:"parameters_flat" json:"parameters_flat" yaml:"parameters_flat" toml:"parameters_flat"`
	BlacklistStr             string         `gorm:"column:blacklist_flat" mapstructure:"blacklist_flat" storm:"blacklist_flat" json:"blacklist_flat" yaml:"blacklist_flat" toml:"blacklist_flat"`
	WhitelistStr             string         `gorm:"column:whitelist" mapstructure:"whitelist" storm:"whitelist" json:"whitelist" yaml:"whitelist" toml:"whitelist"`
	PathsStr                 string         `gorm:"paths" mapstructure:"paths_flat" storm:"paths_flat" json:"paths_flat" yaml:"paths_flat" toml:"paths_flat"`
	MappingStr               string         `gorm:"mapping" mapstructure:"mapping_flat" storm:"mapping_flat" json:"mapping_flat" yaml:"mapping_flat" toml:"mapping_flat"`
	HoldersStr               string         `gorm:"holders" mapstructure:"holders_flat" storm:"holders_flat" json:"holders_flat" yaml:"holders_flat" toml:"holders_flat"`
	URLKeysStr               string         `gorm:"urls_keys" mapstructure:"urls_keys_flat" storm:"urls_keys_flat" json:"urls_keys_flat" yaml:"urls_keys_flat" toml:"urls_keys_flat"`
	ConcurrentCalls          int            `gorm:"column:concurrent_calls" default:"1" storm:"concurrent_calls" json:"concurrent_calls" yaml:"concurrent_calls" toml:"concurrent_calls"`
	Timeout                  time.Duration  `gorm:"column:timeout" default:"3s" storm:"timeout" json:"timeout" yaml:"timeout" toml:"timeout"`
	Paths                    []ProviderPath `gorm:"many2many:gateway_service_backend_paths;" json:"paths,omitempty" yaml:"paths,omitempty" toml:"paths,omitempty"`
}
