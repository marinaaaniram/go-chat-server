package chat

import (
	"context"
	"log"

	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4/pgxpool"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/marinaaaniram/go-chat-server/internal/model"
	"github.com/marinaaaniram/go-chat-server/internal/repository"
	modelRepo "github.com/marinaaaniram/go-chat-server/internal/repository/message/model"
)

const (
	tableName = "message"

	idColumn        = "id"
	chatId          = "chat_id"
	sentBy          = "sent_by"
	text            = "text"
	createdAtColumn = "created_at"
	updatedAtColumn = "updated_at"
)

type repo struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) repository.MessageRepository {
	return &repo{db: db}
}

func (r *repo) SendMessage(ctx context.Context, message *model.Message) error {
	builderInsert := sq.Insert(tableName).
		PlaceholderFormat(sq.Dollar).
		Columns(chatId, sentBy, text).
		Values(message.ChatId, message.SentBy, message.Text).
		Suffix("RETURNING id")

	query, args, err := builderInsert.ToSql()
	if err != nil {
		return status.Errorf(codes.Internal, "Failed to build query: %v", err)
	}

	var messageRepo modelRepo.Message
	err = r.db.QueryRow(ctx, query, args...).Scan(&messageRepo.ID)
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
