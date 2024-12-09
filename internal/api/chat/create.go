package chat

import (
	"context"

	desc "github.com/marinaaaniram/go-chat-server/pkg/chat_v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

// Create Chat in desc layer
func (i *Implementation) Create(ctx context.Context, _ *emptypb.Empty) (*desc.CreateResponse, error) {
	chatId, err := i.chatService.Create(ctx)
	if err != nil {
		return nil, err
	}

	i.channels[chatId] = make(chan *desc.Message, 100)

	return &desc.CreateResponse{
		Id: chatId,
	}, nil
}
