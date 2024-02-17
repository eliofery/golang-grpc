package server

import (
	"github.com/eliofery/golang-fullstack/internal/interceptor"
	"go.uber.org/fx"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
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
		grpc.UnaryInterceptor(interceptor.Validate),
	)

	return GRPC{
		Server:           server,
		ServiceRegistrar: server,
	}
}
