package api

import (
	"context"

	"github.com/eliofery/golang-grpc/internal/core/server/grpc/interceptor"
	"google.golang.org/protobuf/types/known/emptypb"
)

// Logout ...
func (a *api) Logout(ctx context.Context, _ *emptypb.Empty) (*emptypb.Empty, error) {
	if !interceptor.IsAuthenticated(ctx) {
		return nil, interceptor.ErrNotAuthenticated
	}

	if err := a.deniedtokenService.Create(ctx, interceptor.UserToken(ctx)); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
