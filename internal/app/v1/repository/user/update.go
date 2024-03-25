package user

import (
	"context"
	"log/slog"

	"github.com/Masterminds/squirrel"
	"github.com/eliofery/golang-grpc/internal/app/v1/dto"
	"github.com/eliofery/golang-grpc/internal/app/v1/model"
	"github.com/eliofery/golang-grpc/internal/core"
	"github.com/eliofery/golang-grpc/internal/core/database/postgres"
)

// Update ...
func (r *repository) Update(ctx context.Context, user *dto.UserUpdate) (*model.User, error) {
	op := "app.v1.repository.user.Update"

	values := squirrel.Eq{"updated_at": user.UpdatedAt}

	if user.FirstName.Valid {
		values["first_name"] = user.FirstName
	}

	if user.LastName.Valid {
		values["last_name"] = user.LastName
	}

	if user.Email.Valid {
		values["email"] = user.Email
	}

	if user.NewPassword.Valid {
		values["password"] = user.NewPassword
	}

	if user.RoleID.Valid {
		values["role_id"] = user.RoleID
	}

	query, args, err := r.pgQb.
		Update("users").
		SetMap(values).
		Where(squirrel.Eq{"id": user.ID}).
		Suffix("RETURNING users.id AS user_id, first_name, last_name, email, password, role_id, created_at, updated_at").
		ToSql()

	if err != nil {
		r.logger.Debug(op, slog.String("err", err.Error()))
		return nil, core.ErrInternal
	}

	q := postgres.Query{
		Name:     op,
		QueryRaw: query,
	}

	var updateUser model.User
	if err = r.db.ScanOneContext(ctx, &updateUser, q, args...); err != nil {
		r.logger.Debug(op, slog.String("err", err.Error()))
		return nil, core.ErrInternal
	}

	role, err := r.roleRepository.GetByID(ctx, updateUser.Role.ID)
	if err != nil {
		r.logger.Debug(op, slog.String("err", err.Error()))
		return nil, core.ErrInternal
	}
	updateUser.Role = *role

	return &updateUser, nil
}
