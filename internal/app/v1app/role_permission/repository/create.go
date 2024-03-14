package repository

import (
	"context"
	"log/slog"

	"github.com/eliofery/golang-grpc/internal/app/v1app/role/dto"
	"github.com/eliofery/golang-grpc/internal/app/v1app/role_permission/model"
	"github.com/eliofery/golang-grpc/internal/core/database/postgres"
)

func (r *repository) Create(ctx context.Context, role *dto.Role) error {
	op := "v1.role_permission.repository.Create"

	for _, permissionID := range role.Permissions {
		qb := r.pgQb.
			Insert(model.TableName).
			Columns(model.ColumnRoleID, model.ColumnPermissionID).
			Values(role.ID, permissionID)

		query, args, err := qb.ToSql()
		if err != nil {
			r.logger.Debug(op, slog.String("err", err.Error()))
			return errCreate
		}

		q := postgres.Query{
			Name:     op,
			QueryRaw: query,
		}

		if _, err = r.db.ExecContext(ctx, q, args...); err != nil {
			r.logger.Debug(op, slog.String("err", err.Error()))
			return errCreate
		}
	}

	return nil
}
