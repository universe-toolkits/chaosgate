package rules

import "net/http"

type Matcher interface {
	Match(*http.Request) bool
}

type DefaultMatcher struct {
	Host       string
	PathPrefix string
	Method     string
}

func (m DefaultMatcher) Match(r *http.Request) bool {

	if m.Host != "" && r.Host != m.Host {
		return false
	}

	if m.Method != "" && r.Method != m.Method {
		return false
	}

	if m.PathPrefix != "" && len(r.URL.Path) >= len(m.PathPrefix) {
		return r.URL.Path[:len(m.PathPrefix)] == m.PathPrefix
	}

	return true
}
