package repository

import (
	"context"
	"errors"
	"fmt"
	"log/slog"

	"github.com/Masterminds/squirrel"
	roleModel "github.com/eliofery/golang-grpc/internal/app/v1app/role/model"
	"github.com/eliofery/golang-grpc/internal/app/v1app/user/model"
	"github.com/eliofery/golang-grpc/internal/core/database/postgres"
	"github.com/jackc/pgx/v5"
)

// GetByID ...
func (r *repository) GetByID(ctx context.Context, id int64) (*model.User, error) {
	op := "v1.user.repository.GetByID"

	qb := r.pgQb.
		Select(model.ColumnAsID, model.ColumnFirstName, model.ColumnLastName, model.ColumnEmail, model.ColumnPassword, roleModel.ColumnAsID, roleModel.ColumnName, model.ColumnCreatedAt, model.ColumnUpdatedAt).
		From(model.TableName).
		InnerJoin(fmt.Sprintf("%s ON %s = %s", roleModel.TableName, roleModel.ColumnID, roleModel.ColumnAliasID)).
		Where(squirrel.Eq{model.ColumnID: id})

	query, args, err := qb.ToSql()
	if err != nil {
		r.logger.Debug(op, slog.String("err", err.Error()))
		return nil, errGetByID
	}

	q := postgres.Query{
		Name:     op,
		QueryRaw: query,
	}

	var user model.User
	if err = r.db.ScanOneContext(ctx, &user, q, args...); err != nil {
		r.logger.Debug(op, slog.String("err", err.Error()))

		if errors.Is(err, pgx.ErrNoRows) {
			return nil, errNotFound
		}

		return nil, errGetByID
	}

	return &user, nil
}
