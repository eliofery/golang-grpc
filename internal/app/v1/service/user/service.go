package user

import (
	"context"

	"github.com/eliofery/golang-grpc/internal/app/v1/repository/role"
	"github.com/eliofery/golang-grpc/internal/app/v1/repository/setting"
	"github.com/eliofery/golang-grpc/internal/app/v1/repository/user"
	"github.com/eliofery/golang-grpc/internal/core/authorize/token"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/eliofery/golang-grpc/internal/app/v1/dto"
	"github.com/eliofery/golang-grpc/internal/app/v1/model"
	"github.com/eliofery/golang-grpc/internal/app/v1/repository/auth"
	"github.com/eliofery/golang-grpc/internal/core/authorize/access"
	"github.com/eliofery/golang-grpc/internal/core/authorize/password"
	"github.com/eliofery/golang-grpc/internal/core/database/postgres"
	"github.com/eliofery/golang-grpc/internal/core/database/redis"
	"github.com/eliofery/golang-grpc/internal/core/pagination"
	"github.com/eliofery/golang-grpc/pkg/eslog"
)

const (
	createPermission = "create_users"
	readPermission   = "read_users"
	updatePermission = "update_users"
	deletePermission = "delete_users"
)

var (
	errPasswordsDoNotMatch = status.Error(codes.InvalidArgument, "passwords do not match")
)

// Service ...
type Service interface {
	Create(ctx context.Context, user *dto.User) (int64, error)
	Delete(ctx context.Context, reqID int64) error
	GetAll(ctx context.Context, page uint64) ([]model.User, error)
	GetByID(ctx context.Context, reqID int64) (*model.User, error)
	Update(ctx context.Context, reqUser *dto.UserUpdate) (*model.User, error)
}

type service struct {
	logger          eslog.Logger
	txPGManager     postgres.TxManager
	txRedisManager  redis.TxManager
	passwordManager password.Manager
	accessManager   access.Manager
	redis           redis.DB
	pagination      pagination.Pagination
	tokenManager    token.Manager

	userRepository    user.Repository
	settingRepository setting.Repository
	authRepository    auth.Repository
	roleRepository    role.Repository
}

// New ...
func New(
	logger eslog.Logger,
	txPGManager postgres.TxManager,
	txRedisManager redis.TxManager,
	passwordManager password.Manager,
	accessManager access.Manager,
	redis redis.Client,
	pagination pagination.Pagination,
	tokenManager token.Manager,

	userRepository user.Repository,
	settingRepository setting.Repository,
	authRepository auth.Repository,
	roleRepository role.Repository,
) Service {
	return &service{
		logger:          logger,
		txPGManager:     txPGManager,
		txRedisManager:  txRedisManager,
		passwordManager: passwordManager,
		accessManager:   accessManager,
		redis:           redis.DB(),
		pagination:      pagination,
		tokenManager:    tokenManager,

		userRepository:    userRepository,
		settingRepository: settingRepository,
		authRepository:    authRepository,
		roleRepository:    roleRepository,
	}
}
