package server

import (
	"context"
	"log/slog"
	"net/http"

	"github.com/eliofery/golang-fullstack/internal/libs/config"
	"github.com/eliofery/golang-fullstack/pkg/eslog"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"go.uber.org/fx"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// NewREST ...
func NewREST(config *config.Config) (*runtime.ServeMux, string, []grpc.DialOption) {
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	return mux, config.GRPCAddress(), opts
}

// InvokeREST ...
func InvokeREST(lc fx.Lifecycle, mux *runtime.ServeMux, config *config.Config, logger *eslog.Logger) {
	server := &http.Server{
		Addr:         config.RESTAddress(),
		Handler:      allowCORS(mux),
		ReadTimeout:  config.Read,
		WriteTimeout: config.Write,
		IdleTimeout:  config.Idle,
	}

	lc.Append(fx.Hook{
		OnStart: func(_ context.Context) error {
			errCh := make(chan error)
			go func() {
				logger.Info("REST server start", slog.String("address", config.Server.RESTAddress()))
				if err := server.ListenAndServe(); err != nil {
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
			return server.Shutdown(ctx)
		},
	})
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
