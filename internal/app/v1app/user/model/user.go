package model

import (
	"database/sql"
	"time"
)

// User table description
const (
	TableName = "users"

	ColumnID        = "id"
	ColumnFirstName = "first_name"
	ColumnLastName  = "last_name"
	ColumnEmail     = "email"
	ColumnPassword  = "password"
	ColumnRoleID    = "role_id"
	ColumnCreatedAt = "created_at"
	ColumnUpdatedAt = "updated_at"
)

// User ...
type User struct {
	ID        int64          `db:"id"`
	FirstName sql.NullString `db:"first_name"`
	LastName  sql.NullString `db:"last_name"`
	Email     string         `db:"email"`
	Password  string         `db:"password"`
	RoleID    int64          `db:"role_id"`
	CreatedAt time.Time      `db:"created_at"`
	UpdatedAt sql.NullTime   `db:"updated_at"`
}
