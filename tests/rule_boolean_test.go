package go_validate

import (
	"encoding/json"
	"testing"

	govalidate "github.com/xloss/go-validate"
	"github.com/xloss/go-validate/rules"
)

func TestRuleBoolean(t *testing.T) {
	type testRequest struct {
		Boolean1    bool `json:"boolean1"`
		Boolean2    bool `json:"boolean2"`
		Invalidate1 bool `json:"invalidate1"`
		Invalidate2 bool `json:"invalidate2"`
	}

	var data map[string]any
	r1text := `{"boolean1": null, "boolean2": true, "invalidate1": "1", "invalidate2": 1.5}`
	_ = json.Unmarshal([]byte(r1text), &data)

	_, errors := govalidate.Run[testRequest](data, map[string][]any{
		"boolean1":    {&rules.Boolean{}},
		"boolean2":    {"boolean"},
		"invalidate1": {&rules.Boolean{}},
		"invalidate2": {"boolean"},
	})

	if len(errors) != 2 {
		t.Errorf("expected 1 validation error, got %v", errors)
	}

	for _, err := range errors {
		switch err.Attribute {
		case "invalidate1":
			if err.Name != "boolean" {
				t.Errorf("Errors invalidate1.Name should be boolean")
			}
		case "invalidate2":
			if err.Name != "boolean" {
				t.Errorf("Errors invalidate2.Name should be boolean")
			}
		default:
			t.Errorf("Errors should be 2")
		}
	}
}
