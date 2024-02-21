package server

import (
	"github.com/eliofery/golang-fullstack/internal/app"
	"github.com/eliofery/golang-fullstack/internal/core/server/grpc"
	"github.com/eliofery/golang-fullstack/internal/core/server/rest"
	"go.uber.org/fx"
)

// NewModule ...
func NewModule() fx.Option {
	return fx.Module("server",
		fx.Provide(
			NewTransport,
		),
		fx.Options(
			app.NewModule(),
			grpc.NewModule(),
			rest.NewModule(),
		),
	)
}
