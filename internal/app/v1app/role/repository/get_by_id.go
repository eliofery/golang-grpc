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

// GetByID ...
func (r *repository) GetByID(ctx context.Context, id int64) (*model.Role, error) {
	op := "v1.role.repository.GetByID"

	qb := r.pgQb.
		Select(model.ColumnAsID, model.ColumnName).
		From(model.TableName).
		Where(squirrel.Eq{model.ColumnID: id})

	query, args, err := qb.ToSql()
	if err != nil {
		r.logger.Debug(op, slog.String("err", err.Error()))
		return nil, errGetByID
	}

	q := postgres.Query{
		Name:     op,
		QueryRaw: query,
	}

	var role model.Role
	if err = r.db.ScanOneContext(ctx, &role, q, args...); err != nil {
		r.logger.Debug(op, slog.String("err", err.Error()))

		if errors.Is(err, pgx.ErrNoRows) {
			return nil, errNotFound
		}

		return nil, errGetByID
	}

	return &role, nil
}
