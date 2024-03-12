package repository

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/Masterminds/squirrel"
	"github.com/eliofery/golang-grpc/internal/app/v1app/permission/dto"
	"github.com/eliofery/golang-grpc/internal/app/v1app/permission/model"
	"github.com/eliofery/golang-grpc/internal/core/database/postgres"
)

func (r *repository) Update(ctx context.Context, permission *dto.Update) (*model.Permission, error) {
	op := "v1.permission.repository.Update"

	values := make(squirrel.Eq)
	if permission.Name.Valid {
		values[model.ColumnName] = permission.Name
	}

	if permission.Description.Valid {
		values[model.ColumnDescription] = permission.Description
	}

	qb := r.pgQb.
		Update(model.TableName).
		SetMap(values).
		Where(squirrel.Eq{model.ColumnID: permission.ID}).
		Suffix(
			fmt.Sprintf("RETURNING %v, %v, %v", model.ColumnID, model.ColumnName, model.ColumnDescription),
		)

	query, args, err := qb.ToSql()
	if err != nil {
		r.logger.Debug(op, slog.String("err", err.Error()))
		return nil, errUpdate
	}

	q := postgres.Query{
		Name:     op,
		QueryRaw: query,
	}

	var updatePermission model.Permission
	if err = r.db.ScanOneContext(ctx, &updatePermission, q, args...); err != nil {
		r.logger.Debug(op, slog.String("err", err.Error()))
		return nil, errUpdate
	}

	return &updatePermission, nil
}
