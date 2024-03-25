package api

import (
	"context"

	desc "github.com/eliofery/golang-grpc/pkg/api/app/v1/user"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (a *api) DeleteUser(ctx context.Context, req *desc.DeleteUserRequest) (*emptypb.Empty, error) {
	if err := a.userService.Delete(ctx, req.Id); err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}
