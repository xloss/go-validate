package go_validate

import "github.com/xloss/go-validate/rules"

type Rule interface {
	GetName() string
	GetValues() map[string]any
	Validate(value any) bool
}

func nameToRule(name string) Rule {
	switch name {
	case "required":
		return rules.Required{}
	case "string":
		return rules.String{}
	case "integer":
		return rules.Integer{}
	case "numeric":
		return rules.Numeric{}
	case "boolean":
		return rules.Boolean{}
	}

	return nil
}
