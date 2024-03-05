package repository

import (
	"context"

	"github.com/Masterminds/squirrel"
	"github.com/eliofery/golang-grpc/internal/app/v1app/user/dto"
	"github.com/eliofery/golang-grpc/internal/app/v1app/user/model"
	"github.com/eliofery/golang-grpc/internal/core/database/postgres"
	"github.com/eliofery/golang-grpc/internal/core/pagination"
	"github.com/eliofery/golang-grpc/pkg/eslog"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	errCreate               = status.Error(codes.Internal, "failed to create user")
	errUpdate               = status.Error(codes.Internal, "failed to update user")
	errDelete               = status.Error(codes.Internal, "failed to delete user")
	errExists               = status.Error(codes.AlreadyExists, "user already exists")
	errRoleID               = status.Error(codes.InvalidArgument, "role not exists")
	errGetByEmail           = status.Error(codes.Internal, "failed to get user by email")
	errGetByID              = status.Error(codes.Internal, "failed to get user by id")
	errGetAll               = status.Error(codes.Internal, "failed to get users")
	errWrongLoginOrPassword = status.Error(codes.NotFound, "wrong login or password")
	errNotFound             = status.Error(codes.NotFound, "user not found")
)

// Repository ...
type Repository interface {
	GetByID(context.Context, int64) (*model.User, error)
	GetByEmail(context.Context, string) (*model.User, error)
	GetAll(context.Context, *pagination.Pagination) ([]model.User, error)
	Create(context.Context, *dto.User) (int64, error)
	Update(context.Context, *dto.Update) (*model.User, error)
	Delete(context.Context, int64) error
}

type repository struct {
	db     postgres.DB
	pgQb   squirrel.StatementBuilderType
	logger *eslog.Logger
}

// New ...
func New(pg postgres.Client, logger *eslog.Logger) Repository {
	return &repository{
		db:     pg.DB(),
		pgQb:   pg.QB(),
		logger: logger,
	}
}
