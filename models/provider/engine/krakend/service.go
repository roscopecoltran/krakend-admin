package krakend

import (
	"time"
)

// ServiceConfig defines the krakend service
type Service struct {
	// set of endpoint definitions
	Endpoints []Endpoint `mapstructure:"endpoints"`
	// Processors configuration for handling data received or cached
	Processors Processor `mapstructure:"processors"`
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
