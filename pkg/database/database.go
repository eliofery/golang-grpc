package database

import "context"

// Database ...
type Database interface {
	Connect(ctx context.Context) error
}
