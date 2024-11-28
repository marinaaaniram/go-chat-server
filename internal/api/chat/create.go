package chat

import (
	"context"

	"go-chat-server/internal/converter"
	"go-chat-server/internal/errors"
	desc "go-chat-server/pkg/chat_v1"
)

// Create Chat in desc layer
func (i *Implementation) Create(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	if req == nil {
		return nil, errors.ErrPointerIsNil("req")
	}

	chatId, err := i.chatService.Create(ctx, converter.FromDescCreateToChat(req))
	if err != nil {
		return nil, err
	}

	return converter.FromChatIdToDescCreate(chatId), nil
}
