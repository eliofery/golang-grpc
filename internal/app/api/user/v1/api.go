package userv1

import (
	userv1 "github.com/eliofery/golang-fullstack/internal/app/service/user/v1"
	desc "github.com/eliofery/golang-fullstack/pkg/api/user/v1"
)

// API ...
type API struct {
	desc.UnimplementedUserV1ServiceServer
	userService userv1.Service
}

// New ...
func New(userService userv1.Service) desc.UserV1ServiceServer {
	return &API{
		userService: userService,
	}
}
