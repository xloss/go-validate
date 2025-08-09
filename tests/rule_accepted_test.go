package go_validate

import (
	"encoding/json"
	"testing"

	govalidate "github.com/xloss/go-validate"
	"github.com/xloss/go-validate/rules"
)

func TestRuleAccepted(t *testing.T) {
	type testRequest struct {
		Accepted1 bool   `json:"accepted1"`
		Accepted2 int    `json:"accepted2"`
		Accepted3 string `json:"accepted3"`
		Accepted4 string `json:"accepted4"`
		Accepted5 string `json:"accepted5"`
		Accepted6 string `json:"accepted6"`
	}

	var data map[string]any
	r1text := `{"accepted1": true, "accepted2": 1, "accepted3": "yes", "accepted4": "on", "accepted5": "true", "accepted6": "1"}`
	_ = json.Unmarshal([]byte(r1text), &data)

	_, errors := govalidate.Run[testRequest](data, map[string][]any{
		"accepted1": {&rules.Accepted{}},
		"accepted2": {&rules.Accepted{}},
		"accepted3": {&rules.Accepted{}},
		"accepted4": {"accepted"},
		"accepted5": {"accepted"},
		"accepted6": {"accepted"},
	})

	if len(errors) != 0 {
		t.Errorf("Errors should be 0")
	}

	r2text := `{"accepted1": false, "accepted2": ""}`
	_ = json.Unmarshal([]byte(r2text), &data)

	_, errors = govalidate.Run[testRequest](data, map[string][]any{
		"accepted1": {&rules.Accepted{}},
		"accepted2": {&rules.Accepted{}},
	})

	if len(errors) != 2 {
		t.Errorf("Errors should be 2")
	}

	for _, err := range errors {
		switch err.Attribute {
		case "accepted1":
			fallthrough
		case "accepted2":
			if err.Name != "accepted" {
				t.Errorf("Error should be accepted")
			}
		default:
			t.Errorf("Error should be accepted")
		}
	}
}
