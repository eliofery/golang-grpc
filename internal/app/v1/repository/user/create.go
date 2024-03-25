package user

import (
	"context"
	"errors"
	"log/slog"

	"github.com/eliofery/golang-grpc/internal/app/v1/dto"
	"github.com/eliofery/golang-grpc/internal/core"
	"github.com/eliofery/golang-grpc/internal/core/database/postgres"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5/pgconn"
)

// Create ...
func (r *repository) Create(ctx context.Context, user *dto.User) (int64, error) {
	op := "app.v1.repository.user.Create"

	qb := r.pgQb.
		Insert("users").
		Columns("first_name", "last_name", "email", "password", "role_id").
		Values(user.FirstName, user.LastName, user.Email, user.Password, user.Role.ID).
		Suffix("RETURNING id")

	query, args, err := qb.ToSql()
	if err != nil {
		r.logger.Debug(op, slog.String("err", err.Error()))
		return 0, core.ErrInternal
	}

	q := postgres.Query{
		Name:     op,
		QueryRaw: query,
	}

	var id int64
	if err = r.db.QueryRowContext(ctx, q, args...).Scan(&id); err != nil {
		r.logger.Debug(op, slog.String("err", err.Error()))

		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code == pgerrcode.UniqueViolation {
				return 0, errExists
			}
		}

		return 0, core.ErrInternal
	}

	return id, nil
}
