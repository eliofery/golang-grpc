package authorize

import (
	"github.com/eliofery/golang-grpc/internal/core/authorize/access"
	"github.com/eliofery/golang-grpc/internal/core/authorize/password"
	"github.com/eliofery/golang-grpc/internal/core/authorize/token"
	"go.uber.org/fx"
)

// NewModule ...
func NewModule() fx.Option {
	return fx.Module("authorize",
		fx.Options(
			token.NewModule(),
			password.NewModule(),
			access.NewModule(),
		),
	)
}
