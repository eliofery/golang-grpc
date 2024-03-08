package api

import (
	"context"

	"github.com/eliofery/golang-grpc/internal/app/v1app/user/converter"
	"github.com/eliofery/golang-grpc/internal/core/server/grpc/interceptor"
	desc "github.com/eliofery/golang-grpc/pkg/api/user/v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

// SignIn ...
func (a *api) SignIn(ctx context.Context, req *desc.SignInRequest) (*emptypb.Empty, error) {
	if interceptor.IsAuthenticated(ctx) {
		return nil, interceptor.ErrAlreadyAuthenticated
	}

	if err := a.userService.SignIn(ctx, converter.FromSignInRequestToUserDTO(req)); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
