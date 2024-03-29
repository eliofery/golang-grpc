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
			NewClient,
			NewTransactionManager,
		),
		fx.Invoke(
			func(lc fx.Lifecycle, pgClient Client) {
				lc.Append(fx.Hook{
					OnStart: func(ctx context.Context) error {
						if err := pgClient.DB().Ping(ctx); err != nil {
							return err
						}

						return pgClient.Migrate()
					},
					OnStop: func(_ context.Context) error {
						return pgClient.Close()
					},
				})
			},
		),
	)
}
