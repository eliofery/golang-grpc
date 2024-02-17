package main

import (
	"go.uber.org/fx"

	"github.com/eliofery/golang-fullstack/internal/app"
	"github.com/eliofery/golang-fullstack/internal/app/logger"
	"github.com/eliofery/golang-fullstack/internal/app/postgres"
	"github.com/eliofery/golang-fullstack/internal/app/server"
)

func main() {
	fx.New(
		fx.WithLogger(logger.WithLogger),
		fx.Provide(
			app.NewContext,
			app.NewCli,
			app.NewConfig,
			logger.NewConfig,
			logger.New,
		),
		fx.Options(
			postgres.NewModule(),
			server.NewModule(),
		),
	).Run()
}
