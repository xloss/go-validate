package rules

type Confirmed struct {
	name string
}

func (r *Confirmed) GetName() string {
	return r.name
}

func (r *Confirmed) GetValues() map[string]any {
	return map[string]any{}
}

func (r *Confirmed) Validate(field string, value any, data map[string]any) bool {
	r.name = "confirmed"

	if value == nil {
		return true
	}

	if data[field+"_confirmation"] == nil {
		return false
	}

	if data[field+"_confirmation"] != value {
		return false
	}

	return true
}
