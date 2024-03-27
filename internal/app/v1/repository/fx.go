package repository

import (
	"github.com/eliofery/golang-grpc/internal/app/v1/repository/auth"
	"github.com/eliofery/golang-grpc/internal/app/v1/repository/permission"
	"github.com/eliofery/golang-grpc/internal/app/v1/repository/role"
	"github.com/eliofery/golang-grpc/internal/app/v1/repository/setting"
	"github.com/eliofery/golang-grpc/internal/app/v1/repository/user"
	"go.uber.org/fx"
)

// NewModule ...
func NewModule() fx.Option {
	return fx.Module("repository",
		fx.Provide(
			setting.New,
			user.New,
			auth.New,
			role.New,
			permission.New,
		),
	)
}
