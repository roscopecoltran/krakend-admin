package openapi

import (
	// 	"time"
	// "github.com/qor/media/oss"
	"github.com/jinzhu/gorm"
)

// OpenAPI is the root swagger document itself
type OpenAPI struct {
	gorm.Model `json:"-" yaml:"-" toml:"-"`

	// SwaggerVersion, as of this writing, is usually "2.0"
	SwaggerVersion string `gorm:"column:swagger" json:"swagger,omitempty" yaml:"swagger,omitempty" toml:"swagger,omitempty"`
	// Info is a section defining basic information about the API
	Info Info `gorm:"column:info" json:"info,omitempty" yaml:"info,omitempty" toml:"info,omitempty"`
	// Host is the FQDN of the server for API requests
	Host string `gorm:"column:host" json:"host,omitempty" yaml:"host,omitempty" toml:"host,omitempty"`
	// Schemes is a list of accepted schemes for API requests (http, https)
	Schemes []string `gorm:"column:schemes" json:"schemes,omitempty" yaml:"schemes,omitempty" toml:"schemes,omitempty"`
	// BasePath is a path to append to the Host in all API requests, if any
	BasePath string `gorm:"column:basePath" json:"basePath,omitempty" yaml:"basePath,omitempty" toml:"basePath,omitempty"`
	// Produces is a list of generated mime types for responses
	Produces []string `gorm:"column:produces" json:"produces,omitempty" yaml:"produces,omitempty" toml:"produces,omitempty"`
	// Paths is a map of string paths with Path specifications for API requests
	Paths map[string]map[string]Request `gorm:"column:paths" json:"paths,omitempty" yaml:"paths,omitempty" toml:"paths,omitempty"`
	// Definitions define type definitions for use in requests and responses
	Definitions map[string]Definition `gorm:"column:definitions" json:"definitions,omitempty" yaml:"definitions,omitempty" toml:"definitions,omitempty"`
	// Security indicates the required security definition for operations performed
	Security []map[string][]string `gorm:"column:security" json:"security,omitempty" yaml:"security,omitempty" toml:"security,omitempty"`
	// SecurityDefinitions specifies security requirements
	SecurityDefinitions map[string]SecurityDefinition `gorm:"column:securityDefinitions" json:"securityDefinitions,omitempty" yaml:"securityDefinitions,omitempty" toml:"securityDefinitions,omitempty"`
}

// Info is an OpenAPI section for basic information about the API
type Info struct {
	gorm.Model `json:"-" yaml:"-" toml:"-"`

	// Title is the short name of the API
	Title string `gorm:"column:title" json:"title,omitempty" yaml:"title,omitempty" toml:"title,omitempty"`
	// Description is a description of the API
	Description string `gorm:"column:description" json:"description,omitempty" yaml:"description,omitempty" toml:"description,omitempty"`
	// Version is the API's current version
	Version string `gorm:"column:version" json:"version,omitempty" yaml:"version,omitempty" toml:"version,omitempty"`
}

// Request is a specification of request parameters and documentation for an HTTP verb and its request
type Request struct {
	gorm.Model `json:"-" yaml:"-" toml:"-"`

	// Summary is a title for the HTTP request
	Summary string `gorm:"column:summary" json:"summary,omitempty" yaml:"summary,omitempty" toml:"summary,omitempty"`
	// Description is a description of the HTTP request
	Description string `gorm:"column:description" json:"description,omitempty" yaml:"description,omitempty" toml:"description,omitempty"`
	// Parameters is a list of parameters for the HTTP request
	Parameters []Parameter `gorm:"column:parameters" json:"parameters,omitempty" yaml:"parameters,omitempty" toml:"parameters,omitempty"`
	// Tags is an array of tags for the request
	Tags []string `gorm:"column:tags" json:"tags,omitempty" yaml:"tags,omitempty" toml:"tags,omitempty"`
	// Responses houses possible responses to a request
	Responses map[string]Response `gorm:"column:responses" json:"responses,omitempty" yaml:"responses,omitempty" toml:"responses,omitempty"`
}

// Parameter is an HTTP request parameter
type Parameter struct {
	gorm.Model `json:"-" yaml:"-" toml:"-"`

	// Name is a name for the parameter
	Name string `gorm:"column:name" json:"name,omitempty" yaml:"name,omitempty" toml:"name,omitempty"`
	// In defines where the parameter is specified. One of "query", "header", "path", "formData", or "body"
	In string `gorm:"column:in" json:"in,omitempty" yaml:"in,omitempty" toml:"in,omitempty"`
	// Description is a text description of the parameter
	Description string `gorm:"column:description" json:"description,omitempty" yaml:"description,omitempty" toml:"description,omitempty"`
	// Required indicates whether the parameter is required or not
	Required bool `gorm:"column:required" json:"required,omitempty" yaml:"required,omitempty" toml:"required,omitempty"`
	// Type defines the type of data being inputted and is not needed if Query=="body". The value MUST
	// be one of "string", "number", "integer", "boolean", "array" or "file". If type is "file", the consumes
	// MUST be either "multipart/form-data", " application/x-www-form-urlencoded" or both and the parameter
	// MUST be in "formData".
	Type string `gorm:"column:type" json:"type,omitempty" yaml:"type,omitempty" toml:"type,omitempty"`
	// Format further specifies the format of the data in the parameter and is not needed if Query=="body".
	// See https://github.com/OAI/OpenAPI-Specification/blob/master/versions/2.0.md#dataTypeFormat
	Format string `gorm:"column:format" json:"format,omitempty" yaml:"format,omitempty" toml:"format,omitempty"`
	// Schema declares the schema for body parameters
	Schema *Schema `gorm:"column:schema" json:"schema,omitempty" yaml:"schema,omitempty" toml:"schema,omitempty"`
}

// Response is an individual response for an HTTP request
type Response struct {
	gorm.Model `json:"-" yaml:"-" toml:"-"`

	// Description is a short description of the response
	Description string `gorm:"column:description" json:"description,omitempty" yaml:"description,omitempty" toml:"description,omitempty"`
	// Schema is the Schema for the returned response
	Schema Schema `gorm:"column:schema" json:"schema,omitempty" yaml:"schema,omitempty" toml:"schema,omitempty"`
}

// Schema is a response schema definition
type Schema struct {
	gorm.Model `json:"-" yaml:"-" toml:"-"`

	// Type is the type of response returned
	Type string `gorm:"column:type" json:"type,omitempty" yaml:"type,omitempty" toml:"type,omitempty"`
	// Items is an item type reference for the response
	Items *ItemRef `gorm:"column:items" json:"items,omitempty" yaml:"items,omitempty" toml:"items,omitempty"`
	// Ref is a reference (if the type is not an array)
	Ref string `json:"$ref,omitempty"`
}

// ItemRef is a reference to the type of item used in a request or a response
type ItemRef struct {
	gorm.Model `json:"-" yaml:"-" toml:"-"`

	// Ref indicates a reference to a definition (if any)
	Ref string `json:"$ref,omitempty"`
	// Type indicates the type of items for an array
	Type string `gorm:"column:type" json:"type,omitempty" yaml:"type,omitempty" toml:"type,omitempty"`
}

// Definition is a type definition for an HTTP request or response
type Definition struct {
	gorm.Model `json:"-" yaml:"-" toml:"-"`

	// Type is the type of data being used (usually "object", "string", "integer", "array")
	Type string `gorm:"column:type" json:"type,omitempty" yaml:"type,omitempty" toml:"type,omitempty"`
	// Properties is a list of properties of an object
	Properties map[string]Property `gorm:"column:properties" json:"properties,omitempty" yaml:"properties,omitempty" toml:"properties,omitempty"`
	// Required is an array of strings representing the names of the required parameters
	Required []string `gorm:"column:required" json:"required,omitempty" yaml:"required,omitempty" toml:"required,omitempty"`
}

// Property is a definition of a property of a Definition
type Property struct {
	gorm.Model `json:"-" yaml:"-" toml:"-"`

	// Type is a data type. See https://github.com/OAI/OpenAPI-Specification/blob/master/versions/2.0.md#data-types
	Type string `gorm:"column:type" json:"type,omitempty" yaml:"type,omitempty" toml:"type,omitempty"`
	// Format is the data type format for an Property
	Format string `gorm:"column:format" json:"format,omitempty" yaml:"format,omitempty" toml:"format,omitempty"`
	// Description is a description of the property
	Description string `gorm:"column:description" json:"description,omitempty" yaml:"description,omitempty" toml:"description,omitempty"`
	// Items is an option ItemRef for the items in the Property if the Property references another item type
	Items *ItemRef `gorm:"column:items" json:"items,omitempty" yaml:"items,omitempty" toml:"items,omitempty"`
	// Ref is a reference for the type of object if the Type=="object"
	Ref string `json:"$ref,omitempty"`
}

type SecurityDefinition struct {
	gorm.Model `json:"-" yaml:"-" toml:"-"`

	// Type determines the type of security used
	Type string `gorm:"column:type" json:"type,omitempty" yaml:"type,omitempty" toml:"type,omitempty"`
	// Name is the name of the parameter used for security enforcement
	Name string `gorm:"column:name" json:"name,omitempty" yaml:"name,omitempty" toml:"name,omitempty"`
	// In indicates one of "query" or "header" depending on where the security is verfied in the request
	In string `gorm:"column:in" json:"in,omitempty" yaml:"in,omitempty" toml:"in,omitempty"`
}

// Definer is an interface for objects that can define Definition specifications
type Definer interface {
	OpenAPIDefinitions() map[string]Definition
}
