package server

import (
	"context"
	"log/slog"
	"net"

	"github.com/eliofery/golang-fullstack/internal/libs/config"
	"github.com/eliofery/golang-fullstack/pkg/eslog"
	"go.uber.org/fx"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
)

// NewGRPC ...
func NewGRPC() (*grpc.Server, grpc.ServiceRegistrar) {
	server := grpc.NewServer(grpc.Creds(insecure.NewCredentials()))
	reflection.Register(server)

	return server, server
}

// InvokeGRPC ...
func InvokeGRPC(lc fx.Lifecycle, server *grpc.Server, config *config.Config, logger *eslog.Logger) {
	lc.Append(fx.Hook{
		OnStart: func(_ context.Context) error {
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
}
