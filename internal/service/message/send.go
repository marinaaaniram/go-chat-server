package message

import (
	"context"

	"github.com/marinaaaniram/go-chat-server/internal/model"
)

func (s *serv) SendMessage(ctx context.Context, chat *model.Message) error {
	err := s.messageRepository.SendMessage(ctx, chat)
	if err != nil {
		return err
	}

	return nil
}
