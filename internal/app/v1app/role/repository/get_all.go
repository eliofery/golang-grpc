package repository

import (
	"context"
	"log/slog"

	"github.com/eliofery/golang-grpc/internal/app/v1app/role/model"
	"github.com/eliofery/golang-grpc/internal/core/database/postgres"
	"github.com/eliofery/golang-grpc/internal/core/pagination"
)

// GetAll ...
func (r *repository) GetAll(ctx context.Context, pagination *pagination.Pagination) ([]model.Role, error) {
	op := "v1.role.repository.GetAll"

	qb := r.pgQb.
		Select(model.ColumnAsID, model.ColumnName).
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

	var roles []model.Role
	if err = r.db.ScanAllContext(ctx, &roles, q); err != nil {
		r.logger.Debug(op, slog.String("err", err.Error()))
		return nil, errGetAll
	}

	if len(roles) == 0 {
		return nil, errNotFound
	}

	return roles, nil
}
