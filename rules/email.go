package rules

import (
	"net/mail"
)

type Email struct {
	name   string
	values map[string]any
}

func (r *Email) GetName() string {
	return r.name
}

func (r *Email) GetValues() map[string]any {
	return r.values
}

func (r *Email) Validate(value any) bool {
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
