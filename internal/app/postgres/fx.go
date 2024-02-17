package postgres

import (
	"context"

	"go.uber.org/fx"
)

// NewModule ...
func NewModule() fx.Option {
	return fx.Module("postgres",
		fx.Provide(
			NewConfig,
			New,
		),
		fx.Invoke(
			func(lc fx.Lifecycle, postgres *Postgres) {
				lc.Append(fx.Hook{
					OnStart: func(_ context.Context) error {
						if err := postgres.Migrate(); err != nil {
							return err
						}

						return nil
					},
					OnStop: func(_ context.Context) error {
						return postgres.Close()
					},
				})
			},
		),
	)
}
