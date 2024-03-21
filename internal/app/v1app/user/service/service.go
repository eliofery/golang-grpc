package service

import (
	"context"
	"log/slog"

	deniedToken "github.com/eliofery/golang-grpc/internal/app/v1app/denied_token/repository"
	role "github.com/eliofery/golang-grpc/internal/app/v1app/role/repository"
	rolePermission "github.com/eliofery/golang-grpc/internal/app/v1app/role_permission/repository"
	setting "github.com/eliofery/golang-grpc/internal/app/v1app/setting/repository"
	"github.com/eliofery/golang-grpc/internal/app/v1app/user/dto"
	"github.com/eliofery/golang-grpc/internal/app/v1app/user/model"
	user "github.com/eliofery/golang-grpc/internal/app/v1app/user/repository"
	"github.com/eliofery/golang-grpc/internal/core/database/postgres"
	"github.com/eliofery/golang-grpc/internal/core/jwt"
	"github.com/eliofery/golang-grpc/internal/core/pagination"
	"github.com/eliofery/golang-grpc/internal/core/server/grpc/interceptor"
	"github.com/eliofery/golang-grpc/pkg/eslog"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	errWrongAuth        = status.Error(codes.PermissionDenied, "wrong login or password")
	errWrongOldPassword = status.Error(codes.InvalidArgument, "wrong old password")
	errPasswordLong     = status.Error(codes.InvalidArgument, "password is too long")
)

const (
	createPermission = "create_users"
	readPermission   = "read_users"
	updatePermission = "update_users"
	deletePermission = "delete_users"
)

// Service ...
type Service interface {
	SignUp(context.Context, *dto.User) (int64, error)
	SignIn(context.Context, *dto.User) error
	Logout(context.Context, string) error
	GetByID(context.Context, int64, int64) (*model.User, error)
	GetAll(context.Context, uint64) ([]model.User, error)
	Create(context.Context, *dto.User) (int64, error)
	Update(context.Context, *dto.Update, int64) (*model.User, error)
	Delete(context.Context, int64, *interceptor.UserData) error
}

type service struct {
	tokenManager *jwt.TokenManager
	logger       *eslog.Logger
	txManager    postgres.TxManager
	pagination   *pagination.Pagination

	settingRepository        setting.Repository
	deniedTokenRepository    deniedToken.Repository
	roleRepository           role.Repository
	rolePermissionRepository rolePermission.Repository
	userRepository           user.Repository
}

// New ...
func New(
	tokenManager *jwt.TokenManager,
	logger *eslog.Logger,
	txManager postgres.TxManager,
	pagination *pagination.Pagination,

	settingRepository setting.Repository,
	deniedTokenRepository deniedToken.Repository,
	roleRepository role.Repository,
	rolePermissionRepository rolePermission.Repository,
	userRepository user.Repository,
) Service {
	return &service{
		tokenManager: tokenManager,
		logger:       logger,
		txManager:    txManager,
		pagination:   pagination,

		settingRepository:        settingRepository,
		deniedTokenRepository:    deniedTokenRepository,
		roleRepository:           roleRepository,
		rolePermissionRepository: rolePermissionRepository,
		userRepository:           userRepository,
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
