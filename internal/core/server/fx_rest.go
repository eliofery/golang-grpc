package server

import (
	"context"
	"log/slog"
	"net/http"

	"github.com/eliofery/golang-fullstack/pkg/eslog"
	"go.uber.org/fx"
)

// NewRESTModule ...
func NewRESTModule() fx.Option {
	return fx.Module("rest",
		fx.Invoke(
			func(lc fx.Lifecycle, handler http.Handler, config *Config, logger *eslog.Logger) {
				httpserv := &http.Server{
					Addr:         config.RESTAddress(),
					Handler:      handler,
					ReadTimeout:  config.Read,
					WriteTimeout: config.Write,
					IdleTimeout:  config.Idle,
				}

				lc.Append(fx.Hook{
					OnStart: func(_ context.Context) error {
						errCh := make(chan error)
						go func() {
							logger.Info("REST server start", slog.String("address", config.RESTAddress()))
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
