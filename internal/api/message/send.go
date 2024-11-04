package message

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/marinaaaniram/go-chat-server/internal/converter"
	desc "github.com/marinaaaniram/go-chat-server/pkg/message_v1"
)

// Update Message in desc layer
func (i *Implementation) Send(ctx context.Context, req *desc.SendRequest) (*emptypb.Empty, error) {
	err := i.messageService.Send(ctx, converter.FromDescToMessage(req))
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}
