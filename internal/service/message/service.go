package message

import (
	"github.com/marinaaaniram/go-chat-server/internal/repository"
	"github.com/marinaaaniram/go-chat-server/internal/service"
)

type serv struct {
	messageRepository repository.MessageRepository
}

// Create Message service
func NewMessageService(messageRepository repository.MessageRepository) service.MessageService {
	return &serv{
		messageRepository: messageRepository,
	}
}
