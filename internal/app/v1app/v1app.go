package v1app

import (
	deniedtoken "github.com/eliofery/golang-grpc/internal/app/v1app/denied_token"
	"github.com/eliofery/golang-grpc/internal/app/v1app/role"
	"github.com/eliofery/golang-grpc/internal/app/v1app/setting"
	"github.com/eliofery/golang-grpc/internal/app/v1app/user"
	"go.uber.org/fx"
)

// NewModule ...
func NewModule() fx.Option {
	return fx.Module("app_v1",
		setting.NewModule(),
		deniedtoken.NewModule(),
		role.NewModule(),
		user.NewModule(),
	)
}
