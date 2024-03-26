package auth

import (
	"context"

	"github.com/eliofery/golang-grpc/internal/app/v1/dto"
	"github.com/eliofery/golang-grpc/internal/app/v1/model"
	"github.com/eliofery/golang-grpc/internal/core"
)

// SignIn ...
func (s *service) SignIn(ctx context.Context, user *dto.User) (*model.Token, error) {
	if s.accessManager.IsAuth(ctx) {
		return nil, core.ErrAccessDenied
	}

	findUser, err := s.userRepository.GetByEmail(ctx, user.Email)
	if err != nil {
		return nil, err
	}

	if err = s.passwordManager.CompareHashAndPassword(findUser.Password, user.Password); err != nil {
		return nil, err
	}

	user.Role.Permissions, err = s.permissionRepository.GetByRoleID(ctx, findUser.Role.ID)
	if err != nil {
		return nil, err
	}

	expires := s.tokenManager.RefreshTokenExpires()
	cacheUser := &model.UserCache{
		ID:          findUser.ID,
		Permissions: user.PermissionsNames(),
	}

	var token *model.Token
	_, err = s.txRedisManager.Committed(ctx, func(ctx context.Context) error {
		var errTx error

		if errTx = s.authRepository.CreateUserCache(ctx, cacheUser, expires); errTx != nil {
			return errTx
		}

		token, errTx = s.authRepository.CreateUserSession(ctx, findUser.ID, expires)
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
