package go_validate

import (
	"strings"

	"github.com/xloss/go-validate/rules"
)

type Rule interface {
	GetName() string
	GetValues() map[string]any
	Validate(field string, value any, data map[string]any) bool
}

type Rewriter interface {
	Rewrite(value any) (func(typ string) any, bool)
}

func nameToRule(rule string) Rule {
	params := strings.SplitN(rule, ":", 2)

	switch params[0] {
	case "required":
		return &rules.Required{}
	case "string":
		return &rules.String{}
	case "integer":
		return &rules.Integer{}
	case "numeric":
		return &rules.Numeric{}
	case "boolean":
		return &rules.Boolean{}
	case "min":
		if len(params) != 2 {
			errRule := &rules.Error{}
			errRule.AddParams(rule)

			return errRule
		}

		r := rules.Min{}
		err := r.AddParams(params[1])
		if err != nil {
			errRule := &rules.Error{}
			errRule.AddParams(rule)
			errRule.AddError(err)

			return errRule
		}

		return &r
	case "max":
		if len(params) != 2 {
			errRule := &rules.Error{}
			errRule.AddParams(rule)

			return errRule
		}

		r := rules.Max{}
		err := r.AddParams(params[1])
		if err != nil {
			errRule := &rules.Error{}
			errRule.AddParams(rule)
			errRule.AddError(err)

			return errRule
		}

		return &r
	case "domain":
		return &rules.Domain{}
	case "date":
		return &rules.Date{}
	case "date_format":
		if len(params) != 2 {
			errRule := &rules.Error{}
			errRule.AddParams(rule)

			return errRule
		}

		r := rules.Date{}
		err := r.AddParams(params[1])
		if err != nil {
			errRule := &rules.Error{}
			errRule.AddParams(rule)
			errRule.AddError(err)

			return errRule
		}

		return &r
	case "email":
		return &rules.Email{}
	case "confirmed":
		return &rules.Confirmed{}
	case "accepted":
		return &rules.Accepted{}
	case "uuid":
		return &rules.UUID{}
	case "array":
		return &rules.Array{}
	case "json":
		return &rules.JSON{}
	}

	errRule := &rules.Error{}
	errRule.AddParams(rule)
	return errRule
}
