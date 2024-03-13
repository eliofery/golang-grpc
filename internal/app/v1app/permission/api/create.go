package api

import (
	"context"

	"github.com/eliofery/golang-grpc/internal/app/v1app/permission/converter"
	"github.com/eliofery/golang-grpc/internal/core/server/grpc/interceptor"
	desc "github.com/eliofery/golang-grpc/pkg/api/permission/v1"
)

func (a *api) Create(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	if !interceptor.IsAuthenticated(ctx) {
		return nil, interceptor.ErrNotAuthenticated
	}

	id, err := a.permissionService.Create(ctx, converter.FromCreateRequestToPermissionDTO(req))
	if err != nil {
		return nil, err
	}

	return &desc.CreateResponse{Id: id}, nil
}
