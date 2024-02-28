package repository

import (
	"context"
	"errors"
	"log/slog"

	"github.com/Masterminds/squirrel"
	"github.com/eliofery/golang-grpc/internal/app/v1app/role/model"
	"github.com/eliofery/golang-grpc/internal/core/database/postgres"
	"github.com/jackc/pgx/v5"
)

// Delete ...
func (r *repository) Delete(ctx context.Context, id int64) error {
	op := "v1.role.repository.Delete"

	qb := r.pgQb.
		Delete(model.TableName).
		Where(squirrel.Eq{"id": id})

	query, args, err := qb.ToSql()
	if err != nil {
		r.logger.Debug(op, slog.String("err", err.Error()))
		return errDelete
	}

	q := postgres.Query{
		Name:     op,
		QueryRaw: query,
	}

	if _, err = r.db.ExecContext(ctx, q, args...); err != nil {
		r.logger.Debug(op, slog.String("err", err.Error()))

		if errors.Is(err, pgx.ErrNoRows) {
			return errNotFound
		}

		return errDelete
	}

	return nil
}
