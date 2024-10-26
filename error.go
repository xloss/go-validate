package go_validate

type Error struct {
	Attribute string         `json:"attribute,omitempty"`
	Name      string         `json:"name,omitempty"`
	Values    map[string]any `json:"values,omitempty"`
}
