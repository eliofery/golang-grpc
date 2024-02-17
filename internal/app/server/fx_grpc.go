package server

import (
	"context"
	"log/slog"
	"net"

	"github.com/eliofery/golang-fullstack/internal/app/server/v1api"
	"github.com/eliofery/golang-fullstack/pkg/eslog"
	"go.uber.org/fx"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// NewGRPCModule ...
func NewGRPCModule() fx.Option {
	return fx.Module("grpc",
		fx.Provide(
			func(server *GRPC) grpc.ServiceRegistrar {
				return server.GRPC()
			},
		),
		fx.Invoke(
			v1api.RegisterServiceServers...,
		//v2api.RegisterServiceServers...,
		),
		fx.Invoke(
			func(lc fx.Lifecycle, server *GRPC, config *Config, logger *eslog.Logger) {
				lc.Append(fx.Hook{
					OnStart: func(_ context.Context) error {
						reflection.Register(server.grpc)

						list, err := net.Listen("tcp", config.GRPCAddress())
						if err != nil {
							return err
						}

						errCh := make(chan error)
						go func() {
							logger.Info("GRPC server start", slog.String("address", config.GRPCAddress()))
							if err = server.grpc.Serve(list); err != nil {
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
						server.grpc.Stop()

						return nil
					},
				})
			},
		),
	)
}
