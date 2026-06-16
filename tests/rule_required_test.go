package go_validate

import (
	"encoding/json"
	"testing"

	govalidate "github.com/xloss/go-validate"
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

	_, errors := govalidate.Run[testRequest](data, map[string][]any{
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

func TestRuleRequiredEmptyValues(t *testing.T) {
	type testRequest struct {
		StringVal string         `json:"string_val"`
		ArrayVal  []string       `json:"array_val"`
		MapVal    map[string]any `json:"map_val"`
	}

	data := map[string]any{
		"string_val": "",
		"array_val":  []any{},
		"map_val":    map[string]any{},
	}

	_, errors := govalidate.Run[testRequest](data, map[string][]any{
		"string_val": {"required"},
		"array_val":  {"required"},
		"map_val":    {"required"},
	})

	if len(errors) != 3 {
		t.Errorf("Errors should be 3. %+v", errors)
	}

	for _, err := range errors {
		if err.Name != "required" {
			t.Errorf("Error should be required")
		}
	}
}

func TestRuleRequiredZeroValues(t *testing.T) {
	type testRequest struct {
		IntVal  int  `json:"int_val"`
		BoolVal bool `json:"bool_val"`
	}

	data := map[string]any{
		"int_val":  float64(0),
		"bool_val": false,
	}

	r, errors := govalidate.Run[testRequest](data, map[string][]any{
		"int_val":  {"required"},
		"bool_val": {"required"},
	})

	if len(errors) != 0 {
		t.Errorf("Errors should be 0. %+v", errors)
	}

	if r.IntVal != 0 {
		t.Errorf("IntVal should be 0")
	}

	if r.BoolVal != false {
		t.Errorf("BoolVal should be false")
	}
}
