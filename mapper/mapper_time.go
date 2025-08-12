package proto_mapper

import (
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"
)

// FromTimestampProto преобразует protobuf Timestamp в указатель на time.Time.
// Если входящее значение nil или невалидно, возвращает nil.
func FromTimestampProto(value *timestamppb.Timestamp) *time.Time {
	if value == nil || !value.IsValid() {
		return nil
	}

	_t := value.AsTime()
	return &_t
}

// ToTimestampProto преобразует указатель на time.Time в protobuf Timestamp.
// Если входящий указатель nil или время является нулевым, возвращает nil.
func ToTimestampProto(value *time.Time) *timestamppb.Timestamp {
	if value == nil || value.IsZero() {
		return nil
	}

	return timestamppb.New(*value)
}
