package repository

import (
	"context"
	"log/slog"

	"github.com/Masterminds/squirrel"
	"github.com/eliofery/golang-grpc/internal/app/v1app/role/dto"
	"github.com/eliofery/golang-grpc/internal/app/v1app/role/model"
	"github.com/eliofery/golang-grpc/internal/core/database/postgres"
)

// Update ...
func (r *repository) Update(ctx context.Context, role *dto.Role) (*model.Role, error) {
	op := "v1.role.repository.Update"

	qb := r.pgQb.
		Update(model.TableName).
		SetMap(squirrel.Eq{model.ColumnName: role.Name}).
		Where(squirrel.Eq{model.ColumnID: role.ID}).
		Suffix("RETURNING id, name")

	query, args, err := qb.ToSql()
	if err != nil {
		r.logger.Debug(op, slog.String("err", err.Error()))
		return nil, errUpdate
	}

	q := postgres.Query{
		Name:     op,
		QueryRaw: query,
	}

	var roleUpdate model.Role
	if err = r.db.ScanOneContext(ctx, &roleUpdate, q, args...); err != nil {
		r.logger.Debug(op, slog.String("err", err.Error()))
		return nil, errUpdate
	}

	return &roleUpdate, nil
}
