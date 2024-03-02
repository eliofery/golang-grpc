package setting

import (
	"github.com/eliofery/golang-grpc/internal/app/v1app/setting/repository"
	"go.uber.org/fx"
)

// NewModule ...
func NewModule() fx.Option {
	return fx.Module("setting_v1",
		fx.Provide(
			repository.New,
			//service.New,
			//api.New,
		),
		//fx.Invoke(),
	)
}
