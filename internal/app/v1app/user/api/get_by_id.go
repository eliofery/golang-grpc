package api

import (
	"context"

	"github.com/eliofery/golang-grpc/internal/app/v1app/user/converter"
	"github.com/eliofery/golang-grpc/internal/core/server/grpc/interceptor"
	desc "github.com/eliofery/golang-grpc/pkg/api/user/v1"
)

// GetByID ...
func (a *api) GetByID(ctx context.Context, req *desc.GetByIDRequest) (*desc.GetByIDResponse, error) {
	if !interceptor.IsAuthenticated(ctx) {
		return nil, interceptor.ErrNotAuthenticated
	}

	user, err := a.userService.GetByID(ctx, req.GetId(), interceptor.UserID(ctx))
	if err != nil {
		return nil, err
	}

	return converter.FromUserModelToGetByIDResponse(user), nil
}
