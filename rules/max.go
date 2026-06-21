package rules

import (
	"strconv"
)

type Max struct {
	name   string
	values map[string]any

	Max int
}

func (r *Max) GetName() string {
	return r.name
}

func (r *Max) GetValues() map[string]any {
	return r.values
}

func (r *Max) AddParams(params string) error {
	value, err := strconv.Atoi(params)
	if err != nil {
		return err
	}

	r.Max = value

	return nil
}

func (r *Max) Validate(_ string, value any, _ map[string]any) bool {
	r.name = "max.numeric"
	r.values = map[string]any{
		"max": r.Max,
	}

	switch v := value.(type) {
	case nil:
		return true
	case string:
		r.name = "max.string"
		return len(v) <= r.Max
	case int:
		return v <= r.Max
	case int8:
		return int(v) <= r.Max
	case int16:
		return int(v) <= r.Max
	case int32:
		return int64(v) <= int64(r.Max)
	case int64:
		return v <= int64(r.Max)
	case uint:
		if r.Max < 0 {
			return false
		}

		return v <= uint(r.Max)
	case uint8:
		if r.Max < 0 {
			return false
		}

		return uint(v) <= uint(r.Max)
	case uint16:
		if r.Max < 0 {
			return false
		}

		return uint(v) <= uint(r.Max)
	case uint32:
		if r.Max < 0 {
			return false
		}

		return uint64(v) <= uint64(r.Max)
	case uint64:
		if r.Max < 0 {
			return false
		}

		return v <= uint64(r.Max)
	case float32:
		return v <= float32(r.Max)
	case float64:
		return v <= float64(r.Max)
	case []any:
		r.name = "max.array"
		return len(v) <= r.Max
	case []string:
		r.name = "max.array"
		return len(v) <= r.Max
	case []int:
		r.name = "max.array"
		return len(v) <= r.Max
	case []int8:
		r.name = "max.array"
		return len(v) <= r.Max
	case []int16:
		r.name = "max.array"
		return len(v) <= r.Max
	case []int32:
		r.name = "max.array"
		return len(v) <= r.Max
	case []int64:
		r.name = "max.array"
		return len(v) <= r.Max
	case []uint:
		r.name = "max.array"
		return len(v) <= r.Max
	case []uint8:
		r.name = "max.array"
		return len(v) <= r.Max
	case []uint16:
		r.name = "max.array"
		return len(v) <= r.Max
	case []uint32:
		r.name = "max.array"
		return len(v) <= r.Max
	case []uint64:
		r.name = "max.array"
		return len(v) <= r.Max
	case []float32:
		r.name = "max.array"
		return len(v) <= r.Max
	case []float64:
		r.name = "max.array"
		return len(v) <= r.Max
	case []bool:
		r.name = "max.array"
		return len(v) <= r.Max
	case map[string]any:
		r.name = "max.array"
		return len(v) <= r.Max
	case map[int]any:
		r.name = "max.array"
		return len(v) <= r.Max
	case map[float64]any:
		r.name = "max.array"
		return len(v) <= r.Max
	case map[string]string:
		r.name = "max.array"
		return len(v) <= r.Max
	case map[string]int:
		r.name = "max.array"
		return len(v) <= r.Max
	case map[string]float64:
		r.name = "max.array"
		return len(v) <= r.Max
	case map[string]bool:
		r.name = "max.array"
		return len(v) <= r.Max
	default:
		r.name = "max.error"
		return false
	}
}
