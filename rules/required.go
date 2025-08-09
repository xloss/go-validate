package rules

type Required struct {
	name string
}

func (r *Required) GetName() string {
	return r.name
}

func (r *Required) GetValues() map[string]any {
	return map[string]any{}
}

func (r *Required) Validate(_ string, value any, _ map[string]any) bool {
	r.name = "required"

	return value != nil
}
