package model

// Role permissions table description
const (
	TableName = "role_permissions"

	ColumnRoleID       = "role_id"
	ColumnPermissionID = "permission_id"
)

// RolePermission ...
type RolePermission struct {
	RoleID       int64 `db:"role_id"`
	PermissionID int64 `db:"permission_id"`
}
