package main

import (
	"log"

	"github.com/eliofery/golang-fullstack/docs/cli"
	"github.com/eliofery/golang-fullstack/internal/libs/config"
	"github.com/eliofery/golang-fullstack/internal/libs/logger"
	"github.com/eliofery/golang-fullstack/internal/libs/server"
	"github.com/eliofery/golang-fullstack/pkg/database/postgres"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
)

func main() {
	fx.New(
		fx.WithLogger(func(config *config.Config) fxevent.Logger {
			return logger.New(config)
		}),

		fx.Provide(
			cli.New,
			config.New,
			logger.New,
			postgres.New,
			server.NewGRPC,
			server.NewREST,
		),

		fx.Invoke(
			func(grpc *server.GRPC, rest *server.REST) {
				go grpc.Init()
				go func() {
					err := grpc.Run()
					if err != nil {
						log.Fatal(err)
					}
				}()

				go func() {
					err := rest.Init()
					if err != nil {
						log.Fatal(err)
					}
				}()
				go func() {
					err := rest.Run()
					if err != nil {
						log.Fatal(err)
					}
				}()
			},
		),
	).Run()
}
