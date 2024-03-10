package service

import (
	"context"
	"log/slog"

	rolev1 "github.com/eliofery/golang-grpc/internal/app/v1app/role/repository"
	settingv1 "github.com/eliofery/golang-grpc/internal/app/v1app/setting/repository"
	"github.com/eliofery/golang-grpc/internal/app/v1app/user/dto"
	"github.com/eliofery/golang-grpc/internal/app/v1app/user/model"
	userv1 "github.com/eliofery/golang-grpc/internal/app/v1app/user/repository"
	"github.com/eliofery/golang-grpc/internal/core/database/postgres"
	"github.com/eliofery/golang-grpc/internal/core/jwt"
	"github.com/eliofery/golang-grpc/internal/core/pagination"
	"github.com/eliofery/golang-grpc/pkg/eslog"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	errWrongAuth        = status.Error(codes.PermissionDenied, "wrong login or password")
	errWrongOldPassword = status.Error(codes.InvalidArgument, "wrong old password")
	errPasswordLong     = status.Error(codes.PermissionDenied, "password is too long")
)

// Service ...
type Service interface {
	SignUp(context.Context, int64) error
	SignIn(context.Context, *dto.User) error
	GetByID(context.Context, int64) (*model.User, error)
	GetAll(context.Context, uint64) ([]model.User, error)
	Create(context.Context, *dto.User) (int64, error)
	Update(context.Context, *dto.Update) (*model.User, error)
	Delete(context.Context, int64) error
}

type service struct {
	tokenManager *jwt.TokenManager
	logger       *eslog.Logger
	txManager    postgres.TxManager
	pagination   *pagination.Pagination

	settingRepository settingv1.Repository
	roleRepository    rolev1.Repository
	userRepository    userv1.Repository
}

// New ...
func New(
	tokenManager *jwt.TokenManager,
	logger *eslog.Logger,
	txManager postgres.TxManager,
	pagination *pagination.Pagination,

	settingRepository settingv1.Repository,
	roleRepository rolev1.Repository,
	userRepository userv1.Repository,
) Service {
	return &service{
		tokenManager: tokenManager,
		logger:       logger,
		txManager:    txManager,
		pagination:   pagination,

		settingRepository: settingRepository,
		userRepository:    userRepository,
		roleRepository:    roleRepository,
	}
}

// generateFromPassword ...
func (s *service) generateFromPassword(password string) (string, error) {
	op := "v1.user.service.generateFromPassword"

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		s.logger.Debug(op, slog.String("err", err.Error()))
		return "", errPasswordLong
	}

	return string(hashedPassword), nil
}

// compareHashAndPassword ...
func (s *service) compareHashAndPassword(hashedPassword, password string) error {
	op := "v1.user.service.compareHashAndPassword"

	if err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)); err != nil {
		s.logger.Debug(op, slog.String("err", err.Error()))
		return errWrongAuth
	}

	return nil
}
