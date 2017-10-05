package openapi

import (
	"github.com/emicklei/go-restful-swagger12"
	"sort"
)

func AscendingByResourcePath(apiDeclarationMap *swagger.ApiDeclarationList) {
	// First get the offset and resource path
	paths := make([]string, len(apiDeclarationMap.List))
	for k, apiDeclaration := range apiDeclarationMap.List {
		paths[k] = apiDeclaration.ResourcePath
	}

	// Next, sort the values
	sort.Strings(paths)

	// Then, create our new list
	newList := make([]swagger.ApiDeclaration, len(apiDeclarationMap.List))

	// Iterate over some stuff and populate the new list...
	for i, resourcePath := range paths {
		for _, apiDeclaration := range apiDeclarationMap.List {
			if resourcePath == apiDeclaration.ResourcePath {
				newList[i] = apiDeclaration
				break
			}
		}
	}

	// Finally, override list definition with our newly sorted list.
	apiDeclarationMap.List = newList
}
