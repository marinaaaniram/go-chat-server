package chat

import (
	"context"
)

// Create chat in service layer
func (s *serv) Create(ctx context.Context) (int64, error) {
	chatId, err := s.chatRepository.Create(ctx)
	if err != nil {
		return 0, err
	}

	return chatId, nil
}
