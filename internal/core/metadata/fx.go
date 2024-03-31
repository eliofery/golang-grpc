package metadata

import "go.uber.org/fx"

// NewModule ...
func NewModule() fx.Option {
	return fx.Module("metadata",
		fx.Provide(
			New,
		),
	)
}
