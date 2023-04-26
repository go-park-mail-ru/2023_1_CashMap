package utils

import (
	"depeche/internal/delivery/dto"
	"html"
	"reflect"
)

func Escaping[T *dto.SignUp | *dto.EditProfile | *dto.PostCreate | *dto.NewMessageDTO](obj T) T {
	value := reflect.ValueOf(obj).Elem()
	for i := 0; i < value.NumField(); i++ {
		valueField := value.Field(i)
		f := valueField.Interface()
		if valueStr, ok := f.(string); ok {
			escapeStr := html.EscapeString(valueStr)
			valueField.SetString(escapeStr)
		}
	}
	return obj
}
