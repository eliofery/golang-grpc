package api

import (
	"context"

	desc "github.com/eliofery/golang-grpc/pkg/api/app/v1/auth"
	"google.golang.org/protobuf/types/known/emptypb"
)

// Logout ...
func (a *api) Logout(ctx context.Context, req *desc.LogoutRequest) (*emptypb.Empty, error) {
	if err := a.authService.Logout(ctx, req); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
