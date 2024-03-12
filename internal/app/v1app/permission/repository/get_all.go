package repository

import (
	"context"
	"log/slog"

	"github.com/eliofery/golang-grpc/internal/app/v1app/permission/model"
	"github.com/eliofery/golang-grpc/internal/core/database/postgres"
	"github.com/eliofery/golang-grpc/internal/core/pagination"
)

// GetAll ...
func (r *repository) GetAll(ctx context.Context, pagination *pagination.Pagination) ([]model.Permission, error) {
	op := "v1.permission.repository.GetAll"

	qb := r.pgQb.
		Select(model.ColumnID, model.ColumnName, model.ColumnDescription).
		From(model.TableName).
		Limit(pagination.Limit()).
		Offset(pagination.Offset())

	query, _, err := qb.ToSql()
	if err != nil {
		r.logger.Debug(op, slog.String("err", err.Error()))
		return nil, errGetAll
	}

	q := postgres.Query{
		Name:     op,
		QueryRaw: query,
	}

	var permissions []model.Permission
	if err = r.db.ScanAllContext(ctx, &permissions, q); err != nil {
		r.logger.Debug(op, slog.String("err", err.Error()))
		return nil, errGetAll
	}

	if len(permissions) == 0 {
		return nil, errNotFound
	}

	return permissions, nil
}
