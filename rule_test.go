package go_validate

import "testing"

func TestRule_Setup(t *testing.T) {
	rule := Rule{}

	rule.Setup("required")
	if rule.Name != "required" || len(rule.Params) != 0 {
		t.Errorf("problem parsing without params, rule: %#v", rule)
	}

	rule.Setup("min:1")
	if rule.Name != "min" || len(rule.Params) != 1 || rule.Params[0] != "1" {
		t.Errorf("problem parsing with single param, rule: %#v", rule)
	}

	rule.Setup("between:1,5")
	if rule.Name != "between" || len(rule.Params) != 2 || rule.Params[0] != "1" || rule.Params[1] != "5" {
		t.Errorf("problem parsing with two params, rule: %#v", rule)
	}
}
