package model

import (
	"database/sql"
	"time"
)

// Internal Message model
type Message struct {
	ID        int64
	ChatId    int64
	SentBy    string
	Text      string
	CreatedAt time.Time
	UpdatedAt sql.NullTime
}
