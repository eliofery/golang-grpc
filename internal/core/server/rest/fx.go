package rest

import (
	"context"
	"log/slog"
	"net/http"

	"github.com/eliofery/golang-fullstack/pkg/eslog"
	"go.uber.org/fx"
)

// NewModule ...
func NewModule() fx.Option {
	return fx.Module("rest",
		fx.Provide(
			NewConfig,
			NewOption,
			New,
			NewMiddleware,
		),
		fx.Invoke(
			func(lc fx.Lifecycle, middlewares http.Handler, config *Config, logger *eslog.Logger) {
				httpserv := &http.Server{
					Addr:         config.Address(),
					Handler:      middlewares,
					ReadTimeout:  config.Read,
					WriteTimeout: config.Write,
					IdleTimeout:  config.Idle,
				}

				lc.Append(fx.Hook{
					OnStart: func(_ context.Context) error {
						errCh := make(chan error)
						go func() {
							logger.Info("REST server start", slog.String("address", config.Address()))
							if err := httpserv.ListenAndServe(); err != nil {
								errCh <- err
							}
						}()

						select {
						case err := <-errCh:
							return err
						default:
							return nil
						}
					},
					OnStop: func(ctx context.Context) error {
						return httpserv.Shutdown(ctx)
					},
				})
			},
		),
	)
}
