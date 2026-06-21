package go_validate

import (
	"encoding/json"
	"testing"

	govalidate "github.com/xloss/go-validate"
	"github.com/xloss/go-validate/rules"
)

func TestRuleMax(t *testing.T) {
	type testRequest struct {
		MaxInt   int            `json:"max_int"`
		MaxStr   string         `json:"max_str"`
		MaxFloat float64        `json:"max_float"`
		MaxArray []int          `json:"max_array"`
		MaxUInt8 uint8          `json:"max_uint8"`
		MaxMap   map[string]int `json:"max_map"`
	}

	var data map[string]any
	r1text := `{"max_int": 5, "max_str": "str", "max_float": 5.0, "max_array": [1,2,3,4,5], "max_uint8": 2, "max_map": {"a": 1}}`
	_ = json.Unmarshal([]byte(r1text), &data)

	r, errors := govalidate.Run[testRequest](data, map[string][]any{
		"max_int":   {&rules.Max{Max: 6}},
		"max_str":   {&rules.Max{Max: 4}},
		"max_float": {&rules.Max{Max: 10}},
		"max_array": {&rules.Max{Max: 8}},
		"max_uint8": {&rules.Max{Max: 2}},
		"max_map":   {&rules.Max{Max: 1}},
	})

	if len(errors) != 0 {
		t.Errorf("Errors should be 0")
	}

	if r.MaxInt != 5 {
		t.Errorf("MaxInt should be 5")
	}
	if r.MaxStr != "str" {
		t.Errorf("MaxStr should be str")
	}
	if r.MaxFloat != 5.0 {
		t.Errorf("MaxFloat should be 5.0")
	}
	if len(r.MaxArray) != 5 {
		t.Errorf("MaxArray should be 5")
	}
	if r.MaxUInt8 != 2 {
		t.Errorf("MaxUInt8 should be 2")
	}

	_, errors = govalidate.Run[testRequest](data, map[string][]any{
		"max_int":   {&rules.Max{Max: 2}},
		"max_str":   {&rules.Max{Max: 2}},
		"max_float": {&rules.Max{Max: 2}},
		"max_array": {&rules.Max{Max: 2}},
		"max_uint8": {&rules.Max{Max: 1}},
	})

	if len(errors) != 5 {
		t.Errorf("Errors should be 5")
	}

	for _, err := range errors {
		switch err.Attribute {
		case "max_int":
			if err.Name != "max.numeric" {
				t.Errorf("Errors max_int.Name should be max.numeric")
			}
			if err.Values["max"].(int) != 2 {
				t.Errorf("Errors max_int.Values[max] should be 2")
			}
		case "max_str":
			if err.Name != "max.string" {
				t.Errorf("Errors max_str.Name should be max.string")
			}
			if err.Values["max"].(int) != 2 {
				t.Errorf("Errors max_str.Values[max] should be 2")
			}
		case "max_float":
			if err.Name != "max.numeric" {
				t.Errorf("Errors max_float.Name should be max.numeric")
			}
			if err.Values["max"].(int) != 2 {
				t.Errorf("Errors max_float.Values[max] should be 2")
			}
		case "max_array":
			if err.Name != "max.array" {
				t.Errorf("Errors max_array.Name should be max.array")
			}
			if err.Values["max"].(int) != 2 {
				t.Errorf("Errors max_array.Values[max] should be 2")
			}
		case "max_uint8":
			if err.Name != "max.numeric" {
				t.Errorf("Errors max_uint8.Name should be max.numeric")
			}
			if err.Values["max"].(int) != 1 {
				t.Errorf("Errors max_uint8.Values[max] should be 1")
			}
		default:
			t.Errorf("Errors should be 5")
		}
	}

	r, errors = govalidate.Run[testRequest](data, map[string][]any{
		"max_int": {"max:7"},
	})

	if len(errors) != 0 {
		t.Errorf("Errors should be 0")
	}

	if r.MaxInt != 5 {
		t.Errorf("MaxInt should be 5")
	}

	_, errors = govalidate.Run[testRequest](data, map[string][]any{
		"max_int":   {"max:2"},
		"max_str":   {"max"},
		"max_float": {"max:2,1"},
	})

	if len(errors) != 3 {
		t.Errorf("Errors should be 3")
	}

	for _, err := range errors {
		switch err.Attribute {
		case "max_int":
			if err.Name != "max.numeric" {
				t.Errorf("Errors max_int.Name should be max.numeric")
			}
			if err.Values["max"].(int) != 2 {
				t.Errorf("Errors max_int.Values[max] should be 2")
			}
		case "max_str":
			if err.Name != "error" {
				t.Errorf("Errors max_str.Name should be error")
			}
			if err.Values["rule"].(string) != "max" {
				t.Errorf("Errors max_str.Values[rule] should be max")
			}
		case "max_float":
			if err.Name != "error" {
				t.Errorf("Errors max_float.Name should be error")
			}
			if err.Values["rule"].(string) != "max:2,1" {
				t.Errorf("Errors max_float.Values[rule] should be max:2,1")
			}
			if err.Values["error"].(error).Error() != "strconv.Atoi: parsing \"2,1\": invalid syntax" {
				t.Errorf("Errors max_float.Values[error] should be error text")
			}
		default:
			t.Errorf("Errors should be 3")
		}
	}
}

func TestRuleMaxIntegerTypes(t *testing.T) {
	type testRequest struct {
		Int8Val   int8   `json:"int8_val"`
		UInt8Val  uint8  `json:"uint8_val"`
		Int16Val  int16  `json:"int16_val"`
		UInt16Val uint16 `json:"uint16_val"`
		Int32Val  int32  `json:"int32_val"`
		UInt32Val uint32 `json:"uint32_val"`
		Int64Val  int64  `json:"int64_val"`
		UInt64Val uint64 `json:"uint64_val"`
	}

	data := map[string]any{
		"int8_val":   int8(10),
		"uint8_val":  uint8(10),
		"int16_val":  int16(10),
		"uint16_val": uint16(10),
		"int32_val":  int32(10),
		"uint32_val": uint32(10),
		"int64_val":  int64(10),
		"uint64_val": uint64(10),
	}

	r, errors := govalidate.Run[testRequest](data, map[string][]any{
		"int8_val":   {"max:15"},
		"uint8_val":  {"max:15"},
		"int16_val":  {"max:15"},
		"uint16_val": {"max:15"},
		"int32_val":  {"max:15"},
		"uint32_val": {"max:15"},
		"int64_val":  {"max:15"},
		"uint64_val": {"max:15"},
	})

	if len(errors) != 0 {
		t.Errorf("Errors should be 0. %+v", errors)
	}

	if r.Int8Val != 10 {
		t.Errorf("Int8Val should be 10")
	}
	if r.UInt8Val != 10 {
		t.Errorf("UInt8Val should be 10")
	}
	if r.Int16Val != 10 {
		t.Errorf("Int16Val should be 10")
	}
	if r.UInt16Val != 10 {
		t.Errorf("UInt16Val should be 10")
	}
	if r.Int32Val != 10 {
		t.Errorf("Int32Val should be 10")
	}
	if r.UInt32Val != 10 {
		t.Errorf("UInt32Val should be 10")
	}
	if r.Int64Val != 10 {
		t.Errorf("Int64Val should be 10")
	}
	if r.UInt64Val != 10 {
		t.Errorf("UInt64Val should be 10")
	}
}

func TestRuleMaxIntegerTypesErrors(t *testing.T) {
	type testRequest struct {
		Int8Val int8  `json:"int8_val"`
		UIntVal uint  `json:"uint_val"`
		IntVal  int64 `json:"int_val"`
	}

	data := map[string]any{
		"int8_val": int8(2),
		"uint_val": uint(2),
		"int_val":  int64(2),
	}

	_, errors := govalidate.Run[testRequest](data, map[string][]any{
		"int8_val": {"max:1"},
		"uint_val": {"max:1"},
		"int_val":  {"max:1"},
	})

	if len(errors) != 3 {
		t.Errorf("Errors should be 3. %+v", errors)
	}

	for _, err := range errors {
		if err.Name != "max.numeric" {
			t.Errorf("Error should be max.numeric")
		}
	}
}
