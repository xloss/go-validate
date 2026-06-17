package rules

import "encoding/json"

type JSON struct {
	name string
}

func (r *JSON) GetName() string {
	return r.name
}

func (r *JSON) GetValues() map[string]any {
	return map[string]any{}
}

func (r *JSON) Validate(_ string, value any, _ map[string]any) bool {
	r.name = "json"

	if value == nil {
		return true
	}

	j, ok := value.(string)
	if !ok {
		return false
	}

	return json.Valid([]byte(j))
}
