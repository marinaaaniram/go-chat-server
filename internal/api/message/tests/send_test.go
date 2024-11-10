package tests

import (
	"context"
	"fmt"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/marinaaaniram/go-chat-server/internal/api/message"
	"github.com/marinaaaniram/go-chat-server/internal/errors"
	"github.com/marinaaaniram/go-chat-server/internal/model"
	"github.com/marinaaaniram/go-chat-server/internal/service"
	serviceMocks "github.com/marinaaaniram/go-chat-server/internal/service/mocks"
	desc "github.com/marinaaaniram/go-chat-server/pkg/message_v1"
)

func TestCreate(t *testing.T) {
	t.Parallel()
	type messageServiceMockFunc func(mc *minimock.Controller) service.MessageService

	type args struct {
		ctx context.Context
		req *desc.SendRequest
	}

	var (
		ctx = context.Background()
		mc  = minimock.NewController(t)

		chatId = gofakeit.Int64()
		sendBy = gofakeit.Name()
		text   = gofakeit.Sentence(10)

		serviceErr = fmt.Errorf("service error")

		serviceReq = &model.Message{
			ChatId: chatId,
			SentBy: sendBy,
			Text:   text,
		}

		req = &desc.SendRequest{
			ChatId: chatId,
			SentBy: sendBy,
			Text:   text,
		}

		res = &emptypb.Empty{}
	)

	defer t.Cleanup(mc.Finish)

	tests := []struct {
		name               string
		args               args
		want               *emptypb.Empty
		err                error
		messageServiceMock messageServiceMockFunc
	}{
		{
			name: "success case",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: res,
			err:  nil,
			messageServiceMock: func(mc *minimock.Controller) service.MessageService {
				mock := serviceMocks.NewMessageServiceMock(mc)
				mock.SendMock.Expect(ctx, serviceReq).Return(nil)
				return mock
			},
		},
		{
			name: "api nil pointer",
			args: args{
				ctx: ctx,
				req: nil,
			},
			want: nil,
			err:  errors.ErrPointerIsNil("req"),
			messageServiceMock: func(mc *minimock.Controller) service.MessageService {
				return nil
			},
		},
		{
			name: "service error case",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: nil,
			err:  serviceErr,
			messageServiceMock: func(mc *minimock.Controller) service.MessageService {
				mock := serviceMocks.NewMessageServiceMock(mc)
				mock.SendMock.Expect(ctx, serviceReq).Return(serviceErr)
				return mock
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			messageServiceMock := tt.messageServiceMock(mc)
			api := message.NewMessageImplementation(messageServiceMock)

			newID, err := api.Send(tt.args.ctx, tt.args.req)
			require.Equal(t, tt.err, err)
			require.Equal(t, tt.want, newID)
		})
	}
}
