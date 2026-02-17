package rules

import (
	"math/rand"
	"strconv"
)

type RandomStatusAction struct {
	Statuses []int
	Next     Action
}

func (a *RandomStatusAction) Execute(ctx *Context) {
	if len(a.Statuses) == 0 {
		a.Next.Execute(ctx)
		return
	}
	status := a.Statuses[rand.Intn(len(a.Statuses))]
	ctx.Write(status, "Injected error: "+strconv.Itoa(status), nil)
}
