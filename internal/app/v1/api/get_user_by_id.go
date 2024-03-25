package api

import (
	"context"

	"github.com/eliofery/golang-grpc/internal/app/v1/converter"
	desc "github.com/eliofery/golang-grpc/pkg/api/app/v1/user"
)

// GetUserByID ...
func (a *api) GetUserByID(ctx context.Context, req *desc.GetUserByIDRequest) (*desc.GetUserByIDResponse, error) {
	user, err := a.userService.GetByID(ctx, req.GetId())
	if err != nil {
		return nil, err
	}

	return converter.UserToGetUserByIDResponse(user), nil
}
