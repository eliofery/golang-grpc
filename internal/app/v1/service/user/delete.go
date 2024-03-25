package user

import (
	"context"

	"github.com/eliofery/golang-grpc/internal/core"
)

// Delete ...
func (s *service) Delete(ctx context.Context, reqID int64) error {
	user, err := s.accessManager.User(ctx)
	if err != nil {
		return err
	}

	if reqID != user.ID && !s.accessManager.IsAccess(ctx, deletePermission) {
		return core.ErrAccessDenied
	}

	_, err = s.txRedisManager.Committed(ctx, func(ctx context.Context) error {
		var errTx error

		if _, errTx = s.authRepository.DeleteUserSession(ctx, user.ID, "."); errTx != nil {
			return core.ErrInternal
		}

		if _, errTx = s.authRepository.DeleteUserCache(ctx, user.ID, "."); errTx != nil {
			return core.ErrInternal
		}

		if errTx = s.userRepository.Delete(ctx, reqID); errTx != nil {
			return errTx
		}

		return nil
	})
	if err != nil {
		return err
	}

	return nil
}
