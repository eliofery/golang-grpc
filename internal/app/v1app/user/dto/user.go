package dto

import (
	"database/sql"
)

// User ...
type User struct {
	ID        int64
	FirstName sql.NullString
	LastName  sql.NullString
	Email     string
	Password  string
	RoleID    int64
}
