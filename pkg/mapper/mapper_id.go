package proto_mapper

import (
	"github.com/google/uuid"
)

// ValidateID проверяет, что строка является валидным UUID.
// Возвращает true, если UUID валиден, иначе false.
func ValidateID(s string) bool {
	err := uuid.Validate(s)
	return err == nil
}

// ValidateIDs проверяет срез строк на валидность каждой строки как UUID.
// Возвращает false при первом невалидном UUID, иначе true.
func ValidateIDs(s []string) bool {
	valid := true
	for _, v := range s {
		err := uuid.Validate(v)
		if err != nil {
			valid = false
			break
		}
	}
	return valid
}

// FromIDProto конвертирует указатель на строку (proto UUID) в указатель на uuid.UUID.
// Возвращает nil, если строка пустая или невалидная.
func FromIDProto(s *string) *uuid.UUID {
	if s == nil || *s == "" {
		return nil
	}

	parsed, err := uuid.Parse(*s)
	if err != nil {
		return nil
	}

	return &parsed
}

// FromIDsProto конвертирует срез строк (proto UUID) в срез uuid.UUID.
// Возвращает nil, если хоть один UUID невалидный.
func FromIDsProto(s []string) []uuid.UUID {
	result := make([]uuid.UUID, len(s))
	for i, v := range s {
		parsed, err := uuid.Parse(v)
		if err != nil {
			return nil
		}
		result[i] = parsed
	}

	return result
}

// ToIDProto конвертирует указатель на uuid.UUID в строку.
// Возвращает пустую строку, если указатель nil.
func ToIDProto(u *uuid.UUID) string {
	if u == nil {
		return ""
	}

	return u.String()
}

// ToOptionalIDProto конвертирует указатель на uuid.UUID в указатель на строку.
// Возвращает nil, если указатель nil.
func ToOptionalIDProto(u *uuid.UUID) *string {
	if u == nil {
		return nil
	}

	res := u.String()

	return &res
}

// ToIDsProto конвертирует срез uuid.UUID в срез строк.
func ToIDsProto(u []uuid.UUID) []string {
	result := make([]string, len(u))
	for i, v := range u {
		result[i] = ToIDProto(&v)
	}
	return result
}
