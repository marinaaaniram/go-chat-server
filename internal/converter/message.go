package converter

import (
	"github.com/marinaaaniram/go-chat-server/internal/model"
	desc "github.com/marinaaaniram/go-chat-server/pkg/message_v1"
)

func FromDescToMessage(message *desc.SendRequest) *model.Message {
	return &model.Message{
		ChatId: message.ChatId,
		SentBy: message.SentBy,
		Text:   message.Text,
	}
}
