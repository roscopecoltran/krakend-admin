package main

import (
	// "errors"
	// "fmt"

	"github.com/Machiel/slugify"
)

func replaceSlugify(input string, replacementMap map[rune]string) string {
	replacementMap = map[rune]string{
		'a': "hello",
		'b': "hi",
	}
	replacementMapSlugifier := slugify.New(slugify.Configuration{
		ReplacementMap: replacementMap,
	})
	output := replacementMapSlugifier.Slugify(input)
	return output
}
