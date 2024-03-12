package model

import (
	"database/sql"
)

// Permission table description
const (
	TableName = "permissions"

	ColumnID          = "id"
	ColumnName        = "name"
	ColumnDescription = "description"
)

// Permission ...
type Permission struct {
	ID          int64          `db:"id"`
	Name        string         `db:"name"`
	Description sql.NullString `db:"description"`
}
