package main

import (
	"github.com/eliofery/golang-grpc/internal/app"
	"github.com/eliofery/golang-grpc/internal/core"
	"github.com/eliofery/golang-grpc/internal/core/authorize"
	"github.com/eliofery/golang-grpc/internal/core/database"
	"github.com/eliofery/golang-grpc/internal/core/logger"
	"github.com/eliofery/golang-grpc/internal/core/metadata"
	"github.com/eliofery/golang-grpc/internal/core/pagination"
	"github.com/eliofery/golang-grpc/internal/core/server"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		fx.WithLogger(logger.WithLogger),
		fx.Provide(
			core.NewContext,
			core.NewCli,
			core.NewConfig,
		),
		fx.Options(
			logger.NewModule(),
			database.NewModule(),
			authorize.NewModule(),
			metadata.NewModule(),
			pagination.NewModule(),
			app.NewModule(),
			server.NewModule(),
		),
	).Run()
}
