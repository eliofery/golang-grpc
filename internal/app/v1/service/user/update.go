package user

import (
	"context"

	"github.com/eliofery/golang-grpc/internal/app/v1/dto"
	"github.com/eliofery/golang-grpc/internal/app/v1/model"
	"github.com/eliofery/golang-grpc/internal/core"
)

// Update ...
func (s *service) Update(ctx context.Context, reqUser *dto.UserUpdate) (*model.User, error) {
	user, err := s.accessManager.User(ctx)
	if err != nil {
		return nil, core.ErrAccessDenied
	}

	if reqUser.ID != user.ID && !s.accessManager.IsAccess(ctx, updatePermission) {
		return nil, core.ErrAccessDenied
	}

	findUser, err := s.userRepository.GetByID(ctx, reqUser.ID)
	if err != nil {
		return nil, err
	}

	if reqUser.OldPassword.Valid && reqUser.NewPassword.Valid {
		err = s.passwordManager.CompareHashAndPassword(findUser.Password, reqUser.OldPassword.String)
		if err != nil {
			return nil, errPasswordsDoNotMatch
		}

		reqUser.NewPassword.String, err = s.passwordManager.GenerateFromPassword(reqUser.NewPassword.String)
		if err != nil {
			return nil, err
		}
	}

	updateUser, err := s.userRepository.Update(ctx, reqUser)
	if err != nil {
		return nil, err
	}

	var userDTO dto.User
	userDTO.Role = updateUser.Role
	expires := s.tokenManager.RefreshTokenExpires()
	cacheUser := &model.UserCache{
		ID:          updateUser.ID,
		Permissions: userDTO.PermissionsNames(),
	}
	if err = s.authRepository.CreateUserCache(ctx, cacheUser, expires); err != nil {
		return nil, err
	}

	return updateUser, nil
}
