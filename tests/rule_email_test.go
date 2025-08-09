package go_validate

import (
	"encoding/json"
	"testing"

	govalidate "github.com/xloss/go-validate"
	"github.com/xloss/go-validate/rules"
)

func TestEmailDomain(t *testing.T) {
	type testRequest struct {
		Mail1 string `json:"mail1"`
		Mail2 string `json:"mail2"`
		Mail3 string `json:"mail3"`
	}

	var data map[string]any
	r1text := `{"mail1": "mail@localhost", "mail2": "mail@example.com", "mail3": "name@example.com"}`
	_ = json.Unmarshal([]byte(r1text), &data)

	r, errors := govalidate.Run[testRequest](data, map[string][]any{
		"mail1": {&rules.Email{}},
		"mail2": {&rules.Email{}},
		"mail3": {"email"},
	})

	if len(errors) != 0 {
		t.Errorf("Errors should be 0")
	}

	if r.Mail1 != "mail@localhost" {
		t.Errorf("Email should be mail@localhost")
	}
	if r.Mail2 != "mail@example.com" {
		t.Errorf("Email should be mail@example.com")
	}
	if r.Mail3 != "name@example.com" {
		t.Errorf("Email should be name@example.com")
	}

	r2text := `{"mail1": "localhost", "mail2": "example.com", "mail3": "mail@example.com"}`
	_ = json.Unmarshal([]byte(r2text), &data)

	r, errors = govalidate.Run[testRequest](data, map[string][]any{
		"mail1": {&rules.Email{}},
		"mail2": {&rules.Email{}},
		"mail3": {&rules.Email{}},
	})

	if len(errors) != 2 {
		t.Errorf("Errors should be 2")
	}

	for _, err := range errors {
		switch err.Attribute {
		case "mail1":
			fallthrough
		case "mail2":
			if err.Name != "email" {
				t.Errorf("Error should be email")
			}
		default:
			t.Errorf("Error should be email")
		}
	}
}
