package main

import (
	"github.com/eliofery/golang-fullstack/internal/core/jwt"
	"go.uber.org/fx"

	"github.com/eliofery/golang-fullstack/internal/core"
	"github.com/eliofery/golang-fullstack/internal/core/database"
	"github.com/eliofery/golang-fullstack/internal/core/logger"
	"github.com/eliofery/golang-fullstack/internal/core/server"
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
			jwt.NewModule(),
			server.NewModule(),
		),
	).Run()
}
