package api

import (
	"context"

	"github.com/eliofery/golang-grpc/internal/core/server/grpc/interceptor"
	desc "github.com/eliofery/golang-grpc/pkg/api/denied_token/v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

// Create ...
func (a *api) Create(ctx context.Context, req *desc.CreateRequest) (*emptypb.Empty, error) {
	if !interceptor.IsAuthenticated(ctx) {
		return nil, interceptor.ErrNotAuthenticated
	}

	if err := a.tokenService.Create(ctx, req.GetToken()); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
