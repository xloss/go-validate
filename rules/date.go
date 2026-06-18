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

	format := r.Format
	if format == "" {
		format = time.RFC3339Nano
	}

	r.values = map[string]any{
		"format": format,
	}

	if value == nil {
		return true
	}

	d, ok := value.(string)
	if !ok {
		return false
	}

	_, err := time.Parse(format, d)
	if err != nil {
		return false
	}

	return true
}
