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

	switch v := value.(type) {
	case nil:
		return true
	case string:
		r.name = "min.string"
		return len(v) >= r.Min
	case int:
		return v >= r.Min
	case int8:
		return int(v) >= r.Min
	case int16:
		return int(v) >= r.Min
	case int32:
		return int64(v) >= int64(r.Min)
	case int64:
		return v >= int64(r.Min)
	case uint:
		if r.Min < 0 {
			return true
		}

		return v >= uint(r.Min)
	case uint8:
		if r.Min < 0 {
			return true
		}

		return uint(v) >= uint(r.Min)
	case uint16:
		if r.Min < 0 {
			return true
		}

		return uint(v) >= uint(r.Min)
	case uint32:
		if r.Min < 0 {
			return true
		}

		return uint64(v) >= uint64(r.Min)
	case uint64:
		if r.Min < 0 {
			return true
		}

		return v >= uint64(r.Min)
	case float32:
		return v >= float32(r.Min)
	case float64:
		return v >= float64(r.Min)
	case []any:
		r.name = "min.array"
		return len(v) >= r.Min
	case []string:
		r.name = "min.array"
		return len(v) >= r.Min
	case []int:
		r.name = "min.array"
		return len(v) >= r.Min
	case []int8:
		r.name = "min.array"
		return len(v) >= r.Min
	case []int16:
		r.name = "min.array"
		return len(v) >= r.Min
	case []int32:
		r.name = "min.array"
		return len(v) >= r.Min
	case []int64:
		r.name = "min.array"
		return len(v) >= r.Min
	case []uint:
		r.name = "min.array"
		return len(v) >= r.Min
	case []uint8:
		r.name = "min.array"
		return len(v) >= r.Min
	case []uint16:
		r.name = "min.array"
		return len(v) >= r.Min
	case []uint32:
		r.name = "min.array"
		return len(v) >= r.Min
	case []uint64:
		r.name = "min.array"
		return len(v) >= r.Min
	case []float32:
		r.name = "min.array"
		return len(v) >= r.Min
	case []float64:
		r.name = "min.array"
		return len(v) >= r.Min
	case []bool:
		r.name = "min.array"
		return len(v) >= r.Min
	case map[string]any:
		r.name = "min.array"
		return len(v) >= r.Min
	case map[int]any:
		r.name = "min.array"
		return len(v) >= r.Min
	case map[float64]any:
		r.name = "min.array"
		return len(v) >= r.Min
	case map[string]string:
		r.name = "min.array"
		return len(v) >= r.Min
	case map[string]int:
		r.name = "min.array"
		return len(v) >= r.Min
	case map[string]float64:
		r.name = "min.array"
		return len(v) >= r.Min
	case map[string]bool:
		r.name = "min.array"
		return len(v) >= r.Min
	default:
		r.name = "min.error"
		return false
	}
}
