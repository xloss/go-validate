package rules

type Boolean struct {
	name string
}

func (r *Boolean) GetName() string {
	return r.name
}

func (r *Boolean) GetValues() map[string]any {
	return map[string]any{}
}

func (r *Boolean) Validate(value any) bool {
	r.name = "boolean"

	if value == nil {
		return true
	}

	_, ok := value.(bool)

	return ok
}
