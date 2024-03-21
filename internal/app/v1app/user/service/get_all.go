package service

import (
	"context"

	"github.com/eliofery/golang-grpc/internal/app/v1app/user/model"
	"github.com/eliofery/golang-grpc/internal/core/server/grpc/interceptor"
)

// GetAll ...
func (s *service) GetAll(ctx context.Context, page uint64) ([]model.User, error) {
	if !interceptor.IsAccess(ctx, readPermission) {
		return nil, interceptor.ErrAccessDenied
	}

	return s.userRepository.GetAll(ctx, s.pagination.SetOffset(page))
}
