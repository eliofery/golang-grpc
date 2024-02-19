package grpc

import (
	"go.uber.org/fx"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

// Option ...
type Option struct {
	fx.Out

	Options []grpc.ServerOption
}

// NewOption ...
func NewOption(
	interceptors []grpc.UnaryServerInterceptor,
	transportCredentials credentials.TransportCredentials,
) Option {
	options := []grpc.ServerOption{
		grpc.Creds(transportCredentials),
		grpc.ChainUnaryInterceptor(interceptors...),
	}

	return Option{
		Options: options,
	}
}
