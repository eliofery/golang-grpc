package grpc

import (
	"context"

	"github.com/bufbuild/protovalidate-go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

// NewInterceptor ...
func NewInterceptor() []grpc.UnaryServerInterceptor {
	return []grpc.UnaryServerInterceptor{
		validate(),
	}
}

// validate ...
func validate() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, _ *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
		v, err := protovalidate.New()
		if err != nil {
			return nil, err
		}

		if msg, ok := req.(proto.Message); ok {
			if err = v.Validate(msg); err != nil {
				return nil, status.Error(codes.InvalidArgument, err.Error())
			}
		}

		return handler(ctx, req)
	}
}