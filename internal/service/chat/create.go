package chat

import (
	"context"

	"go-chat-server/internal/errors"
	"go-chat-server/internal/model"
)

// Create chat in service layer
func (s *serv) Create(ctx context.Context, chat *model.Chat) (int64, error) {
	if chat == nil {
		return 0, errors.ErrPointerIsNil("chat")
	}

	chatId, err := s.chatRepository.Create(ctx, chat)
	if err != nil {
		return 0, err
	}

	return chatId, nil
}
