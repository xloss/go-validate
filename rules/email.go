package rules

import (
	"net/mail"
)

type Email struct {
	name string
}

func (r *Email) GetName() string {
	return r.name
}

func (r *Email) GetValues() map[string]any {
	return map[string]any{}
}

func (r *Email) Validate(_ string, value any, _ map[string]any) bool {
	r.name = "email"

	if value == nil {
		return false
	}

	switch value.(type) {
	case string:
	default:
		return false
	}

	v := value.(string)

	_, err := mail.ParseAddress(v)
	if err != nil {
		return false
	}

	return true
}
