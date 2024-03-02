package model

// Setting table description
const (
	TableName = "settings"

	ColumnID    = "id"
	ColumnName  = "name"
	ColumnValue = "value"
)

// Setting ...
type Setting struct {
	ID    int64  `db:"id"`
	Name  string `db:"name"`
	Value string `db:"value"`
}
