## Usage

```go
import govalidate "github.com/xloss/go-validate"

type apiRequest struct {
    IntegerValue  int     `json:"integer_value"`
    StringValue   string  `json:"string_value"`
    NumericValue  float64 `json:"numeric_value"`
    BooleanValue  bool    `json:"boolean_value"`
    RequiredValue string  `json:"required_value"`
}

jsonText := `{"integer_value": 1, "string_value": "b", "numeric_value": 3.1, "boolean_value": true, "required_value": "r"}`

var data map[string]any
errUnmarshal := json.Unmarshal([]byte(jsonText), &data)
if errUnmarshal != nil {
    // ...
}

fieldRules := map[string][]any{
    "integer_value":  {&rules.Integer{}},
    "string_value":   {"string"},
    "numeric_value":  {&rules.Numeric{}},
    "boolean_value":  {"boolean"},
    "required_value": {&rules.Required{}, &rules.String{}},
}

r, errors := govalidate.Run[apiRequest](data, fieldRules)
if len(errors) != 0 {
    // ...
}
```

## Available Validations

| Rule            | Rule (string) | Description                                                  | Error Structure                                                                       |
|-----------------|---------------|--------------------------------------------------------------|---------------------------------------------------------------------------------------|
| &rules.Required | required      | Required field                                               | Name: "required"                                                                      |         
| &rules.Integer  | integer       | Integer                                                      | Name: "integer"                                                                       |
| &rules.Numeric  | numeric       | Floating-point number                                        | Name: "numeric"                                                                       |
| &rules.String   | string        | String                                                       | Name: "string"                                                                        |
| &rules.Boolean  | boolean       | Boolean type                                                 | Name: "boolean"                                                                       |
| &rules.Min      | min:value     | Minimum value check for string size, number, or array length | Name: "min.numeric", "min.string", "min.array", Values["min"]: minimum required value |
| &rules.Domain   | domain        | Validation for a valid domain name                           | Name: "domain"                                                                        |
| &rules.Date     | date          | Date Validation in RFC3339 format                            | Name: "date"                                                                          |

## Description

```go
func Run[T interface{}](data map[string]any, fieldRules map[string][]any) (*T, []Error)
```

```go
type Error struct {
    Attribute string            `json:"attribute,omitempty"` // Field name
    Name      string            `json:"name,omitempty"`      // Error name
    Values    map[string]any    `json:"values,omitempty"` // Additional fields
}
```

## Custom Rule Interface

```go
type Rule interface {
    GetName() string // Returns the field name for Error.Name
    GetValues() map[string]any // Returns additional values for Error.Values
    Validate(value any) bool   // Validation result for the value
}
```