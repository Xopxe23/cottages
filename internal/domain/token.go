package domain

import "time"

type RefreshToken struct {
	ID        string //uuid
	UserID    string //uuid
	Token     string
	ExpiresAt time.Time
}
