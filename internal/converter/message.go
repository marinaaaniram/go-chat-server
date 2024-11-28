package converter

import (
	"go-chat-server/internal/model"
	desc "go-chat-server/pkg/message_v1"
)

// Convert desc SendRequest fields to internal Message model
func FromDescToMessage(message *desc.SendRequest) *model.Message {
	if message == nil {
		return nil
	}

	return &model.Message{
		ChatId: message.ChatId,
		SentBy: message.SentBy,
		Text:   message.Text,
	}
}
