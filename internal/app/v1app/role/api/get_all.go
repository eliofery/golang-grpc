package api

import (
	"context"

	"github.com/eliofery/golang-grpc/internal/app/v1app/role/converter"
	"github.com/eliofery/golang-grpc/internal/core/server/grpc/interceptor"
	desc "github.com/eliofery/golang-grpc/pkg/api/role/v1"
)

// GetAll ...
func (a *api) GetAll(ctx context.Context, req *desc.GetAllRequest) (*desc.GetAllResponse, error) {
	if !interceptor.IsAuthenticated(ctx) {
		return nil, interceptor.ErrNotAuthenticated
	}

	users, err := a.roleService.GetAll(ctx, req.GetPage())
	if err != nil {
		return nil, err
	}

	return converter.FromRolesModelToGetAllResponse(users), nil
}
