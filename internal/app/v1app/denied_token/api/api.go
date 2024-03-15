package api

import (
	deniedToken "github.com/eliofery/golang-grpc/internal/app/v1app/denied_token/service"
	desc "github.com/eliofery/golang-grpc/pkg/api/denied_token/v1"
)

// api ...
type api struct {
	desc.UnimplementedDeniedTokenV1ServiceServer
	deniedTokenService deniedToken.Service
}

// New ...
func New(
	deniedTokenService deniedToken.Service,
) desc.DeniedTokenV1ServiceServer {
	return &api{
		deniedTokenService: deniedTokenService,
	}
}
