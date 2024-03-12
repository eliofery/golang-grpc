package repository

import (
	"context"
	"errors"
	"log/slog"

	"github.com/Masterminds/squirrel"
	"github.com/eliofery/golang-grpc/internal/app/v1app/permission/model"
	"github.com/eliofery/golang-grpc/internal/core/database/postgres"
	"github.com/jackc/pgx/v5"
)

func (r *repository) GetByID(ctx context.Context, id int64) (*model.Permission, error) {
	op := "v1.permission.repository.GetByID"

	qb := r.pgQb.
		Select(model.ColumnID, model.ColumnName, model.ColumnDescription).
		From(model.TableName).
		Where(squirrel.Eq{model.ColumnID: id})

	query, args, err := qb.ToSql()
	if err != nil {
		return nil, errGetByID
	}

	q := postgres.Query{
		Name:     op,
		QueryRaw: query,
	}

	var permission model.Permission
	if err = r.db.ScanOneContext(ctx, &permission, q, args...); err != nil {
		r.logger.Debug(op, slog.String("err", err.Error()))

		if errors.Is(err, pgx.ErrNoRows) {
			return nil, errNotFound
		}

		return nil, errGetByID
	}

	return &permission, nil
}
