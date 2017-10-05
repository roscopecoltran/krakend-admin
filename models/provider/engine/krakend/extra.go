package krakend

type ExtraConfig map[string]interface{}

// ConfigGetter is a function for parsing ExtraConfig into a previously know type
type ConfigGetter func(ExtraConfig) interface{}

// DefaultConfigGetter is the Default implementation for ConfigGetter, it just returns the KrakendExtraConfig map.
func DefaultConfigGetter(extra ExtraConfig) interface{} { return extra }
