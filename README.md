## Использование

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

## Доступные валидации

| Правило (Rule)  | Правило (string) | Описание                                                           | Структура Error                                                                              |
|-----------------|------------------|--------------------------------------------------------------------|----------------------------------------------------------------------------------------------|
| &rules.Required | required         | Обязательное поле                                                  | Name: "required"                                                                             |         
| &rules.Integer  | integer          | Число                                                              | Name: "integer"                                                                              |
| &rules.Numeric  | numeric          | Число с плавающей точкой                                           | Name: "numeric"                                                                              |
| &rules.String   | string           | Строка                                                             | Name: "string"                                                                               |
| &rules.Boolean  | boolean          | Логический тип                                                     | Name: "boolean"                                                                              |
| &rules.Min      | min:value        | Проверка на минимальное значение размера строки, числа или массива | Name: "min.numeric", "min.string", "min.array", Values["min"]: минимально требуемое значение |

## Описание

```go
func Run[T interface{}](data map[string]any, fieldRules map[string][]any) (*T, []Error)
```

```go
type Error struct {
Attribute string            `json:"attribute,omitempty"` // Название поля
Name      string            `json:"name,omitempty"`      // Название ошибки
Values    map[string]any    `json:"values,omitempty"` // Дополнительные поля
}
```

## Интерфейс для своих привил

```go
type Rule interface {
GetName() string // отдача названия поля для Error.Name
GetValues() map[string]any // отдача дополнительных значений для Error.Values
Validate(value any) bool    // результат валидация значения
}
```