package repository

import (
	"context"
	"errors"
	"log/slog"

	"github.com/eliofery/golang-grpc/internal/app/v1app/permission/dto"
	"github.com/eliofery/golang-grpc/internal/app/v1app/permission/model"
	"github.com/eliofery/golang-grpc/internal/core/database/postgres"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5/pgconn"
)

func (r *repository) Create(ctx context.Context, permission *dto.Permission) (int64, error) {
	op := "v1.permission.repository.Create"

	qb := r.pgQb.
		Insert(model.TableName).
		Columns(model.ColumnName, model.ColumnDescription).
		Values(permission.Name, permission.Description).
		Suffix("RETURNING id")

	query, args, err := qb.ToSql()
	if err != nil {
		return 0, errCreate
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

		return 0, errCreate
	}

	return id, nil
}
