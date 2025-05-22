package rules

type Error struct {
	name   string
	values map[string]any
}

func (r *Error) GetName() string {
	return "error"
}

func (r *Error) GetValues() map[string]any {
	return r.values
}

func (r *Error) AddParams(params string) {
	if r.values == nil {
		r.values = map[string]any{}
	}

	r.values["rule"] = params
}

func (r *Error) AddError(err error) {
	if r.values == nil {
		r.values = map[string]any{}
	}

	r.values["error"] = err
}

func (r *Error) Validate(_ any) bool {
	return false
}
