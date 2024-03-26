package api

import (
	"go.uber.org/fx"
)

// NewModule ...
func NewModule() fx.Option {
	return fx.Module("api",
		fx.Provide(
			New,
		),
	)
}
