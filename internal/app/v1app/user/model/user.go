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
	ColumnCreatedAt = "created_at"
	ColumnUpdatedAt = "updated_at"
)

// User ...
type User struct {
	ID        int64 `db:"id"`
	UserInfo  `db:""`
	CreatedAt time.Time    `db:"created_at"`
	UpdatedAt sql.NullTime `db:"updated_at"`
}
