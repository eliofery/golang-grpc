package userv1

import (
	"context"

	userv1 "github.com/eliofery/golang-fullstack/internal/app/repository/user/v1"

	desc "github.com/eliofery/golang-fullstack/pkg/api/user/v1"
)

// Service ...
type Service interface {
	Get(context.Context, *desc.GetRequest) (*desc.GetResponse, error)
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
