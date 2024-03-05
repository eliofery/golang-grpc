package repository

import (
	"context"
	"errors"
	"log/slog"

	"github.com/eliofery/golang-grpc/internal/app/v1app/user/dto"
	"github.com/eliofery/golang-grpc/internal/app/v1app/user/model"
	"github.com/eliofery/golang-grpc/internal/core/database/postgres"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5/pgconn"
)

// Create ...
func (r *repository) Create(ctx context.Context, user *dto.User) (int64, error) {
	op := "v1.user.repository.User"

	qb := r.pgQb.
		Insert(model.TableName).
		Columns(model.ColumnFirstName, model.ColumnLastName, model.ColumnEmail, model.ColumnPassword, model.ColumnRoleID).
		Values(user.FirstName, user.LastName, user.Email, user.Password, user.RoleID).
		Suffix("RETURNING id")

	query, args, err := qb.ToSql()
	if err != nil {
		r.logger.Debug(op, slog.String("err", err.Error()))
		return 0, errCreate
	}

	q := postgres.Query{
		Name:     op,
		QueryRaw: query,
	}

	var id int64
	if err = r.db.QueryRowContext(ctx, q, args...).Scan(&id); err != nil {
		r.logger.Debug(op, slog.String("err", err.Error()))

		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			switch pgErr.Code {
			case pgerrcode.UniqueViolation:
				return 0, errExists
			case pgerrcode.ForeignKeyViolation:
				return 0, errRoleID
			}
		}

		return 0, errCreate
	}

	return id, nil
}
