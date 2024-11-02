package converter

import (
	"github.com/marinaaaniram/go-chat-server/internal/repository/chat/model"
	modelRepo "github.com/marinaaaniram/go-chat-server/internal/repository/chat/model"
)

func FromRepoToChat(chat *modelRepo.Chat) *model.Chat {
	return &model.Chat{
		ID:        chat.ID,
		Usernames: chat.Usernames,
		CreatedAt: chat.CreatedAt,
		UpdatedAt: chat.UpdatedAt,
	}
}
