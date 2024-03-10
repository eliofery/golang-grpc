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

	if err := a.userService.Delete(ctx, req.GetId()); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
