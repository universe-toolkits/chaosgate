package rules

import "net/http"

type HeaderMatcher struct {
	Key   string
	Value string
}

func (m HeaderMatcher) Match(r *http.Request) bool {
	v := r.Header.Get(m.Key)
	return v == m.Value
}
