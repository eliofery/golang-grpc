package migrations

import (
	"embed"
)

// EmbedMigration ...
//
//go:embed sql/*.sql
var EmbedMigration embed.FS
