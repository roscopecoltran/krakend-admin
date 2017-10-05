package restful

import (
	"fmt"
	"path/filepath"
	"strings"

	// "github.com/roscopecoltran/krakend-admin/config"
	"github.com/roscopecoltran/krakend-admin/models/provider/restful/openapi"
	"github.com/roscopecoltran/krakend-admin/utils/download"
)

func ParseSpecs(apis []SpecsRegistryAPI, ignoreList []string) error {

	for _, api := range apis {
		var contentURL string
		var isSkip bool
		for _, v := range api.Versions {
			contentURL = fmt.Sprintf("%s", v.SwaggerYamlURL)
			if len(ignoreList) > 0 {
				for _, pattern := range ignoreList {
					if strings.Contains(contentURL, pattern) == true {
						fmt.Println("ParseSpecs(...): |- skipping url ", contentURL, " with pattern ", pattern)
						isSkip = true
					}
				}
			}
			if !isSkip {
				fileName, fullPath, fullName := download.GetFilePathFromURL(contentURL)
				filePathAbs, err := filepath.Abs(fullName)
				if err != nil {
					continue
				}
				fmt.Println("ParseSpecs(...): |- fileName: ", fileName, ", fullPath: ", fullPath, ", fullName: ", fullName, "filePathAbsolute: ", filePathAbs)
				// openapiSplit(filePathAbsolute, prefixPathAbsolute)
				openapi.ConvertToProtobuf(filePathAbs, fullPath, true)
			}
		}
	}
	return nil
}
