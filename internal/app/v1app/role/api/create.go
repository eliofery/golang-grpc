package api

import (
	"context"

	"github.com/eliofery/golang-grpc/internal/app/v1app/role/converter"
	"github.com/eliofery/golang-grpc/internal/core/server/grpc/interceptor"
	desc "github.com/eliofery/golang-grpc/pkg/api/role/v1"
)

// Create ...
func (a *api) Create(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	if !interceptor.IsAuthenticated(ctx) {
		return nil, interceptor.ErrNotAuthenticated
	}

	id, err := a.roleService.Create(ctx, converter.FromCreateRequestToRoleDTO(req))
	if err != nil {
		return nil, err
	}

	return &desc.CreateResponse{Id: id}, nil
}
