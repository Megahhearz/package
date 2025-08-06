package interceptor

import "github.com/Megahhearz/package/pkg/logger"

// GRPCInterceptor представляет собой структуру, содержащую перехватчики (interceptors) для gRPC-сервера.
// На данный момент в ней хранится только логгер, но в будущем может быть расширена метриками, трассировкой и т.п.
type GRPCInterceptor struct {
	logger logger.Logger // Логгер используется для логирования в interceptor'ах
}

// New создаёт новый экземпляр GRPCInterceptor с переданным логгером.
// Это основной способ инициализации перехватчиков и внедрения зависимостей.
func New(logger logger.Logger) *GRPCInterceptor {
	return &GRPCInterceptor{
		logger: logger,
	}
}
