package chat

import (
	"context"

	"github.com/marinaaaniram/go-chat-server/internal/model"
)

func (s *serv) Delete(ctx context.Context, chat *model.Chat) error {
	err := s.chatRepository.Delete(ctx, chat)
	if err != nil {
		return err
	}

	return nil
}
