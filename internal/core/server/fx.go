package server

import (
	"github.com/eliofery/golang-fullstack/internal/v1app"
	"go.uber.org/fx"
)

// NewModule ...
func NewModule() fx.Option {
	return fx.Module("server",
		fx.Provide(
			NewConfig,
			NewGRPC,
			NewREST,
		),
		fx.Options(
			v1app.NewModule(),
			//v2app.NewModule(),

			NewGRPCModule(),
			NewRESTModule(),
		),
	)
}
