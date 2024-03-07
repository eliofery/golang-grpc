package repository

import (
	"context"
	"errors"
	"log/slog"

	"github.com/Masterminds/squirrel"
	"github.com/eliofery/golang-grpc/internal/app/v1app/user/model"
	"github.com/eliofery/golang-grpc/internal/core/database/postgres"
	"github.com/jackc/pgx/v5"
)

// GetByEmail ...
func (r *repository) GetByEmail(ctx context.Context, email string) (*model.User, error) {
	op := "v1.user.repository.GetByEmail"

	qb := r.pgQb.
		Select(model.ColumnID, model.ColumnPassword).
		From(model.TableName).
		Where(squirrel.Eq{"email": email})

	query, args, err := qb.ToSql()
	if err != nil {
		r.logger.Debug(op, slog.String("err", err.Error()))
		return nil, errGetByEmail
	}

	q := postgres.Query{
		Name:     op,
		QueryRaw: query,
	}

	var user model.User
	if err = r.db.ScanOneContext(ctx, &user, q, args...); err != nil {
		r.logger.Debug(op, slog.String("err", err.Error()))

		if errors.Is(err, pgx.ErrNoRows) {
			return nil, errWrongLoginOrPassword
		}

		return nil, errGetByEmail
	}

	return &user, nil
}
