package rules

type ForwardAction struct {
	Target string
}

func (a *ForwardAction) Execute(ctx *Context) {
	ctx.ForwardTo(a.Target)
}
