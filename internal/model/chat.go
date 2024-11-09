package model

import (
	"time"
)

// Internal Chat model
type Chat struct {
	ID        int64
	Usernames []string
	CreatedAt time.Time
	UpdatedAt *time.Time
}
