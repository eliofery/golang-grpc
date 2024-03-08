package api

import (
	"context"

	"github.com/eliofery/golang-grpc/internal/app/v1app/user/converter"
	"github.com/eliofery/golang-grpc/internal/core/server/grpc/interceptor"
	desc "github.com/eliofery/golang-grpc/pkg/api/user/v1"
)

// GetAll ...
func (a *api) GetAll(ctx context.Context, req *desc.GetAllRequest) (*desc.GetAllResponse, error) {
	if !interceptor.IsAuthenticated(ctx) {
		return nil, interceptor.ErrNotAuthenticated
	}

	users, err := a.userService.GetAll(ctx, req.GetPage())
	if err != nil {
		return nil, err
	}

	return converter.FromUsersModelToGetAllResponse(users), nil
}
