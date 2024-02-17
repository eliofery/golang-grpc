package interceptor

import (
	"context"

	"github.com/bufbuild/protovalidate-go"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/proto"
)

// Validate ...
func Validate(
	ctx context.Context,
	req any,
	_ *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (any, error) {
	v, err := protovalidate.New()
	if err != nil {
		return nil, err
	}

	if msg, ok := req.(proto.Message); ok {
		if err = v.Validate(msg); err != nil {
			return nil, err
		}
	}

	return handler(ctx, req)
}
