package rules

type Numeric struct {
	name string
}

func (r Numeric) GetName() string {
	return r.name
}

func (r Numeric) GetValues() map[string]any {
	return map[string]any{}
}

func (r Numeric) Validate(value any) bool {
	r.name = "numeric"

	if value == nil {
		return true
	}

	_, ok := value.(float64)

	return ok
}
