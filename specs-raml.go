package main

import (
	"fmt"

	"github.com/go-raml/raml"
	"github.com/k0kubun/pp"
)

func parseRAML(ramlFilePath string) {
	if apiDefinition, err := raml.ParseFile(ramlFilePath); err != nil {
		fmt.Printf("Failed parsing RAML file %s:\n  %s", ramlFilePath, err.Error())
	} else {
		fmt.Printf("Successfully parsed RAML file %s!\n\n", ramlFilePath)
		pp.Print(apiDefinition)
	}
}
