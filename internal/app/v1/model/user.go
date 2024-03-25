package model

import (
	"database/sql"
	"time"
)

// User ...
type User struct {
	ID        int64          `db:"user_id"`
	FirstName sql.NullString `db:"first_name"`
	LastName  sql.NullString `db:"last_name"`
	Email     string         `db:"email"`
	Password  string         `db:"password"`
	Role      Role           `db:""`
	CreatedAt time.Time      `db:"created_at"`
	UpdatedAt sql.NullTime   `db:"updated_at"`
}
