package api

import (
	deniedtokenv1 "github.com/eliofery/golang-grpc/internal/app/v1app/denied_token/service"
	userv1 "github.com/eliofery/golang-grpc/internal/app/v1app/user/service"
	desc "github.com/eliofery/golang-grpc/pkg/api/user/v1"
)

// api ...
type api struct {
	desc.UnimplementedUserV1ServiceServer
	userService        userv1.Service
	deniedtokenService deniedtokenv1.Service
}

// New ...
func New(
	userService userv1.Service,
	deniedtokenService deniedtokenv1.Service,
) desc.UserV1ServiceServer {
	return &api{
		userService:        userService,
		deniedtokenService: deniedtokenService,
	}
}
