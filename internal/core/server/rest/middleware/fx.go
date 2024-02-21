package middleware

import "go.uber.org/fx"

// NewModule ...
func NewModule() fx.Option {
	return fx.Module("middleware",
		fx.Provide(
			NewConfig,
			New,
		),
	)
}
