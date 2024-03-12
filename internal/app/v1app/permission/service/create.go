package service

import (
	"context"

	"github.com/eliofery/golang-grpc/internal/app/v1app/permission/dto"
)

func (s *service) Create(ctx context.Context, permission *dto.Permission) (int64, error) {
	return s.permissionRepository.Create(ctx, permission)
}
