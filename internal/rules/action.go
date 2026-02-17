package rules

type Action interface {
	Execute(*Context)
}
