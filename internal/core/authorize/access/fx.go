package access

import "go.uber.org/fx"

// NewModule ...
func NewModule() fx.Option {
	return fx.Module("access",
		fx.Provide(
			New,
		),
	)
}
