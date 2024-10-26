package go_validate

import (
	"encoding/json"
	"github.com/xloss/go-validate/rules"
	"testing"
)

func TestRuleRequired(t *testing.T) {
	type testRequest struct {
		StringVal string  `json:"string_val"`
		StringPtr *string `json:"string_ptr"`
	}

	var data map[string]any
	r1text := `{"string_val": null, "string_ptr": "str2"}`
	_ = json.Unmarshal([]byte(r1text), &data)

	_, errors := Run[testRequest](data, map[string][]any{
		"string_val":  {rules.Required{}},
		"string_val2": {"required"},
	})

	if len(errors) != 2 {
		t.Errorf("expected 1 validation error, got %v", errors)
	}
}
