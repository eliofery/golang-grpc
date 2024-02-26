package deniedtoken

import (
	"github.com/eliofery/golang-grpc/internal/app/v1app/denied_token/api"
	"github.com/eliofery/golang-grpc/internal/app/v1app/denied_token/repository"
	"github.com/eliofery/golang-grpc/internal/app/v1app/denied_token/service"
	desc "github.com/eliofery/golang-grpc/pkg/api/denied_token/v1"
	"go.uber.org/fx"
)

// NewModule ...
func NewModule() fx.Option {
	return fx.Module("denied_token_v1",
		fx.Provide(
			repository.New,
			service.New,
			api.New,
		),
		fx.Invoke(
			desc.RegisterDeniedTokenV1ServiceServer,
			desc.RegisterDeniedTokenV1ServiceHandlerFromEndpoint,
		),
	)
}
