package db

import (
	"reflect"
	"strings"

	"github.com/athom/suitecase"
)

func getVarTypeName(variable interface{}) string {
	if t := reflect.TypeOf(variable); t.Kind() == reflect.Ptr {
		return "*" + t.Elem().Name()
	} else {
		return t.Name()
	}
}

func ExtractMenuSectionName(camelInputStr string) string {
	if camelInputStr == "" {
		return "Global"
	}
	convStr := suitecase.ToSnakeCase(camelInputStr)
	parts := strings.Split(convStr, "_")
	// lastPart := info[len(parts)-1]
	if len(parts) > 0 {
		return parts[0]
	} else {
		return "Global"
	}
}
