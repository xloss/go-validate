package go_validate

import (
	"encoding/json"
	"math"
	"reflect"
	"strconv"
	"strings"
	"time"
)

func Run[T any](data map[string]any, fieldRules map[string][]any) (*T, []Error) {
	errors := make([]Error, 0)

	var (
		request T
		typeOf  = reflect.TypeOf(request)
		valueOf = reflect.ValueOf(&request).Elem()
	)

	if typeOf == nil || typeOf.Kind() == reflect.Ptr {
		errors = append(errors, Error{
			Name: "type",
			Values: map[string]any{
				"name": "Run type parameter must be a non-pointer type",
			},
		})

		return nil, errors
	}

	validateFields(fieldRules, data, &errors)

	if len(errors) > 0 {
		return nil, errors
	}

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
			v, ok := data.(string)
			if !ok {
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

			t, err := time.Parse(time.RFC3339Nano, v)
			if err != nil {
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
		if isRawMessage(typ) {
			return rawMessageValue(typ, data, value, field, errors)
		}

		sliceValue(typ, data, value, field, errors)
	case reflect.Map:
		mapValue(typ, data, value, field, errors)
	case reflect.Ptr:
		n := reflect.New(typ.Elem())
		if ok := setValue(typ.Elem(), data, n.Elem(), field, errors); ok {
			value.Set(n)
		}
	case reflect.Interface:
		dataValue := reflect.ValueOf(data)
		if !dataValue.IsValid() {
			return true
		}

		if dataValue.Type().AssignableTo(typ) || dataValue.Type().Implements(typ) {
			value.Set(dataValue)
			return true
		}

		*errors = append(*errors, Error{
			Attribute: field,
			Name:      "format",
			Values: map[string]any{
				"name":    typ.String(),
				"reflect": typ.Kind().String(),
			},
		})

		return false
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

	switch typ.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return convertIntValue(typ, data)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return convertUintValue(typ, data)
	case reflect.Float32:
		if f, ok := data.(float64); ok && (f > math.MaxFloat32 || f < -math.MaxFloat32) {
			return v, false
		}
	}

	if !v.CanConvert(typ) {
		return v, false
	}

	return v.Convert(typ), true
}

func convertIntValue(typ reflect.Type, data any) (reflect.Value, bool) {
	v := reflect.ValueOf(data)

	switch n := data.(type) {
	case float64:
		if math.Trunc(n) != n {
			return v, false
		}

		bits := typ.Bits()
		minValue := -math.Pow(2, float64(bits-1))
		maxExclusive := math.Pow(2, float64(bits-1))

		if n < minValue || n >= maxExclusive {
			return v, false
		}

		return reflect.ValueOf(int64(n)).Convert(typ), true
	}

	if !v.IsValid() || !v.CanConvert(typ) {
		return v, false
	}

	bits := typ.Bits()
	minValue := int64(-1) << (bits - 1)
	maxValue := int64(1)<<(bits-1) - 1

	switch v.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		n := v.Int()
		if n < minValue || n > maxValue {
			return v, false
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		n := v.Uint()
		if n > uint64(maxValue) {
			return v, false
		}
	default:
		return v, false
	}

	return v.Convert(typ), true
}

func convertUintValue(typ reflect.Type, data any) (reflect.Value, bool) {
	v := reflect.ValueOf(data)

	switch n := data.(type) {
	case float64:
		if math.Trunc(n) != n || n < 0 {
			return v, false
		}

		bits := typ.Bits()
		maxExclusive := math.Pow(2, float64(bits))

		if n >= maxExclusive {
			return v, false
		}

		return reflect.ValueOf(uint64(n)).Convert(typ), true
	}

	if !v.IsValid() || !v.CanConvert(typ) {
		return v, false
	}

	bits := typ.Bits()
	maxValue := uint64(math.MaxUint64)
	if bits < 64 {
		maxValue = uint64(1)<<bits - 1
	}

	switch v.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		n := v.Int()
		if n < 0 || uint64(n) > maxValue {
			return v, false
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		n := v.Uint()
		if n > maxValue {
			return v, false
		}
	default:
		return v, false
	}

	return v.Convert(typ), true
}

func structValue(t reflect.Type, v reflect.Value, data map[string]any, f string, errors *[]Error) {
	for i := 0; i < v.NumField(); i++ {
		value := v.Field(i)
		field := t.Field(i)

		if !value.CanSet() {
			continue
		}

		jsonName := field.Tag.Get("json")
		if jsonName != "" {
			jsonName = strings.Split(jsonName, ",")[0]
		}

		if jsonName == "-" {
			continue
		}

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
	mapData, ok := data.(map[string]any)
	if !ok {
		*errors = append(*errors, Error{
			Attribute: field,
			Name:      "format",
			Values: map[string]any{
				"name":    typ.String(),
				"reflect": typ.Kind().String(),
			},
		})

		return
	}

	if typ.Key().Kind() != reflect.String {
		*errors = append(*errors, Error{
			Attribute: field,
			Name:      "format",
			Values: map[string]any{
				"name":    typ.String(),
				"reflect": typ.Kind().String(),
			},
		})

		return
	}

	m := reflect.MakeMap(typ)

	for k, v := range mapData {
		nv := reflect.New(typ.Elem())
		setValue(typ.Elem(), v, nv.Elem(), field+"."+k, errors)
		m.SetMapIndex(reflect.ValueOf(k).Convert(typ.Key()), nv.Elem())
	}

	value.Set(m)
}

func validateFields(fieldRules map[string][]any, data map[string]any, errors *[]Error) {
	for field, rules := range fieldRules {
		r := make([]Rule, 0)
		for _, v := range rules {
			switch ruleValue := v.(type) {
			case Rule:
				r = append(r, ruleValue)
			case string:
				rule := nameToRule(ruleValue)
				if rule != nil {
					r = append(r, rule)
				}
			default:
				if rule := valueToRule(v); rule != nil {
					r = append(r, rule)
				}
			}
		}

		for _, rule := range r {
			if ok := rule.Validate(field, data[field], data); !ok {
				*errors = append(*errors, Error{
					Attribute: field,
					Name:      rule.GetName(),
					Values:    rule.GetValues(),
				})
			}
		}
	}
}

func valueToRule(value any) Rule {
	v := reflect.ValueOf(value)
	if !v.IsValid() {
		return nil
	}

	if v.Kind() == reflect.Ptr {
		return nil
	}

	ptr := reflect.New(v.Type())
	ptr.Elem().Set(v)

	rule, ok := ptr.Interface().(Rule)
	if !ok {
		return nil
	}

	return rule
}

func isRawMessage(typ reflect.Type) bool {
	return typ.PkgPath() == "encoding/json" && typ.Name() == "RawMessage"
}

func rawMessageValue(typ reflect.Type, data any, value reflect.Value, field string, errors *[]Error) bool {
	switch v := data.(type) {
	case json.RawMessage:
		value.Set(reflect.ValueOf(v))
		return true
	case []byte:
		value.Set(reflect.ValueOf(json.RawMessage(v)))
		return true
	default:
		b, err := json.Marshal(v)
		if err != nil {
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

		value.Set(reflect.ValueOf(json.RawMessage(b)))
		return true
	}
}
