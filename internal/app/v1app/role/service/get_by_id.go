package service

import (
	"context"

	"github.com/eliofery/golang-grpc/internal/app/v1app/role/model"
	"github.com/eliofery/golang-grpc/internal/core/server/grpc/interceptor"
)

// GetByID ...
func (s *service) GetByID(ctx context.Context, id int64) (*model.Role, error) {
	if !interceptor.IsAccess(ctx, readPermission) {
		return nil, interceptor.ErrAccessDenied
	}

	return s.roleRepository.GetByID(ctx, id)
}
