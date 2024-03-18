package model

import (
	"database/sql"
)

// Permission table description
const (
	TableName = "permissions"

	ColumnID          = "permissions.id"
	ColumnName        = "name"
	ColumnDescription = "description"

	ColumnAliasID = "permission_id"
	ColumnAsID    = "permissions.id AS permission_id"
)

// Permission ...
type Permission struct {
	ID          int64          `db:"permission_id"`
	Name        string         `db:"name"`
	Description sql.NullString `db:"description"`
}
