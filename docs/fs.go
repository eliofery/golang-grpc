package docs

import (
	"embed"
)

// EmbedMigration ...
//
//go:embed migrations/*.sql
var EmbedMigration embed.FS
