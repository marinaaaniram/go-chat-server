package converter

import (
	"github.com/marinaaaniram/go-chat-server/internal/model"
	desc "github.com/marinaaaniram/go-chat-server/pkg/chat_v1"
)

func FromDescToMessage(message *desc.SendMessageRequest) *model.Message {
	return &model.Message{
		ChatId: message.ChatId,
		SentBy: message.SentBy,
		Text:   message.Text,
	}
}
