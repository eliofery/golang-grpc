package model

import (
	"database/sql"
	"time"

	roleModel "github.com/eliofery/golang-grpc/internal/app/v1app/role/model"
)

// User table description
const (
	TableName = "users"

	ColumnID        = "users.id"
	ColumnFirstName = "first_name"
	ColumnLastName  = "last_name"
	ColumnEmail     = "email"
	ColumnPassword  = "password"
	ColumnRoleID    = "role_id"
	ColumnCreatedAt = "created_at"
	ColumnUpdatedAt = "updated_at"

	ColumnAliasID = "user_id"
	ColumnAsID    = "users.id AS user_id"
)

// User ...
type User struct {
	ID        int64           `db:"user_id"`
	FirstName sql.NullString  `db:"first_name"`
	LastName  sql.NullString  `db:"last_name"`
	Email     string          `db:"email"`
	Password  string          `db:"password"`
	Role      *roleModel.Role `db:""`
	CreatedAt time.Time       `db:"created_at"`
	UpdatedAt sql.NullTime    `db:"updated_at"`
}
