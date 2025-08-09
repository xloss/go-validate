package go_validate

import (
	"strings"

	"github.com/xloss/go-validate/rules"
)

type Rule interface {
	GetName() string
	GetValues() map[string]any
	Validate(value any) bool
}

func nameToRule(rule string) Rule {
	params := strings.Split(rule, ":")

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
	case "domain":
		return &rules.Domain{}
	case "date":
		return &rules.Date{}
	case "email":
		return &rules.Email{}
	}

	return nil
}
