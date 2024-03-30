package access

import (
	"context"

	"github.com/eliofery/golang-grpc/internal/core"
	"github.com/eliofery/golang-grpc/pkg/eslog"
)

type key string

const userKey key = "user"

// Manager ...
type Manager interface {
	WithUser(ctx context.Context, user *User) context.Context
	User(ctx context.Context) (*User, error)
	IsAuth(ctx context.Context) bool
	IsAccess(ctx context.Context, needPermission string) bool
}

// access ...
type access struct {
	logger eslog.Logger
}

// New ...
func New(logger eslog.Logger) Manager {
	return &access{
		logger: logger,
	}
}

// WithUser ...
func (a *access) WithUser(ctx context.Context, user *User) context.Context {
	return context.WithValue(ctx, userKey, user)
}

// User ...
func (a *access) User(ctx context.Context) (*User, error) {
	user, ok := ctx.Value(userKey).(*User)
	if ok {
		return user, nil
	}

	return nil, core.ErrAccessDenied
}

// IsAuth ...
func (a *access) IsAuth(ctx context.Context) bool {
	user, err := a.User(ctx)
	if err != nil {
		return false
	}

	return user != nil
}

func (a *access) IsAccess(ctx context.Context, needPermission string) bool {
	user, err := a.User(ctx)
	if err != nil {
		return false
	}

	for _, permission := range user.Permissions {
		if permission == needPermission {
			return true
		}
	}

	return false
}
