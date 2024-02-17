package server

import (
	"github.com/eliofery/golang-fullstack/internal/app/server/v1api"
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
			v1api.NewModule(),
			//v2api.NewModule(),

			NewGRPCModule(),
			NewRESTModule(),
		),
	)
}
