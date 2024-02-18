package server

import (
	"net/http"

	"github.com/eliofery/golang-fullstack/internal/core/server/middleware"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"go.uber.org/fx"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// REST ...
type REST struct {
	fx.Out

	Mux         *runtime.ServeMux
	Handler     http.Handler
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
		Handler:     middleware.ChainMiddlewares(mux),
		Opts:        opts,
		GRPCAddress: config.GRPCAddress(),
	}
}
