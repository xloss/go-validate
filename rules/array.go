package rules

import "reflect"

type Array struct {
	name string
}

func (r *Array) GetName() string {
	return r.name
}

func (r *Array) GetValues() map[string]any {
	return map[string]any{}
}

func (r *Array) Validate(_ string, value any, _ map[string]any) bool {
	r.name = "array"

	if value == nil {
		return true
	}

	v := reflect.ValueOf(value)

	switch v.Kind() {
	case reflect.Array, reflect.Slice, reflect.Map:
		return true
	default:
		return false
	}
}
