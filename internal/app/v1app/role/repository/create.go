package repository

import (
	"context"
	"errors"
	"log/slog"

	"github.com/eliofery/golang-grpc/internal/app/v1app/role/dto"
	"github.com/eliofery/golang-grpc/internal/app/v1app/role/model"
	"github.com/eliofery/golang-grpc/internal/core/database/postgres"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5/pgconn"
)

func (r *repository) Create(ctx context.Context, role *dto.Role) (int64, error) {
	op := "v1.role.repository.Create"

	qb := r.pgQb.
		Insert(model.TableName).
		Columns(model.ColumnName).
		Values(role.Name).
		Suffix("RETURNING id")

	query, args, err := qb.ToSql()
	if err != nil {
		r.logger.Debug(op, slog.String("err", err.Error()))
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
