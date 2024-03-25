package api

import (
	"context"

	"github.com/eliofery/golang-grpc/internal/app/v1/converter"
	desc "github.com/eliofery/golang-grpc/pkg/api/app/v1/user"
)

func (a *api) GetUsers(ctx context.Context, req *desc.GetUsersRequest) (*desc.GetUsersResponse, error) {
	users, err := a.userService.GetAll(ctx, req.GetPage())
	if err != nil {
		return nil, err
	}

	return converter.UsersToGetAllResponse(users), nil
}
