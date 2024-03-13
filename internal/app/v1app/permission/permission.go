package permission

import (
	"github.com/eliofery/golang-grpc/internal/app/v1app/permission/api"
	"github.com/eliofery/golang-grpc/internal/app/v1app/permission/repository"
	"github.com/eliofery/golang-grpc/internal/app/v1app/permission/service"
	desc "github.com/eliofery/golang-grpc/pkg/api/permission/v1"
	"go.uber.org/fx"
)

// NewModule ...
func NewModule() fx.Option {
	return fx.Module("permission_v1",
		fx.Provide(
			repository.New,
			service.New,
			api.New,
		),
		fx.Invoke(
			desc.RegisterPermissionV1ServiceServer,
			desc.RegisterPermissionV1ServiceHandlerFromEndpoint,
		),
	)
}
