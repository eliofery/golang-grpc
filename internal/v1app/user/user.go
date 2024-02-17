package user

import (
	"github.com/eliofery/golang-fullstack/internal/v1app/user/api"
	"github.com/eliofery/golang-fullstack/internal/v1app/user/repository"
	"github.com/eliofery/golang-fullstack/internal/v1app/user/service"
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
	)
}
