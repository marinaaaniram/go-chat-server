package converter

import (
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/marinaaaniram/go-chat-server/internal/model"
	desc "github.com/marinaaaniram/go-chat-server/pkg/chat_v1"
)

// Convert Chat internal model to desc model
func FromChatToDesc(chat *model.Chat) *desc.Chat {
	if chat == nil {
		return nil
	}

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

// Convert desc CreateRequest fields to internal Chat model
func FromDescCreateToChat(req *desc.CreateRequest) *model.Chat {
	if req == nil {
		return nil
	}

	return &model.Chat{
		Usernames: req.GetUsernames(),
	}
}

// Convert desc DeleteRequest fields to internal Chat model
func FromDescDeleteToChat(req *desc.DeleteRequest) *model.Chat {
	if req == nil {
		return nil
	}

	return &model.Chat{
		ID: req.GetId(),
	}
}
