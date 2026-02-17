package rules

type MockAction struct {
	Status  int
	Body    string
	Headers map[string]string
}

func (a *MockAction) Execute(ctx *Context) {
	ctx.Write(a.Status, a.Body, a.Headers)
}
