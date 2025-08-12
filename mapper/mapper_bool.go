package proto_mapper

// ToBoolProto конвертирует указатель на bool в обычный bool.
// Если указатель nil, возвращает false по умолчанию.
// Это полезно, когда протобаф-структура ожидает неуказательное значение.
func ToBoolProto(value *bool) bool {
	if value == nil {
		return false
	}
	return *value
}

// FromBoolProto конвертирует обычный bool в указатель на bool.
// Это позволяет создавать указатели для передачи в протобаф, где bool - optional.
func FromBoolProto(value bool) *bool {
	return &value
}

// ToOptionalBoolProto просто возвращает указатель на bool без изменений.
// Используется, если нужно передать optional bool в протобаф без изменения.
func ToOptionalBoolProto(value *bool) *bool {
	return value
}

// FromOptionalBoolProto возвращает указатель на bool без изменений.
// Аналогично, используется для optional bool.
func FromOptionalBoolProto(value *bool) *bool {
	return value
}
