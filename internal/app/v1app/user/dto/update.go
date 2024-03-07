package dto

import (
	"database/sql"
	"time"
)

// Update ...
type Update struct {
	ID          int64
	FirstName   sql.NullString
	LastName    sql.NullString
	Email       sql.NullString
	OldPassword sql.NullString
	NewPassword sql.NullString
	UpdatedAt   time.Time
}
