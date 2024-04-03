package server

import (
	"github.com/eliofery/golang-grpc/internal/core/server/grpc"
	"github.com/eliofery/golang-grpc/internal/core/server/rest"
	"go.uber.org/fx"
)

// NewModule ...
func NewModule() fx.Option {
	return fx.Module("server",
		fx.Provide(
			NewTransport,
		),
		fx.Options(
			grpc.NewModule(),
			rest.NewModule(),
		),
	)
}
