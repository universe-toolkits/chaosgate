package rules

import "net/http"

type Rule struct {
	Name       string
	Percentage float64
	Matchers   []Matcher
	Action     Action
}

func (r Rule) Match(req *http.Request) bool {
	for _, m := range r.Matchers {
		if !m.Match(req) {
			return false
		}
	}
	return true
}
