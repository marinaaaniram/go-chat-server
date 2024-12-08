package repository

import (
	"context"

	"github.com/marinaaaniram/go-chat-server/internal/model"
)

// Describe Chat repository interface
type ChatRepository interface {
	Create(ctx context.Context) (int64, error)
	Delete(ctx context.Context, chatId int64) error
}

// Describe Message repository interface
type MessageRepository interface {
	Send(ctx context.Context, message *model.Message) error
}
