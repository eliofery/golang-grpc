package api

import (
	"context"

	"github.com/eliofery/golang-grpc/internal/app/v1app/role/converter"
	"github.com/eliofery/golang-grpc/internal/core/server/grpc/interceptor"
	desc "github.com/eliofery/golang-grpc/pkg/api/role/v1"
)

// GetByID ...
func (a *api) GetByID(ctx context.Context, req *desc.GetByIDRequest) (*desc.GetByIDResponse, error) {
	if !interceptor.IsAuthenticated(ctx) {
		return nil, interceptor.ErrNotAuthenticated
	}

	role, err := a.roleService.GetByID(ctx, req.GetId())
	if err != nil {
		return nil, err
	}

	return converter.FromRoleModelToGetByIDResponse(role), nil
}
