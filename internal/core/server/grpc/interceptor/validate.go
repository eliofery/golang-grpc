package interceptor

import (
	"context"
	"log/slog"

	"github.com/bufbuild/protovalidate-go"
	"github.com/eliofery/golang-grpc/pkg/eslog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

// validate ...
func validate(logger *eslog.Logger) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, _ *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
		v, err := protovalidate.New()
		if err != nil {
			logger.Debug("interceptor.Validate", slog.String("err", err.Error()))
			return nil, err
		}

		if msg, ok := req.(proto.Message); ok {
			if err = v.Validate(msg); err != nil {
				logger.Debug("interceptor.Validate", slog.String("err", err.Error()))
				return nil, status.Error(codes.InvalidArgument, err.Error())
			}
		}

		return handler(ctx, req)
	}
}
