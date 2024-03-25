package dto

import (
	"database/sql"
	"time"
)

// UserUpdate ...
type UserUpdate struct {
	ID          int64
	FirstName   sql.NullString
	LastName    sql.NullString
	Email       sql.NullString
	OldPassword sql.NullString
	NewPassword sql.NullString
	RoleID      sql.NullInt64
	UpdatedAt   time.Time
}
