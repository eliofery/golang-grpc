package auth

import (
	"context"

	user2 "github.com/eliofery/golang-grpc/internal/app/v1/dto"
	auth2 "github.com/eliofery/golang-grpc/internal/app/v1/model"
	"github.com/eliofery/golang-grpc/internal/app/v1/repository/auth"
	"github.com/eliofery/golang-grpc/internal/app/v1/repository/permission"
	"github.com/eliofery/golang-grpc/internal/app/v1/repository/role"
	"github.com/eliofery/golang-grpc/internal/app/v1/repository/setting"
	"github.com/eliofery/golang-grpc/internal/app/v1/repository/user"
	"github.com/eliofery/golang-grpc/internal/core/authorize/access"
	"github.com/eliofery/golang-grpc/internal/core/authorize/password"
	"github.com/eliofery/golang-grpc/internal/core/authorize/token"
	"github.com/eliofery/golang-grpc/internal/core/database/postgres"
	"github.com/eliofery/golang-grpc/internal/core/database/redis"
	desc "github.com/eliofery/golang-grpc/pkg/api/app/v1/auth"
	"github.com/eliofery/golang-grpc/pkg/eslog"
)

//const (
//	createAccess = "create_users"
//	readAccess   = "read_users"
//	updateAccess = "update_users"
//	deleteAccess = "delete_users"
//)

// Service ...
type Service interface {
	SignUp(ctx context.Context, user *user2.User) (*auth2.Token, error)
	SignIn(ctx context.Context, user *user2.User) (*auth2.Token, error)
	Logout(ctx context.Context, req *desc.LogoutRequest) error
}

type service struct {
	logger          eslog.Logger
	txPGManager     postgres.TxManager
	txRedisManager  redis.TxManager
	tokenManager    token.Manager
	passwordManager password.Manager
	accessManager   access.Manager
	redis           redis.DB

	authRepository       auth.Repository
	settingRepository    setting.Repository
	userRepository       user.Repository
	roleRepository       role.Repository
	permissionRepository permission.Repository
}

// New ...
func New(
	logger eslog.Logger,
	txPGManager postgres.TxManager,
	txRedisManager redis.TxManager,
	tokenManager token.Manager,
	passwordManager password.Manager,
	accessManager access.Manager,
	redis redis.Client,

	settingRepository setting.Repository,
	authRepository auth.Repository,
	userRepository user.Repository,
	roleRepository role.Repository,
	permissionRepository permission.Repository,
) Service {
	return &service{
		logger:          logger,
		txPGManager:     txPGManager,
		txRedisManager:  txRedisManager,
		tokenManager:    tokenManager,
		passwordManager: passwordManager,
		accessManager:   accessManager,
		redis:           redis.DB(),

		settingRepository:    settingRepository,
		authRepository:       authRepository,
		userRepository:       userRepository,
		roleRepository:       roleRepository,
		permissionRepository: permissionRepository,
	}
}
