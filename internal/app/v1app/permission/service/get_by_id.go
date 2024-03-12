package service

import (
	"context"

	"github.com/eliofery/golang-grpc/internal/app/v1app/permission/model"
)

func (s *service) GetByID(ctx context.Context, id int64) (*model.Permission, error) {
	return s.permissionRepository.GetByID(ctx, id)
}
