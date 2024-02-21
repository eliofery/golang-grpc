package grpc

import (
	"go.uber.org/fx"
	"google.golang.org/grpc"
)

// GRPC ...
type GRPC struct {
	fx.Out

	*grpc.Server
	grpc.ServiceRegistrar
}

// New ...
func New(options []grpc.ServerOption) GRPC {
	server := grpc.NewServer(options...)

	return GRPC{
		Server:           server,
		ServiceRegistrar: server,
	}
}
