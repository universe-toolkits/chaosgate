package rules

type DropAction struct{}

func (a *DropAction) Execute(ctx *Context) {
	// intentionally no response
}
