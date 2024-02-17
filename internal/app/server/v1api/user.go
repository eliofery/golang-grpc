package v1api

import (
	apiUserV1 "github.com/eliofery/golang-fullstack/internal/api/user/v1"
	repositoryUserV1 "github.com/eliofery/golang-fullstack/internal/repository/user/v1"
	serviceUserV1 "github.com/eliofery/golang-fullstack/internal/service/user/v1"
	"go.uber.org/fx"
)

// NewUserModule ...
func NewUserModule() fx.Option {
	return fx.Module("user_v1",
		fx.Provide(
			repositoryUserV1.New,
			serviceUserV1.New,
			apiUserV1.New,
		),
	)
}
