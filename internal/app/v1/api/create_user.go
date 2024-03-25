package api

import (
	"context"

	"github.com/eliofery/golang-grpc/internal/app/v1/converter"
	desc "github.com/eliofery/golang-grpc/pkg/api/app/v1/user"
)

func (a *api) CreateUser(ctx context.Context, req *desc.CreateUserRequest) (*desc.CreateUserResponse, error) {
	id, err := a.userService.Create(ctx, converter.CreateUserRequestToUser(req))
	if err != nil {
		return nil, err
	}

	return converter.IDToCreateUserRequest(id), nil
}
