package convert

import (
	"strconv"
	"strings"
	// "github.com/coocood/qbs"
	// "github.com/serenize/snaker"
	// strcase "github.com/stoewer/go-strcase"
)

func Underscore(camelcase string) string {
	returnstring := ""
	for _, letter := range camelcase {
		if IsConsonant(letter) {
			returnstring += "_"
			runetostring, _ := strconv.Unquote(strconv.QuoteRuneToASCII(letter))
			returnstring += strings.ToLower(runetostring)
		} else {
			runetostring, _ := strconv.Unquote(strconv.QuoteRuneToASCII(letter))
			returnstring += runetostring
		}
	}
	return strings.Trim(returnstring, "_")
}

func IsConsonant(s rune) bool {
	return strings.ContainsRune("ABCDEFGHIJKLMNOPQRSTUVWXYZ", s)
}

/*
func ConvertCase(input string) string {
	qbs.ColumnNameToFieldName = snaker.SnakeToCamel
	qbs.FieldNameToColumnName = snaker.CamelToSnake
	return qbs.CamelToSnake(input)

	return strcase.SnakeCase(input)
}
*/
