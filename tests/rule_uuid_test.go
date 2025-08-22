package go_validate

import (
	"encoding/json"
	"testing"

	govalidate "github.com/xloss/go-validate"
	"github.com/xloss/go-validate/rules"
)

func TestRuleUUID(t *testing.T) {
	type testRequest struct {
		UUID        string `json:"uuid"`
		Invalidate1 string `json:"invalidate1"`
	}

	var data map[string]any
	r1text := `{"uuid": "3bc6eb8e-0cb1-4f8e-9589-633df2d0e14e", "invalidate1": "1234567890"}`
	_ = json.Unmarshal([]byte(r1text), &data)

	_, errors := govalidate.Run[testRequest](data, map[string][]any{
		"string1":     {&rules.UUID{}},
		"invalidate1": {"uuid"},
	})

	if len(errors) != 1 {
		t.Errorf("expected 1 validation error, got %v", errors)
	}

	for _, err := range errors {
		switch err.Attribute {
		case "invalidate1":
			if err.Name != "uuid" {
				t.Errorf("Errors invalidate1.Name should be uuid")
			}
		default:
			t.Errorf("Errors should be 1")
		}
	}
}
