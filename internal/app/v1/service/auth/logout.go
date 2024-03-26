package auth

import (
	"context"

	"github.com/eliofery/golang-grpc/internal/core"
	desc "github.com/eliofery/golang-grpc/pkg/api/app/v1/auth"
)

// Logout ...
func (s *service) Logout(ctx context.Context, req *desc.LogoutRequest) error {
	user, err := s.accessManager.User(ctx)
	if err != nil {
		return core.ErrAccessDenied
	}

	_, err = s.txRedisManager.Committed(ctx, func(ctx context.Context) error {
		var errTx error

		count, errTx := s.authRepository.DeleteUserSession(ctx, user.ID, req.RefreshToken)
		if errTx != nil || count == 0 {
			return core.ErrTokenNotValid
		}

		sessions, errTx := s.authRepository.GetUserSessions(ctx, user.ID)
		if errTx != nil {
			return core.ErrInternal
		}

		if len(sessions) == 0 {
			if _, errTx = s.authRepository.DeleteUserSession(ctx, user.ID, "."); errTx != nil {
				return core.ErrInternal
			}

			if _, errTx = s.authRepository.DeleteUserCache(ctx, user.ID, "."); errTx != nil {
				return core.ErrInternal
			}
		}

		return nil
	})
	if err != nil {
		return err
	}

	return nil
}
