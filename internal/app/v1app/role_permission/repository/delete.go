package repository

import (
	"context"
	"log/slog"

	"github.com/Masterminds/squirrel"
	"github.com/eliofery/golang-grpc/internal/app/v1app/role_permission/model"
	"github.com/eliofery/golang-grpc/internal/core/database/postgres"
)

func (r *repository) Delete(ctx context.Context, id int64) error {
	op := "v1.role_permission.repository.Delete"

	qb := r.pgQb.
		Delete(model.TableName).
		Where(squirrel.Eq{model.ColumnRoleID: id})

	query, args, err := qb.ToSql()
	if err != nil {
		r.logger.Debug(op, slog.String("err", err.Error()))
		return errCreate
	}

	q := postgres.Query{
		Name:     op,
		QueryRaw: query,
	}

	rows, err := r.db.ExecContext(ctx, q, args...)
	if err != nil {
		r.logger.Debug(op, slog.String("err", err.Error()))
		return errCreate
	}

	if rows.RowsAffected() == 0 {
		return errNotFound
	}

	return nil
}
