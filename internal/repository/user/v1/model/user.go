package model

import (
	"database/sql"
	"time"
)

// User ...
type User struct {
	ID        int64
	Name      string
	CreatedAt time.Time
	UpdatedAt sql.NullTime
}
