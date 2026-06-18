package go_validate

import (
	"encoding/json"
	"testing"
	"time"

	govalidate "github.com/xloss/go-validate"
	"github.com/xloss/go-validate/rules"
)

func TestRuleDate(t *testing.T) {
	type testRequest struct {
		Date1 string    `json:"date1"`
		Date2 time.Time `json:"date2"`
	}

	var data map[string]any
	text := `{"date1": "2025-04-03T02:01:00-01:00", "date2": "2025-04-03T02:01:00-01:00"}`
	_ = json.Unmarshal([]byte(text), &data)

	r, errors := govalidate.Run[testRequest](data, map[string][]any{
		"date1": {&rules.Date{}},
		"date2": {"date"},
	})

	if len(errors) != 0 {
		t.Errorf("Errors should be 0")
	}

	if r.Date1 != "2025-04-03T02:01:00-01:00" {
		t.Errorf("Date1 should be 2025-04-03T02:01:00-01:00")
	}
	if r.Date2.GoString() == "2025-04-03T02:01:00-01:00" {
		t.Errorf("Date1 should be 2025-04-03T02:01:00-01:00")
	}

	text = `{"date1": "2025-04-03 02:01:00", "date2": "2025-04-03"}`
	_ = json.Unmarshal([]byte(text), &data)

	r, errors = govalidate.Run[testRequest](data, map[string][]any{
		"date1": {&rules.Date{}},
		"date2": {&rules.Date{}},
	})

	if len(errors) != 2 {
		t.Errorf("Errors should be 2")
	}

	for _, err := range errors {
		switch err.Attribute {
		case "date1":
			fallthrough
		case "date2":
			if err.Name != "date" {
				t.Errorf("Error should be date")
			}
		default:
			t.Errorf("Error should be date")
		}
	}

}

func TestRuleDateFormat(t *testing.T) {
	type testRequest struct {
		Date1 string `json:"date1"`
		Date2 string `json:"date2"`
		Date3 string `json:"date3"`
	}

	var data map[string]any
	text := `{"date1": "2025-04-03", "date2": "2025-04-03 02:01:00", "date3": "2025-04-03T02:01:00-01:00"}`
	_ = json.Unmarshal([]byte(text), &data)

	r, errors := govalidate.Run[testRequest](data, map[string][]any{
		"date1": {&rules.Date{Format: "2006-01-02"}},
		"date2": {&rules.Date{Format: "2006-01-02 15:04:05"}},
		"date3": {"date_format:2006-01-02T15:04:05Z07:00"},
	})

	if len(errors) != 0 {
		t.Errorf("Errors should be 0")
	}

	if r.Date1 != "2025-04-03" {
		t.Errorf("Date1 should be 2025-04-03")
	}
	if r.Date2 != "2025-04-03 02:01:00" {
		t.Errorf("Date2 should be 2025-04-03 02:01:00")
	}
	if r.Date3 != "2025-04-03T02:01:00-01:00" {
		t.Errorf("Date3 should be 2025-04-03T02:01:00-01:00")
	}

	text = `{"date1": "2025-04-03T02:01:00-01:00", "date2": "2025-04-03", "date3": "2025-04-03 02:01:00"}`
	_ = json.Unmarshal([]byte(text), &data)

	_, errors = govalidate.Run[testRequest](data, map[string][]any{
		"date1": {&rules.Date{Format: "2006-01-02"}},
		"date2": {&rules.Date{Format: "2006-01-02 15:04:05"}},
		"date3": {"date_format:2006-01-02T15:04:05Z07:00"},
	})

	if len(errors) != 3 {
		t.Errorf("Errors should be 3")
	}

	for _, err := range errors {
		switch err.Attribute {
		case "date1":
			if err.Name != "date" {
				t.Errorf("Errors date1.Name should be date")
			}
			if err.Values["format"].(string) != "2006-01-02" {
				t.Errorf("Errors date1.Values[format] should be 2006-01-02")
			}
		case "date2":
			if err.Name != "date" {
				t.Errorf("Errors date2.Name should be date")
			}
			if err.Values["format"].(string) != "2006-01-02 15:04:05" {
				t.Errorf("Errors date2.Values[format] should be 2006-01-02 15:04:05")
			}
		case "date3":
			if err.Name != "date" {
				t.Errorf("Errors date3.Name should be date")
			}
			if err.Values["format"].(string) != "2006-01-02T15:04:05Z07:00" {
				t.Errorf("Errors date3.Values[format] should be 2006-01-02T15:04:05Z07:00")
			}
		default:
			t.Errorf("Errors should be 3")
		}
	}
}

func TestRuleDateFormatStringSyntaxError(t *testing.T) {
	type testRequest struct {
		Date string `json:"date"`
	}

	data := map[string]any{
		"date": "2025-04-03",
	}

	_, errors := govalidate.Run[testRequest](data, map[string][]any{
		"date": {"date_format:"},
	})

	if len(errors) != 1 {
		t.Errorf("Errors should be 1")
	}

	for _, err := range errors {
		if err.Attribute != "date" {
			t.Errorf("Error attribute should be date")
		}
		if err.Name != "error" {
			t.Errorf("Error name should be error")
		}
		if err.Values["rule"].(string) != "date_format:" {
			t.Errorf("Error Values[rule] should be date:")
		}
		if err.Values["error"].(error).Error() != "date format is required" {
			t.Errorf("Error Values[error] should be date format is required")
		}
	}
}
