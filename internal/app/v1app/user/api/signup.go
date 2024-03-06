package api

import (
	"context"

	"github.com/eliofery/golang-grpc/internal/app/v1app/user/converter"
	"github.com/eliofery/golang-grpc/internal/core/server/grpc/interceptor"
	desc "github.com/eliofery/golang-grpc/pkg/api/user/v1"
)

// SignUp ...
func (a *api) SignUp(ctx context.Context, req *desc.SignUpRequest) (*desc.SignUpResponse, error) {
	if interceptor.IsAuthenticated(ctx) {
		return nil, interceptor.ErrAlreadyAuthenticated
	}

	id, err := a.userService.Create(ctx, converter.FromSignUpRequestToUserDTO(req))
	if err != nil {
		return nil, err
	}

	if err = a.userService.SignUp(ctx, id); err != nil {
		return nil, err
	}

	return &desc.SignUpResponse{Id: id}, nil
}
