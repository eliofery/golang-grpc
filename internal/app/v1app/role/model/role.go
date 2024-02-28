package model

// Role table description
const (
	TableName = "roles"

	ColumnID   = "id"
	ColumnName = "name"
)

// Role ...
type Role struct {
	ID   int64  `db:"id"`
	Name string `db:"name"`
}
