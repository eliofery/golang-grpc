package model

import "database/sql"

// UserInfo ...
type UserInfo struct {
	FirstName sql.NullString `db:"first_name"`
	LastName  sql.NullString `db:"last_name"`
	Email     string         `db:"email"`
	Password  string         `db:"password"`
}
