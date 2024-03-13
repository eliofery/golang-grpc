package api

import (
	"context"

	"github.com/eliofery/golang-grpc/internal/app/v1app/permission/converter"
	"github.com/eliofery/golang-grpc/internal/core/server/grpc/interceptor"
	desc "github.com/eliofery/golang-grpc/pkg/api/permission/v1"
)

func (a *api) Update(ctx context.Context, req *desc.UpdateRequest) (*desc.UpdateResponse, error) {
	if !interceptor.IsAuthenticated(ctx) {
		return nil, interceptor.ErrNotAuthenticated
	}

	permission, err := a.permissionService.Update(ctx, converter.FromUpdateRequestToUpdateDTO(req))
	if err != nil {
		return nil, err
	}

	return converter.FromPermissionModelToUpdateResponse(permission), nil
}
