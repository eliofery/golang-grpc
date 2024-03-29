package redis

import (
	"context"

	"go.uber.org/fx"
)

// NewModule ...
func NewModule() fx.Option {
	return fx.Module("redis",
		fx.Provide(
			NewConfig,
			NewClient,
			NewTransactionManager,
		),
		fx.Invoke(
			func(lc fx.Lifecycle, redis Client) {
				lc.Append(fx.Hook{
					OnStart: func(ctx context.Context) error {
						return redis.DB().Ping(ctx)
					},
					OnStop: func(_ context.Context) error {
						return redis.Close()
					},
				})
			},
		),
	)
}
