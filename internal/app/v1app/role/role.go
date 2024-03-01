package role

import (
	"github.com/eliofery/golang-grpc/internal/app/v1app/role/api"
	"github.com/eliofery/golang-grpc/internal/app/v1app/role/repository"
	"github.com/eliofery/golang-grpc/internal/app/v1app/role/service"
	desc "github.com/eliofery/golang-grpc/pkg/api/role/v1"
	"go.uber.org/fx"
)

// NewModule ...
func NewModule() fx.Option {
	return fx.Module("role_v1",
		fx.Provide(
			repository.New,
			service.New,
			api.New,
		),
		fx.Invoke(
			desc.RegisterRoleV1ServiceServer,
			desc.RegisterRoleV1ServiceHandlerFromEndpoint,
		),
	)
}
