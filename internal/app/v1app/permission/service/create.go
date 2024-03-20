package service

import (
	"context"

	"github.com/eliofery/golang-grpc/internal/app/v1app/permission/dto"
	"github.com/eliofery/golang-grpc/internal/core/server/grpc/interceptor"
)

// Create ...
func (s *service) Create(ctx context.Context, permission *dto.Permission) (int64, error) {
	if !interceptor.IsAccess(ctx, createPermission) {
		return 0, interceptor.ErrAccessDenied
	}

	return s.permissionRepository.Create(ctx, permission)
}
