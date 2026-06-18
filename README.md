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
		"integer_value":  {rules.Integer{}},
		"string_value":   {rules.String{}},
		"numeric_value":  {rules.Numeric{}},
		"boolean_value":  {rules.Boolean{}},
		"required_value": {rules.Required{}, rules.String{}},
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

| Case | Result | Errors |
|---|---|---|
| Rule validation fails | `nil` | validation errors |
| Rule validation succeeds, conversion succeeds | `*T` | empty slice |
| Rule validation succeeds, conversion has errors | `*T` | `format` errors |

Conversion errors use the name `format`.

Always check `errors` before using the returned result:

```go
request, errors := govalidate.Run[apiRequest](data, fieldRules)
if len(errors) != 0 {
    return
}

_ = request
```

## Rule Syntax

Rules can be passed as rule values:

```go
fieldRules := map[string][]any{
    "name": {rules.Required{}, rules.String{}},
    "age":  {rules.Integer{}, rules.Min{Min: 18}},
}
```

This is the preferred Go-style syntax.

Rules can also be passed as strings:

```go
fieldRules := map[string][]any{
    "name": {"required", "string"},
    "age":  {"integer", "min:18"},
}
```

String rules are provided as a Laravel-like shorthand and can be useful when migrating from Laravel-style validation.

Both forms can be mixed:

```go
fieldRules := map[string][]any{
    "name": {"required", rules.String{}},
}
```

### Preferred Syntax

Prefer this:

```go
fieldRules := map[string][]any{
    "email": {rules.Required{}, rules.Email{}},
}
```

Instead of this:

```go
fieldRules := map[string][]any{
    "email": {"required", "email"},
}
```

The value syntax is more idiomatic in Go, works better with IDE autocomplete, avoids string typos, and matches the way custom rules are defined.

### Rule Values and Rule Pointers

Both value and pointer rule instances are supported:

```go
fieldRules := map[string][]any{
    "name": {rules.Required{}, rules.String{}},
}
```

```go
fieldRules := map[string][]any{
    "name": {&rules.Required{}, &rules.String{}},
}
```

Prefer rule values unless you have a specific reason to use pointers.

Rule values are copied internally into fresh rule instances. This makes them safer to use and avoids accidental shared state.

Avoid reusing the same rule pointer across concurrent validations unless your custom rule is safe for concurrent use.

## Nested Fields and Wildcards

Rules can be applied to nested fields using dot notation.

Example input:

```json
{
  "user": {
    "name": "John",
    "age": 30
  }
}
```

Example struct:

```go
type Request struct {
	User struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	} `json:"user"`
}
```

Example rules:

```go
fieldRules := map[string][]any{
	"user":      {rules.Required{}},
	"user.name": {rules.Required{}, rules.String{}},
	"user.age":  {rules.Required{}, rules.Integer{}},
}
```

If validation fails, the error attribute contains the full path:

```text
user.name
```

### Wildcards

Use `*` to validate every item in an array.

Example input:

```json
{
  "items": [
    {"name": "A", "qty": 1},
    {"name": "B", "qty": 2}
  ]
}
```

Example struct:

```go
type Item struct {
	Name string `json:"name"`
	Qty  int    `json:"qty"`
}

type Request struct {
	Items []Item `json:"items"`
}
```

Example rules:

```go
fieldRules := map[string][]any{
	"items":        {rules.Required{}, rules.Array{}},
	"items.*.name": {rules.Required{}, rules.String{}},
	"items.*.qty":  {rules.Required{}, rules.Integer{}},
}
```

If validation fails for an item, the error attribute contains the item index:

```text
items.0.name
items.1.qty
```

Nested wildcard rules validate existing items. If the parent field is required, validate it separately:

```go
fieldRules := map[string][]any{
	"items":        {rules.Required{}, rules.Array{}},
	"items.*.name": {rules.Required{}, rules.String{}},
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

This means that if a field is missing or has a `nil` value, rules such as `email`, `domain`, `integer`, `string`, `numeric`, `array`, `json`, and others do not fail.

Use `required` when the field must be present.

```go
fieldRules := map[string][]any{
	"email": {rules.Required{}, rules.Email{}},
}
```

Examples:

| Input | Rules | Result |
|---|---|---|
| missing field | `rules.Email{}` | valid |
| missing field | `rules.Required{}, rules.Email{}` | invalid |
| `null` | `rules.Email{}` | valid |
| `null` | `rules.Required{}, rules.Email{}` | invalid |
| `""` | `rules.Email{}` | invalid |
| `"user@example.com"` | `rules.Email{}` | valid |

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

For nested values and wildcard rules, the attribute name uses dot notation:

```
text
user.name
items.0.name
```

## Built-in Rules

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

Preferred syntax:

```go
rules.Required{}
```

String syntax:

```go
"required"
```

### `string`

The value must be a string.

Missing or `nil` values are valid unless `required` is also used.

Preferred syntax:

```go
rules.String{}
```

String syntax:

```go
"string"
```

### `integer`

The value must be an integer.

For JSON input, numbers are usually `float64`, so `1` is valid and `1.5` is invalid.

Preferred syntax:

```go
rules.Integer{}
```

String syntax:

```go
"integer"
```

### `numeric`

The value must be a number.

For JSON input, numeric values are represented as `float64`.

Preferred syntax:

```go
rules.Numeric{}
```

String syntax:

```go
"numeric"
```

### `boolean`

The value must be a boolean.

Preferred syntax:

```go
rules.Boolean{}
```

String syntax:

```go
"boolean"
```

### `array`

The value must be an array, slice, or map.

This rule follows Laravel-like naming. In Go terms it accepts:

- arrays;
- slices;
- maps.

Missing or `nil` values are valid unless `required` is also used.

Preferred syntax:

```go
rules.Array{}
```

String syntax:

```go
"array"
```

Examples:

```go
fieldRules := map[string][]any{
	"items": {rules.Required{}, rules.Array{}},
}
```

### `json`

The value must be a string containing valid JSON.

Valid examples:

```text
{"a":1}
[1,2,3]
"hello"
123
true
null
```

Invalid examples:

```text
hello
{bad json}
{"a":}
```

Important: the rule validates JSON strings. It does not validate already decoded Go values such as `map[string]any` or `[]any`.

Missing or `nil` values are valid unless `required` is also used.

Preferred syntax:

```go
rules.JSON{}
```

String syntax:

```go
"json"
```

Example:

```go
fieldRules := map[string][]any{
	"metadata": {rules.JSON{}},
}
```

### `min`

Checks the minimum value.

Supported value types:

| Type | Check | Error name |
|---|---|---|
| number | value must be greater than or equal to min | `min.numeric` |
| string | string length must be greater than or equal to min | `min.string` |
| array/slice/map | length must be greater than or equal to min | `min.array` |
| unsupported type | invalid | `min.error` |

Preferred syntax:

```go
rules.Min{Min: 5}
```

String syntax:

```go
"min:5"
```

Example:

```go
fieldRules := map[string][]any{
	"name": {rules.Required{}, rules.String{}, rules.Min{Min: 3}},
	"age":  {rules.Integer{}, rules.Min{Min: 18}},
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

Preferred syntax:

```go
rules.Domain{}
```

String syntax:

```go
"domain"
```

### `date`

The value must be a valid date string.

By default, the rule validates dates using `time.RFC3339Nano`.

Valid default examples:

```text
2026-06-16T12:30:00Z
2026-06-16T12:30:00.123456789Z
2026-06-16T12:30:00+03:00
```

Preferred syntax:

```go
rules.Date{}
```

String syntax:

```go
"date"
```

#### Custom date format

A custom date format can be provided through the `Format` field.

The format must be a Go time layout.

Preferred syntax:

```go
rules.Date{Format: "2006-01-02"}
```

String syntax:

```go
"date_format:2006-01-02"
```

Examples:

```go
fieldRules := map[string][]any{
	"birthday": {rules.Date{Format: "2006-01-02"}},
	"starts_at": {"date_format:2006-01-02 15:04:05"},
}
```

Valid values for these examples:

```text
2026-06-16
2026-06-16 12:30:00
```

Common Go date layouts:

| Expected value | Go layout |
|---|---|
| `2026-06-16` | `2006-01-02` |
| `16.06.2026` | `02.01.2006` |
| `2026-06-16 12:30:00` | `2006-01-02 15:04:05` |
| `2026-06-16T12:30:00Z` | `2006-01-02T15:04:05Z07:00` |

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

Preferred syntax:

```go
rules.Email{}
```

String syntax:

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

Preferred syntax:

```go
rules.Confirmed{}
```

String syntax:

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

Preferred syntax:

```go
rules.Accepted{}
```

String syntax:

```go
"accepted"
```

### `uuid`

The value must be a valid UUID.

Preferred syntax:

```go
rules.UUID{}
```

String syntax:

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
- `any` / `interface{}`;
- pointers;
- structs;
- slices;
- maps with string-like keys;
- `time.Time`;
- `json.RawMessage`.

### `any` / `interface{}`

Fields of type `any` keep the original decoded value.

Example:

```go
type Request struct {
	Data map[string]any `json:"data"`
}
```

For JSON input, numbers inside `any` values remain `float64`, because this is how `encoding/json` decodes numbers into `map[string]any`.

Example:

```json
{
  "data": {
    "count": 1
  }
}
```

Result:

```go
request.Data["count"] == float64(1)
```

### `json.RawMessage`

`json.RawMessage` fields are supported.

Example:

```go
type Request struct {
	Raw json.RawMessage `json:"raw"`
}
```

For input:

```json
{
  "raw": {
    "a": 1,
    "b": true
  }
}
```

The field will contain JSON bytes for the `raw` value.

Important: if the original input was already decoded into `map[string]any`, the original byte-for-byte JSON formatting is not preserved. The value is converted back to JSON using `json.Marshal`.

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
package main

import "strings"

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
	"name": {StartsWithA{}},
}
```

Pointer usage is also supported:

```go
fieldRules := map[string][]any{
	"name": {&StartsWithA{}},
}
```

Prefer value usage when possible.

## Unknown Rules

Unknown string rules are treated as validation errors.

Example:

```go
fieldRules := map[string][]any{
	"email": {"emial"},
}
```

This will return an error instead of silently ignoring the rule.

Using rule values avoids this kind of typo:

```go
fieldRules := map[string][]any{
	"email": {rules.Email{}},
}
```

## Laravel-like String Rules

String rules are supported for convenience and for easier migration from Laravel-style validation.

Example:

```go
fieldRules := map[string][]any{
    "name":  {"required", "string", "min:3"},
    "email": {"required", "email"},
    "date":  {"required", "date_format:2006-01-02"},
}
```

The preferred Go-style version is:

```go
fieldRules := map[string][]any{
    "name":  {rules.Required{}, rules.String{}, rules.Min{Min: 3}},
    "email": {rules.Required{}, rules.Email{}},
    "date":  {rules.Required{}, rules.Date{Format: "2006-01-02"}},
}
```

## License

See [LICENSE.md](LICENSE.md).
```