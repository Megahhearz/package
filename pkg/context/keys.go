package shared_context

// ContextKeyEnum — тип для ключей в context.Context,
// чтобы избежать конфликтов с другими ключами в глобальном пространстве имён.
type ContextKeyEnum string

// Константы для ключей контекста.
const (
	ContextKeyEnumXRequestID ContextKeyEnum = "X_REQUEST_ID"
)

// String возвращает строковое представление ключа.
func (c ContextKeyEnum) String() string {
	return string(c)
}
