# go-validate

A small Go validation library for validating `map[string]any` input and converting it into a typed struct.

The library is designed mainly for data decoded from JSON into `map[string]any`.

## Installation

```bash
go get github.com/xloss/go-validate
```
## Quick Start

```go
package main

import (
	"encoding/json"
	"fmt"

	govalidate "github.com/xloss/go-validate"
	"github.com/xloss/go-validate/rules"
)

type apiRequest struct {
	IntegerValue  int     `json:"integer_value"`
	StringValue   string  `json:"string_value"`
	NumericValue  float64 `json:"numeric_value"`
	BooleanValue  bool    `json:"boolean_value"`
	RequiredValue string  `json:"required_value"`
}

func main() {
	jsonText := `{
		"integer_value": 1,
		"string_value": "hello",
		"numeric_value": 3.14,
		"boolean_value": true,
		"required_value": "value"
	}`

	var data map[string]any
	if err := json.Unmarshal([]byte(jsonText), &data); err != nil {
		panic(err)
	}

	fieldRules := map[string][]any{
		"integer_value":  {"integer"},
		"string_value":   {"string"},
		"numeric_value":  {&rules.Numeric{}},
		"boolean_value":  {"boolean"},
		"required_value": {"required", "string"},
	}

	request, errors := govalidate.Run[apiRequest](data, fieldRules)
	if len(errors) != 0 {
		fmt.Printf("validation errors: %+v\n", errors)
		return
	}

	fmt.Printf("%+v\n", request)
}
```
## API
```go
func Run[T any](data map[string]any, fieldRules map[string][]any) (*T, []Error)
```
`Run` does two things:

1. Validates input values using the provided rules.
2. Converts the input map into a typed struct.

There are two kinds of errors:

1. Rule validation errors.
2. Conversion errors.

If rule validation fails, `Run` returns:
```go
nil, []Error
```
In this case, the target struct is not created.

If rule validation succeeds, `Run` converts the input into the target struct and returns:
```go
*T, []Error
```
The returned error slice may still contain conversion errors. Conversion errors use the name `format`.

This means the result can be non-nil even when `errors` is not empty:
```go
request, errors := govalidate.Run[apiRequest](data, fieldRules)
if len(errors) != 0 {
	// request may be nil or non-nil.
	// Always check errors before using the result.
}
```

## Important Notes

### `T` must be a non-pointer type

Use:
```go
request, errors := govalidate.Run[apiRequest](data, rules)
```
Do not use:
```go
request, errors := govalidate.Run[*apiRequest](data, rules)
```
Pointer type parameters are rejected.

### Input is expected to be `map[string]any`

The library is designed for input like this:
```go
var data map[string]any
err := json.Unmarshal(jsonBytes, &data)
```
This means JSON numbers are usually represented as `float64`.

### Validation and conversion are separate steps

Rules validate the original value from `map[string]any`.

After validation succeeds, the library converts values into the target struct fields.

If a nested value cannot be converted, the result field keeps its zero value and a `format` error is added.

## Optional Fields

Most validation rules are optional by default.

This means that if a field is missing or has a `nil` value, rules such as `email`, `domain`, `integer`, `string`, `numeric`, and others do not fail.

Use `required` when the field must be present.
```go
fieldRules := map[string][]any{
	"email": {"required", "email"},
}
```
Examples:

| Input | Rules | Result |
|---|---|---|
| missing field | `email` | valid |
| missing field | `required`, `email` | invalid |
| `null` | `email` | valid |
| `null` | `required`, `email` | invalid |
| `""` | `email` | invalid |
| `"user@example.com"` | `email` | valid |

## Error Format
```go
type Error struct {
	Attribute string         `json:"attribute,omitempty"`
	Name      string         `json:"name,omitempty"`
	Values    map[string]any `json:"values,omitempty"`
}
```
Example:
```json
{
  "attribute": "email",
  "name": "email"
}
```
For nested values, the attribute name uses dot notation:
```text
items.0.name
```
## Built-in Rules

Rules can be passed as strings:
```go
fieldRules := map[string][]any{
	"name": {"required", "string"},
}
```
Or as rule instances:
```go
fieldRules := map[string][]any{
	"name": {&rules.Required{}, &rules.String{}},
}
```
Both forms can be mixed.

### `required`

The field must be present and not empty.

Invalid values:

- missing field;
- `nil`;
- empty string;
- empty array/slice;
- empty map.

Valid zero values:

- `0`;
- `false`.

String form:
```go
"required"
```
### `string`

The value must be a string.

Missing or `nil` values are valid unless `required` is also used.

String form:
```go
"string"
```
### `integer`

The value must be an integer.

For JSON input, numbers are usually `float64`, so `1` is valid and `1.5` is invalid.

String form:
```go
"integer"
```
### `numeric`

The value must be a number.

For JSON input, numeric values are represented as `float64`.

String form:
```go
"numeric"
```
### `boolean`

The value must be a boolean.

String form:
```go
"boolean"
```
### `min`

Checks the minimum value.

String form:
```go
"min:5"
```
Supported value types:

| Type | Check | Error name |
|---|---|---|
| number | value must be greater than or equal to min | `min.numeric` |
| string | string length must be greater than or equal to min | `min.string` |
| array/slice/map | length must be greater than or equal to min | `min.array` |
| unsupported type | invalid | `min.error` |

Example:
```go
fieldRules := map[string][]any{
	"name": {"required", "string", "min:3"},
	"age":  {"integer", "min:18"},
}
```
### `domain`

The value must be a valid domain/host name.

Examples of valid values:
```text
localhost
example.com
a.com
пример.испытание
例子.测试
```
Examples of invalid values:
```text
.example.com
example.com.
-example.com
Example.COM
```
Notes:

- The rule supports IDN domains.
- Single-label domains such as `localhost` are valid.
- A trailing dot is not accepted.
- The rule does not normalize input.
- Domain names should be passed in lowercase form.

String form:
```go
"domain"
```
### `date`

The value must be a valid RFC3339/RFC3339Nano date string.

Example:
```text
2026-06-16T12:30:00Z
```
String form:
```go
"date"
```
### `email`

The value must be an email address.

Examples of valid values:
```text
user@example.com
mail@localhost
```
Display-name addresses are not accepted:
```text
John Doe <john@example.com>
```
String form:
```go
"email"
```
### `confirmed`

The field must have a matching confirmation field.

For a field named:
```text
password
```
The confirmation field must be:
```text
password_confirmation
```
String form:
```go
"confirmed"
```
### `accepted`

The value must be one of:
```text
yes
on
1
true
```
Also accepted:
```go
true
float64(1)
```
Missing or `nil` values are valid unless `required` is also used.

String form:
```go
"accepted"
```
### `uuid`

The value must be a valid UUID.

String form:
```go
"uuid"
```
## Type Conversion

After validation, `Run` converts the input map into the target struct.

Supported target field types include:

- integers;
- unsigned integers;
- floats;
- booleans;
- strings;
- pointers;
- structs;
- slices;
- maps with string-like keys;
- `time.Time`.

### JSON tags

The library uses `json` tags to map input fields to struct fields.

Supported examples:
```go
Name string `json:"name"`
Name string `json:"name,omitempty"`
Name string `json:",omitempty"`
Secret string `json:"-"`
```
### Numeric conversion

Numeric conversion is strict for integer types.

Valid:
```json
{"age": 18}
```
Invalid:
```json
{"age": 18.5}
```
Integer overflow is rejected.

For example, this is invalid for `int8`:
```json
{"value": 128}
```
Negative values are rejected for unsigned integer fields.

### Maps

Maps are supported for string-like keys:
```go
map[string]int
```
Custom string key types are also supported:
```go
type Key string

type Request struct {
	Values map[Key]int `json:"values"`
}
```
Maps with non-string keys are not supported.

## Custom Rules

You can create custom validation rules by implementing the `Rule` interface.
```go
type Rule interface {
	GetName() string
	GetValues() map[string]any
	Validate(field string, value any, data map[string]any) bool
}
```
Example:
```go
type StartsWithA struct {
	name string
}

func (r *StartsWithA) GetName() string {
	return r.name
}

func (r *StartsWithA) GetValues() map[string]any {
	return map[string]any{}
}

func (r *StartsWithA) Validate(_ string, value any, _ map[string]any) bool {
	r.name = "starts_with_a"

	if value == nil {
		return true
	}

	v, ok := value.(string)
	if !ok {
		return false
	}

	return strings.HasPrefix(v, "A")
}
```
Usage:
```go
fieldRules := map[string][]any{
	"name": {&StartsWithA{}},
}
```
### Rule instances and concurrency

Rule instances may keep internal state, for example the error name or error values.

Do not reuse the same rule instance across concurrent validations unless your rule is safe for concurrent use.

Prefer creating new rule instances for each validation call.

## Unknown Rules

Unknown string rules are treated as validation errors.

Example:
```go
fieldRules := map[string][]any{
	"email": {"emial"},
}
```
This will return an error instead of silently ignoring the rule.

## License

See [LICENSE.md](LICENSE.md).