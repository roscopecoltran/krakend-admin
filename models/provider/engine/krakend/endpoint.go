package krakend

import (
	"time"
)

type Endpoint struct {
	// url pattern to be registered and exposed to the world
	Endpoint string `mapstructure:"endpoint"`
	// HTTP method of the endpoint (GET, POST, PUT, etc)
	Method string `mapstructure:"method"`
	// set of definitions of the backends to be linked to this endpoint
	Backend []Backend `mapstructure:"backend"`
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
	Processors Processor `mapstructure:"processors"`
	// Indexes configuration for matching, searching or caching behaviours at endpoint level
	// Indexes IndexConfig `mapstructure:"indexes"`
	// Entities configuration for graph datastore indexing
	Leafs []Leaf `mapstructure:"leafs"`
	// Endpoint Extra configuration for customized behaviour
	ExtraConfig ExtraConfig `mapstructure:"extra_config"`
	// run endpoint in debug mode
	Debug bool `mapstructure:"debug"`
}
