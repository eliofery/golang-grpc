package role

import (
	"context"
	"errors"
	"log/slog"

	"github.com/Masterminds/squirrel"
	"github.com/eliofery/golang-grpc/internal/app/v1/model"
	"github.com/eliofery/golang-grpc/internal/core"
	"github.com/eliofery/golang-grpc/internal/core/database/postgres"
	"github.com/jackc/pgx/v5"
)

// GetByID ...
func (r *repository) GetByID(ctx context.Context, roleID int64) (*model.Role, error) {
	op := "app.v1.repository.role.GetByID"

	query, args, err := r.pgQb.
		Select(
			"roles.id AS role_id",
			"roles.name AS role_name",
		).
		From("roles").
		Where(squirrel.Eq{"id": roleID}).
		ToSql()

	if err != nil {
		r.logger.Debug(op, slog.String("err", err.Error()))
		return nil, core.ErrInternal
	}

	q := postgres.Query{
		Name:     op,
		QueryRaw: query,
	}

	var role model.Role
	if err = r.db.ScanOneContext(ctx, &role, q, args...); err != nil {
		r.logger.Debug(op, slog.String("err", err.Error()))

		if errors.Is(err, pgx.ErrNoRows) {
			return nil, errNotFound
		}

		return nil, core.ErrInternal
	}

	if role.Permissions, err = r.permissionRepository.GetByRoleID(ctx, roleID); err != nil {
		r.logger.Debug(op, slog.String("err", err.Error()))
		role.Permissions = nil
	}

	return &role, nil
}
