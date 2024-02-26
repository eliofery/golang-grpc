package api

import (
	deniedtokenv1 "github.com/eliofery/golang-grpc/internal/app/v1app/denied_token/service"
	desc "github.com/eliofery/golang-grpc/pkg/api/denied_token/v1"
)

// api ...
type api struct {
	desc.UnimplementedDeniedTokenV1ServiceServer
	tokenService deniedtokenv1.Service
}

// New ...
func New(
	tokenService deniedtokenv1.Service,
) desc.DeniedTokenV1ServiceServer {
	return &api{
		tokenService: tokenService,
	}
}
