package chat

import (
	"context"

	"github.com/marinaaaniram/go-chat-server/internal/converter"
	"github.com/marinaaaniram/go-chat-server/internal/errors"
	desc "github.com/marinaaaniram/go-chat-server/pkg/chat_v1"
)

// Create Chat in desc layer
func (i *Implementation) Create(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	if req == nil {
		return nil, errors.ErrPointerIsNil("req")
	}

	chatDesc, err := i.chatService.Create(ctx, converter.FromDescCreateToChat(req))
	if err != nil {
		return nil, err
	}

	return &desc.CreateResponse{
		Chat: chatDesc,
	}, nil
}
