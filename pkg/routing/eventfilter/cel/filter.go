package cel

import (
	"context"

	cloudevents "github.com/cloudevents/sdk-go/v2"
	"github.com/google/cel-go/cel"
	"github.com/tidwall/gjson"

	"github.com/zeiss/typhoon/pkg/routing/eventfilter"
)

// ConditionalFilter structure holds both CEL Program and variable definitions
// so that it can be evaluated with the new variable values
type ConditionalFilter struct {
	Expression *cel.Program
	Variables  []Variable
}

// Variable contains the meta data required to parse event payload and execute
// CEL Program
type Variable struct {
	Name string
	Path string
	Type string
}

// Filter parses Event payload values defined as the expression variables, asserts their types,
// and executes CEL Program. If expression result is true, Event passes the filter.
func (c *ConditionalFilter) Filter(ctx context.Context, event cloudevents.Event) eventfilter.FilterResult {
	vars := make(map[string]interface{})

	for _, v := range c.Variables {
		switch v.Type {
		case "bool":
			vars[v.Name] = gjson.GetBytes(event.Data(), v.Path).Bool()
		case "int64":
			vars[v.Name] = gjson.GetBytes(event.Data(), v.Path).Int()
		case "uint64":
			vars[v.Name] = gjson.GetBytes(event.Data(), v.Path).Uint()
		case "double":
			vars[v.Name] = gjson.GetBytes(event.Data(), v.Path).Float()
		case "string":
			vars[v.Name] = gjson.GetBytes(event.Data(), v.Path).String()
		}
	}

	pass, err := eval(*c.Expression, vars)
	if err != nil || pass {
		return eventfilter.PassFilter
	}

	return eventfilter.FailFilter
}

// eval evaluates precompiled Expression with passed variables
func eval(program cel.Program, vars map[string]interface{}) (bool, error) {
	out, _, err := program.Eval(vars)
	return out.Value().(bool), err
}
