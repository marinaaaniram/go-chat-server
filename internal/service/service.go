package service

import (
	"context"

	"github.com/marinaaaniram/go-chat-server/internal/model"
)

// Describe Chat service interface
type ChatService interface {
	Create(ctx context.Context) (int64, error)
	Delete(ctx context.Context, chatId int64) error
	SendMessage(ctx context.Context, chat *model.Message) error
}
