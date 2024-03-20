package service

import (
	"context"

	deniedToken "github.com/eliofery/golang-grpc/internal/app/v1app/denied_token/repository"
	"github.com/eliofery/golang-grpc/internal/core/jwt"
)

const (
	createPermission = "create_denied_tokens"
	//readPermission   = "read_denied_tokens"
	//updatePermission = "update_denied_tokens"
	//deletePermission = "delete_denied_tokens"
)

// Service ...
type Service interface {
	Create(context.Context, string) error
}

type service struct {
	tokenManager *jwt.TokenManager

	deniedTokenRepository deniedToken.Repository
}

// New ...
func New(
	tokenManager *jwt.TokenManager,

	tokenRepository deniedToken.Repository,
) Service {
	return &service{
		tokenManager: tokenManager,

		deniedTokenRepository: tokenRepository,
	}
}
