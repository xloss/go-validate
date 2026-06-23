package rules

import (
	"errors"
	"time"
)

type Date struct {
	name   string
	values map[string]any

	Format string
}

func (r *Date) GetName() string {
	return r.name
}

func (r *Date) GetValues() map[string]any {
	return r.values
}

func (r *Date) AddParams(params string) error {
	if params == "" {
		return errors.New("date format is required")
	}

	r.Format = params

	return nil
}

func (r *Date) Validate(_ string, value any, _ map[string]any) bool {
	r.name = "date"

	if r.Format == "" {
		r.Format = time.RFC3339Nano
	}

	r.values = map[string]any{
		"format": r.Format,
	}

	if value == nil {
		return true
	}

	d, ok := value.(string)
	if !ok {
		return false
	}

	_, err := time.Parse(r.Format, d)
	if err != nil {
		return false
	}

	return true
}

func (r *Date) Rewrite(value any) (func(typ string) any, bool) {
	return func(typ string) any {
		switch v := value.(type) {
		case func(typ string) any:
			value = v(typ)
		}

		v, ok := value.(string)
		if !ok {
			return value
		}

		if typ == "time.Time" {
			t, _ := time.Parse(r.Format, v)

			return t
		}

		return value
	}, true
}
