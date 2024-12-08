package chat

import (
	"sync"

	"github.com/marinaaaniram/go-chat-server/internal/service"
	desc "github.com/marinaaaniram/go-chat-server/pkg/chat_v1"
)

type Chat struct {
	streams map[string]desc.ChatV1_ConnectChatServer
	m       sync.RWMutex
}

type Implementation struct {
	desc.UnimplementedChatV1Server
	chatService service.ChatService

	chats  map[int64]*Chat
	mxChat sync.RWMutex

	channels  map[int64]chan *desc.Message
	mxChannel sync.RWMutex
}

// Create Chat implementation
func NewChatImplementation(chatService service.ChatService) *Implementation {
	return &Implementation{
		chatService: chatService,
		chats:       make(map[int64]*Chat),
		channels:    make(map[int64]chan *desc.Message),
	}
}
