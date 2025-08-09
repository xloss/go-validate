package go_validate

import (
	"fmt"
	"testing"

	govalidate "github.com/xloss/go-validate"
	"github.com/xloss/go-validate/rules"
)

func TestRuleError(t *testing.T) {

	type testRequest struct {
		Error string `json:"field"`
	}

	var data map[string]any

	errRule := &rules.Error{}
	errRule.AddParams("param:1")
	errRule.AddError(fmt.Errorf("error"))

	_, errs := govalidate.Run[testRequest](data, map[string][]any{
		"field": {errRule},
	})

	if len(errs) != 1 {
		t.Errorf("expected 1 validation error, got %v", errs)
	}

	if errs[0].Name != "error" {
		t.Errorf("expected error validation error, got %v", errs[0].Name)
	}

	if len(errs[0].Values) != 2 {
		t.Errorf("expected 2 validation error values, got %v", errs[0].Values)
	}

	if errs[0].Values["error"].(error).Error() != "error" {
		t.Errorf("expected error validation error, got %v", errs[0].Values["error"])
	}
	if errs[0].Values["rule"] != "param:1" {
		t.Errorf("expected rule validation error, got %v", errs[0].Values["rule"])
	}
}
