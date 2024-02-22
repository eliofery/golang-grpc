package jwt

import "go.uber.org/fx"

// NewModule ...
func NewModule() fx.Option {
	return fx.Module("jwt",
		fx.Provide(
			NewConfig,
			New,
		),
	)
}
