package user

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
func (r *repository) GetByID(ctx context.Context, id int64) (*model.User, error) {
	op := "app.v1.repository.user.GetByID"

	query, args, err := r.pgQb.
		Select(
			"users.id AS user_id",
			"first_name",
			"last_name",
			"email",
			"password",
			"roles.id AS role_id",
			"roles.name AS role_name",
			"created_at",
			"updated_at",
		).
		From("users").
		InnerJoin("roles ON roles.id = users.role_id").
		Where(squirrel.Eq{"users.id": id}).
		ToSql()

	if err != nil {
		r.logger.Debug(op, slog.String("err", err.Error()))
		return nil, core.ErrInternal
	}

	q := postgres.Query{
		Name:     op,
		QueryRaw: query,
	}

	var user model.User
	if err = r.db.ScanOneContext(ctx, &user, q, args...); err != nil {
		r.logger.Debug(op, slog.String("err", err.Error()))

		if errors.Is(err, pgx.ErrNoRows) {
			return nil, errNotFound
		}

		return nil, core.ErrInternal
	}

	user.Role.Permissions, err = r.permissionRepository.GetByRoleID(ctx, user.Role.ID)
	if err != nil {
		r.logger.Debug(op, slog.String("err", err.Error()))
		return nil, core.ErrInternal
	}

	return &user, nil
}
