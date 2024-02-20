package service

import (
	"context"

	"github.com/eliofery/golang-fullstack/internal/app/v1app/user/model"
	"golang.org/x/crypto/bcrypt"
)

// SignUp ...
func (s service) SignUp(ctx context.Context, userInfo *model.UserInfo) (*int64, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userInfo.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	userInfo.Password = string(hashedPassword)

	id, err := s.userRepository.Create(ctx, userInfo)
	if err != nil {
		return nil, err
	}

	return id, nil
}
