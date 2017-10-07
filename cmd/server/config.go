package main

import (
	"os"

	"github.com/qor/help"
	i18n_database "github.com/qor/i18n/backends/database"
	"github.com/qor/media/asset_manager"

	"github.com/roscopecoltran/krakend-admin/models/classify"
	"github.com/roscopecoltran/krakend-admin/models/locale"
	"github.com/roscopecoltran/krakend-admin/models/machine"

	"github.com/roscopecoltran/krakend-admin/models/provider"
	"github.com/roscopecoltran/krakend-admin/models/provider/engine/gateway"
	"github.com/roscopecoltran/krakend-admin/models/provider/engine/krakend"
	"github.com/roscopecoltran/krakend-admin/models/provider/restful"
)

type Config struct {
	Debug    bool   `default:"false" storm:"debug" json:"debug" yaml:"debug" toml:"debug"`
	Instance string `storm:"instance" json:"instance" yaml:"instance" toml:"instance"`

	Patterns struct {
		Ignore []string `json:"ignore" yaml:"ignore" toml:"ignore"`
	} `json:"patterns" yaml:"patterns" toml:"patterns"`

	//Env        []EnvConfig      `gorm:"many2many:gateway_local_envs;" storm:"env" json:"env" yaml:"env" file:"environment" toml:"env"`
	Service   []gateway.Service   `storm:"service" json:"service" yaml:"service" toml:"service"`
	Backends  []gateway.Export    `json:"backends" yaml:"backends" file:"backends" toml:"backends"`
	Providers []provider.Provider `json:"providers" yaml:"providers" file:"providers" toml:"providers"`
	Krakend   krakend.Service     `json:"krakend" yaml:"krakend" file:"krakend" toml:"krakend"`
	Specs     restful.Specs       `json:"specs" yaml:"specs" file:"specs" toml:"specs"`
}

const (
	_sep = string(os.PathSeparator)
)

var defaultConfigFiles = []string{
	"shared/testdata/patterns.yaml",
	"shared/testdata/service.yaml",
	"shared/testdata/downloads.yaml",
	"shared/testdata/specs.yaml",
	"shared/testdata/file_types.yaml",
	"shared/testdata/engines.yaml",
	"shared/testdata/specs/api-gurus.yaml",
	"shared/testdata/providers.yaml",
}

var defaultModels = []interface{}{
	&i18n_database.Translation{},
	&classify.Category{}, &classify.Tag{}, &classify.Topic{},
	&gateway.Backend{}, &gateway.Endpoint{}, &gateway.Service{},
}

var defaultModelsFull = []interface{}{

	&i18n_database.Translation{},
	&help.QorHelpEntry{},
	&asset_manager.AssetManager{},

	&restful.Specs{}, &restful.SpecsRegistry{},
	// &SpecsRegistryAPIVersion{}, &SpecsRegistryAPIVersionInfo{}, &SpecsRegistryAPIVersionInfoContact{}, &SpecsRegistryAPIVersionInfoXOrigin{}, &SpecsRegistryAPIVersionInfoXLogo{},

	&local.File{}, &local.Environment{}, &local.EnvironmentVariable{},
	&locales.Country{}, &locales.Currency{},

	&classify.Category{}, &classify.Tag{}, &classify.Topic{},

	&provider.Provider{}, &provider.CategoryToKeyword{}, &provider.CountryToHostname{}, &provider.QueryPattern{}, &provider.Path{}, &provider.Path{}, &provider.PathExtraBlocks{}, &provider.URL{},

	&gateway.Backend{}, &gateway.Endpoint{}, &gateway.Service{},
	&gateway.Header{}, &gateway.QueryString{}, &gateway.Parameter{}, &gateway.Blacklist{}, &gateway.Whitelist{},
	&gateway.UrlKey{}, &gateway.Holder{}, &gateway.Mapping{},

	&local.TaskQueue{}, &local.TaskQueueRunner{},
}
