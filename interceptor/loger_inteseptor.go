package interceptor

import (
	"context"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

// LoggerInterceptor — unary-интерсептор, логирующий входящие запросы, их продолжительность и ошибки, если они есть.
func (i *GRPCInterceptor) LoggerInterceptor() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req any,
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler) (any, error) {

		start := time.Now() // Засекаем время начала запроса

		// Пытаемся извлечь x-request-id из метаданных gRPC-контекста
		md, _ := metadata.FromIncomingContext(ctx)
		requestID := ""
		if ids := md.Get("x-request-id"); len(ids) > 0 {
			requestID = ids[0]
		}

		// Логируем начало запроса
		i.logger.Info("Запрос начат", "метод", info.FullMethod, "x-request-id", requestID)

		// Вызываем основной обработчик запроса
		resp, err := handler(ctx, req)

		// Вычисляем длительность выполнения запроса
		duration := time.Since(start)

		// Логируем результат выполнения: с ошибкой или успешно
		if err != nil {
			i.logger.Error("Запрос завершился с ошибкой", "метод", info.FullMethod, err)
		} else {
			i.logger.Info("Запрос успешно завершен", "метод", info.FullMethod, "x-request-id", requestID, "длительность", duration)
		}

		return resp, err
	}
}
