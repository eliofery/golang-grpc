package v1api

import (
	desc "github.com/eliofery/golang-fullstack/pkg/api/user/v1"
	"go.uber.org/fx"
)

// RegisterServiceServers ...
var RegisterServiceServers = []any{
	desc.RegisterUserV1ServiceServer,
}

// RegisterServiceHandlerFromEndpoints ...
var RegisterServiceHandlerFromEndpoints = []any{
	desc.RegisterUserV1ServiceHandlerFromEndpoint,
}

// NewModule ...
func NewModule() fx.Option {
	return fx.Module("api_v1",
		NewUserModule(),
		//NewOtherModule(),
	)
}
