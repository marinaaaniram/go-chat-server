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
	"github.com/marinaaaniram/go-chat-server/internal/service"
	serviceMocks "github.com/marinaaaniram/go-chat-server/internal/service/mocks"
	desc "github.com/marinaaaniram/go-chat-server/pkg/chat_v1"
)

func TestApiChatCreate(t *testing.T) {
	t.Parallel()
	type chatServiceMockFunc func(mc *minimock.Controller) service.ChatService

	type args struct {
		ctx context.Context
	}

	var (
		ctx = context.Background()
		mc  = minimock.NewController(t)

		id = gofakeit.Int64()

		serviceErr = fmt.Errorf("Service error")

		res = &desc.CreateResponse{
			Id: id,
		}
	)

	defer t.Cleanup(mc.Finish)

	tests := []struct {
		name            string
		args            args
		want            *desc.CreateResponse
		err             error
		chatServiceMock chatServiceMockFunc
	}{
		{
			name: "Success case",
			args: args{
				ctx: ctx,
			},
			want: res,
			err:  nil,
			chatServiceMock: func(mc *minimock.Controller) service.ChatService {
				mock := serviceMocks.NewChatServiceMock(mc)
				mock.CreateMock.Expect(ctx).Return(id, nil)
				return mock
			},
		},
		{
			name: "Service error case",
			args: args{
				ctx: ctx,
			},
			want: nil,
			err:  serviceErr,
			chatServiceMock: func(mc *minimock.Controller) service.ChatService {
				mock := serviceMocks.NewChatServiceMock(mc)
				mock.CreateMock.Expect(ctx).Return(0, serviceErr)
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

			newID, err := api.Create(tt.args.ctx, &emptypb.Empty{})
			require.Equal(t, tt.err, err)
			require.Equal(t, tt.want, newID)
		})
	}
}
