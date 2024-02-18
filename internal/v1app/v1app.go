package v1app

import (
	"github.com/eliofery/golang-fullstack/internal/v1app/user"
	"go.uber.org/fx"
)

// NewModule ...
func NewModule() fx.Option {
	return fx.Module("api_v1",
		user.NewUserModule(),
		//NewOtherModule(),
	)
}
