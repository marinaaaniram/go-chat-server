package message

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"

	"go-chat-server/internal/converter"
	"go-chat-server/internal/errors"
	desc "go-chat-server/pkg/message_v1"
)

// Update Message in desc layer
func (i *Implementation) Send(ctx context.Context, req *desc.SendRequest) (*emptypb.Empty, error) {
	if req == nil {
		return nil, errors.ErrPointerIsNil("req")
	}

	err := i.messageService.Send(ctx, converter.FromDescToMessage(req))
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}
