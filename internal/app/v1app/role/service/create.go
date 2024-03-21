package service

import (
	"context"

	"github.com/eliofery/golang-grpc/internal/app/v1app/role/dto"
	"github.com/eliofery/golang-grpc/internal/core/server/grpc/interceptor"
)

// Create ...
func (s *service) Create(ctx context.Context, role *dto.Role) (int64, error) {
	if !interceptor.IsAccess(ctx, createPermission) {
		return 0, interceptor.ErrAccessDenied
	}

	err := s.txManager.ReadCommitted(ctx, func(ctx context.Context) error {
		var errTx error

		role.ID, errTx = s.roleRepository.Create(ctx, role.Name)
		if errTx != nil {
			return errTx
		}

		if errTx = s.rolePermissionRepository.Create(ctx, role); errTx != nil {
			return errTx
		}

		return nil
	})
	if err != nil {
		return 0, err
	}

	return role.ID, nil
}
