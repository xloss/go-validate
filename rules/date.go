package rules

import "time"

type Date struct {
	name   string
	values map[string]any
}

func (r *Date) GetName() string {
	return r.name
}

func (r *Date) GetValues() map[string]any {
	return r.values
}

func (r *Date) Validate(value any) bool {
	r.name = "date"

	if value == nil {
		return false
	}

	switch value.(type) {
	case string:
	default:
		return false
	}

	_, err := time.Parse(time.RFC3339Nano, value.(string))
	if err != nil {
		return false
	}

	return true
}
