package rules

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	"github.com/PaesslerAG/jsonpath"
)

type JSONPathMatcher struct {
	Expression string
	Expected   interface{}
}

func (m JSONPathMatcher) Match(r *http.Request) bool {
	if r.Body == nil {
		return false
	}

	bodyBytes, _ := io.ReadAll(r.Body)
	r.Body = io.NopCloser(bytes.NewReader(bodyBytes)) // reset body

	var data interface{}
	if err := json.Unmarshal(bodyBytes, &data); err != nil {
		return false
	}

	res, err := jsonpath.Get(m.Expression, data)
	if err != nil {
		return false
	}

	return res == m.Expected
}
