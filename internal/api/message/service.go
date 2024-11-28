package message

import (
	"go-chat-server/internal/service"
	desc "go-chat-server/pkg/message_v1"
)

type Implementation struct {
	desc.UnimplementedMessageV1Server
	messageService service.MessageService
}

// Create Message implementation
func NewMessageImplementation(messageService service.MessageService) *Implementation {
	return &Implementation{
		messageService: messageService,
	}
}
