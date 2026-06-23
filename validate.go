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

type rewriteValue func(typ string) any

func setValue(typ reflect.Type, data any, value reflect.Value, field string, errors *[]Error) bool {
	if data == nil {
		return true
	}

	switch d := data.(type) {
	case rewriteValue:
		data = d(typ.String())
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
			switch v := data.(type) {
			case string:
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

			case time.Time:
				value.Set(reflect.ValueOf(v))

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

		validateField(strings.Split(field, "."), data, "", r, errors)
	}
}

func validateField(parts []string, current any, path string, rules []Rule, errors *[]Error) {
	if len(parts) == 0 {
		return
	}

	part := parts[0]

	if part == "*" {
		switch value := current.(type) {
		case []any:
			for i, item := range value {
				nextPath := joinPath(path, strconv.Itoa(i))
				validateField(parts[1:], item, nextPath, rules, errors)
			}
		case map[string]any:
			for key, item := range value {
				nextPath := joinPath(path, key)
				validateField(parts[1:], item, nextPath, rules, errors)
			}
		}

		return
	}

	currentMap, okType := current.(map[string]any)
	if !okType {
		applyRules(part, nil, map[string]any{}, joinPath(path, part), rules, errors)
		return
	}

	if len(parts) == 1 {
		value := currentMap[part]
		applyRules(part, value, currentMap, joinPath(path, part), rules, errors)
		return
	}

	next, okPart := currentMap[part]
	if !okPart {
		validateField(parts[1:], nil, joinPath(path, part), rules, errors)
		return
	}

	validateField(parts[1:], next, joinPath(path, part), rules, errors)
}

func applyRules(field string, value any, data map[string]any, attribute string, rules []Rule, errors *[]Error) {
	var rewrite func(typ string) any = nil

	for _, rule := range rules {
		if ok := rule.Validate(field, value, data); !ok {
			*errors = append(*errors, Error{
				Attribute: attribute,
				Name:      rule.GetName(),
				Values:    rule.GetValues(),
			})

			continue
		}

		if value == nil {
			continue
		}

		switch r := rule.(type) {
		case Rewriter:
			if rewrite == nil {
				rw, enable := r.Rewrite(value)
				if enable {
					rewrite = rw
				}
			} else {
				rw, enable := r.Rewrite(rewrite)

				if enable {
					rewrite = rw
				}
			}
		}
	}

	if rewrite != nil {
		data[field] = rewriteValue(rewrite)
	}
}

func joinPath(path string, part string) string {
	if path == "" {
		return part
	}

	return path + "." + part
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
