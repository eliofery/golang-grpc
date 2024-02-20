package user

import (
	"github.com/eliofery/golang-fullstack/internal/app/v1app/user/api"
	"github.com/eliofery/golang-fullstack/internal/app/v1app/user/repository"
	"github.com/eliofery/golang-fullstack/internal/app/v1app/user/service"
	desc "github.com/eliofery/golang-fullstack/pkg/api/user/v1"
	"go.uber.org/fx"
)

// NewUserModule ...
func NewUserModule() fx.Option {
	return fx.Module("user_v1",
		fx.Provide(
			repository.New,
			service.New,
			api.New,
		),
		fx.Invoke(
			desc.RegisterUserV1ServiceServer,
			desc.RegisterUserV1ServiceHandlerFromEndpoint,
		),
	)
}
