package main

import (
	//"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
	"strings"

	"github.com/roscopecoltran/configor"
)

func parseSpecsOpenAPI(apis []SpecsRegistryAPI) error {
	for _, api := range apis {
		var contentURL string
		for _, v := range api.Versions {
			contentURL = fmt.Sprintf("%s\n", v.SwaggerYamlURL)
			fileName, fullPath, fullName := getFilePathFromURL(contentURL)
			fmt.Println(" |- fileName: ", fileName, ", fullPath: ", fullPath, ", fullName: ", fullName)
			filePathAbsolute, err := filepath.Abs(fullName)
			if err != nil {
				continue
			}
			// openapiSplit(filePathAbsolute, prefixPathAbsolute)
			openapi2Protobuf(filePathAbsolute, fullPath, true)
		}
	}
	return nil
}

func openapiSplit(swaggerFile string, prefixPath string) {
	indexData, err := ioutil.ReadFile(swaggerFile)
	if err != nil {
		fmt.Println(err)
		return
	}
	responsesData := concatYamlFilesFromDir(*responsesDir)
	pathsData := concatYamlFilesFromDir(*pathsDir)
	definitionsData := concatYamlFilesFromDir(*definitionsDir)
	// print out the specification
	result := make([]string, 0)
	result = append(result, string(indexData)+"\n")
	if len(responsesData) != 0 {
		result = append(result, fmt.Sprintf("responses:\n%s\n", strings.Join(responsesData, "\n")))
	}
	if len(pathsData) != 0 {
		result = append(result, fmt.Sprintf("paths:\n%s\n", strings.Join(pathsData, "\n")))
	}
	if len(definitionsData) != 0 {
		result = append(result, fmt.Sprintf("definitions:\n%s\n", strings.Join(definitionsData, "\n")))
	}

	// resultStr := strings.Join(result, "\n")

	if len(responsesData) > 0 {
		if err := configor.Dump(responsesData, "swagger-response", "yaml,json", prefixPath); err != nil {
			log.Fatal("ERROR while dumping the swagger responses data:", err.Error())
		}
	}

	if len(pathsData) > 0 {
		if err := configor.Dump(pathsData, "swagger-paths", "yaml,json", prefixPath); err != nil {
			log.Fatal("ERROR while dumping the swagger paths data:", err.Error())
		}
	}

	if len(definitionsData) > 0 {
		if err := configor.Dump(definitionsData, "swagger-defs", "yaml,json", prefixPath); err != nil {
			log.Fatal("ERROR while dumping the swagger definitions data:", err.Error())
		}
	}

	/*
		switch strings.ToLower(format) {
		case "yaml":
			fmt.Println(resultStr)
			break

		case "json":
			yamlParsed, err := yaml2json.BytesToYAMLDoc([]byte(resultStr))
			if err != nil {
				printError(err)
			}
			jsonRaw, err := yaml2json.YAMLToJSON(yamlParsed)
			if err != nil {
				fmt.Println(yamlParsed)
				printError(err)
			}
			json, err := jsonRaw.MarshalJSON()
			if err != nil {
				printError(err)
			}
			fmt.Println(string(json))
			break

		default:
			fmt.Println("format not supported")
			return
		}
	*/
}
