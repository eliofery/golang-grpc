package server

import (
	"context"
	"log/slog"
	"net/http"

	"github.com/eliofery/golang-fullstack/internal/app/server/v1api"
	"github.com/eliofery/golang-fullstack/pkg/eslog"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"go.uber.org/fx"
	"google.golang.org/grpc"
)

// NewRESTModule ...
func NewRESTModule() fx.Option {
	return fx.Module("rest",
		fx.Provide(
			func(server *REST, config *Config) (*runtime.ServeMux, string, []grpc.DialOption) {
				return server.Mux(), config.GRPCAddress(), server.Opts()
			},
		),
		fx.Invoke(
			v1api.RegisterServiceHandlerFromEndpoints...,
		//v2api.RegisterServiceHandlerFromEndpoints...,
		),
		fx.Invoke(
			func(lc fx.Lifecycle, server *REST, config *Config, logger *eslog.Logger) {
				httpserv := &http.Server{
					Addr:         config.RESTAddress(),
					Handler:      allowCORS(server.mux),
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

// allowCORS ...
func allowCORS(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		handler.ServeHTTP(w, r)
	})
}
