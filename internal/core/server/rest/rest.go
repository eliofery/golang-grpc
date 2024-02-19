package rest

import (
	"github.com/eliofery/golang-fullstack/internal/core/server/grpc"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"go.uber.org/fx"
)

// REST ...
type REST struct {
	fx.Out

	Mux         *runtime.ServeMux
	GRPCAddress string
}

// New ...
func New(config *grpc.Config) REST {
	return REST{
		Mux:         runtime.NewServeMux(),
		GRPCAddress: config.Address(),
	}
}
