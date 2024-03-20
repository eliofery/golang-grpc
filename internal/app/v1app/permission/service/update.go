package service

import (
	"context"

	"github.com/eliofery/golang-grpc/internal/app/v1app/permission/dto"
	"github.com/eliofery/golang-grpc/internal/app/v1app/permission/model"
	"github.com/eliofery/golang-grpc/internal/core/server/grpc/interceptor"
)

// Update ...
func (s *service) Update(ctx context.Context, permission *dto.Update) (*model.Permission, error) {
	if !interceptor.IsAccess(ctx, updatePermission) {
		return nil, interceptor.ErrAccessDenied
	}

	return s.permissionRepository.Update(ctx, permission)
}
