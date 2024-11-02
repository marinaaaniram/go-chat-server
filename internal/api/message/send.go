package message

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/marinaaaniram/go-chat-server/internal/converter"
	desc "github.com/marinaaaniram/go-chat-server/pkg/message_v1"
)

func (i *Implementation) SendMessage(ctx context.Context, req *desc.SendRequest) (*emptypb.Empty, error) {
	err := i.messageService.SendMessage(ctx, converter.FromDescToMessage(req))
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}