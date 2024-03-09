package model

// Role table description
const (
	TableName = "roles"

	ColumnID   = "roles.id"
	ColumnName = "name"

	ColumnAliasID = "role_id"
	ColumnAsID    = "roles.id AS " + ColumnAliasID
)

// Role ...
type Role struct {
	ID   int64  `db:"role_id"`
	Name string `db:"name"`
}
