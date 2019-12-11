package util

import (
	"regexp"
)

// MCRegexp model
type MCRegexp struct {
	Regexp     *regexp.Regexp
	Match      []string
	CompileStr string
	SearchStr  string
}

// NewMCRegExp initialize regular expresion object
func NewMCRegExp(cs string, ss string) *MCRegexp {
	r := regexp.MustCompile(cs)
	match := r.FindAllString(ss, -1)
	return &MCRegexp{
		Regexp:     r,
		Match:      match,
		CompileStr: cs,
		SearchStr:  ss,
	}
}

// Count return count result
func (mcregexp *MCRegexp) Count() int {
	return len(mcregexp.Match)
}

// Results return match string result
func (mcregexp *MCRegexp) Results() []string {
	return mcregexp.Match
}
