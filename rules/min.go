package rules

import (
	"strconv"
)

type Min struct {
	name   string
	values map[string]any

	Min int
}

func (r *Min) GetName() string {
	return r.name
}

func (r *Min) GetValues() map[string]any {
	return r.values
}

func (r *Min) AddParams(params string) error {
	value, err := strconv.Atoi(params)
	if err != nil {
		return err
	}

	r.Min = value

	return nil
}

func (r *Min) Validate(_ string, value any, _ map[string]any) bool {
	r.name = "min.numeric"
	r.values = map[string]any{
		"min": r.Min,
	}

	switch value.(type) {
	case nil:
		return true
	case string:
		r.name = "min.string"
		return len(value.(string)) >= r.Min
	case int8:
		return value.(int) >= r.Min
	case uint8:
		return value.(int) >= r.Min
	case int16:
		return value.(int) >= r.Min
	case uint16:
		return value.(int) >= r.Min
	case int32:
		return value.(int) >= r.Min
	case uint32:
		return value.(int) >= r.Min
	case int64:
		return value.(int) >= r.Min
	case uint64:
		return value.(int) >= r.Min
	case int:
		return value.(int) >= r.Min
	case uint:
		return value.(int) >= r.Min
	case float32:
		return value.(float32) >= float32(r.Min)
	case float64:
		return value.(float64) >= float64(r.Min)
	case []any:
		r.name = "min.array"
		return len(value.([]any)) >= r.Min
	case map[string]any:
		r.name = "min.array"
		return len(value.(map[string]any)) >= r.Min
	case map[int]any:
		r.name = "min.array"
		return len(value.(map[int]any)) >= r.Min
	case map[float64]any:
		r.name = "min.array"
		return len(value.(map[float64]any)) >= r.Min
	default:
		r.name = "min.error"
		return false
	}
}
