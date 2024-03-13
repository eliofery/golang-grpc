package api

import (
	"context"

	"github.com/eliofery/golang-grpc/internal/core/server/grpc/interceptor"
	v1 "github.com/eliofery/golang-grpc/pkg/api/permission/v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (a *api) Delete(ctx context.Context, req *v1.DeleteRequest) (*emptypb.Empty, error) {
	if !interceptor.IsAuthenticated(ctx) {
		return nil, interceptor.ErrNotAuthenticated
	}

	if err := a.permissionService.Delete(ctx, req.GetId()); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
