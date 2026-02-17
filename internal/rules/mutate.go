package rules

import (
	"encoding/json"
	"math/rand"
)

type MutateAction struct {
	Mutations []Mutation
	Next      Action
}

type Mutation struct {
	Field      string
	TypeChange string // string, int, bool
	Random     bool
}

func (a *MutateAction) Execute(ctx *Context) {
	// Read body
	body := ctx.Body() // می‌تونی اضافه کنیم در Context متد Body() که request یا response body رو برگردونه
	var data map[string]interface{}
	json.Unmarshal(body, &data)

	// Apply mutations
	for _, m := range a.Mutations {
		if m.Random {
			switch m.TypeChange {
			case "int":
				data[m.Field] = rand.Intn(1000)
			case "string":
				data[m.Field] = "random_value"
			case "bool":
				data[m.Field] = rand.Intn(2) == 0
			}
		} else {
			data[m.Field] = nil
		}
	}

	newBody, _ := json.Marshal(data)
	ctx.Write(200, string(newBody), nil)
}
