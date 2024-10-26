package go_validate

import (
	"encoding/json"
	"github.com/xloss/go-validate/rules"
	"testing"
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

	_, errors := Run[testRequest](data, map[string][]any{
		"float1":      {rules.Numeric{}},
		"float2":      {"numeric"},
		"invalidate1": {rules.Numeric{}},
		"invalidate2": {"numeric"},
	})

	if len(errors) != 2 {
		t.Errorf("expected 1 validation error, got %v", errors)
	}
}
