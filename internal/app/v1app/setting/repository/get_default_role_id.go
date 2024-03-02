package repository

import (
	"context"
	"errors"
	"log/slog"
	"strconv"

	"github.com/Masterminds/squirrel"
	"github.com/eliofery/golang-grpc/internal/app/v1app/setting/model"
	"github.com/eliofery/golang-grpc/internal/core/database/postgres"
	"github.com/jackc/pgx/v5"
)

const defaultNameRoleID = "default_role_id"

// GetDefaultRoleID ...
func (r *repository) GetDefaultRoleID(ctx context.Context) (int64, error) {
	op := "v1.setting.repository.GetDefaultRoleID"

	qb := r.pgQb.
		Select(model.ColumnValue).
		From(model.TableName).
		Where(squirrel.Eq{model.ColumnName: defaultNameRoleID})

	query, args, err := qb.ToSql()
	if err != nil {
		r.logger.Debug(op, slog.String("err", err.Error()))
		return 0, errGetDefaultRole
	}

	q := postgres.Query{
		Name:     op,
		QueryRaw: query,
	}

	var setting model.Setting
	if err = r.db.ScanOneContext(ctx, &setting, q, args...); err != nil {
		r.logger.Debug(op, slog.String("err", err.Error()))

		if errors.Is(err, pgx.ErrNoRows) {
			return 0, errNotFound
		}

		return 0, errGetDefaultRole
	}

	defaultRoleID, err := strconv.ParseInt(setting.Value, 10, 64)
	if err != nil {
		r.logger.Error(op, slog.String("err", err.Error()))
		return 0, errGetDefaultRole
	}

	return defaultRoleID, nil
}
