package service

import (
	"context"

	"github.com/eliofery/golang-grpc/internal/core/server/grpc/interceptor"
)

// Delete ...
func (s *service) Delete(ctx context.Context, reqID int64, user *interceptor.UserData) error {
	if reqID != user.ID && !interceptor.IsAccess(ctx, deletePermission) {
		return interceptor.ErrAccessDenied
	}

	return s.txManager.ReadCommitted(ctx, func(ctx context.Context) error {
		var errTx error

		if errTx = s.userRepository.Delete(ctx, reqID); errTx != nil {
			return errTx
		}

		if reqID == user.ID {
			if errTx = s.deniedTokenRepository.Create(ctx, user.Token); errTx != nil {
				return errTx
			}
		}

		return nil
	})
}
