package converter

import (
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/marinaaaniram/go-chat-server/internal/model"
	desc "github.com/marinaaaniram/go-chat-server/pkg/chat_v1"
)

func FromChatToDesc(chat *model.Chat) *desc.Chat {
	var updatedAt *timestamppb.Timestamp
	if chat.UpdatedAt.Valid {
		updatedAt = timestamppb.New(chat.UpdatedAt.Time)
	}

	return &desc.Chat{
		Id:        chat.ID,
		Usernames: chat.Usernames,
		CreatedAt: timestamppb.New(chat.CreatedAt),
		UpdatedAt: updatedAt,
	}
}

func FromDescCreateToChat(req *desc.CreateChatRequest) *model.Chat {
	return &model.Chat{
		Usernames: req.GetUsernames(),
	}
}

func FromDescDeleteToChat(req *desc.DeleteChatRequest) *model.Chat {
	return &model.Chat{
		ID: req.GetId(),
	}
}
