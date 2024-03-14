package repository

import (
	"context"

	"github.com/Masterminds/squirrel"
	"github.com/eliofery/golang-grpc/internal/app/v1app/role/dto"
	"github.com/eliofery/golang-grpc/internal/core/database/postgres"
	"github.com/eliofery/golang-grpc/pkg/eslog"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	errCreate   = status.Error(codes.Internal, "failed to create role permissions")
	errGet      = status.Error(codes.Internal, "failed to get role permissions")
	errNotFound = status.Error(codes.NotFound, "role permissions not found")
)

// Repository ...
type Repository interface {
	Create(context.Context, *dto.Role) error
	Delete(ctx context.Context, id int64) error
	GetPermissionsByRoleID(ctx context.Context, id int64) ([]string, error)
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
