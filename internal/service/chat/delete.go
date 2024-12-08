package chat

import (
	"context"
)

// Delete chat in service layer
func (s *serv) Delete(ctx context.Context, chatId int64) error {
	err := s.chatRepository.Delete(ctx, chatId)
	if err != nil {
		return err
	}

	return nil
}
