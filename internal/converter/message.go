package converter

import (
	"github.com/marinaaaniram/go-chat-server/internal/model"
	desc "github.com/marinaaaniram/go-chat-server/pkg/chat_v1"
)

// Convert desc SendRequest fields to internal Message model
func FromDescToMessage(message *desc.SendMessageRequest) *model.Message {
	if message == nil {
		return nil
	}

	return &model.Message{
		ChatId:   message.ChatId,
		Username: message.Message.Username,
		Text:     message.Message.Text,
	}
}
