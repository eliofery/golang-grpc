package api

import (
	user "github.com/eliofery/golang-grpc/internal/app/v1app/user/service"
	desc "github.com/eliofery/golang-grpc/pkg/api/user/v1"
)

// api ...
type api struct {
	desc.UnimplementedUserV1ServiceServer
	userService user.Service
}

// New ...
func New(
	userService user.Service,
) desc.UserV1ServiceServer {
	return &api{
		userService: userService,
	}
}
