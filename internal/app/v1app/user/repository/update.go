package repository

import (
	"context"
	"log/slog"

	"github.com/Masterminds/squirrel"
	"github.com/eliofery/golang-grpc/internal/app/v1app/user/dto"
	"github.com/eliofery/golang-grpc/internal/app/v1app/user/model"
	"github.com/eliofery/golang-grpc/internal/core/database/postgres"
)

// Update ...
func (r *repository) Update(ctx context.Context, user *dto.Update) (*model.User, error) {
	op := "v1.user.repository.Update"

	values := squirrel.Eq{
		model.ColumnUpdatedAt: user.UpdatedAt,
	}

	if user.FirstName.Valid {
		values[model.ColumnFirstName] = user.FirstName
	}

	if user.LastName.Valid {
		values[model.ColumnLastName] = user.LastName
	}

	if user.Email.Valid {
		values[model.ColumnEmail] = user.Email
	}

	if user.NewPassword.Valid {
		values[model.ColumnPassword] = user.NewPassword
	}

	qb := r.pgQb.
		Update(model.TableName).
		SetMap(values).
		Where(squirrel.Eq{model.ColumnID: user.ID}).
		Suffix("RETURNING id, first_name, last_name, email, created_at, updated_at")

	query, args, err := qb.ToSql()
	if err != nil {
		r.logger.Debug(op, slog.String("err", err.Error()))
		return nil, errUpdate
	}

	q := postgres.Query{
		Name:     op,
		QueryRaw: query,
	}

	var updateUser model.User
	if err = r.db.ScanOneContext(ctx, &updateUser, q, args...); err != nil {
		r.logger.Debug(op, slog.String("err", err.Error()))
		return nil, errUpdate
	}

	return &updateUser, nil
}
