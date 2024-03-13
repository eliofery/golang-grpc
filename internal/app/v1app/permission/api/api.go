package api

import (
	permissionv1 "github.com/eliofery/golang-grpc/internal/app/v1app/permission/service"
	desc "github.com/eliofery/golang-grpc/pkg/api/permission/v1"
)

// api ...
type api struct {
	desc.UnimplementedPermissionV1ServiceServer
	permissionService permissionv1.Service
}

// New ...
func New(
	permissionService permissionv1.Service,
) desc.PermissionV1ServiceServer {
	return &api{
		permissionService: permissionService,
	}
}
