package go_validate

import (
	"reflect"
	"strconv"
	"time"
)

func Run[T interface{}](data map[string]any, fieldRules map[string][]any) (*T, []Error) {
	errors := make([]Error, 0)

	validateFields(fieldRules, data, &errors)

	if len(errors) > 0 {
		return nil, errors
	}

	var (
		request T
		typeOf  = reflect.TypeOf(request)
		valueOf = reflect.ValueOf(&request).Elem()
	)

	setValue(typeOf, data, valueOf, "", &errors)

	return &request, errors
}

func setValue(typ reflect.Type, data any, value reflect.Value, field string, errors *[]Error) bool {
	if data == nil {
		return true
	}

	switch typ.Kind() {
	case reflect.Int:
		fallthrough
	case reflect.Int8:
		fallthrough
	case reflect.Int16:
		fallthrough
	case reflect.Int32:
		fallthrough
	case reflect.Int64:
		fallthrough
	case reflect.Uint:
		fallthrough
	case reflect.Uint8:
		fallthrough
	case reflect.Uint16:
		fallthrough
	case reflect.Uint32:
		fallthrough
	case reflect.Uint64:
		fallthrough
	case reflect.Float32:
		fallthrough
	case reflect.Float64:
		fallthrough
	case reflect.Bool:
		fallthrough
	case reflect.String:
		if s, ok := convertValue(typ, data); ok {
			value.Set(s)
		} else {
			*errors = append(*errors, Error{
				Attribute: field,
				Name:      "format",
				Values: map[string]any{
					"name":    typ.String(),
					"reflect": typ.Kind().String(),
				},
			})

			return false
		}
	case reflect.Struct:
		if typ.String() == "time.Time" {
			t, err := time.Parse(time.RFC3339Nano, data.(string))
			if err != nil {
				*errors = append(*errors, Error{
					Attribute: field,
					Name:      "format",
					Values: map[string]any{
						"name":    typ.String(),
						"reflect": typ.Kind().String(),
					},
				})
			}
			value.Set(reflect.ValueOf(t))
		} else {
			if s, ok := data.(map[string]any); ok {
				structValue(typ, value, s, field, errors)
			} else {
				*errors = append(*errors, Error{
					Attribute: field,
					Name:      "format",
					Values: map[string]any{
						"name":    typ.String(),
						"reflect": typ.Kind().String(),
					},
				})

				return false
			}
		}
	case reflect.Slice:
		sliceValue(typ, data, value, field, errors)
	case reflect.Map:
		mapValue(typ, data, value, field, errors)
	case reflect.Ptr:
		n := reflect.New(typ.Elem())
		if ok := setValue(typ.Elem(), data, n.Elem(), field, errors); ok {
			value.Set(n)
		}
	default:
		*errors = append(*errors, Error{
			Attribute: field,
			Name:      "format",
			Values: map[string]any{
				"name":    typ.String(),
				"reflect": typ.Kind().String(),
			},
		})

		return false
	}

	return true
}

func convertValue(typ reflect.Type, data any) (reflect.Value, bool) {
	v := reflect.ValueOf(data)

	if !v.IsValid() {
		return reflect.Zero(typ), true
	}

	if !v.CanConvert(typ) {
		return v, false
	}

	return v.Convert(typ), true
}

func structValue(t reflect.Type, v reflect.Value, data map[string]any, f string, errors *[]Error) {
	for i := 0; i < v.NumField(); i++ {
		value := v.Field(i)
		field := t.Field(i)

		jsonName := field.Tag.Get("json")
		if jsonName == "" {
			jsonName = field.Name
		}

		d := data[jsonName]

		eName := jsonName
		if f != "" {
			eName = f + "." + jsonName
		}

		setValue(field.Type, d, value, eName, errors)
	}
}

func sliceValue(typ reflect.Type, data any, value reflect.Value, field string, errors *[]Error) {
	if sliceData, conv := data.([]interface{}); conv {
		s := reflect.MakeSlice(typ, len(sliceData), len(sliceData))

		for i, el := range sliceData {
			setValue(typ.Elem(), el, s.Index(i), field+"."+strconv.Itoa(i), errors)
		}

		value.Set(s)
	} else {
		*errors = append(*errors, Error{
			Attribute: field,
			Name:      "format",
			Values: map[string]any{
				"name":    typ.String(),
				"reflect": typ.Kind().String(),
			},
		})
	}
}

func mapValue(typ reflect.Type, data any, value reflect.Value, field string, errors *[]Error) {
	m := reflect.MakeMap(typ)

	for k, v := range data.(map[string]any) {
		nv := reflect.New(typ.Elem())
		setValue(typ.Elem(), v, nv.Elem(), field+"."+k, errors)
		m.SetMapIndex(reflect.ValueOf(k), nv.Elem())
	}

	value.Set(m)
}

func validateFields(fieldRules map[string][]any, data map[string]any, errors *[]Error) {
	for field, rules := range fieldRules {
		r := make([]Rule, 0)
		for _, v := range rules {
			switch v.(type) {
			case Rule:
				r = append(r, v.(Rule))
			case string:
				rule := nameToRule(v.(string))

				if rule != nil {
					r = append(r, rule)
				}
			}
		}

		for _, rule := range r {
			if ok := rule.Validate(data[field]); !ok {
				*errors = append(*errors, Error{
					Attribute: field,
					Name:      rule.GetName(),
					Values:    rule.GetValues(),
				})
			}
		}
	}
}
