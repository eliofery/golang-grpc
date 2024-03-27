package model

import (
	"database/sql"
)

// Permission ...
type Permission struct {
	ID          int64          `db:"permission_id"`
	Name        string         `db:"permission_name"`
	Description sql.NullString `db:"permission_description"`
}
