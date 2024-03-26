package app

import (
	v1 "github.com/eliofery/golang-grpc/internal/app/v1"
	"go.uber.org/fx"
)

// NewModule ...
func NewModule() fx.Option {
	return fx.Module("app",
		v1.NewModule(),
		//v2.NewModule(),
	)
}
