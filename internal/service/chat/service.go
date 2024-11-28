package chat

import (
	"go-chat-server/internal/repository"
	"go-chat-server/internal/service"
)

type serv struct {
	chatRepository repository.ChatRepository
}

// Create Chat service
func NewChatService(chatRepository repository.ChatRepository) service.ChatService {
	return &serv{
		chatRepository: chatRepository,
	}
}
