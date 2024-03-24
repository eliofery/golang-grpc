package model

// Role ...
type Role struct {
	ID          int64        `db:"role_id"`
	Name        string       `db:"role_name"`
	Permissions []Permission `db:""`
}
