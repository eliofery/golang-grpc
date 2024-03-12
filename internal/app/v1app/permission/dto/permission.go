package dto

import (
	"database/sql"
)

// Permission ...
type Permission struct {
	ID          int64
	Name        string
	Description sql.NullString
}
