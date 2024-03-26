package auth

import (
	"context"

	"github.com/eliofery/golang-grpc/internal/app/v1/dto"
	"github.com/eliofery/golang-grpc/internal/app/v1/model"
	"github.com/eliofery/golang-grpc/internal/core"
)

// SignUp ...
func (s *service) SignUp(ctx context.Context, user *dto.User) (*model.Token, error) {
	if s.accessManager.IsAuth(ctx) {
		return nil, core.ErrAccessDenied
	}

	var err error

	user.Role.ID, err = s.settingRepository.GetDefaultRoleByID(ctx)
	if err != nil {
		return nil, err
	}

	user.Password, err = s.passwordManager.GenerateFromPassword(user.Password)
	if err != nil {
		return nil, err
	}

	user.ID, err = s.userRepository.Create(ctx, user)
	if err != nil {
		return nil, err
	}

	user.Role.Permissions, err = s.permissionRepository.GetByRoleID(ctx, user.Role.ID)
	if err != nil {
		return nil, err
	}

	expires := s.tokenManager.RefreshTokenExpires()
	cacheUser := &model.UserCache{
		ID:          user.ID,
		Permissions: user.PermissionsNames(),
	}

	var token *model.Token
	_, err = s.txRedisManager.Committed(ctx, func(ctx context.Context) error {
		var errTx error

		if errTx = s.authRepository.CreateUserCache(ctx, cacheUser, expires); errTx != nil {
			return errTx
		}

		token, errTx = s.authRepository.CreateUserSession(ctx, user.ID, expires)
		if errTx != nil {
			return errTx
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return token, nil
}
