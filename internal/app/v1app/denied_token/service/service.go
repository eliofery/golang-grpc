package service

import (
	"context"

	tokenv1 "github.com/eliofery/golang-grpc/internal/app/v1app/denied_token/repository"
	"github.com/eliofery/golang-grpc/internal/core/jwt"
)

// Service ...
type Service interface {
	Create(context.Context, string) error
}

type service struct {
	tokenManager *jwt.TokenManager

	tokenRepository tokenv1.Repository
}

// New ...
func New(
	tokenManager *jwt.TokenManager,

	tokenRepository tokenv1.Repository,
) Service {
	return &service{
		tokenManager: tokenManager,

		tokenRepository: tokenRepository,
	}
}
