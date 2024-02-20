package api

import (
	"context"

	"github.com/eliofery/golang-fullstack/internal/app/v1app/user/converter"
	desc "github.com/eliofery/golang-fullstack/pkg/api/user/v1"
)

// SignUp ...
func (s *API) SignUp(ctx context.Context, req *desc.SignUpRequest) (*desc.SignUpResponse, error) {
	id, err := s.userService.SignUp(ctx, converter.ToUserInfoFromDesc(req))
	if err != nil {
		return nil, err
	}

	return &desc.SignUpResponse{Id: *id}, nil
}
