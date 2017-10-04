package convert

import "testing"

type testpair struct {
	camelcase  string
	underscore string
}
type constestpair struct {
	letter  rune
	boolean bool
}

var constests = []constestpair{
	{'A', true},
	{'i', false},
}
var tests = []testpair{
	{"UnderScore", "under_score"},
	{"NewTest", "new_test"},
	{"CamelCase", "camel_case"},
}

func TestUnderscore(t *testing.T) {
	for _, pair := range tests {
		v := Underscore(pair.camelcase)
		if v != pair.underscore {
			t.Error(
				"For", pair.camelcase,
				"expected", pair.underscore,
				"got", v,
			)
		}
	}
}

func TestIsConsonant(t *testing.T) {
	for _, pair := range constests {
		v := IsConsonant(pair.letter)
		if v != pair.boolean {
			t.Error(
				"For", pair.letter,
				"expected", pair.boolean,
				"got", v,
			)
		}
	}
}
