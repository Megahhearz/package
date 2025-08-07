package proto_mapper

// ToInt64Proto преобразует указатель на int64 в значение int64.
// Если указатель nil, возвращает 0.
func ToInt64Proto(value *int64) int64 {
	if value == nil {
		return 0
	}

	return *value
}

// FromInt64Proto преобразует значение int64 в указатель на int64.
func FromInt64Proto(value int64) *int64 {
	return &value
}

// ToOptionalInt64Proto возвращает тот же указатель *int64 без изменений.
// Используется для передачи опциональных значений.
func ToOptionalInt64Proto(value *int64) *int64 {
	return value
}

// FromOptionalInt64Proto возвращает тот же указатель *int64 без изменений.
// Используется для приема опциональных значений.
func FromOptionalInt64Proto(value *int64) *int64 {
	return value
}
