package service

import (
	"context"

	"go-chat-server/internal/model"
)

// Describe Chat service interface
type ChatService interface {
	Create(ctx context.Context, chat *model.Chat) (int64, error)
	Delete(ctx context.Context, chat *model.Chat) error
}

// Describe Mesaage service interface
type MessageService interface {
	Send(ctx context.Context, chat *model.Message) error
}
