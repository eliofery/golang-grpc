package service

import (
	"context"

	"github.com/eliofery/golang-grpc/internal/app/v1app/permission/dto"
	"github.com/eliofery/golang-grpc/internal/app/v1app/permission/model"
)

func (s *service) Update(ctx context.Context, permission *dto.Update) (*model.Permission, error) {
	return s.permissionRepository.Update(ctx, permission)
}
