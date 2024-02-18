package logger

import (
	"go.uber.org/fx"

	"github.com/eliofery/golang-fullstack/pkg/eslog"
)

// NewModule ...
func NewModule() fx.Option {
	return fx.Module("logger",
		fx.Provide(
			NewConfig,
			New,
			eslog.New,
		),
	)
}
