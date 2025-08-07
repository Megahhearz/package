package shared_context

import (
	"context"
)

// XRequestIDFromContext извлекает из контекста значение X-Request-ID,
// которое обычно используется для трассировки запросов и логирования.
// Возвращает пустую строку, если в контексте нет значения с этим ключом.
func XRequestIDFromContext(ctx context.Context) string {
	// Извлекаем значение по ключу ContextKeyEnumXRequestID из контекста.
	// Приводим его к типу string. Если приведение неудачно — возвращается пустая строка.
	xRequestID, _ := ctx.Value(ContextKeyEnumXRequestID).(string)
	return xRequestID
}
