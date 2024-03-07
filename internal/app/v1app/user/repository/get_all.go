package repository

import (
	"context"
	"log/slog"

	"github.com/eliofery/golang-grpc/internal/app/v1app/user/model"
	"github.com/eliofery/golang-grpc/internal/core/database/postgres"
	"github.com/eliofery/golang-grpc/internal/core/pagination"
)

// GetAll ...
func (r *repository) GetAll(ctx context.Context, pagination *pagination.Pagination) ([]model.User, error) {
	op := "v1.user.repository.GetAll"

	qb := r.pgQb.
		Select(model.ColumnID, model.ColumnFirstName, model.ColumnLastName, model.ColumnEmail, model.ColumnPassword, model.ColumnCreatedAt, model.ColumnUpdatedAt).
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

	var users []model.User
	if err = r.db.ScanAllContext(ctx, &users, q); err != nil {
		r.logger.Debug(op, slog.String("err", err.Error()))
		return nil, errGetAll
	}

	if len(users) == 0 {
		return nil, errNotFound
	}

	return users, nil
}
