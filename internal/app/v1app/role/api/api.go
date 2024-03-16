package api

import (
	role "github.com/eliofery/golang-grpc/internal/app/v1app/role/service"
	desc "github.com/eliofery/golang-grpc/pkg/api/role/v1"
)

// api ...
type api struct {
	desc.UnimplementedRoleV1ServiceServer
	roleService role.Service
}

// New ...
func New(
	roleService role.Service,
) desc.RoleV1ServiceServer {
	return &api{
		roleService: roleService,
	}
}
