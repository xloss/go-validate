package go_validate

import (
	"encoding/json"
	"testing"

	govalidate "github.com/xloss/go-validate"
	"github.com/xloss/go-validate/rules"
)

func TestRuleConfirmed(t *testing.T) {
	type testRequest struct {
		String string `json:"string"`
		Int    int    `json:"int"`
	}

	var data map[string]any
	r1text := `{"string": "str1", "string_confirmation": "str2", "int": 1, "int_confirmation": 2}`
	_ = json.Unmarshal([]byte(r1text), &data)

	_, errors := govalidate.Run[testRequest](data, map[string][]any{
		"string": {&rules.Confirmed{}},
		"int":    {&rules.Confirmed{}},
	})

	if len(errors) != 2 {
		t.Errorf("expected 2 validation error, got %v", errors)
	}

	for _, err := range errors {
		switch err.Attribute {
		case "string":
			if err.Name != "confirmed" {
				t.Errorf("Errors invalidate1.Name should be confirmed")
			}
		case "int":
			if err.Name != "confirmed" {
				t.Errorf("Errors invalidate2.Name should be confirmed")
			}
		default:
			t.Errorf("Errors should be 2")
		}
	}

	r1text = `{"string": "strABC", "string_confirmation": "strABC", "int": 2222, "int_confirmation": 2222}`
	_ = json.Unmarshal([]byte(r1text), &data)

	r, errors2 := govalidate.Run[testRequest](data, map[string][]any{
		"string": {&rules.Confirmed{}},
		"int":    {&rules.Confirmed{}},
	})

	if len(errors2) != 0 {
		t.Errorf("expected 0 validation error, got %v", errors2)
	}

	if r.String != "strABC" {
		t.Errorf("r.String should be strABC, got %v", r.String)
	}
	if r.Int != 2222 {
		t.Errorf("r.Int should be 2222, got %v", r.Int)
	}
}
