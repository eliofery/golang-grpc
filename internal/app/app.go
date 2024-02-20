package app

import (
	"go.uber.org/fx"

	"github.com/eliofery/golang-fullstack/internal/app/v1app"
)

// NewModule ...
func NewModule() fx.Option {
	return fx.Module("app",
		v1app.NewModule(),
		//v2app.NewModule(),
	)
}
