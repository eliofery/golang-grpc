package rest

import (
	"go.uber.org/fx"
	"google.golang.org/grpc/credentials"

	"google.golang.org/grpc"
)

// Option ...
type Option struct {
	fx.Out

	Options []grpc.DialOption
}

// NewOption ...
func NewOption(transportCredentials credentials.TransportCredentials) Option {
	options := []grpc.DialOption{
		grpc.WithTransportCredentials(transportCredentials),
	}

	return Option{
		Options: options,
	}
}
