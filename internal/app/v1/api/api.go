package api

import (
	"github.com/eliofery/golang-grpc/internal/app/v1/service/auth"
	"github.com/eliofery/golang-grpc/internal/app/v1/service/user"
	desc "github.com/eliofery/golang-grpc/pkg/api/app/v1"
)

// api ...
type api struct {
	desc.UnimplementedAppServiceServer
	authService auth.Service
	userService user.Service
}

// New ...
func New(
	authService auth.Service,
	userService user.Service,
) desc.AppServiceServer {
	return &api{
		authService: authService,
		userService: userService,
	}
}
