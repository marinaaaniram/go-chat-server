package chat

import (
	"context"

	"github.com/marinaaaniram/go-chat-server/internal/errors"
	"github.com/marinaaaniram/go-chat-server/internal/model"
)

// Delete chat in service layer
func (s *serv) Delete(ctx context.Context, chat *model.Chat) error {
	if chat == nil {
		return errors.ErrPointerIsNil("chat")
	}

	err := s.chatRepository.Delete(ctx, chat)
	if err != nil {
		return err
	}

	return nil
}
