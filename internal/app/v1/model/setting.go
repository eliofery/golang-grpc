package model

const (
	// DefaultNameRoleID ...
	DefaultNameRoleID = "default_role_id"
)

// Setting ...
type Setting struct {
	ID    int64  `db:"id"`
	Name  string `db:"name"`
	Value string `db:"value"`
}
