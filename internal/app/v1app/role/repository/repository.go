package repository

import (
	"context"

	"github.com/Masterminds/squirrel"
	"github.com/eliofery/golang-grpc/internal/app/v1app/role/dto"
	"github.com/eliofery/golang-grpc/internal/app/v1app/role/model"
	"github.com/eliofery/golang-grpc/internal/core/database/postgres"
	"github.com/eliofery/golang-grpc/internal/core/pagination"
	"github.com/eliofery/golang-grpc/pkg/eslog"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	errCreate   = status.Error(codes.Internal, "failed to create role")
	errUpdate   = status.Error(codes.Internal, "failed to update role")
	errDelete   = status.Error(codes.Internal, "failed to delete role")
	errExists   = status.Error(codes.AlreadyExists, "role already exists")
	errNotFound = status.Error(codes.NotFound, "role not found")
	errGetAll   = status.Error(codes.Internal, "failed to get roles")
	errGetByID  = status.Error(codes.Internal, "failed to get role by id")
)

// Repository ...
type Repository interface {
	Create(context.Context, *dto.Role) (int64, error)
	GetByID(context.Context, int64) (*model.Role, error)
	Update(context.Context, *dto.Role) (*model.Role, error)
	Delete(context.Context, int64) error
	GetAll(ctx context.Context, pagination *pagination.Pagination) ([]model.Role, error)
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
