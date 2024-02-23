package logger

import (
	"go.uber.org/fx"

	"github.com/eliofery/golang-grpc/pkg/eslog"
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
