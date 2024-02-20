package service

import (
	"context"

	"github.com/eliofery/golang-fullstack/internal/app/v1app/user/model"
	userv1 "github.com/eliofery/golang-fullstack/internal/app/v1app/user/repository"
)

// Service ...
type Service interface {
	SignUp(context.Context, *model.UserInfo) (*int64, error)
}

type service struct {
	userRepository userv1.Repository
}

// New ...
func New(userRepository userv1.Repository) Service {
	return &service{
		userRepository: userRepository,
	}
}
