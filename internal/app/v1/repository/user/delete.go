package user

import (
	"context"
	"log/slog"

	"github.com/Masterminds/squirrel"
	"github.com/eliofery/golang-grpc/internal/core"
	"github.com/eliofery/golang-grpc/internal/core/database/postgres"
)

// Delete ...
func (r *repository) Delete(ctx context.Context, id int64) error {
	op := "app.v1.repository.user.Delete"

	qb := r.pgQb.
		Delete("users").
		Where(squirrel.Eq{"id": id})

	query, args, err := qb.ToSql()
	if err != nil {
		r.logger.Debug(op, slog.String("err", err.Error()))
		return core.ErrInternal
	}

	q := postgres.Query{
		Name:     op,
		QueryRaw: query,
	}

	rows, err := r.db.ExecContext(ctx, q, args...)
	if err != nil {
		r.logger.Debug(op, slog.String("err", err.Error()))
		return core.ErrInternal
	}

	if rows.RowsAffected() == 0 {
		return errNotFound
	}

	return nil
}
