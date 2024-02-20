package repository

import (
	"context"
	"errors"

	"github.com/eliofery/golang-fullstack/internal/app/v1app/user/model"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5/pgconn"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Create ...
func (r *repository) Create(ctx context.Context, userInfo *model.UserInfo) (*int64, error) {
	qb := r.pgQb().
		Insert(model.TableName).
		PlaceholderFormat(r.format).
		Columns(model.ColumnFirstName, model.ColumnLastName, model.ColumnEmail, model.ColumnPassword).
		Values(userInfo.FirstName, userInfo.LastName, userInfo.Email, userInfo.Password).
		Suffix("RETURNING id")

	query, args, err := qb.ToSql()
	if err != nil {
		return nil, err
	}

	var id int64
	if err = r.conn.QueryRow(ctx, query, args...).Scan(&id); err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code == pgerrcode.UniqueViolation {
				return nil, status.Error(codes.AlreadyExists, "email already exists")
			}
		}

		return nil, errors.New("failed to create user")
	}

	return &id, nil
}
