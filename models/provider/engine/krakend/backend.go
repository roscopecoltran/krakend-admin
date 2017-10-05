package krakend

import (
	"time"
	// "github.com/roscopecoltran/krakend/encoding"
)

// Backend defines how krakend should connect to the backend service (the API resource to consume)
// and how it should process the received response
type Backend struct {
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
	Processors Processor `mapstructure:"processors"`
	// Indexes configuration for matching, searching or caching behaviours at Backend level
	// Indexes IndexConfig `mapstructure:"indexes"`
	// Entities configuration for graph datastore indexing
	Leafs []Path `mapstructure:"leafs"`
	// Backend Extra configuration for customized behaviours
	ExtraConfig ExtraConfig `mapstructure:"extra_config"`
	// Execute backend in debug mode
	Debug bool `default:"false" json:"debug" yaml:"debug" toml:"debug"`

	// Provider Name
	//ProviderName string `mapstructure:"provider_name"`
	// Provider Hook
	//ProviderHook string `mapstructure:"provider_hook"`
	// Provider Name
	//ProviderSettings ProviderConfig `mapstructure:"provider_settings"`
}
