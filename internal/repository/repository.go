package repository

import (
	"context"

	"github.com/marinaaaniram/go-chat-server/internal/model"
)

type ChatRepository interface {
	Create(ctx context.Context, chat *model.Chat) (*model.Chat, error)
	Delete(ctx context.Context, chat *model.Chat) error
}

type MessageRepository interface {
	Send(ctx context.Context, message *model.Message) error
}
