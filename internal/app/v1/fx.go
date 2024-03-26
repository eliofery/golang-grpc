package v1

import (
	"github.com/eliofery/golang-grpc/internal/app/v1/api"
	"github.com/eliofery/golang-grpc/internal/app/v1/repository"
	"github.com/eliofery/golang-grpc/internal/app/v1/service"
	desc "github.com/eliofery/golang-grpc/pkg/api/app/v1"
	"go.uber.org/fx"
)

// NewModule ...
func NewModule() fx.Option {
	return fx.Module("v1",
		fx.Options(
			repository.NewModule(),
			service.NewModule(),
			api.NewModule(),
		),
		fx.Invoke(
			desc.RegisterAppServiceServer,
			desc.RegisterAppServiceHandlerFromEndpoint,
		),
	)
}
