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
		return true
	}

	v, ok := value.(string)
	if !ok {
		return false
	}

	if v == "" {
		return false
	}

	addr, err := mail.ParseAddress(v)
	if err != nil {
		return false
	}

	return addr.Address == v
}
