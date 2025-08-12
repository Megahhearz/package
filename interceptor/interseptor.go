package interceptor

import (
	"time"

	"github.com/Megahhearz/shared/logger"
	"github.com/sony/gobreaker"
)

// GRPCInterceptor представляет собой структуру, содержащую перехватчики (interceptors) для gRPC-сервера.
// На данный момент в ней хранится только логгер, но в будущем может быть расширена метриками, трассировкой и т.п.
type GRPCInterceptor struct {
	logger  logger.Logger // Логгер используется для логирования в interceptor'ах
	gobreak *gobreaker.CircuitBreaker
}

// New создаёт новый экземпляр GRPCInterceptor с переданным логгером.
// Это основной способ инициализации перехватчиков и внедрения зависимостей.
func New(logger logger.Logger) *GRPCInterceptor {
	return &GRPCInterceptor{
		logger:  logger,
		gobreak: NewServerBreaker(),
	}
}

func NewServerBreaker() *gobreaker.CircuitBreaker {
	settings := gobreaker.Settings{
		Name:        "ServerBreaker",
		Timeout:     5 * time.Second, // Время до Half-Open
		MaxRequests: 2,               // Сколько пускаем в Half-Open
		ReadyToTrip: func(c gobreaker.Counts) bool {
			failRatio := float64(c.TotalFailures) / float64(c.Requests)
			return c.Requests >= 10 && failRatio > 0.5
		},
	}

	return gobreaker.NewCircuitBreaker(settings)

}
