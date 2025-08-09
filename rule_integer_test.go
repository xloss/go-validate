package go_validate

import (
	"encoding/json"
	"testing"

	"github.com/xloss/go-validate/rules"
)

func TestRuleInteger(t *testing.T) {
	type testRequest struct {
		Integer1    int `json:"integer1"`
		Integer2    int `json:"integer2"`
		Invalidate1 int `json:"invalidate1"`
		Invalidate2 int `json:"invalidate2"`
	}

	var data map[string]any
	r1text := `{"integer1": null, "integer2": 1, "invalidate1": "1", "invalidate2": 2.5}`
	_ = json.Unmarshal([]byte(r1text), &data)

	_, errors := Run[testRequest](data, map[string][]any{
		"integer1":    {&rules.Integer{}},
		"integer2":    {"integer"},
		"invalidate1": {&rules.Integer{}},
		"invalidate2": {"integer"},
	})

	if len(errors) != 2 {
		t.Errorf("expected 2 validation error, got %v", errors)
	}

	for _, err := range errors {
		switch err.Attribute {
		case "invalidate1":
			if err.Name != "integer" {
				t.Errorf("Errors invalidate1.Name should be integer")
			}
		case "invalidate2":
			if err.Name != "integer" {
				t.Errorf("Errors invalidate2.Name should be integer")
			}
		default:
			t.Errorf("Errors should be 2")
		}
	}

	_, errors = Run[testRequest](map[string]any{
		"integer1": int8(5),
		"integer2": 5.6,
	}, map[string][]any{
		"integer1": {&rules.Integer{}},
		"integer2": {&rules.Integer{}},
	})

	if len(errors) != 1 {
		t.Errorf("expected 1 validation error, got %v", errors)
	}

	for _, err := range errors {
		switch err.Attribute {
		case "integer2":
			if err.Name != "integer" {
				t.Errorf("Errors integer2.Name should be integer")
			}
		default:
			t.Errorf("Errors should be 1")
		}
	}
}
