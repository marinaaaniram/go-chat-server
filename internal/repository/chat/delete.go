package chat

import (
	"context"
	"log"

	sq "github.com/Masterminds/squirrel"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/marinaaaniram/go-chat-server/internal/client/db"
	"github.com/marinaaaniram/go-chat-server/internal/model"
)

// Delete chat in repository layer
func (r *repo) Delete(ctx context.Context, chat *model.Chat) error {
	builderSelect := sq.Select("COUNT(*)").
		From(tableName).
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{idColumn: chat.ID})

	selectQuery, args, err := builderSelect.ToSql()
	if err != nil {
		return status.Errorf(codes.Internal, "Failed to build select query: %v", err)
	}

	selectQ := db.Query{
		Name:     "chat_repository.SelectId",
		QueryRaw: selectQuery,
	}

	var count int
	err = r.db.DB().QueryRowContext(ctx, selectQ, args...).Scan(&count)
	if err != nil {
		return status.Errorf(codes.Internal, "Failed to select chat: %v", err)
	}

	if count == 0 {
		return status.Errorf(codes.NotFound, "Chat with id %d not found", chat.ID)
	}

	builderDelete := sq.Delete(tableName).
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{idColumn: chat.ID})

	query, args, err := builderDelete.ToSql()
	if err != nil {
		return status.Errorf(codes.Internal, "Failed to build delete query: %v", err)
	}

	q := db.Query{
		Name:     "chat_repository.Delete",
		QueryRaw: query,
	}

	_, err = r.db.DB().ExecContext(ctx, q, args...)
	if err != nil {
		return status.Errorf(codes.Internal, "Failed to delete chat: %v", err)
	}

	log.Printf("Chat with id %d deleted", chat.ID)

	return nil
}
