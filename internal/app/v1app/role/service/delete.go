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

	err := s.txManager.ReadCommitted(ctx, func(ctx context.Context) error {
		var errTx error

		if errTx = s.rolePermissionRepository.Delete(ctx, id); errTx != nil {
			return errTx
		}

		if errTx = s.roleRepository.Delete(ctx, id); errTx != nil {
			return errTx
		}

		return nil
	})
	if err != nil {
		return err
	}

	return nil
}
