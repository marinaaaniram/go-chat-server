package chat

import (
	"github.com/marinaaaniram/go-chat-server/internal/repository"
	"github.com/marinaaaniram/go-chat-server/internal/service"
)

type serv struct {
	chatRepository    repository.ChatRepository
	messageRepository repository.MessageRepository
}

// Create Chat service
func NewChatService(chatRepository repository.ChatRepository, messageRepository repository.MessageRepository) service.ChatService {
	return &serv{
		chatRepository:    chatRepository,
		messageRepository: messageRepository,
	}
}
