package permission

import (
	"context"
	"errors"
	"log/slog"

	"github.com/Masterminds/squirrel"
	permissionModel "github.com/eliofery/golang-grpc/internal/app/v1/model"
	"github.com/eliofery/golang-grpc/internal/core"
	"github.com/eliofery/golang-grpc/internal/core/database/postgres"
	"github.com/jackc/pgx/v5"
)

// GetByRoleID ...
func (r *repository) GetByRoleID(ctx context.Context, roleID int64) ([]permissionModel.Permission, error) {
	op := "app.v1.repository.permission.GetByRoleID"

	query, args, err := r.pgQb.
		Select(
			"permissions.id AS permission_id",
			"permissions.name AS permission_name",
			"permissions.description AS permission_description",
		).
		From("permissions").
		InnerJoin("roles_permissions ON permissions.id = roles_permissions.permission_id").
		Where(squirrel.Eq{"roles_permissions.role_id": roleID}).
		ToSql()

	if err != nil {
		r.logger.Debug(op, slog.String("err", err.Error()))
		return nil, core.ErrInternal
	}

	q := postgres.Query{
		Name:     op,
		QueryRaw: query,
	}

	var permissions []permissionModel.Permission
	if err = r.db.ScanAllContext(ctx, &permissions, q, args...); err != nil {
		r.logger.Debug(op, slog.String("err", err.Error()))

		if errors.Is(err, pgx.ErrNoRows) {
			return nil, errNotFound
		}

		return nil, core.ErrInternal
	}

	return permissions, nil
}
