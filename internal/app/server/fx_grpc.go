package server

import (
	"context"
	"log/slog"
	"net"

	"google.golang.org/grpc"

	"github.com/eliofery/golang-fullstack/internal/app/server/v1api"
	"github.com/eliofery/golang-fullstack/pkg/eslog"
	"go.uber.org/fx"
	"google.golang.org/grpc/reflection"
)

// NewGRPCModule ...
func NewGRPCModule() fx.Option {
	return fx.Module("grpc",
		fx.Invoke(
			v1api.RegisterServiceServers...,
		//v2api.RegisterServiceServers...,
		),
		fx.Invoke(
			func(lc fx.Lifecycle, server *grpc.Server, config *Config, logger *eslog.Logger) {
				lc.Append(fx.Hook{
					OnStart: func(_ context.Context) error {
						reflection.Register(server)

						list, err := net.Listen("tcp", config.GRPCAddress())
						if err != nil {
							return err
						}

						errCh := make(chan error)
						go func() {
							logger.Info("GRPC server start", slog.String("address", config.GRPCAddress()))
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
