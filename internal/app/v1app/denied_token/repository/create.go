package repository

import (
	"context"
	"errors"
	"log/slog"

	"github.com/eliofery/golang-grpc/internal/app/v1app/denied_token/model"
	"github.com/eliofery/golang-grpc/internal/core/database/postgres"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5/pgconn"
)

// Create ...
func (r *repository) Create(ctx context.Context, token string) error {
	op := "v1.token.repository.Create"

	qb := r.pgQb.
		Insert(model.TableName).
		Columns(model.ColumnToken).
		Values(token)

	query, args, err := qb.ToSql()
	if err != nil {
		r.logger.Debug(op, slog.String("err", err.Error()))
		return errCreate
	}

	q := postgres.Query{
		Name:     op,
		QueryRaw: query,
	}

	if _, err = r.db.ExecContext(ctx, q, args...); err != nil {
		r.logger.Debug(op, slog.String("err", err.Error()))

		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code == pgerrcode.UniqueViolation {
				return errExists
			}
		}

		return errCreate
	}

	return nil
}
