package user

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"sort"

	"github.com/eliofery/golang-grpc/internal/app/v1/model"
	"github.com/eliofery/golang-grpc/internal/core"
	"github.com/eliofery/golang-grpc/internal/core/database/postgres"
	"github.com/eliofery/golang-grpc/internal/core/pagination"
	"github.com/jackc/pgx/v5"
)

// GetAll ...
func (r *repository) GetAll(ctx context.Context, pagination pagination.Pagination) ([]model.User, error) {
	op := "app.v1.repository.user.GetAll"

	usersQb := r.pgQb.
		Select("users.id AS user_id", "first_name", "last_name", "email", "password", "role_id", "created_at", "updated_at").
		From("users").
		OrderBy("users.id").
		Limit(pagination.Limit()).
		Offset(pagination.Offset())

	rolesQb := r.pgQb.
		Select("roles.id AS role_id", "roles.name AS role_name", "permissions.id AS permission_id", "permissions.name AS permission_name", "permissions.description AS permission_description").
		From("roles").
		InnerJoin("roles_permissions ON roles_permissions.role_id = roles.id").
		InnerJoin("permissions ON permissions.id = roles_permissions.permission_id").
		Where("roles.id = u.role_id")

	queryRoles, args, err := rolesQb.ToSql()
	if err != nil {
		r.logger.Debug(op, slog.String("err", err.Error()))
		return nil, core.ErrInternal
	}

	query, _, err := r.pgQb.
		Select("user_id", "u.email", "u.first_name", "u.last_name", "u.password", "u.created_at", "u.updated_at", "r.role_id", "role_name", "permission_id", "permission_name", "permission_description").
		FromSelect(usersQb, "u").
		JoinClause(fmt.Sprintf("LEFT JOIN LATERAL (%s)", queryRoles), args...).
		Suffix("AS r ON true").
		ToSql()

	if err != nil {
		r.logger.Debug(op, slog.String("err", err.Error()))
		return nil, core.ErrInternal
	}

	q := postgres.Query{
		Name:     op,
		QueryRaw: query,
	}

	rows, err := r.db.QueryContext(ctx, q)
	if err != nil {
		r.logger.Debug(op, slog.String("err", err.Error()))

		if errors.Is(err, pgx.ErrNoRows) {
			return nil, errNotFound
		}

		return nil, core.ErrInternal
	}
	defer rows.Close()

	usersMap := make(map[int64]*model.User)
	for rows.Next() {
		var user model.User
		var role model.Role
		var permission model.Permission

		if err = rows.Scan(
			&user.ID,
			&user.Email,
			&user.FirstName,
			&user.LastName,
			&user.Password,
			&user.CreatedAt,
			&user.UpdatedAt,
			&role.ID,
			&role.Name,
			&permission.ID,
			&permission.Name,
			&permission.Description,
		); err != nil {
			r.logger.Debug(op, slog.String("err", err.Error()))
			return nil, core.ErrInternal
		}

		if value, ok := usersMap[user.ID]; ok {
			value.Role.Permissions = append(value.Role.Permissions, permission)
		} else {
			user.Role = role
			user.Role.Permissions = append(user.Role.Permissions, permission)
			usersMap[user.ID] = &user
		}
	}

	var users []model.User
	for _, user := range usersMap {
		users = append(users, *user)
	}

	sort.Slice(users, func(i, j int) bool {
		return users[i].ID < users[j].ID
	})

	return users, nil
}
