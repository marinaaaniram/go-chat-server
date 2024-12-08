package chat

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/marinaaaniram/go-chat-server/internal/converter"
	"github.com/marinaaaniram/go-chat-server/internal/errors"
	desc "github.com/marinaaaniram/go-chat-server/pkg/chat_v1"
)

// Update Message in desc layer
func (i *Implementation) Send(ctx context.Context, req *desc.SendMessageRequest) (*emptypb.Empty, error) {
	if req == nil {
		return nil, errors.ErrPointerIsNil("req")
	}

	err := i.chatService.SendMessage(ctx, converter.FromDescToMessage(req))
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}
