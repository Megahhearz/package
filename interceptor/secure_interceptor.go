package interceptor

import (
	"context"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *GRPCInterceptor) UnarySecureInterceptor() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (resp interface{}, err error) {

		result, execErr := i.gobreak.Execute(func() (interface{}, error) {
			return handler(ctx, req)
		})

		if execErr != nil {
			i.logger.Error("Откленено/Ошибка безнес локиги", "метод", info.FullMethod, err)
			return nil, status.Error(codes.Unavailable, fmt.Sprintf("service overloaded: %v", execErr))
		}
		return result, nil
	}
}
