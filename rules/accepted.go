package rules

type Accepted struct {
	name string
}

func (r *Accepted) GetName() string {
	return r.name
}

func (r *Accepted) GetValues() map[string]any {
	return map[string]any{}
}

func (r *Accepted) Validate(_ string, value any, _ map[string]any) bool {
	r.name = "accepted"

	if value == nil {
		return true
	}

	return value == true || value == float64(1) || value == "yes" || value == "on" || value == "true" || value == "1"
}
