package repository

import (
	"context"

	"github.com/marinaaaniram/go-chat-server/internal/model"
)

// Describe Chat repository interface
type ChatRepository interface {
	Create(ctx context.Context, chat *model.Chat) (*model.Chat, error)
	Delete(ctx context.Context, chat *model.Chat) error
}

// Describe Message repository interface
type MessageRepository interface {
	Send(ctx context.Context, message *model.Message) error
}
