package go_validate

import (
	"encoding/json"
	"testing"

	govalidate "github.com/xloss/go-validate"
	"github.com/xloss/go-validate/rules"
)

func TestRuleNumeric(t *testing.T) {
	type testRequest struct {
		Float1      float64 `json:"float1"`
		Float2      float64 `json:"float2"`
		Invalidate1 float64 `json:"invalidate1"`
		Invalidate2 float64 `json:"invalidate2"`
	}

	var data map[string]any
	r1text := `{"float1": null, "float2": 1.5, "invalidate1": "1", "invalidate2": true}`
	_ = json.Unmarshal([]byte(r1text), &data)

	_, errors := govalidate.Run[testRequest](data, map[string][]any{
		"float1":      {&rules.Numeric{}},
		"float2":      {"numeric"},
		"invalidate1": {&rules.Numeric{}},
		"invalidate2": {"numeric"},
	})

	if len(errors) != 2 {
		t.Errorf("expected 1 validation error, got %v", errors)
	}

	for _, err := range errors {
		switch err.Attribute {
		case "invalidate1":
			if err.Name != "numeric" {
				t.Errorf("Errors invalidate1.Name should be numeric")
			}
		case "invalidate2":
			if err.Name != "numeric" {
				t.Errorf("Errors invalidate2.Name should be numeric")
			}
		default:
			t.Errorf("Errors should be 2")
		}
	}
}
