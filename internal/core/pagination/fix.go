package pagination

import (
	"go.uber.org/fx"
)

// NewModule ...
func NewModule() fx.Option {
	return fx.Module("pagination",
		fx.Provide(
			NewConfig,
			New,
		),
	)
}
