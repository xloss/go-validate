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
