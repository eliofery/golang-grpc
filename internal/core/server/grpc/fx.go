package grpc

import (
	"context"
	"log/slog"
	"net"

	"github.com/eliofery/golang-grpc/internal/core/server/grpc/interceptor"
	"github.com/eliofery/golang-grpc/pkg/eslog"
	"go.uber.org/fx"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// NewModule ...
func NewModule() fx.Option {
	return fx.Module("grpc",
		fx.Provide(
			NewConfig,
			NewOption,
			New,
		),
		fx.Options(
			interceptor.NewModule(),
		),
		fx.Invoke(
			func(lc fx.Lifecycle, server *grpc.Server, config *Config, logger eslog.Logger) {
				lc.Append(fx.Hook{
					OnStart: func(_ context.Context) error {
						reflection.Register(server)

						list, err := net.Listen("tcp", config.Address())
						if err != nil {
							return err
						}

						errCh := make(chan error)
						go func() {
							logger.Info("GRPC server start", slog.String("address", config.Address()))
							if err = server.Serve(list); err != nil {
								errCh <- err
							}
						}()

						select {
						case err = <-errCh:
							return err
						default:
							return nil
						}
					},
					OnStop: func(_ context.Context) error {
						server.Stop()

						return nil
					},
				})
			},
		),
	)
}
