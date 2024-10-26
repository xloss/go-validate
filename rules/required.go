package rules

type Required struct {
	name string
}

func (r Required) GetName() string {
	return r.name
}

func (r Required) GetValues() map[string]any {
	return map[string]any{}
}

func (r Required) Validate(value any) bool {
	r.name = "required"

	return value != nil
}
