package validate

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"reflect"
	"strings"
)

func ValidAndTrim(data interface{}) error {
	v := reflect.ValueOf(data)
	if v.Kind() == reflect.Ptr { // Проверяем, что data — это указатель
		v = v.Elem() // Получаем значение, на которое указывает указатель
	}
	if v.Kind() != reflect.Struct {
		return fmt.Errorf("data must be a struct or a pointer to a struct")
	}

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		if field.Kind() == reflect.String && field.CanSet() { // Проверяем, что поле можно изменить
			trimmed := strings.TrimSpace(field.String())
			field.SetString(trimmed)
		}
	}

	validate := validator.New()
	return validate.Struct(v.Interface()) // Валидируем исходную структуру
}
