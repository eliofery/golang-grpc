package service

import (
	"context"

	"github.com/eliofery/golang-grpc/internal/app/v1app/permission/model"
	"github.com/eliofery/golang-grpc/internal/core/server/grpc/interceptor"
)

// GetByID ...
func (s *service) GetByID(ctx context.Context, id int64) (*model.Permission, error) {
	if !interceptor.IsAccess(ctx, readPermission) {
		return nil, interceptor.ErrAccessDenied
	}

	return s.permissionRepository.GetByID(ctx, id)
}
