package go_validate

import (
	"encoding/json"
	"testing"

	govalidate "github.com/xloss/go-validate"
	"github.com/xloss/go-validate/rules"
)

func TestRuleDomain(t *testing.T) {
	type testRequest struct {
		Host1 string `json:"host1"`
		Host2 string `json:"host2"`
		Host3 string `json:"host3"`
		Host4 string `json:"host4"`
		Host5 string `json:"host5"`
		Host6 string `json:"host6"`
		Host7 string `json:"host7"`
		Host8 string `json:"host8"`
	}

	var data map[string]any
	r1text := `{"host1": "localhost", "host2": "example.com", "host3": "пример.испытание", "host4": "مثال.إختبار", "host5": "例子.测试", "host6": "παράδειγμα.δοκιμή", "host7": "उदाहरण.परीक्षा", "host8": "例え.テスト"}`
	_ = json.Unmarshal([]byte(r1text), &data)

	r, errors := govalidate.Run[testRequest](data, map[string][]any{
		"host1": {&rules.Domain{}},
		"host2": {&rules.Domain{}},
		"host3": {&rules.Domain{}},
		"host4": {&rules.Domain{}},
		"host5": {"domain"},
		"host6": {"domain"},
		"host7": {"domain"},
		"host8": {"domain"},
	})

	if len(errors) != 0 {
		t.Errorf("Errors should be 0")
	}

	if r.Host1 != "localhost" {
		t.Errorf("Host should be localhost")
	}
	if r.Host2 != "example.com" {
		t.Errorf("Host should be example.com")
	}
	if r.Host3 != "пример.испытание" {
		t.Errorf("Host should be пример.испытание")
	}
	if r.Host4 != "مثال.إختبار" {
		t.Errorf("Host should be مثال.إختبار")
	}
	if r.Host5 != "例子.测试" {
		t.Errorf("Host should be 例子.测试")
	}
	if r.Host6 != "παράδειγμα.δοκιμή" {
		t.Errorf("Host should be παράδειγμα.δοκιμή")
	}
	if r.Host7 != "उदाहरण.परीक्षा" {
		t.Errorf("Host should be उदाहरण.परीक्षा")
	}
	if r.Host8 != "例え.テスト" {
		t.Errorf("Host should be 例え.テスト")
	}

	r2text := `{"host1": ".localhost", "host2": "example.com.", "host3": "examp|e.com", "host4": "-example.com"}`
	_ = json.Unmarshal([]byte(r2text), &data)

	r, errors = govalidate.Run[testRequest](data, map[string][]any{
		"host1": {&rules.Domain{}},
		"host2": {&rules.Domain{}},
		"host3": {&rules.Domain{}},
		"host4": {&rules.Domain{}},
	})

	if len(errors) != 4 {
		t.Errorf("Errors should be 4")
	}

	for _, err := range errors {
		switch err.Attribute {
		case "host1":
			fallthrough
		case "host2":
			fallthrough
		case "host3":
			fallthrough
		case "host4":
			if err.Name != "domain" {
				t.Errorf("Error should be domain")
			}
		default:
			t.Errorf("Error should be domain")
		}
	}
}
