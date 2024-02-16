package main

import (
	"context"

	"github.com/eliofery/golang-fullstack/docs/cli"
	apiUserV1 "github.com/eliofery/golang-fullstack/internal/app/api/user/v1"
	repositoryUserV1 "github.com/eliofery/golang-fullstack/internal/app/repository/user/v1"
	serviceUserV1 "github.com/eliofery/golang-fullstack/internal/app/service/user/v1"
	"github.com/eliofery/golang-fullstack/internal/libs/config"
	"github.com/eliofery/golang-fullstack/internal/libs/database/postgres"
	"github.com/eliofery/golang-fullstack/internal/libs/logger"
	"github.com/eliofery/golang-fullstack/internal/libs/server"
	serverV1 "github.com/eliofery/golang-fullstack/pkg/api/user/v1"
	"go.uber.org/fx"
)

func contextNew() context.Context {
	return context.Background()
}

func main() {
	fx.New(
		fx.WithLogger(logger.WithLogger),

		fx.Provide(
			contextNew,
			cli.New,
			config.New,
			logger.New,
			postgres.New,
			server.NewGRPC,
			server.NewREST,

			repositoryUserV1.New,
			serviceUserV1.New,
			apiUserV1.New,

			//repositoryUserV2.New,
			//serviceUserV2.New,
			//fx.Annotate(
			//    apiUserV2.New,
			//    fx.As(new(serverV2.UserV2ServiceServer)),
			//),
		),

		fx.Invoke(
			serverV1.RegisterUserV1ServiceServer,
			serverV1.RegisterUserV1ServiceHandlerFromEndpoint,

			//serverV2.RegisterUserV2ServiceServer,
			//serverV2.RegisterUserV2ServiceHandlerFromEndpoint,

			server.InvokeGRPC,
			server.InvokeREST,
		),
	).Run()
}
