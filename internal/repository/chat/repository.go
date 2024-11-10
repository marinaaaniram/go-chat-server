package chat

import (
	"github.com/marinaaaniram/go-chat-server/internal/repository"
	"github.com/marinaaaniram/go-common-platform/pkg/db"
)

const (
	tableName = "chat"

	idColumn        = "id"
	usernamesColumn = "usernames"
	createdAtColumn = "created_at"
	updatedAtColumn = "updated_at"
)

type repo struct {
	db db.Client
}

// Create Chat repository
func NewChatRepository(db db.Client) repository.ChatRepository {
	return &repo{db: db}
}
