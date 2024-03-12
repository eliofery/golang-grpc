package repository

import (
	"context"

	"github.com/Masterminds/squirrel"
	"github.com/eliofery/golang-grpc/internal/app/v1app/permission/dto"
	"github.com/eliofery/golang-grpc/internal/app/v1app/permission/model"
	"github.com/eliofery/golang-grpc/internal/core/database/postgres"
	"github.com/eliofery/golang-grpc/internal/core/pagination"
	"github.com/eliofery/golang-grpc/pkg/eslog"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	errCreate   = status.Error(codes.Internal, "failed to create permission")
	errExists   = status.Error(codes.AlreadyExists, "permission already exists")
	errGetByID  = status.Error(codes.Internal, "failed to get permission by id")
	errNotFound = status.Error(codes.NotFound, "permission not found")
	errUpdate   = status.Error(codes.Internal, "failed to update permission")
	errDelete   = status.Error(codes.Internal, "failed to delete permission")
	errGetAll   = status.Error(codes.Internal, "failed to get permissions")
)

// Repository ...
type Repository interface {
	Create(context.Context, *dto.Permission) (int64, error)
	GetByID(context.Context, int64) (*model.Permission, error)
	Update(context.Context, *dto.Update) (*model.Permission, error)
	Delete(context.Context, int64) error
	GetAll(context.Context, *pagination.Pagination) ([]model.Permission, error)
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
