package rules

import (
	"net/http"
	"regexp"
)

type RegexMatcher struct {
	Pattern *regexp.Regexp
	Target  string // path/header/body
}

func (m RegexMatcher) Match(r *http.Request) bool {
	var val string
	switch m.Target {
	case "path":
		val = r.URL.Path
	case "method":
		val = r.Method
	default:
		return false
	}

	return m.Pattern.MatchString(val)
}
