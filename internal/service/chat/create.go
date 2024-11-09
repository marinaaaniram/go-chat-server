package chat

import (
	"context"

	"github.com/marinaaaniram/go-chat-server/internal/converter"
	"github.com/marinaaaniram/go-chat-server/internal/errors"
	"github.com/marinaaaniram/go-chat-server/internal/model"
	desc "github.com/marinaaaniram/go-chat-server/pkg/chat_v1"
)

// Create chat in service layer
func (s *serv) Create(ctx context.Context, chat *model.Chat) (*desc.Chat, error) {
	if chat == nil {
		return nil, errors.ErrPointerIsNil("chat")
	}

	chatObj, err := s.chatRepository.Create(ctx, chat)
	if err != nil {
		return nil, err
	}

	return converter.FromChatToDesc(chatObj), nil
}
