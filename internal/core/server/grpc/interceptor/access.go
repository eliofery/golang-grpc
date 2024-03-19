package interceptor

import "context"

// IsAuthenticated ...
func IsAuthenticated(ctx context.Context) bool {
	return User(ctx) != nil
}

// UserID ...
func UserID(ctx context.Context) int64 {
	if !IsAuthenticated(ctx) {
		return 0
	}

	return User(ctx).ID
}

// UserPermissions ...
func UserPermissions(ctx context.Context) []string {
	if !IsAuthenticated(ctx) {
		return nil
	}

	return User(ctx).Permissions
}

// UserToken ...
func UserToken(ctx context.Context) string {
	if !IsAuthenticated(ctx) {
		return ""
	}

	return User(ctx).Token
}

// IsAccess ...
func IsAccess(ctx context.Context, needPermissions ...string) bool {
	access := false
	for _, permission := range UserPermissions(ctx) {
		for _, needPermission := range needPermissions {
			if permission == needPermission {
				access = true
			}
		}
	}

	return access
}
