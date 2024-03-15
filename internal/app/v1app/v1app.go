package v1app

import (
	deniedToken "github.com/eliofery/golang-grpc/internal/app/v1app/denied_token"
	"github.com/eliofery/golang-grpc/internal/app/v1app/permission"
	"github.com/eliofery/golang-grpc/internal/app/v1app/role"
	rolePermission "github.com/eliofery/golang-grpc/internal/app/v1app/role_permission"
	"github.com/eliofery/golang-grpc/internal/app/v1app/setting"
	"github.com/eliofery/golang-grpc/internal/app/v1app/user"
	"go.uber.org/fx"
)

// NewModule ...
func NewModule() fx.Option {
	return fx.Module("app_v1",
		setting.NewModule(),
		deniedToken.NewModule(),
		role.NewModule(),
		user.NewModule(),
		permission.NewModule(),
		rolePermission.NewModule(),
	)
}
