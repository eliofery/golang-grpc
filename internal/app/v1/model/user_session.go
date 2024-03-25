package model

import "time"

// UserSession ...
type UserSession struct {
	UserID       int64
	IPAddress    string
	UserAgent    string
	RefreshToken string
	Fingerprint  string
	Expires      time.Time
}
