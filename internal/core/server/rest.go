package server

import (
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"go.uber.org/fx"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// REST ...
type REST struct {
	fx.Out

	Mux         *runtime.ServeMux
	Opts        []grpc.DialOption
	GRPCAddress string
}

// NewREST ...
func NewREST(config *Config) REST {
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	return REST{
		Mux:         mux,
		Opts:        opts,
		GRPCAddress: config.GRPCAddress(),
	}
}
