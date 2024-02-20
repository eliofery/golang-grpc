package v1app

import (
	"go.uber.org/fx"

	"github.com/eliofery/golang-fullstack/internal/app/v1app/user"
)

// NewModule ...
func NewModule() fx.Option {
	return fx.Module("app_v1",
		user.NewUserModule(),
		//NewOtherModule(),
	)
}
