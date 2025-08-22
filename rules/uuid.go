package rules

import (
	"github.com/google/uuid"
)

type UUID struct {
	name string
}

func (r *UUID) GetName() string {
	return r.name
}

func (r *UUID) GetValues() map[string]any {
	return map[string]any{}
}

func (r *UUID) Validate(_ string, value any, _ map[string]any) bool {
	r.name = "uuid"

	if value == nil {
		return true
	}

	s, ok := value.(string)
	if !ok {
		return false
	}

	err := uuid.Validate(s)
	if err != nil {
		return false
	}

	return true
}
