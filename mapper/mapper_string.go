package proto_mapper

// ToStringProto преобразует указатель на string в значение string.
// Если указатель nil, возвращает пустую строку.
func ToStringProto(value *string) string {
	if value == nil {
		return ""
	}

	return *value
}

// FromStringProto преобразует значение string в указатель на string.
func FromStringProto(value string) *string {
	return &value
}

// ToOptionalStringProto возвращает тот же указатель *string без изменений.
// Используется для передачи опциональных значений.
func ToOptionalStringProto(value *string) *string {
	return value
}

// FromOptionalStringProto возвращает тот же указатель *string без изменений.
// Используется для приема опциональных значений.
func FromOptionalStringProto(value *string) *string {
	return value
}
