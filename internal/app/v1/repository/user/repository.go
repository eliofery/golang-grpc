package user

import (
	"context"

	"github.com/Masterminds/squirrel"
	"github.com/eliofery/golang-grpc/internal/app/v1/dto"
	"github.com/eliofery/golang-grpc/internal/app/v1/model"
	"github.com/eliofery/golang-grpc/internal/app/v1/repository/permission"
	"github.com/eliofery/golang-grpc/internal/app/v1/repository/role"
	"github.com/eliofery/golang-grpc/internal/core/database/postgres"
	"github.com/eliofery/golang-grpc/internal/core/pagination"
	"github.com/eliofery/golang-grpc/pkg/eslog"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	errExists   = status.Error(codes.AlreadyExists, "user already exists")
	errNotFound = status.Error(codes.NotFound, "user not found")
)

// Repository ...
type Repository interface {
	GetByEmail(ctx context.Context, email string) (*model.User, error)
	Create(ctx context.Context, user *dto.User) (int64, error)
	Delete(ctx context.Context, id int64) error
	GetAll(ctx context.Context, pagination pagination.Pagination) ([]model.User, error)
	GetByID(ctx context.Context, id int64) (*model.User, error)
	Update(ctx context.Context, user *dto.UserUpdate) (*model.User, error)
}

type repository struct {
	logger eslog.Logger
	db     postgres.DB
	pgQb   squirrel.StatementBuilderType

	roleRepository       role.Repository
	permissionRepository permission.Repository
}

// New ...
func New(
	logger eslog.Logger,
	pg postgres.Client,

	roleRepository role.Repository,
	permissionRepository permission.Repository,
) Repository {
	return &repository{
		logger: logger,
		db:     pg.DB(),
		pgQb:   pg.QB(),

		roleRepository:       roleRepository,
		permissionRepository: permissionRepository,
	}
}
