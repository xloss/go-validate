package go_validate

import (
	"encoding/json"
	"io"
	"reflect"
	"strconv"
	"strings"
)

type Error struct {
	Attribute string            `json:"attribute,omitempty"`
	Name      string            `json:"name,omitempty"`
	Values    map[string]string `json:"values,omitempty"`
}

func Run[T interface{}](body io.ReadCloser) (*T, []Error) {
	var (
		data    map[string]any
		request T
		errList []Error
	)

	_ = json.NewDecoder(body).Decode(&data)

	typeOf := reflect.TypeOf(request)
	valueOf := reflect.ValueOf(&request).Elem()

	errList = checkList(typeOf, valueOf, data, []string{})

	return &request, errList
}

func checkList(typeOf reflect.Type, valueOf reflect.Value, data map[string]any, path []string) []Error {
	errList := make([]Error, 0)

	for i := 0; i < valueOf.NumField(); i++ {
		value := valueOf.Field(i)
		field := typeOf.Field(i)

		jsonName := field.Tag.Get("json")
		if jsonName == "" {
			jsonName = field.Name
		}

		// Если вложенная структура
		if field.Type.Kind() == reflect.Struct {
			innerPath := append([]string{}, path...)
			// структура в поле
			if field.IsExported() {
				innerPath = append(innerPath, jsonName)
			}

			errs := checkList(field.Type, value, data, innerPath)
			if len(errs) > 0 {
				errList = append(errList, errs...)
			}

			continue
		}

		rules, required := getRules(field)
		errs := validator(data, path, jsonName, rules, required)
		if len(errs) > 0 {
			errList = append(errList, errs...)
			continue
		}

		v := data[jsonName]

		switch field.Type.Kind() {
		case reflect.Int:
			value.SetInt(toInt(v))
		case reflect.Float64:
			value.SetFloat(toFloat(v))
		case reflect.String:
			value.SetString(toString(v))
		case reflect.Bool:
			value.SetBool(toBool(v))
		case reflect.Slice:
			switch field.Type.Elem().String() {
			case "string":
				value.Set(reflect.ValueOf(toSliceString(v)))
			case "int":
				value.Set(reflect.ValueOf(toSliceInt(v)))
			}
		default:
			continue
		}
	}

	return errList
}

func validator(data map[string]any, path []string, name string, rules []Rule, required bool) []Error {
	if len(rules) == 0 {
		return nil
	}

	errList := make([]Error, 0)
	value, exist := getValue(data, append(path, name))
	if required && (!exist || value == "" || value == nil) {
		errList = append(errList, Error{
			Attribute: name,
			Name:      "required",
		})

		return errList
	}

	if !exist {
		return errList
	}

	for _, rule := range rules {
		if rule.Name == "integer" && reflect.TypeOf(value).Kind() != reflect.Float64 {
			errList = append(errList, Error{
				Attribute: name,
				Name:      "integer",
			})
		} else if rule.Name == "float" && reflect.TypeOf(value).Kind() != reflect.Float64 {
			errList = append(errList, Error{
				Attribute: name,
				Name:      "numeric",
			})
		} else if rule.Name == "string" && reflect.TypeOf(value).Kind() != reflect.String {
			errList = append(errList, Error{
				Attribute: name,
				Name:      "string",
			})
		}
	}

	return errList
}

func getValue(data map[string]any, path []string) (any, bool) {
	if len(path) == 0 {
		return nil, false
	}

	key := path[0]

	value, exist := data[key]
	if !exist || value == nil {
		return nil, false
	}

	if len(path) == 1 {
		return value, true
	}

	return getValue(value.(map[string]any), path[1:])
}

func getRules(field reflect.StructField) ([]Rule, bool) {
	var (
		rules    []Rule
		required bool
	)

	if validate, ok := field.Tag.Lookup("validate"); ok {
		ruleList := strings.Split(validate, "|")
		for _, ruleData := range ruleList {
			if strings.ToLower(ruleData) == "required" {
				required = true

				continue
			}

			rule := Rule{}
			rule.Setup(ruleData)

			rules = append(rules, rule)
		}
	}

	return rules, required
}

func toInt(value any) int64 {
	switch value.(type) {
	case int:
		return int64(value.(int))
	case float64:
		return int64(value.(float64))
	case string:
		if v, err := strconv.Atoi(value.(string)); err == nil {
			return int64(v)
		}
	}

	return 0
}

func toFloat(value any) float64 {
	switch value.(type) {
	case int:
		return float64(value.(int))
	case float64:
		return value.(float64)
	case string:
		if v, err := strconv.ParseFloat(value.(string), 64); err == nil {
			return v
		}
	}

	return 0.0
}

func toString(value any) string {
	switch value.(type) {
	case int:
		return strconv.Itoa(value.(int))
	case float64:
		return strconv.FormatFloat(value.(float64), 'g', -1, 64)
	case string:
		return value.(string)
	}

	return ""
}

func toBool(value any) bool {
	switch value.(type) {
	case bool:
		return value.(bool)
	}

	return false
}

func toSliceString(value any) []string {
	r := make([]string, 0)

	if value == nil {
		return r
	}

	for _, v := range value.([]interface{}) {
		switch v.(type) {
		case string:
			r = append(r, v.(string))
		default:
			return r
		}
	}

	return r
}

func toSliceInt(value any) []int {
	r := make([]int, 0)

	if value == nil {
		return r
	}

	switch value.(type) {
	case []interface{}:
	default:
		return r
	}

	for _, v := range value.([]interface{}) {
		switch v.(type) {
		case float64:
			r = append(r, int(v.(float64)))
		default:
			return r
		}
	}

	return r
}
