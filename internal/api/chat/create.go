package chat

import (
	"context"

	"github.com/marinaaaniram/go-chat-server/internal/converter"
	desc "github.com/marinaaaniram/go-chat-server/pkg/chat_v1"
)

func (i *Implementation) Create(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	chatDesc, err := i.chatService.Create(ctx, converter.FromDescCreateToChat(req))
	if err != nil {
		return nil, err
	}

	return &desc.CreateResponse{
		Chat: chatDesc,
	}, nil
}
