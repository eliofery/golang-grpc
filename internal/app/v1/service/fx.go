package service

import (
	"github.com/eliofery/golang-grpc/internal/app/v1/service/auth"
	"github.com/eliofery/golang-grpc/internal/app/v1/service/user"
	"go.uber.org/fx"
)

// NewModule ...
func NewModule() fx.Option {
	return fx.Module("service",
		fx.Provide(
			auth.New,
			user.New,
		),
	)
}
