package sitemap

import (
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/stretchrcom/testify/assert"
)

/*
func Test_(t *testing.T) {

}

func Benchmark_(b *testing.B) {

}
*/

func Test_ParseURL_Normal(t *testing.T) {
	expected := `http://www.venuecom.com?beef=stew&whee=phewü`
	test_string := `http://www.venuecom.com?beef=stew&whee=phew%C3%BC`
	parsed, err := ParseURL(test_string)
	assert.Nil(t, err, fmt.Sprintf("unexpected error from ParseURL: %s", err))
	assert.Equal(t, parsed, expected, fmt.Sprintf("parsed item \"%s\" did not match expected \"%s\"", parsed, expected))
}

func Test_ParseURL_BadValue(t *testing.T) {
	expected := ""
	test_string := `www.venuecom.com?beef=stew&whee=phew%C3%BC`
	parsed, err := ParseURL(test_string)
	assert.NotNil(t, err, fmt.Sprintf("unexpected nil error from ParseURL on parsed item \"%s\"", test_string))
	assert.Equal(t, parsed, expected, fmt.Sprintf("parsed item \"%s\" did not match expected \"%s\"", parsed, expected))
}

func Test_AppendURL_EmptyList(t *testing.T) {
	expected := `www.venuecom.com`
	test_string := `www.venuecom.com`
	parsed, err := AppendURL(test_string, []string{})
	assert.Nil(t, err, fmt.Sprintf("unexpected error from AppendURL: %s", err))
	assert.Equal(t, parsed, expected, fmt.Sprintf("parsed item \"%s\" did not match expected \"%s\"", parsed, expected))
}

func Test_AppendURL_NilList(t *testing.T) {
	expected := `www.venuecom.com`
	test_string := `www.venuecom.com`
	parsed, err := AppendURL(test_string, nil)
	assert.Nil(t, err, fmt.Sprintf("unexpected error from AppendURL: %s", err))
	assert.Equal(t, parsed, expected, fmt.Sprintf("parsed item \"%s\" did not match expected \"%s\"", parsed, expected))
}

func Test_AppendURL_NoSlash(t *testing.T) {
	expected := `www.venuecom.com/end`
	test_string := `www.venuecom.com`
	test_append := `end`
	parsed, err := AppendURL(test_string, []string{test_append})
	assert.Nil(t, err, fmt.Sprintf("unexpected error from AppendURL: %s", err))
	assert.Equal(t, parsed, expected, fmt.Sprintf("parsed item \"%s\" did not match expected \"%s\"", parsed, expected))
}

func Test_AppendURL_Slash(t *testing.T) {
	expected := `www.venuecom.com/end`
	test_string := `www.venuecom.com/`
	test_append := `end`
	parsed, err := AppendURL(test_string, []string{test_append})
	assert.Nil(t, err, fmt.Sprintf("unexpected error from AppendURL: %s", err))
	assert.Equal(t, parsed, expected, fmt.Sprintf("parsed item \"%s\" did not match expected \"%s\"", parsed, expected))
}

func Test_AppendURL_ExtraSlash(t *testing.T) {
	expected := `www.venuecom.com//end`
	test_string := `www.venuecom.com/`
	test_append := `/end`
	parsed, err := AppendURL(test_string, []string{test_append})
	assert.Nil(t, err, fmt.Sprintf("unexpected error from AppendURL: %s", err))
	assert.Equal(t, parsed, expected, fmt.Sprintf("parsed item \"%s\" did not match expected \"%s\"", parsed, expected))
}

func Test_AppendURL_ExtraNoSlash(t *testing.T) {
	expected := `www.venuecom.com//end`
	test_string := `www.venuecom.com`
	test_append := `/end`
	parsed, err := AppendURL(test_string, []string{test_append})
	assert.Nil(t, err, fmt.Sprintf("unexpected error from AppendURL: %s", err))
	assert.Equal(t, parsed, expected, fmt.Sprintf("parsed item \"%s\" did not match expected \"%s\"", parsed, expected))
}

func Test_AppendURL_Multiple(t *testing.T) {
	expected := `www.venuecom.com/end/of/world`
	test_string := `www.venuecom.com`
	test_append := []string{`end`, `of`, `world`}
	parsed, err := AppendURL(test_string, test_append)
	assert.Nil(t, err, fmt.Sprintf("unexpected error from AppendURL: %s", err))
	assert.Equal(t, parsed, expected, fmt.Sprintf("parsed item \"%s\" did not match expected \"%s\"", parsed, expected))
}

func Test_parsePriority_Normal(t *testing.T) {
	expected := 0.75
	test_string := `0.75`
	test_parsePriority(t, expected, test_string)
}

func Test_parsePriority_Ceiling(t *testing.T) {
	expected := 1.0
	test_string := `1.75`
	test_parsePriority(t, expected, test_string)
}

func Test_parsePriority_Floor(t *testing.T) {
	expected := 0.0
	test_string := `-0.75`
	test_parsePriority(t, expected, test_string)
}

func test_parsePriority(t *testing.T, expected float64, test_string string) {
	parsed, err := parsePriority(test_string)
	assert.Nil(t, err, fmt.Sprintf("unexpected error from parsePriority: %s", err))
	assert.Equal(t, parsed, expected, fmt.Sprintf("parsed item \"%s\" did not match expected \"%s\"", parsed, expected))
}

func Test_parsePriority_BadValue(t *testing.T) {
	expected := 1.0
	test_string := `asdf`
	parsed, err := parsePriority(test_string)
	assert.NotNil(t, err, fmt.Sprintf("unexpected nil error from parsePriority on parsed item \"%s\"", test_string))
	assert.Equal(t, parsed, expected, fmt.Sprintf("parsed item \"%s\" did not match expected \"%s\"", parsed, expected))
}

func Test_parseChangeFrequency_Normal(t *testing.T) {
	valid_values := []string{"always", "hourly", "daily", "weekly", "monthly", "yearly", "never"}
	for index := 0; index < len(valid_values); index++ {
		expected := valid_values[index]
		test_string := valid_values[index]
		test_parseChangeFrequency(t, expected, test_string)
	}
}

func Test_parseChangeFrequency_LowerCase(t *testing.T) {
	valid_values := []string{"always", "hourly", "daily", "weekly", "monthly", "yearly", "never"}
	for index := 0; index < len(valid_values); index++ {
		expected := valid_values[index]
		test_string := strings.ToUpper(valid_values[index])
		test_parseChangeFrequency(t, expected, test_string)
	}
}

func Test_parseChangeFrequency_BadValue(t *testing.T) {
	expected := ""
	test_string := `asdf`
	parsed, err := parseChangeFrequency(test_string)
	assert.NotNil(t, err, fmt.Sprintf("unexpected nil error from parseChangeFrequency on parsed item \"%s\"", test_string))
	assert.Equal(t, parsed, expected, fmt.Sprintf("parsed item \"%s\" did not match expected \"%s\"", parsed, expected))
}

func test_parseChangeFrequency(t *testing.T, expected string, test_string string) {
	parsed, err := parseChangeFrequency(test_string)
	assert.Nil(t, err, fmt.Sprintf("unexpected error from parseChangeFrequency: %s", err))
	assert.Equal(t, parsed, expected, fmt.Sprintf("parsed item \"%s\" did not match expected \"%s\"", parsed, expected))
}

func Test_parseLastModified_Normal(t *testing.T) {
	expected := time.Date(2013, time.November, 27, 0, 6, 52, 0, time.FixedZone("-0500", -5*60*60))
	test_string := `2013-11-27T00:06:52-05:00`
	parsed, err := parseLastModified(test_string)
	assert.Nil(t, err, fmt.Sprintf("unexpected error from parseLastModified: %s", err))
	assert.Equal(t, parsed.String(), expected.String(), fmt.Sprintf("parsed item \"%s\" did not match expected \"%s\"", parsed, expected))
}

func Test_parseLastModified_BadValue(t *testing.T) {
	expected := time.Time{}
	test_string := `asdf`
	parsed, err := parseLastModified(test_string)
	assert.NotNil(t, err, fmt.Sprintf("unexpected nil error from parseLastModified on parsed item \"%s\"", test_string))
	assert.Equal(t, parsed.String(), expected.String(), fmt.Sprintf("parsed item \"%s\" did not match expected \"%s\"", parsed, expected))
}
