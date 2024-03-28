package token

import "go.uber.org/fx"

// NewModule ...
func NewModule() fx.Option {
	return fx.Module("token",
		fx.Provide(
			NewConfig,
			New,
		),
	)
}
