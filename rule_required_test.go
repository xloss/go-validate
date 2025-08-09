package go_validate

import (
	"encoding/json"
	"testing"

	"github.com/xloss/go-validate/rules"
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
		"string_val":  {&rules.Required{}},
		"string_val2": {"required"},
	})

	if len(errors) != 2 {
		t.Errorf("expected 2 validation error, got %v", errors)
	}

	for _, err := range errors {
		switch err.Attribute {
		case "string_val":
			if err.Name != "required" {
				t.Errorf("Errors string_val.Name should be required")
			}
		case "string_val2":
			if err.Name != "required" {
				t.Errorf("Errors string_val2.Name should be required")
			}
		default:
			t.Errorf("Errors should be 2, found %s", err.Attribute)
		}
	}
}
