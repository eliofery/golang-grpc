package password

import "go.uber.org/fx"

// NewModule ...
func NewModule() fx.Option {
	return fx.Module("password",
		fx.Provide(
			New,
		),
	)
}
