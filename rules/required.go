package rules

import "reflect"

type Required struct {
	name string
}

func (r *Required) GetName() string {
	return r.name
}

func (r *Required) GetValues() map[string]any {
	return map[string]any{}
}

func (r *Required) Validate(field string, value any, data map[string]any) bool {
	r.name = "required"

	if _, ok := data[field]; !ok {
		return false
	}

	if value == nil {
		return false
	}

	switch v := value.(type) {
	case string:
		return v != ""
	}

	rv := reflect.ValueOf(value)
	switch rv.Kind() {
	case reflect.Slice, reflect.Array, reflect.Map:
		return rv.Len() > 0
	}

	return true
}
