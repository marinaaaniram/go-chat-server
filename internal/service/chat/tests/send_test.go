package tests

import (
	"context"
	"fmt"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"

	"github.com/marinaaaniram/go-chat-server/internal/errors"
	"github.com/marinaaaniram/go-chat-server/internal/model"
	"github.com/marinaaaniram/go-chat-server/internal/repository"
	repoMocks "github.com/marinaaaniram/go-chat-server/internal/repository/mocks"
	"github.com/marinaaaniram/go-chat-server/internal/service/chat"
)

func TestServiceMessageSend(t *testing.T) {
	t.Parallel()
	type chatRepositoryMockFunc func(mc *minimock.Controller) repository.ChatRepository
	type messageRepositoryMockFunc func(mc *minimock.Controller) repository.MessageRepository

	type args struct {
		ctx context.Context
		req *model.Message
	}

	var (
		ctx = context.Background()
		mc  = minimock.NewController(t)

		id = gofakeit.Int64()

		repoErr = fmt.Errorf("Repo error")

		chatId   = gofakeit.Int64()
		username = gofakeit.Name()
		text     = gofakeit.Sentence(10)

		req = &model.Message{
			ChatId:   chatId,
			Username: username,
			Text:     text,
		}
	)
	defer t.Cleanup(mc.Finish)

	tests := []struct {
		name                  string
		args                  args
		want                  int64
		err                   error
		chatRepositoryMock    chatRepositoryMockFunc
		messageRepositoryMock messageRepositoryMockFunc
	}{
		{
			name: "Success case",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: id,
			err:  nil,
			messageRepositoryMock: func(mc *minimock.Controller) repository.MessageRepository {
				mock := repoMocks.NewMessageRepositoryMock(mc)
				mock.SendMock.Expect(ctx, req).Return(nil)
				return mock
			},
		},
		{
			name: "Api nil pointer",
			args: args{
				ctx: ctx,
				req: nil,
			},
			want: 0,
			err:  errors.ErrPointerIsNil("message"),
			messageRepositoryMock: func(mc *minimock.Controller) repository.MessageRepository {
				return nil
			},
		},
		{
			name: "Service error case",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: 0,
			err:  repoErr,
			messageRepositoryMock: func(mc *minimock.Controller) repository.MessageRepository {
				mock := repoMocks.NewMessageRepositoryMock(mc)
				mock.SendMock.Expect(ctx, req).Return(repoErr)
				return mock
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			chatRepoMock := tt.chatRepositoryMock(mc)
			messageRepoMock := tt.messageRepositoryMock(mc)
			service := chat.NewChatService(chatRepoMock, messageRepoMock)

			err := service.SendMessage(tt.args.ctx, tt.args.req)
			require.Equal(t, tt.err, err)
		})
	}
}
