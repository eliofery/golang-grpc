package dto

import (
	"database/sql"
)

// User ...
type User struct {
	FirstName sql.NullString
	LastName  sql.NullString
	Email     string
	Password  string
	RoleID    int64
}
