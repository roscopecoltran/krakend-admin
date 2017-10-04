package main

import (
	"github.com/jinzhu/gorm"
)

/*
  #downloads:
  #  prefix_path: ./shared/specs/registry/apis-guru/
  #  extension: ["yaml", "raml", "json"]
  #  overwrite: false
*/

type SpecsConfig struct {
	gorm.Model `json:"-" yaml:"-" toml:"-"`

	Disabled bool `gorm:"column:disabled" default:"false" storm:"disabled" json:"disabled" yaml:"disabled" toml:"disabled"`
	Debug    bool `gorm:"column:debug" default:"false" example:"false" storm:"debug" json:"debug" yaml:"debug" toml:"debug"`

	Registry  []SpecsRegistryAPI  `json:"registry,omitempty" yaml:"registry,omitempty" toml:"registry,omitempty"`
	Downloads SpecsDownloadConfig `json:"downloads,omitempty" yaml:"downloads,omitempty" toml:"downloads,omitempty"`

	// RAMLs   []RAMLConfig  `json:"ramls,omitempty" yaml:"ramls,omitempty" toml:"ramls,omitempty"`
	// OpenAPIs []OpenAPIConfig `json:"openapis,omitempty" yaml:"openapis,omitempty" toml:"openapis,omitempty"`
}

type SpecsDownloadConfig struct {
	gorm.Model `json:"-" yaml:"-" toml:"-"`
	Disabled   bool `gorm:"column:disabled" default:"false" storm:"disabled" json:"disabled" yaml:"disabled" toml:"disabled"`
	Debug      bool `gorm:"column:debug" default:"false" example:"false" storm:"debug" json:"debug" yaml:"debug" toml:"debug"`

	PrefixPath string `gorm:"column:prefix_path" default:"./shared/specs" json:"prefix_path,omitempty" yaml:"prefix_path,omitempty" toml:"prefix_path,omitempty" `
	Overwrite  bool   `gorm:"column:overwrite" default:"false" json:"overwrite,omitempty" yaml:"overwrite,omitempty" toml:"overwrite,omitempty" `
}

type SpecsRegistry struct {
	gorm.Model `json:"-" yaml:"-" toml:"-"`
	Disabled   bool `gorm:"column:disabled" default:"false" storm:"disabled" json:"disabled" yaml:"disabled" toml:"disabled"`
	Debug      bool `gorm:"column:debug" default:"false" example:"false" storm:"debug" json:"debug" yaml:"debug" toml:"debug"`

	Apis []SpecsRegistryAPI `json:"apis,omitempty" yaml:"apis,omitempty" toml:"apis,omitempty" `
}

type SpecsRegistryAPI struct {
	gorm.Model `json:"-" yaml:"-" toml:"-"`
	Disabled   bool `gorm:"column:disabled" default:"false" storm:"disabled" json:"disabled" yaml:"disabled" toml:"disabled"`
	Debug      bool `gorm:"column:debug" default:"false" example:"false" storm:"debug" json:"debug" yaml:"debug" toml:"debug"`

	Added     string                    `json:"added,omitempty" yaml:"added,omitempty" toml:"added,omitempty" `
	Name      string                    `json:"name,omitempty" yaml:"name,omitempty" toml:"name,omitempty" `
	Preferred string                    `json:"preferred,omitempty" yaml:"preferred,omitempty" toml:"preferred,omitempty" `
	Versions  []SpecsRegistryAPIVersion `json:"versions,omitempty" yaml:"versions,omitempty" toml:"versions,omitempty" `
}

type SpecsRegistryAPIVersion struct {
	gorm.Model `json:"-" yaml:"-" toml:"-"`
	Disabled   bool `gorm:"column:disabled" default:"false" storm:"disabled" json:"disabled" yaml:"disabled" toml:"disabled"`
	Debug      bool `gorm:"column:debug" default:"false" example:"false" storm:"debug" json:"debug" yaml:"debug" toml:"debug"`

	Added          string                      `json:"added,omitempty" yaml:"added,omitempty" toml:"added,omitempty" `
	Info           SpecsRegistryAPIVersionInfo `json:"info,omitempty" yaml:"info,omitempty" toml:"info,omitempty" `
	SwaggerURL     string                      `json:"swaggerUrl,omitempty" yaml:"swaggerUrl,omitempty" toml:"swaggerUrl,omitempty" `
	SwaggerYamlURL string                      `json:"swaggerYamlUrl,omitempty" yaml:"swaggerYamlUrl,omitempty" toml:"swaggerYamlUrl,omitempty" `
	Updated        string                      `json:"updated,omitempty" yaml:"updated,omitempty" toml:"updated,omitempty" `
	Version        string                      `json:"version,omitempty" yaml:"version,omitempty" toml:"version,omitempty" `
	OpenAPI        OpenApi                     `json:"openapi,omitempty" yaml:"openapi,omitempty" toml:"openapi,omitempty" `

	// OpenAPIParams is a list of parameters which the handler can accept
	// OpenAPIParams    []OpenApiParameter         `json:"openAPIParams,omitempty"`
	// OpenAPIResponses map[string]OpenApiResponse `json:"openAPIResponses,omitempty"`
}

type SpecsRegistryAPIVersionInfo struct {
	gorm.Model `json:"-" yaml:"-" toml:"-"`
	Disabled   bool `gorm:"column:disabled" default:"false" storm:"disabled" json:"disabled" yaml:"disabled" toml:"disabled"`
	Debug      bool `gorm:"column:debug" default:"false" example:"false" storm:"debug" json:"debug" yaml:"debug" toml:"debug"`

	Contact             SpecsRegistryAPIVersionInfoContact   `json:"contact,omitempty" yaml:"contact,omitempty" toml:"contact,omitempty" `
	Description         string                               `json:"description,omitempty" yaml:"description,omitempty" toml:"description,omitempty" `
	Title               string                               `json:"title,omitempty" yaml:"title,omitempty" toml:"title,omitempty" `
	Version             string                               `json:"version,omitempty" yaml:"version,omitempty" toml:"version,omitempty" `
	XApisguruCategories []string                             `json:"x-apisguru-categories,omitempty" yaml:"x-apisguru-categories,omitempty" toml:"x-apisguru-categories,omitempty" `
	XLogo               SpecsRegistryAPIVersionInfoXLogo     `json:"x-logo,omitempty" yaml:"x-logo,omitempty" toml:"x-logo,omitempty" `
	XOrigin             []SpecsRegistryAPIVersionInfoXOrigin `json:"x-origin,omitempty" yaml:"x-origin,omitempty" toml:"x-origin,omitempty" `
	XPreferred          bool                                 `json:"x-preferred,omitempty" yaml:"x-preferred,omitempty" toml:"x-preferred,omitempty" `
	XProviderName       string                               `json:"x-providerName,omitempty" yaml:"x-providerName,omitempty" toml:"x-providerName,omitempty" `
}

type SpecsRegistryAPIVersionInfoContact struct {
	gorm.Model `json:"-" yaml:"-" toml:"-"`
	Disabled   bool `gorm:"column:disabled" default:"false" storm:"disabled" json:"disabled" yaml:"disabled" toml:"disabled"`
	Debug      bool `gorm:"column:debug" default:"false" example:"false" storm:"debug" json:"debug" yaml:"debug" toml:"debug"`

	Email string `json:"email,omitempty" yaml:"email,omitempty" toml:"email,omitempty" `
	Name  string `json:"name,omitempty" yaml:"name,omitempty" toml:"name,omitempty" `
	URL   string `json:"url,omitempty" yaml:"url,omitempty" toml:"url,omitempty" `
}

type SpecsRegistryAPIVersionInfoXLogo struct {
	gorm.Model `json:"-" yaml:"-" toml:"-"`
	Disabled   bool `gorm:"column:disabled" default:"false" storm:"disabled" json:"disabled" yaml:"disabled" toml:"disabled"`
	Debug      bool `gorm:"column:debug" default:"false" example:"false" storm:"debug" json:"debug" yaml:"debug" toml:"debug"`

	BackgroundColor string `json:"backgroundColor,omitempty" yaml:"backgroundColor,omitempty" toml:"backgroundColor,omitempty" `
	URL             string `json:"url,omitempty" yaml:"url,omitempty" toml:"url,omitempty" `
}

type SpecsRegistryAPIVersionInfoXOrigin struct {
	gorm.Model `json:"-" yaml:"-" toml:"-"`
	Disabled   bool `gorm:"column:disabled" default:"false" storm:"disabled" json:"disabled" yaml:"disabled" toml:"disabled"`
	Debug      bool `gorm:"column:debug" default:"false" example:"false" storm:"debug" json:"debug" yaml:"debug" toml:"debug"`

	Format  string `json:"format,omitempty" yaml:"format,omitempty" toml:"format,omitempty" `
	URL     string `json:"url,omitempty" yaml:"url,omitempty" toml:"url,omitempty" `
	Version string `json:"version,omitempty" yaml:"version,omitempty" toml:"version,omitempty" `
}
