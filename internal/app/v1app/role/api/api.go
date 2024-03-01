package api

import (
	rolev1 "github.com/eliofery/golang-grpc/internal/app/v1app/role/service"
	desc "github.com/eliofery/golang-grpc/pkg/api/role/v1"
)

// api ...
type api struct {
	desc.UnimplementedRoleV1ServiceServer
	roleService rolev1.Service
}

// New ...
func New(
	roleService rolev1.Service,
) desc.RoleV1ServiceServer {
	return &api{
		roleService: roleService,
	}
}
