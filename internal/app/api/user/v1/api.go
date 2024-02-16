package apiuserv1

import (
	userServiceV1 "github.com/eliofery/golang-fullstack/internal/app/service/user/v1"
	desc "github.com/eliofery/golang-fullstack/pkg/api/user/v1"
)

// API ...
type API struct {
	desc.UnimplementedUserV1ServiceServer
	userService userServiceV1.Service
}

// New ...
func New(
	userService userServiceV1.Service,
) *API {
	return &API{
		userService: userService,
	}
}
