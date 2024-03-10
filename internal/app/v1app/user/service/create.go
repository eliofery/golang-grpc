package service

import (
	"context"

	"github.com/eliofery/golang-grpc/internal/app/v1app/user/dto"
)

// Create ...
func (s *service) Create(ctx context.Context, user *dto.User) (int64, error) {
	defaultRoleID, err := s.settingRepository.GetDefaultRoleID(ctx)
	if err != nil {
		return 0, err
	}

	role, err := s.roleRepository.GetByID(ctx, defaultRoleID)
	if err != nil {
		return 0, err
	}
	user.RoleID = role.ID

	user.Password, err = s.generateFromPassword(user.Password)
	if err != nil {
		return 0, err
	}

	id, err := s.userRepository.Create(ctx, user)
	if err != nil {
		return 0, err
	}

	return id, nil
}
