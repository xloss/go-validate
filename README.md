## Использование

```go
type ApiRequest struct {
Page      int     `json:"page" validate:"integer`
YearFrom  int     `json:"year_from" validate:"integer"`
YearTo    int     `json:"year_to" validate:"integer"`
RatingMin float64 `json:"rating_min" validate:"float"`
}

...

r, err := go_validate.Run[ApiRequest](body)
if len(err) > 0 {
// Обратока полученных ошибок
}
```

## Доступные валидации

| Название | Описание                 | Структура Error  |
|----------|--------------------------|------------------|
| required | Обязательное поле        | Name: "required" |         
| integer  | Число                    | Name: "integer"  |
| numeric  | Число с плавающей точкой | Name: "numeric"  |
| string   | Строка                   | Name: "string"   |
| boolean  | Логический тип           | Name: "boolean"  |

## Описание

```go
func Run[T interface{}](body io.ReadCloser) (*T, []Error)
```

```go
type Error struct {
Attribute string            `json:"attribute,omitempty"` // Название поля
Name      string            `json:"name,omitempty"`      // Название ошибки
Values    map[string]string `json:"values,omitempty"` // Дополнительные поля
}
```