package api

import (
	"context"

	"github.com/eliofery/golang-grpc/internal/app/v1/converter"
	desc "github.com/eliofery/golang-grpc/pkg/api/app/v1/auth"
)

// SignUp ...
func (a *api) SignUp(ctx context.Context, req *desc.SignUpRequest) (*desc.SignUpResponse, error) {
	token, err := a.authService.SignUp(ctx, converter.SignUpRequestToUser(req))
	if err != nil {
		return nil, err
	}

	return converter.TokenToSignUpResponse(token), nil
}
