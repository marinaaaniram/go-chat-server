package chat

import (
	"github.com/marinaaaniram/go-chat-server/internal/repository"
	"github.com/marinaaaniram/go-chat-server/internal/service"
)

type serv struct {
	chatRepository repository.ChatRepository
}

// Create Chat service
func NewService(chatRepository repository.ChatRepository) service.ChatService {
	return &serv{
		chatRepository: chatRepository,
	}
}
