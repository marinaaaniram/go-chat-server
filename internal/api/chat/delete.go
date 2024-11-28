package chat

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"

	"go-chat-server/internal/converter"
	"go-chat-server/internal/errors"
	desc "go-chat-server/pkg/chat_v1"
)

// Delete Chat in desc layer
func (i *Implementation) Delete(ctx context.Context, req *desc.DeleteRequest) (*emptypb.Empty, error) {
	if req == nil {
		return nil, errors.ErrPointerIsNil("req")
	}

	err := i.chatService.Delete(ctx, converter.FromDescDeleteToChat(req))
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}
