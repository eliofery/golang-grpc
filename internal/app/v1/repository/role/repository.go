package role

import (
	"context"

	"github.com/Masterminds/squirrel"
	"github.com/eliofery/golang-grpc/internal/app/v1/model"
	"github.com/eliofery/golang-grpc/internal/app/v1/repository/permission"
	"github.com/eliofery/golang-grpc/internal/core/database/postgres"
	"github.com/eliofery/golang-grpc/pkg/eslog"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	errNotFound = status.Error(codes.NotFound, "role not found")
)

// Repository ...
type Repository interface {
	GetByID(ctx context.Context, roleID int64) (*model.Role, error)
}

type repository struct {
	db     postgres.DB
	pgQb   squirrel.StatementBuilderType
	logger eslog.Logger

	permissionRepository permission.Repository
}

// New ...
func New(
	pg postgres.Client,
	logger eslog.Logger,

	permissionRepository permission.Repository,
) Repository {
	return &repository{
		db:     pg.DB(),
		pgQb:   pg.QB(),
		logger: logger,

		permissionRepository: permissionRepository,
	}
}
