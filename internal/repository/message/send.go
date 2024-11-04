package chat

import (
	"context"
	"log"

	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgconn"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/marinaaaniram/go-chat-server/internal/client/db"
	"github.com/marinaaaniram/go-chat-server/internal/model"
	modelRepo "github.com/marinaaaniram/go-chat-server/internal/repository/message/model"
)

// Send message in repository layer
func (r *repo) Send(ctx context.Context, message *model.Message) error {
	builderInsert := sq.Insert(tableName).
		PlaceholderFormat(sq.Dollar).
		Columns(chatId, sentBy, text).
		Values(message.ChatId, message.SentBy, message.Text).
		Suffix("RETURNING id")

	query, args, err := builderInsert.ToSql()
	if err != nil {
		return status.Errorf(codes.Internal, "Failed to build query: %v", err)
	}

	q := db.Query{
		Name:     "message_repository.Send",
		QueryRaw: query,
	}

	var messageRepo modelRepo.Message
	err = r.db.DB().QueryRowContext(ctx, q, args...).Scan(&messageRepo.ID)
	if err != nil {
		if pgErr, ok := err.(*pgconn.PgError); ok && pgErr.Code == "23503" {
			return status.Errorf(codes.InvalidArgument, "Chat with id %d not found or %v", message.ChatId, err)
		} else {
			return status.Errorf(codes.Internal, "Failed to insert message: %v", err)
		}
	}

	log.Printf("Sent message with id: %d", messageRepo.ID)

	return nil
}
