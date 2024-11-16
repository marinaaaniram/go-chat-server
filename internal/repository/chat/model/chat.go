package model

import (
	"database/sql"
	"time"
)

// Repository Chat model
type Chat struct {
	ID        int64
	Usernames []string
	CreatedAt time.Time
	UpdatedAt sql.NullTime
}
