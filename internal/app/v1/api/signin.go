package api

import (
	"context"

	"github.com/eliofery/golang-grpc/internal/app/v1/converter"
	desc "github.com/eliofery/golang-grpc/pkg/api/app/v1/auth"
)

// SignIn ...
func (a *api) SignIn(ctx context.Context, req *desc.SignInRequest) (*desc.SignInResponse, error) {
	token, err := a.authService.SignIn(ctx, converter.SignInRequestToUser(req))
	if err != nil {
		return nil, err
	}

	return converter.TokenToSignInResponse(token), nil
}
