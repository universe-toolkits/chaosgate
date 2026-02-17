package rules

import "net/http"

type QueryMatcher struct {
	Key   string
	Value string
}

func (m QueryMatcher) Match(r *http.Request) bool {
	v := r.URL.Query().Get(m.Key)
	return v == m.Value
}
