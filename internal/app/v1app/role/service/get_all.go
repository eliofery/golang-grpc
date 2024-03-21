package service

import (
	"context"

	"github.com/eliofery/golang-grpc/internal/app/v1app/role/model"
	"github.com/eliofery/golang-grpc/internal/core/server/grpc/interceptor"
)

// GetAll ...
func (s *service) GetAll(ctx context.Context, page uint64) ([]model.Role, error) {
	if !interceptor.IsAccess(ctx, readPermission) {
		return nil, interceptor.ErrAccessDenied
	}

	return s.roleRepository.GetAll(ctx, s.pagination.SetOffset(page))
}
