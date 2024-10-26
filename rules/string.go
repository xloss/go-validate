package rules

type String struct {
	name string
}

func (r String) GetName() string {
	return r.name
}

func (r String) GetValues() map[string]any {
	return map[string]any{}
}

func (r String) Validate(value any) bool {
	r.name = "string"

	if value == nil {
		return true
	}

	_, ok := value.(string)

	return ok
}
