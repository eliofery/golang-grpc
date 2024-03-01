package api

import (
	"context"

	"github.com/eliofery/golang-grpc/internal/app/v1app/role/converter"
	"github.com/eliofery/golang-grpc/internal/core/server/grpc/interceptor"
	desc "github.com/eliofery/golang-grpc/pkg/api/role/v1"
)

// Update ...
func (a *api) Update(ctx context.Context, req *desc.UpdateRequest) (*desc.UpdateResponse, error) {
	if !interceptor.IsAuthenticated(ctx) {
		return nil, interceptor.ErrNotAuthenticated
	}

	role, err := a.roleService.Update(ctx, converter.FromUpdateRequestToRoleDTO(req))
	if err != nil {
		return nil, err
	}

	return converter.FromRoleModelToUpdateResponse(role), nil
}
