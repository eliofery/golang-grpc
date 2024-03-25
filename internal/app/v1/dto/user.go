package dto

import (
	"database/sql"

	roleModel "github.com/eliofery/golang-grpc/internal/app/v1/model"
)

// User ...
type User struct {
	ID        int64
	FirstName sql.NullString
	LastName  sql.NullString
	Email     string
	Password  string
	Role      roleModel.Role
}

// PermissionsNames ...
func (u *User) PermissionsNames() []string {
	permissions := make([]string, len(u.Role.Permissions))
	for i, permission := range u.Role.Permissions {
		permissions[i] = permission.Name
	}

	return permissions
}
