package rules

type Integer struct {
	name string
}

func (r Integer) GetName() string {
	return r.name
}

func (r Integer) GetValues() map[string]any {
	return map[string]any{}
}

func (r Integer) Validate(value any) bool {
	r.name = "integer"

	if value == nil {
		return true
	}

	switch value.(type) {
	case int:
	case int8:
	case int16:
	case int32:
	case int64:
	case uint:
	case uint8:
	case uint16:
	case uint32:
	case uint64:
	case float64:
		v, okFloat := value.(float64)

		return okFloat && v == float64(int(v))
	default:
		return false
	}

	return true
}
