package chat

import (
	"github.com/marinaaaniram/go-common-platform/pkg/db"

	"go-chat-server/internal/repository"
)

const (
	tableName = "message"

	idColumn        = "id"
	chatId          = "chat_id"
	sentBy          = "sent_by"
	text            = "text"
	createdAtColumn = "created_at"
	updatedAtColumn = "updated_at"
)

type repo struct {
	db db.Client
}

// Create Message repository
func NewMessageRepository(db db.Client) repository.MessageRepository {
	return &repo{db: db}
}
