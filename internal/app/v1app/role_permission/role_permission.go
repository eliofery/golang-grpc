package rolepermission

import (
	"github.com/eliofery/golang-grpc/internal/app/v1app/role_permission/repository"
	"go.uber.org/fx"
)

// NewModule ...
func NewModule() fx.Option {
	return fx.Module("role_permission_v1",
		fx.Provide(
			repository.New,
		),
	)
}
