package api

import (
	"context"

	"github.com/eliofery/golang-grpc/internal/core/server/grpc/interceptor"
	desc "github.com/eliofery/golang-grpc/pkg/api/user/v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

// Delete ...
func (a *api) Delete(ctx context.Context, req *desc.DeleteRequest) (*emptypb.Empty, error) {
	if !interceptor.IsAuthenticated(ctx) {
		return nil, interceptor.ErrNotAuthenticated
	}

	id := interceptor.UserID(ctx, req.GetId())
	if err := a.userService.Delete(ctx, id); err != nil {
		return nil, err
	}

	if err := a.deniedtokenService.Create(ctx, interceptor.UserToken(ctx)); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
