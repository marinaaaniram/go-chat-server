package service

import (
	"context"

	"github.com/marinaaaniram/go-chat-server/internal/model"
	desc "github.com/marinaaaniram/go-chat-server/pkg/chat_v1"
)

type ChatService interface {
	Create(ctx context.Context, chat *model.Chat) (*desc.Chat, error)
	Delete(ctx context.Context, chat *model.Chat) error
}

type MessageService interface {
	Send(ctx context.Context, chat *model.Message) error
}
