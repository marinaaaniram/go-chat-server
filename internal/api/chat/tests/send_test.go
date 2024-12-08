package tests

import (
	"context"
	"fmt"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/marinaaaniram/go-chat-server/internal/api/chat"
	"github.com/marinaaaniram/go-chat-server/internal/errors"
	"github.com/marinaaaniram/go-chat-server/internal/model"
	"github.com/marinaaaniram/go-chat-server/internal/service"
	serviceMocks "github.com/marinaaaniram/go-chat-server/internal/service/mocks"
	desc "github.com/marinaaaniram/go-chat-server/pkg/chat_v1"
)

func TestApiMessageCreate(t *testing.T) {
	t.Parallel()
	type chatServiceMockFunc func(mc *minimock.Controller) service.ChatService

	type args struct {
		ctx context.Context
		req *desc.SendMessageRequest
	}

	var (
		ctx = context.Background()
		mc  = minimock.NewController(t)

		chatId   = gofakeit.Int64()
		username = gofakeit.Name()
		text     = gofakeit.Sentence(10)

		serviceErr = fmt.Errorf("Service error")

		serviceReq = &model.Message{
			ChatId:   chatId,
			Username: username,
			Text:     text,
		}

		req = &desc.SendMessageRequest{
			ChatId: chatId,
			Message: &desc.Message{
				Username: username,
				Text:     text,
			},
		}

		res = &emptypb.Empty{}
	)

	defer t.Cleanup(mc.Finish)

	tests := []struct {
		name            string
		args            args
		want            *emptypb.Empty
		err             error
		chatServiceMock chatServiceMockFunc
	}{
		{
			name: "Success case",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: res,
			err:  nil,
			chatServiceMock: func(mc *minimock.Controller) service.ChatService {
				mock := serviceMocks.NewChatServiceMock(mc)
				mock.SendMessageMock.Expect(ctx, serviceReq).Return(nil)
				return mock
			},
		},
		{
			name: "Api nil pointer",
			args: args{
				ctx: ctx,
				req: nil,
			},
			want: nil,
			err:  errors.ErrPointerIsNil("req"),
			chatServiceMock: func(mc *minimock.Controller) service.ChatService {
				return nil
			},
		},
		{
			name: "Service error case",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: nil,
			err:  serviceErr,
			chatServiceMock: func(mc *minimock.Controller) service.ChatService {
				mock := serviceMocks.NewChatServiceMock(mc)
				mock.SendMessageMock.Expect(ctx, serviceReq).Return(nil)
				return mock
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			chatServiceMock := tt.chatServiceMock(mc)
			api := chat.NewChatImplementation(chatServiceMock)

			newID, err := api.Send(tt.args.ctx, tt.args.req)
			require.Equal(t, tt.err, err)
			require.Equal(t, tt.want, newID)
		})
	}
}
