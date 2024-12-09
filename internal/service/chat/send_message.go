package chat

import (
	"context"

	"github.com/marinaaaniram/go-chat-server/internal/errors"
	"github.com/marinaaaniram/go-chat-server/internal/model"
)

// Send message in service layer
func (s *serv) SendMessage(ctx context.Context, message *model.Message) error {
	if message == nil {
		return errors.ErrPointerIsNil("message")
	}

	err := s.messageRepository.Send(ctx, message)
	if err != nil {
		return err
	}

	return nil
}
