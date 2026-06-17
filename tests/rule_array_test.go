package go_validate

import (
	"encoding/json"
	"testing"

	govalidate "github.com/xloss/go-validate"
	"github.com/xloss/go-validate/rules"
)

func TestRuleArray(t *testing.T) {
	type testRequest struct {
		Array1 []string       `json:"array1"`
		Array2 []int          `json:"array2"`
		Array3 []any          `json:"array3"`
		Map1   map[string]any `json:"map1"`
		Map2   map[string]int `json:"map2"`
	}

	var data map[string]any
	r1text := `{"array1": ["a", "b"], "array2": [1, 2], "array3": [1, "a", true], "map1": {"a": 1}, "map2": {"b": 2}}`
	_ = json.Unmarshal([]byte(r1text), &data)

	r, errors := govalidate.Run[testRequest](data, map[string][]any{
		"array1": {&rules.Array{}},
		"array2": {"array"},
		"array3": {"array"},
		"map1":   {"array"},
		"map2":   {"array"},
	})

	if len(errors) != 0 {
		t.Errorf("Errors should be 0. %+v", errors)
	}

	if len(r.Array1) != 2 || r.Array1[0] != "a" || r.Array1[1] != "b" {
		t.Errorf("Array1 does not match. %+v", r.Array1)
	}

	if len(r.Array2) != 2 || r.Array2[0] != 1 || r.Array2[1] != 2 {
		t.Errorf("Array2 does not match. %+v", r.Array2)
	}

	if len(r.Array3) != 3 {
		t.Errorf("Array3 should have 3 elements. %+v", r.Array3)
	}

	if len(r.Map1) != 1 || r.Map1["a"] != float64(1) {
		t.Errorf("Map1 does not match. %+v", r.Map1)
	}

	if len(r.Map2) != 1 || r.Map2["b"] != 2 {
		t.Errorf("Map2 does not match. %+v", r.Map2)
	}
}

func TestRuleArrayOptional(t *testing.T) {
	type testRequest struct {
		Array []string `json:"array"`
	}

	data := map[string]any{}

	r, errors := govalidate.Run[testRequest](data, map[string][]any{
		"array": {"array"},
	})

	if len(errors) != 0 {
		t.Errorf("Errors should be 0. %+v", errors)
	}

	if r == nil {
		t.Errorf("Result should not be nil")
		return
	}

	if r.Array != nil {
		t.Errorf("Array should be nil")
	}
}

func TestRuleArrayErrors(t *testing.T) {
	type testRequest struct {
		Array1 []string `json:"array1"`
		Array2 []int    `json:"array2"`
		Array3 []any    `json:"array3"`
	}

	var data map[string]any
	r1text := `{"array1": "string", "array2": 123, "array3": true}`
	_ = json.Unmarshal([]byte(r1text), &data)

	_, errors := govalidate.Run[testRequest](data, map[string][]any{
		"array1": {"array"},
		"array2": {"array"},
		"array3": {"array"},
	})

	if len(errors) != 3 {
		t.Errorf("Errors should be 3. %+v", errors)
	}

	for _, err := range errors {
		switch err.Attribute {
		case "array1":
			fallthrough
		case "array2":
			fallthrough
		case "array3":
			if err.Name != "array" {
				t.Errorf("Error should be array")
			}
		default:
			t.Errorf("Unexpected error attribute: %s", err.Attribute)
		}
	}
}

func TestRuleArrayRequired(t *testing.T) {
	type testRequest struct {
		Array []string `json:"array"`
	}

	data := map[string]any{}

	_, errors := govalidate.Run[testRequest](data, map[string][]any{
		"array": {"required", "array"},
	})

	if len(errors) != 1 {
		t.Errorf("Errors should be 1. %+v", errors)
	}

	if len(errors) == 1 && errors[0].Name != "required" {
		t.Errorf("Error should be required")
	}
}
