package openapi

import (
	"fmt"
	"io/ioutil"
	// "log"
	"os"
	"strings"

	"github.com/roscopecoltran/configor"
	"github.com/roscopecoltran/krakend-admin/utils/convert"
)

var (
	ResponsesDir   string = "./shared/specs/responses"
	PathsDir       string = "./shared/specs/paths"
	DefinitionsDir string = "./shared/specs/definitions"
)

func Split(swaggerFile string, prefixPath string) {

	_, err := os.Stat(swaggerFile)
	if err != nil {
		fmt.Printf("Split(...): File doesn't exist, skipping\n", swaggerFile)
		return
	}

	indexData, err := ioutil.ReadFile(swaggerFile)
	if err != nil {
		fmt.Println("error: ", err)
		return
	}
	responsesData := convert.ConcatYamlFilesFromDir(ResponsesDir)
	pathsData := convert.ConcatYamlFilesFromDir(PathsDir)
	definitionsData := convert.ConcatYamlFilesFromDir(DefinitionsDir)
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
			fmt.Println("ERROR while dumping the swagger responses data:", err.Error())
		}
	}
	if len(pathsData) > 0 {
		if err := configor.Dump(pathsData, "swagger-paths", "yaml,json", prefixPath); err != nil {
			fmt.Println("ERROR while dumping the swagger paths data:", err.Error())
		}
	}
	if len(definitionsData) > 0 {
		if err := configor.Dump(definitionsData, "swagger-defs", "yaml,json", prefixPath); err != nil {
			fmt.Println("ERROR while dumping the swagger definitions data:", err.Error())
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
