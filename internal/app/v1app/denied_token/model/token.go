package model

// Token table description
const (
	TableName = "denied_tokens"

	ColumnID    = "id"
	ColumnToken = "token"
)

// DeniedToken ...
type DeniedToken struct {
	ID    int64  `db:"id"`
	Token string `db:"token"`
}
