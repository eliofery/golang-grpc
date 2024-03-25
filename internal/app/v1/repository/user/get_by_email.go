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

// GetByEmail ...
func (r *repository) GetByEmail(ctx context.Context, email string) (*model.User, error) {
	op := "app.v1.repository.user.GetByEmail"

	query, args, err := r.pgQb.
		Select(
			"users.id AS user_id",
			"first_name",
			"last_name",
			"email",
			"password",
			"created_at",
			"updated_at",
			"roles.id AS role_id",
			"roles.name AS role_name",
		).
		From("users").
		InnerJoin("roles ON users.role_id = roles.id").
		Where(squirrel.Eq{"email": email}).
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
			return nil, core.ErrWrongLoginOrPassword
		}

		return nil, core.ErrInternal
	}

	return &user, nil
}
