package chat

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/marinaaaniram/go-chat-server/internal/converter"
	desc "github.com/marinaaaniram/go-chat-server/pkg/chat_v1"
)

// Delete Chat in desc layer
func (i *Implementation) Delete(ctx context.Context, req *desc.DeleteRequest) (*emptypb.Empty, error) {
	err := i.chatService.Delete(ctx, converter.FromDescDeleteToChat(req))
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}
