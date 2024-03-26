package setting

import (
	"context"
	"log/slog"
	"strconv"

	"github.com/Masterminds/squirrel"
	"github.com/eliofery/golang-grpc/internal/app/v1/model"
	"github.com/eliofery/golang-grpc/internal/core"
	"github.com/eliofery/golang-grpc/internal/core/database/postgres"
)

// GetDefaultRoleByID ...
func (r *repository) GetDefaultRoleByID(ctx context.Context) (int64, error) {
	op := "app.v1.repository.setting.GetDefaultRoleByID"

	query, args, err := r.pgQb.
		Select("value").
		From("settings").
		Where(squirrel.Eq{"name": model.DefaultNameRoleID}).
		ToSql()

	if err != nil {
		r.logger.Debug(op, slog.String("err", err.Error()))
		return 0, core.ErrInternal
	}

	q := postgres.Query{
		Name:     op,
		QueryRaw: query,
	}

	var defaultRoleIDStr string
	err = r.db.QueryRowContext(ctx, q, args...).Scan(&defaultRoleIDStr)
	if err != nil {
		r.logger.Debug(op, slog.String("err", err.Error()))
		return 0, core.ErrInternal
	}

	defaultRoleID, err := strconv.ParseInt(defaultRoleIDStr, 10, 64)
	if err != nil {
		r.logger.Debug(op, slog.String("err", err.Error()))
		return 0, core.ErrInternal
	}

	return defaultRoleID, nil
}
