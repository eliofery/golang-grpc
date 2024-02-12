package internal

import (
	"embed"
)

//go:embed migrations/*.sql
var EmbedMigration embed.FS
