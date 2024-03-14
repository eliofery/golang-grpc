package repository

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/Masterminds/squirrel"
	permissionModel "github.com/eliofery/golang-grpc/internal/app/v1app/permission/model"
	"github.com/eliofery/golang-grpc/internal/app/v1app/role_permission/model"
	"github.com/eliofery/golang-grpc/internal/core/database/postgres"
)

// GetPermissionsByRoleID ...
func (r *repository) GetPermissionsByRoleID(ctx context.Context, id int64) ([]string, error) {
	op := "v1.role_permission.repository.GetPermissionsByRoleID"

	qb := r.pgQb.
		Select(permissionModel.ColumnName).
		From(model.TableName).
		InnerJoin(fmt.Sprintf("%s ON %s = %s", permissionModel.TableName, model.ColumnPermissionID, permissionModel.ColumnID)).
		Where(squirrel.Eq{model.ColumnRoleID: id})

	query, args, err := qb.ToSql()
	if err != nil {
		r.logger.Debug(op, slog.String("err", err.Error()))
		return nil, errGet
	}

	q := postgres.Query{
		Name:     op,
		QueryRaw: query,
	}

	var permissions []string
	if err = r.db.ScanAllContext(ctx, &permissions, q, args...); err != nil {
		r.logger.Debug(op, slog.String("err", err.Error()))
		return nil, errGet
	}

	return permissions, nil
}
