package interceptor

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// UnaryPanicRecoveryInterceptor — серверный интерсептор для обработки паник.
// Позволяет предотвратить падение сервера, если внутри gRPC-хендлера произойдёт panic.
// Вместо падения вернётся статус INTERNAL с описанием "внутренняя ошибка сервера".
func (i *GRPCInterceptor) UnaryPanicRecoveryInterceptor() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req any,
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler) (resp any, err error) {
		// Отложенное восстановление после panic, если она возникнет внутри handler-а
		defer func() {
			if r := recover(); r != nil {
				i.logger.Error("Паника перехвачена", "метод", info.FullMethod, err, r)
				err = status.Error(codes.Internal, "внутренняя ошибка сервера")
			}
		}()

		// Выполнение основного обработчика запроса
		return handler(ctx, req)
	}
}
