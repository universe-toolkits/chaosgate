package rules

import (
	"math/rand"
	"sync"
)

type Engine struct {
	mu    sync.RWMutex
	rules []Rule
}

func NewEngine(r []Rule) *Engine {
	return &Engine{rules: r}
}

func (e *Engine) Update(r []Rule) {
	e.mu.Lock()
	defer e.mu.Unlock()
	e.rules = r
}

func (e *Engine) Execute(ctx *Context) {
	e.mu.RLock()
	defer e.mu.RUnlock()

	for _, rule := range e.rules {

		if rule.Percentage > 0 {
			if rand.Float64()*100 > rule.Percentage {
				continue
			}
		}

		if rule.Match(ctx.Request) {
			rule.Action.Execute(ctx)
			return
		}
	}

	ctx.Forward()
}
