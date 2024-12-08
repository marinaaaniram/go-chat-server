package tests

import (
	"context"
	"fmt"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"

	"github.com/marinaaaniram/go-chat-server/internal/model"
	"github.com/marinaaaniram/go-chat-server/internal/repository"
	repoMocks "github.com/marinaaaniram/go-chat-server/internal/repository/mocks"
	"github.com/marinaaaniram/go-chat-server/internal/service/chat"
)

func TestServiceChatCreate(t *testing.T) {
	t.Parallel()
	type chatRepositoryMockFunc func(mc *minimock.Controller) repository.ChatRepository
	type messageRepositoryMockFunc func(mc *minimock.Controller) repository.MessageRepository

	type args struct {
		ctx context.Context
		req *model.Chat
	}

	var (
		ctx = context.Background()
		mc  = minimock.NewController(t)

		id = gofakeit.Int64()

		repoErr = fmt.Errorf("Repo error")
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
			},
			want: id,
			err:  nil,
			chatRepositoryMock: func(mc *minimock.Controller) repository.ChatRepository {
				mock := repoMocks.NewChatRepositoryMock(mc)
				mock.CreateMock.Expect(ctx).Return(id, nil)
				return mock
			},
		},
		{
			name: "Service error case",
			args: args{
				ctx: ctx,
			},
			want: 0,
			err:  repoErr,
			chatRepositoryMock: func(mc *minimock.Controller) repository.ChatRepository {
				mock := repoMocks.NewChatRepositoryMock(mc)
				mock.CreateMock.Expect(ctx).Return(0, repoErr)
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

			newID, err := service.Create(tt.args.ctx)
			require.Equal(t, tt.err, err)
			require.Equal(t, tt.want, newID)
		})
	}
}
