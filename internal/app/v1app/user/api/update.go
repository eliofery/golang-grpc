package api

import (
	"context"

	"github.com/eliofery/golang-grpc/internal/app/v1app/user/converter"
	"github.com/eliofery/golang-grpc/internal/core/server/grpc/interceptor"
	desc "github.com/eliofery/golang-grpc/pkg/api/user/v1"
)

// Update ...
func (a *api) Update(ctx context.Context, req *desc.UpdateRequest) (*desc.UpdateResponse, error) {
	if !interceptor.IsAuthenticated(ctx) {
		return nil, interceptor.ErrNotAuthenticated
	}

	req.Id = interceptor.UserID(ctx, req.GetId())
	user, err := a.userService.Update(ctx, converter.FromUpdateRequestToUpdateDTO(req))
	if err != nil {
		return nil, err
	}

	return converter.FromUserModelToUpdateResponse(user), nil
}
