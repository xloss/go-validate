package go_validate

import (
	"encoding/json"
	"testing"

	govalidate "github.com/xloss/go-validate"
	"github.com/xloss/go-validate/rules"
)

func TestRuleJSON(t *testing.T) {
	type testRequest struct {
		JSON1 string `json:"json1"`
		JSON2 string `json:"json2"`
		JSON3 string `json:"json3"`
		JSON4 string `json:"json4"`
		JSON5 string `json:"json5"`
		JSON6 string `json:"json6"`
	}

	var data map[string]any
	r1text := `{
		"json1": "{\"a\":1,\"b\":true}",
		"json2": "[1,2,3]",
		"json3": "\"hello\"",
		"json4": "123",
		"json5": "true",
		"json6": "null"
	}`
	_ = json.Unmarshal([]byte(r1text), &data)

	r, errors := govalidate.Run[testRequest](data, map[string][]any{
		"json1": {&rules.JSON{}},
		"json2": {"json"},
		"json3": {"json"},
		"json4": {"json"},
		"json5": {"json"},
		"json6": {"json"},
	})

	if len(errors) != 0 {
		t.Errorf("Errors should be 0. %+v", errors)
	}

	if r.JSON1 != `{"a":1,"b":true}` {
		t.Errorf("JSON1 should be object json")
	}

	if r.JSON2 != `[1,2,3]` {
		t.Errorf("JSON2 should be array json")
	}

	if r.JSON3 != `"hello"` {
		t.Errorf("JSON3 should be string json")
	}

	if r.JSON4 != `123` {
		t.Errorf("JSON4 should be number json")
	}

	if r.JSON5 != `true` {
		t.Errorf("JSON5 should be boolean json")
	}

	if r.JSON6 != `null` {
		t.Errorf("JSON6 should be null json")
	}
}

func TestRuleJSONErrors(t *testing.T) {
	type testRequest struct {
		JSON1 string `json:"json1"`
		JSON2 string `json:"json2"`
		JSON3 string `json:"json3"`
		JSON4 string `json:"json4"`
		JSON5 string `json:"json5"`
	}

	var data map[string]any
	r1text := `{
		"json1": "hello",
		"json2": "{\"a\":}",
		"json3": "{bad json}",
		"json4": "",
		"json5": 123
	}`
	_ = json.Unmarshal([]byte(r1text), &data)

	_, errors := govalidate.Run[testRequest](data, map[string][]any{
		"json1": {"json"},
		"json2": {"json"},
		"json3": {"json"},
		"json4": {"json"},
		"json5": {"json"},
	})

	if len(errors) != 5 {
		t.Errorf("Errors should be 5. %+v", errors)
	}

	for _, err := range errors {
		switch err.Attribute {
		case "json1":
			fallthrough
		case "json2":
			fallthrough
		case "json3":
			fallthrough
		case "json4":
			fallthrough
		case "json5":
			if err.Name != "json" {
				t.Errorf("Error should be json")
			}
		default:
			t.Errorf("Unexpected error attribute: %s", err.Attribute)
		}
	}
}

func TestRuleJSONOptional(t *testing.T) {
	type testRequest struct {
		JSON1 string  `json:"json1"`
		JSON2 *string `json:"json2"`
	}

	data := map[string]any{
		"json2": nil,
	}

	r, errors := govalidate.Run[testRequest](data, map[string][]any{
		"json1": {"json"},
		"json2": {"json"},
	})

	if len(errors) != 0 {
		t.Errorf("Errors should be 0. %+v", errors)
	}

	if r == nil {
		t.Errorf("Result should not be nil")
		return
	}

	if r.JSON1 != "" {
		t.Errorf("JSON1 should be empty")
	}

	if r.JSON2 != nil {
		t.Errorf("JSON2 should be nil")
	}
}

func TestRuleJSONRequired(t *testing.T) {
	type testRequest struct {
		JSON string `json:"json"`
	}

	data := map[string]any{}

	_, errors := govalidate.Run[testRequest](data, map[string][]any{
		"json": {"required", "json"},
	})

	if len(errors) != 1 {
		t.Errorf("Errors should be 1. %+v", errors)
	}

	if len(errors) == 1 && errors[0].Name != "required" {
		t.Errorf("Error should be required")
	}
}
