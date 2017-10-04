package main

import (
	"time"
	// "github.com/roscopecoltran/krakend/encoding"
)

type KrakendExtraConfig map[string]interface{}

// ConfigGetter is a function for parsing ExtraConfig into a previously know type
type ConfigGetter func(KrakendExtraConfig) interface{}

// DefaultConfigGetter is the Default implementation for ConfigGetter, it just returns the KrakendExtraConfig map.
func DefaultConfigGetter(extra KrakendExtraConfig) interface{} { return extra }

// Backend defines how krakend should connect to the backend service (the API resource to consume)
// and how it should process the received response
type KrakendBackend struct {
	// the name of the group the response should be moved to. If empty, the response is
	// not changed
	Group string `mapstructure:"group"`
	// HTTP method of the request to send to the backend
	Method string `mapstructure:"method"`
	// Set of hosts of the API
	Host []string `mapstructure:"host"`
	// False if the hostname should be sanitized
	HostSanitizationDisabled bool `mapstructure:"disable_host_sanitize"`
	// URL pattern to use to locate the resource to be consumed
	URLPattern string `mapstructure:"url_pattern"`
	// HTTP headers of the endpoint to forward
	Header []string `mapstructure:"header"`
	// Query Strings to use during the request
	QueryStrings map[string]string `mapstructure:"query_strings"`
	// URL pattern to use to locate the resource to be consumed
	Parameters map[string]string `mapstructure:"parameters"`
	// set of response fields to remove. If empty, the filter id not used
	Blacklist []string `mapstructure:"blacklist"`
	// set of response fields to allow. If empty, the filter id not used
	Whitelist []string `mapstructure:"whitelist"`
	// set of pre-processing rules to filters the response body with wildcards. If empty, the PreProcess id not used
	Paths []map[string]string `mapstructure:"paths"`
	// map of response fields to be renamed and their new names
	Mapping map[string]string `mapstructure:"mapping"`
	// set of variables extracted with custom rules. If empty, the holders are not used
	Holders []map[string]string `mapstructure:"holders"`
	// Tasks/Queue order to fetch and aggregate data
	// TaskQ []TaskQConfig `mapstructure:"taskq"`
	// the encoding format
	Encoding string `mapstructure:"encoding"`
	// the response to process is a collection
	IsCollection bool `mapstructure:"is_collection"`
	// name of the field to extract to the root. If empty, the formater will do nothing
	Target string `mapstructure:"target"`
	// list of keys to be replaced in the URLPattern
	URLKeys []string
	// number of concurrent calls this endpoint must send to the API
	ConcurrentCalls int
	// timeout of this backend
	Timeout time.Duration
	// decoder to use in order to parse the received response from the API
	// Decoder encoding.Decoder
	// Processors configuration for handling data received or cached
	Processors ProcessorConfig `mapstructure:"processors"`
	// Indexes configuration for matching, searching or caching behaviours at Backend level
	// Indexes IndexConfig `mapstructure:"indexes"`
	// Entities configuration for graph datastore indexing
	Entities EntityConfig `mapstructure:"entities"`
	// Backend Extra configuration for customized behaviours
	ExtraConfig KrakendExtraConfig `mapstructure:"extra_config"`
	// Execute backend in debug mode
	Debug bool `default:"false" json:"debug" yaml:"debug" toml:"debug"`

	// Provider Name
	//ProviderName string `mapstructure:"provider_name"`
	// Provider Hook
	//ProviderHook string `mapstructure:"provider_hook"`
	// Provider Name
	//ProviderSettings ProviderConfig `mapstructure:"provider_settings"`
}

// ServiceConfig defines the krakend service
type KrakendServiceConfig struct {
	// set of endpoint definitions
	Endpoints []KrakendEndpointConfig `mapstructure:"endpoints"`
	// Processors configuration for handling data received or cached
	Processors ProcessorConfig `mapstructure:"processors"`
	// Tasks/Queue order to fetch and aggregate data
	// TaskQ []TaskQConfig `mapstructure:"taskq"`
	// defafult timeout
	Timeout time.Duration `mapstructure:"timeout"`
	// default TTL for GET
	CacheTTL time.Duration `mapstructure:"cache_ttl"`
	// default set of hosts
	Host []string `mapstructure:"host"`
	// port to bind the krakend service
	Port int `mapstructure:"port"`
	// version code of the configuration
	Version int `mapstructure:"version"`
	// run krakend in debug mode
	Debug bool `mapstructure:"debug"`
	// run krakend in debug mode
	Admin bool `mapstructure:"admin"`
	// URI Parsing object
	// uriParser URIParser
}

type KrakendEndpointConfig struct {
	// url pattern to be registered and exposed to the world
	Endpoint string `mapstructure:"endpoint"`
	// HTTP method of the endpoint (GET, POST, PUT, etc)
	Method string `mapstructure:"method"`
	// set of definitions of the backends to be linked to this endpoint
	Backend []KrakendBackend `mapstructure:"backend"`
	// HTTP headers of the endpoint to forward
	Header []string `mapstructure:"header"`
	// set of pre-processing rules to filters the response body with wildcards. If empty, the PreProcess id not used
	Paths []map[string]string `mapstructure:"paths"`
	// Tasks/Queue order to fetch and aggregate data
	// TaskQ []TaskQConfig `mapstructure:"taskq"`
	// set of variables extracted with custom rules. If empty, the holders are not used
	Holders []map[string]string `mapstructure:"holders"`
	// list of query string params to be extracted from the URI
	QueryStrings []string `mapstructure:"query_strings"`
	// number of concurrent calls this endpoint must send to the backends
	ConcurrentCalls int `mapstructure:"concurrent_calls"`
	// timeout of this endpoint
	Timeout time.Duration `mapstructure:"timeout"`
	// duration of the cache header
	CacheTTL time.Duration `mapstructure:"cache_ttl"`
	// list of query string params to be extracted from the URI
	QueryString []string `mapstructure:"querystring_params"`
	// Processors configuration for handling data received or cached
	Processors ProcessorConfig `mapstructure:"processors"`
	// Indexes configuration for matching, searching or caching behaviours at endpoint level
	// Indexes IndexConfig `mapstructure:"indexes"`
	// Entities configuration for graph datastore indexing
	Entities EntityConfig `mapstructure:"entities"`
	// Endpoint Extra configuration for customized behaviour
	ExtraConfig KrakendExtraConfig `mapstructure:"extra_config"`
	// run endpoint in debug mode
	Debug bool `mapstructure:"debug"`
}

type ProcessorConfig struct {
	Disabled  bool            `mapstructure:"disabled" default:"false"json:"disabled" yaml:"disabled" toml:"disabled"`
	Debug     bool            `mapstructure:"debug" default:"false"json:"debug" yaml:"debug" toml:"debug"`
	Formatter FormatterConfig `mapstructure:"formatter" json:"formatter" yaml:"formatter" toml:"formatter"`
	Indexer   IndexerConfig   `mapstructure:"indexer" json:"indexer" yaml:"indexer" toml:"indexer"`
	//TaskQ     []TaskQConfig   `mapstructure:"taskq" json:"taskq" yaml:"taskq" toml:"taskq"`
}

type IndexerConfig struct {
	Disabled bool   `mapstructure:"disabled" default:"false" json:"disabled" yaml:"disabled" toml:"disabled"`
	Debug    bool   `mapstructure:"debug" default:"false" json:"debug" yaml:"debug" toml:"debug"`
	Engine   string `mapstructure:"engine" default:"bleve" json:"engine" yaml:"engine" toml:"engine"`
	Format   string `mapstructure:"format" default:"json" json:"format" yaml:"format" toml:"format"`
	//TaskQ    []TaskQConfig `mapstructure:"taskq" json:"taskq" yaml:"taskq" toml:"taskq"`
}

type FormatterConfig struct {
	Disabled bool   `mapstructure:"disabled" default:"false" json:"disabled" yaml:"disabled" toml:"disabled"`
	Debug    bool   `mapstructure:"debug" default:"false" json:"debug" yaml:"debug" toml:"debug"`
	Engine   string `mapstructure:"engine" default:"mxj" json:"engine" yaml:"engine" toml:"engine"`
	Format   string `mapstructure:"format" default:"json" json:"format" yaml:"format" toml:"format"`
	//TaskQ    []TaskQConfig `mapstructure:"taskq" json:"taskq" yaml:"taskq" toml:"taskq"`
}

type EntityConfig struct {
	Disabled bool               `default:"false" json:"disabled" yaml:"disabled" toml:"disabled"`
	Debug    bool               `default:"false" json:"debug" yaml:"debug" toml:"debug"`
	Mappings []EntityMapsConfig `json:"mappings" yaml:"mappings" toml:"mappings"`
}

type EntityMapsConfig struct {
	Disabled bool   `default:"false" json:"disabled" yaml:"disabled" toml:"disabled"`
	Debug    bool   `default:"false" json:"debug" yaml:"debug" toml:"debug"`
	Key      string `json:"key" yaml:"key" toml:"key"`
	Map      string `json:"map" yaml:"map" toml:"map"`
	Type     string `json:"type" yaml:"type" toml:"type"`
	Format   string `json:"format" yaml:"format" toml:"format"`
}
