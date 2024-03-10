package repository

import (
	"context"
	"fmt"
	"log/slog"

	roleModel "github.com/eliofery/golang-grpc/internal/app/v1app/role/model"
	"github.com/eliofery/golang-grpc/internal/app/v1app/user/model"
	"github.com/eliofery/golang-grpc/internal/core/database/postgres"
	"github.com/eliofery/golang-grpc/internal/core/pagination"
)

// GetAll ...
func (r *repository) GetAll(ctx context.Context, pagination *pagination.Pagination) ([]model.User, error) {
	op := "v1.user.repository.GetAll"

	qb := r.pgQb.
		Select(model.ColumnAsID, model.ColumnFirstName, model.ColumnLastName, model.ColumnEmail, model.ColumnPassword, roleModel.ColumnAsID, roleModel.ColumnName, model.ColumnCreatedAt, model.ColumnUpdatedAt).
		From(model.TableName).
		InnerJoin(fmt.Sprintf("%s ON %s = %s", roleModel.TableName, roleModel.ColumnID, roleModel.ColumnAliasID)).
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
