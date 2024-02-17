package main

import (
	"go.uber.org/fx"

	"github.com/eliofery/golang-fullstack/internal/core"
	"github.com/eliofery/golang-fullstack/internal/core/logger"
	"github.com/eliofery/golang-fullstack/internal/core/postgres"
	"github.com/eliofery/golang-fullstack/internal/core/server"
)

func main() {
	fx.New(
		fx.WithLogger(logger.WithLogger),
		fx.Provide(
			core.NewContext,
			core.NewCli,
			core.NewConfig,
			logger.NewConfig,
			logger.New,
		),
		fx.Options(
			postgres.NewModule(),
			server.NewModule(),
		),
	).Run()
}
