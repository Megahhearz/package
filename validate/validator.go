package validate

import "github.com/go-playground/validator"

// Создаём экземпляр валидатора из пакета go-playground/validator.
// Он используется для проверки структур на основе тегов в полях.
var validate = validator.New()

// Validate принимает любой объект (обычно структуру) и проверяет его
// на валидность по заданным в тегах правилам.
// Возвращает ошибку, если валидация не пройдена, иначе nil.
func Validate(r any) error {
	return validate.Struct(r)
}
