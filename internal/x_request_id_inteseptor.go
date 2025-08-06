package interceptor

import (
	"context"
	"log/slog"

	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

// XRequestIDInterceptor — серверный интерсептор для добавления/валидации x-request-id.
// При отсутствии ID он будет сгенерирован. Это важно для трассировки и корреляции запросов между сервисами.
func (i *GRPCInterceptor) XRequestIDInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
		// Извлекаем метаданные из входящего контекста
		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			md = metadata.New(nil)
		}

		// Извлекаем или генерируем x-request-id
		requestID := ""
		if ids := md.Get("x-request-id"); len(ids) > 0 {
			requestID = ids[0]
			if _, err := uuid.Parse(requestID); err != nil {
				slog.Warn("невалидный x-request-id, генерируем новый", "invalid_id", requestID)
				requestID = uuid.New().String()
				md.Set("x-request-id", requestID)
			}
		} else {
			requestID = uuid.New().String()
			md.Set("x-request-id", requestID)
		}

		// Создаём новый контекст с обновлёнными метаданными
		newCtx := metadata.NewOutgoingContext(ctx, md)

		// Передаём управление хендлеру с новым контекстом
		return handler(newCtx, req)
	}
}

// ClientXRequestIDInterceptor — клиентский интерсептор для добавления/валидации x-request-id.
// Используется для проксирования ID между клиентом и сервером при вызове gRPC-клиентов.
func (i *GRPCInterceptor) ClientXRequestIDInterceptor() grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		// Извлекаем метаданные из исходящего контекста клиента
		md, ok := metadata.FromOutgoingContext(ctx)
		if !ok {
			md = metadata.New(nil)
		}

		// Извлекаем или генерируем x-request-id
		requestID := ""
		if ids := md.Get("x-request-id"); len(ids) > 0 {
			requestID = ids[0]
			if _, err := uuid.Parse(requestID); err != nil {
				slog.Warn("невалидный x-request-id, генерируем новый", "invalid_id", requestID)
				requestID = uuid.New().String()
				md.Set("x-request-id", requestID)
			}
		} else {
			requestID = uuid.New().String()
			md.Set("x-request-id", requestID)
		}

		// Обновляем контекст с новым request-id
		newCtx := metadata.NewOutgoingContext(ctx, md)

		// Выполняем вызов RPC с обновлённым контекстом
		return invoker(newCtx, method, req, reply, cc, opts...)
	}
}
