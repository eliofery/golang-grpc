package service

import (
	"context"

	"github.com/eliofery/golang-grpc/internal/core/server/grpc/interceptor"
)

// Delete ...
func (s *service) Delete(ctx context.Context, id int64) error {
	if !interceptor.IsAccess(ctx, deletePermission) {
		return interceptor.ErrAccessDenied
	}

	return s.permissionRepository.Delete(ctx, id)
}
