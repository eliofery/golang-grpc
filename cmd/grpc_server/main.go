package main

import (
	"github.com/eliofery/golang-fullstack/docs/cli"
	"github.com/eliofery/golang-fullstack/internal/app"
	"github.com/eliofery/golang-fullstack/internal/config"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
)

func main() {
	fx.New(
		fx.WithLogger(func(cfg *config.Config) fxevent.Logger {
			return cfg.NewLogger()
		}),

		fx.Provide(
			cli.New,
			config.New,
			app.New,
		),

		fx.Invoke(app.Run),
	).Run()
}
