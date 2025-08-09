package go_validate

import (
	"encoding/json"
	"testing"

	"github.com/xloss/go-validate/rules"
)

func TestRuleMin(t *testing.T) {
	type testRequest struct {
		MinInt   int            `json:"min_int"`
		MinStr   string         `json:"min_str"`
		MinFloat float64        `json:"min_float"`
		MinArray []int          `json:"min_array"`
		MinUInt8 int8           `json:"min_uint8"`
		MinMap   map[string]int `json:"min_map"`
	}

	var data map[string]any
	r1text := `{"min_int": 5, "min_str": "str", "min_float": 5.0, "min_array": [1,2,3,4,5], "min_uint8": 2, "min_map": {"a": 1}}`
	_ = json.Unmarshal([]byte(r1text), &data)

	r, errors := Run[testRequest](data, map[string][]any{
		"min_int":   {&rules.Min{Min: 2}},
		"min_str":   {&rules.Min{Min: 2}},
		"min_float": {&rules.Min{Min: 2}},
		"min_array": {&rules.Min{Min: 2}},
		"min_uint8": {&rules.Min{Min: 2}},
		"min_map":   {&rules.Min{Min: 1}},
	})

	if len(errors) != 0 {
		t.Errorf("Errors should be 0")
	}

	if r.MinInt != 5 {
		t.Errorf("MinInt should be 5")
	}
	if r.MinStr != "str" {
		t.Errorf("MinStr should be str")
	}
	if r.MinFloat != 5.0 {
		t.Errorf("MinFloat should be 5.0")
	}
	if len(r.MinArray) != 5 {
		t.Errorf("MinArray should be 5")
	}
	if r.MinUInt8 != 2 {
		t.Errorf("MinUInt8 should be 2")
	}

	_, errors = Run[testRequest](data, map[string][]any{
		"min_int":   {&rules.Min{Min: 6}},
		"min_str":   {&rules.Min{Min: 4}},
		"min_float": {&rules.Min{Min: 7}},
		"min_array": {&rules.Min{Min: 8}},
		"min_uint8": {&rules.Min{Min: 3}},
	})

	if len(errors) != 5 {
		t.Errorf("Errors should be 5")
	}

	for _, err := range errors {
		switch err.Attribute {
		case "min_int":
			if err.Name != "min.numeric" {
				t.Errorf("Errors min_int.Name should be min.numeric")
			}
			if err.Values["min"].(int) != 6 {
				t.Errorf("Errors min_int.Values[min] should be 6")
			}
		case "min_str":
			if err.Name != "min.string" {
				t.Errorf("Errors min_str.Name should be min.string")
			}
			if err.Values["min"].(int) != 4 {
				t.Errorf("Errors min_str.Values[min] should be 4")
			}
		case "min_float":
			if err.Name != "min.numeric" {
				t.Errorf("Errors min_float.Name should be min.numeric")
			}
			if err.Values["min"].(int) != 7 {
				t.Errorf("Errors min_float.Values[min] should be 7")
			}
		case "min_array":
			if err.Name != "min.array" {
				t.Errorf("Errors min_array.Name should be min.array")
			}
			if err.Values["min"].(int) != 8 {
				t.Errorf("Errors min_array.Values[min] should be 8")
			}
		case "min_uint8":
			if err.Name != "min.numeric" {
				t.Errorf("Errors min_uint8.Name should be min.numeric")
			}
			if err.Values["min"].(int) != 3 {
				t.Errorf("Errors min_uint8.Values[min] should be 3")
			}
		default:
			t.Errorf("Errors should be 5")
		}
	}

	r, errors = Run[testRequest](data, map[string][]any{
		"min_int": {"min:2"},
	})

	if len(errors) != 0 {
		t.Errorf("Errors should be 0")
	}

	if r.MinInt != 5 {
		t.Errorf("MinInt should be 5")
	}

	_, errors = Run[testRequest](data, map[string][]any{
		"min_int":   {"min:10"},
		"min_str":   {"min"},
		"min_float": {"min:2,1"},
	})

	if len(errors) != 3 {
		t.Errorf("Errors should be 3")
	}

	for _, err := range errors {
		switch err.Attribute {
		case "min_int":
			if err.Name != "min.numeric" {
				t.Errorf("Errors min_int.Name should be min.numeric")
			}
			if err.Values["min"].(int) != 10 {
				t.Errorf("Errors min_int.Values[min] should be 10")
			}
		case "min_str":
			if err.Name != "error" {
				t.Errorf("Errors min_str.Name should be error")
			}
			if err.Values["rule"].(string) != "min" {
				t.Errorf("Errors min_str.Values[rule] should be min")
			}
		case "min_float":
			if err.Name != "error" {
				t.Errorf("Errors min_float.Name should be error")
			}
			if err.Values["rule"].(string) != "min:2,1" {
				t.Errorf("Errors min_float.Values[rule] should be min:2,1")
			}
			if err.Values["error"].(error).Error() != "strconv.Atoi: parsing \"2,1\": invalid syntax" {
				t.Errorf("Errors min_float.Values[error] should be error text")
			}
		default:
			t.Errorf("Errors should be 3")
		}
	}
}
