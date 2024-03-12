package service

import (
	"context"

	"github.com/eliofery/golang-grpc/internal/app/v1app/permission/model"
)

// GetAll ...
func (s *service) GetAll(ctx context.Context, page uint64) ([]model.Permission, error) {
	return s.permissionRepository.GetAll(ctx, s.pagination.SetOffset(page))
}
