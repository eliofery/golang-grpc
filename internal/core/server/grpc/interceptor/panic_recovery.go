package interceptor

import (
	"bytes"
	"context"
	"log/slog"
	"runtime/debug"

	"github.com/eliofery/golang-grpc/pkg/eslog"
	"google.golang.org/grpc"
)

// panicRecovery ...
func panicRecovery(logger *eslog.Logger) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, _ *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
		defer func() {
			if r := recover(); r != nil {
				logger.Error("Recovered from panic", slog.Any("err", r))

				buf := new(bytes.Buffer)
				buf.Write(debug.Stack())

				logger.Error("Stack trace", slog.String("err", buf.String()))
			}
		}()

		return handler(ctx, req)
	}
}
