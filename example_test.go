package go_validate

import (
	"encoding/json"
	"github.com/xloss/go-validate/rules"
	"testing"
)

func TestRun(t *testing.T) {
	type apiRequest struct {
		IntegerValue  int     `json:"integer_value"`
		StringValue   string  `json:"string_value"`
		NumericValue  float64 `json:"numeric_value"`
		BooleanValue  bool    `json:"boolean_value"`
		RequiredValue string  `json:"required_value"`
	}

	jsonText := `{"integer_value": 1, "string_value": "b", "numeric_value": 3.1, "boolean_value": true, "required_value": "r"}`

	var data map[string]any
	errUnmarshal := json.Unmarshal([]byte(jsonText), &data)
	if errUnmarshal != nil {
		t.Error(errUnmarshal)
	}

	fieldRules := map[string][]any{
		"integer_value":  {rules.Integer{}},
		"string_value":   {"string"},
		"numeric_value":  {rules.Numeric{}},
		"boolean_value":  {"boolean"},
		"required_value": {rules.Required{}, rules.String{}},
	}

	r, errors := Run[apiRequest](data, fieldRules)
	if len(errors) != 0 {
		t.Error(errors)
	}

	if r.IntegerValue != 1 {
		t.Errorf("IntegerValue should be 1")
	}
	if r.StringValue != "b" {
		t.Errorf("StringValue should be b")
	}
	if r.NumericValue != 3.1 {
		t.Errorf("NumericValue should be 3.1")
	}
	if r.BooleanValue != true {
		t.Errorf("BooleanValue should be true")
	}
	if r.RequiredValue != "r" {
		t.Errorf("RequiredValue should be r")
	}
}
