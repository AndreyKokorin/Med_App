package validate

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"reflect"
	"strings"
)

func ValidAndTrim(data interface{}) error {
	// Проверяем, что data — это указатель на структуру
	v := reflect.ValueOf(data)
	if v.Kind() != reflect.Ptr {
		return fmt.Errorf("data must be a pointer to a struct")
	}
	v = v.Elem() // Получаем значение, на которое указывает указатель
	if v.Kind() != reflect.Struct {
		return fmt.Errorf("data must be a pointer to a struct")
	}

	// Удаляем пробелы из строковых полей
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		if field.Kind() == reflect.String && field.CanSet() { // Проверяем, что поле можно изменить
			trimmed := strings.TrimSpace(field.String())
			field.SetString(trimmed)
		}
	}

	// Валидируем структуру
	validate := validator.New()
	return validate.Struct(data) // Передаём указатель на структуру
}
