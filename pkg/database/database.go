package database

import "context"

// Database ...
type Database interface {
	Connect(ctx context.Context) error
}

// Connect ...
func Connect(ctx context.Context, db Database) error {
	if err := db.Connect(ctx); err != nil {
		return err
	}

	return nil
}
