package api

import (
	"context"

	"github.com/eliofery/golang-grpc/internal/app/v1/converter"
	desc "github.com/eliofery/golang-grpc/pkg/api/app/v1/user"
)

func (a *api) UpdateUser(ctx context.Context, req *desc.UpdateUserRequest) (*desc.UpdateUserResponse, error) {
	user, err := a.userService.Update(ctx, converter.UpdateUserRequestToUser(req))
	if err != nil {
		return nil, err
	}
	_ = user

	return converter.UserToUpdateUserResponse(user), nil
}
