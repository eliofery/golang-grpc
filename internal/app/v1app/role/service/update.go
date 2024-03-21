package service

import (
	"context"

	"github.com/eliofery/golang-grpc/internal/app/v1app/role/dto"
	"github.com/eliofery/golang-grpc/internal/app/v1app/role/model"
	"github.com/eliofery/golang-grpc/internal/core/server/grpc/interceptor"
)

// Update ...
func (s *service) Update(ctx context.Context, role *dto.Role) (*model.Role, error) {
	if !interceptor.IsAccess(ctx, updatePermission) {
		return nil, interceptor.ErrAccessDenied
	}

	var roleUpdate *model.Role
	err := s.txManager.ReadCommitted(ctx, func(ctx context.Context) error {
		var errTx error

		roleUpdate, errTx = s.roleRepository.Update(ctx, role)
		if errTx != nil {
			return errTx
		}

		if errTx = s.rolePermissionRepository.Delete(ctx, role.ID); errTx != nil {
			return errTx
		}

		if errTx = s.rolePermissionRepository.Create(ctx, role); errTx != nil {
			return errTx
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return roleUpdate, nil
}
