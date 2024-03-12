package dto

import (
	"database/sql"
)

// Update ...
type Update struct {
	ID          int64
	Name        sql.NullString
	Description sql.NullString
}
