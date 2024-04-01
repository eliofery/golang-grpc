package interceptor

import "go.uber.org/fx"

// NewModule ...
func NewModule() fx.Option {
	return fx.Module("interceptor",
		fx.Provide(
			NewPanicRecovery,
			NewValidator,
			NewAuth,
			New,
		),
	)
}
