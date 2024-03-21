package service

import (
	"context"

	"github.com/eliofery/golang-grpc/internal/app/v1app/user/dto"
)

// SignUp ...
func (s *service) SignUp(ctx context.Context, user *dto.User) (int64, error) {
	var err error

	user.RoleID, err = s.settingRepository.GetDefaultRoleID(ctx)
	if err != nil {
		return 0, err
	}

	user.Password, err = s.generateFromPassword(user.Password)
	if err != nil {
		return 0, err
	}

	user.ID, err = s.userRepository.Create(ctx, user)
	if err != nil {
		return 0, err
	}

	token, err := s.tokenManager.Generate(user.ID, user.RoleID)
	if err != nil {
		return 0, err
	}

	if err = s.tokenManager.SendAuthHeader(ctx, token); err != nil {
		return 0, err
	}

	return user.ID, nil
}
