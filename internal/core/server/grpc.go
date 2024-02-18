package server

import (
	"go.uber.org/fx"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/eliofery/golang-fullstack/internal/core/server/interceptor"
)

// GRPC ...
type GRPC struct {
	fx.Out

	Server           *grpc.Server
	ServiceRegistrar grpc.ServiceRegistrar
}

// NewGRPC ...
func NewGRPC() GRPC {
	server := grpc.NewServer(
		grpc.Creds(insecure.NewCredentials()),
		interceptor.ChainUnaryInterceptors(),
	)

	return GRPC{
		Server:           server,
		ServiceRegistrar: server,
	}
}
