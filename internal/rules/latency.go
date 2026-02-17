package rules

import (
	"math/rand"
	"time"
)

type LatencyAction struct {
	MinMS int
	MaxMS int
	Next  Action
}

func (a *LatencyAction) Execute(ctx *Context) {
	delay := a.MinMS
	if a.MaxMS > a.MinMS {
		delay += rand.Intn(a.MaxMS - a.MinMS)
	}
	time.Sleep(time.Duration(delay) * time.Millisecond)
	a.Next.Execute(ctx)
}
