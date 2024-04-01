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

// Validator ...
type Validator interface {
	Validate() grpc.UnaryServerInterceptor
}

type validator struct {
	logger eslog.Logger
}

// NewValidator ...
func NewValidator(logger eslog.Logger) Validator {
	return &validator{
		logger: logger,
	}
}

// Validate ...
func (v *validator) Validate() grpc.UnaryServerInterceptor {
	op := "core.server.grpc.interceptor.Validate"

	return func(ctx context.Context, req any, _ *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
		validate, err := protovalidate.New()
		if err != nil {
			v.logger.Debug(op, slog.String("err", err.Error()))
			return nil, err
		}

		if msg, ok := req.(proto.Message); ok {
			if err = validate.Validate(msg); err != nil {
				v.logger.Debug(op, slog.String("err", err.Error()))
				return nil, status.Error(codes.InvalidArgument, err.Error())
			}
		}

		return handler(ctx, req)
	}
}
