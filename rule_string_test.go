package go_validate

import (
	"encoding/json"
	"github.com/xloss/go-validate/rules"
	"testing"
)

func TestRuleString(t *testing.T) {
	type testRequest struct {
		String1     string `json:"string1"`
		String2     string `json:"string2"`
		Invalidate1 string `json:"invalidate1"`
		Invalidate2 string `json:"invalidate2"`
	}

	var data map[string]any
	r1text := `{"string1": null, "string2": "str2", "invalidate1": 1, "invalidate2": 2}`
	_ = json.Unmarshal([]byte(r1text), &data)

	_, errors := Run[testRequest](data, map[string][]any{
		"string1":     {rules.String{}},
		"string2":     {"string"},
		"invalidate1": {rules.String{}},
		"invalidate2": {"string"},
	})

	if len(errors) != 2 {
		t.Errorf("expected 1 validation error, got %v", errors)
	}
}
