package rules

import (
	"time"
)

type DelayAction struct {
	Milliseconds int
	Next         Action
}

func (a *DelayAction) Execute(ctx *Context) {
	time.Sleep(time.Duration(a.Milliseconds) * time.Millisecond)
	a.Next.Execute(ctx)
}
