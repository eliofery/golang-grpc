package api

import (
	"context"

	"github.com/eliofery/golang-grpc/internal/app/v1app/permission/converter"
	"github.com/eliofery/golang-grpc/internal/core/server/grpc/interceptor"
	desc "github.com/eliofery/golang-grpc/pkg/api/permission/v1"
)

// GetAll ...
func (a *api) GetAll(ctx context.Context, req *desc.GetAllRequest) (*desc.GetAllResponse, error) {
	if !interceptor.IsAuthenticated(ctx) {
		return nil, interceptor.ErrNotAuthenticated
	}

	users, err := a.permissionService.GetAll(ctx, req.GetPage())
	if err != nil {
		return nil, err
	}

	return converter.FromPermissionModelToGetAllResponse(users), nil
}
