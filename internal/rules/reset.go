package rules

import (
	"net/http"
)

type ResetConnectionAction struct{}

func (a *ResetConnectionAction) Execute(ctx *Context) {
	hj, ok := ctx.Writer.(http.Hijacker)
	if !ok {
		return
	}
	conn, _, err := hj.Hijack()
	if err != nil {
		return
	}
	conn.Close() // Force reset mid-stream
}
