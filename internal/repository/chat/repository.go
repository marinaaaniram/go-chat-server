package chat

import (
	"context"
	"fmt"
	"log"

	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/lib/pq"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/marinaaaniram/go-chat-server/internal/model"
	"github.com/marinaaaniram/go-chat-server/internal/repository"
	converterRepo "github.com/marinaaaniram/go-chat-server/internal/repository/chat/converter"
	modelRepo "github.com/marinaaaniram/go-chat-server/internal/repository/chat/model"
)

const (
	tableName = "chat"

	idColumn        = "id"
	usernamesColumn = "usernames"
	createdAtColumn = "created_at"
	updatedAtColumn = "updated_at"
)

type repo struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) repository.ChatRepository {
	return &repo{db: db}
}

func (r *repo) Create(ctx context.Context, chat *model.Chat) (*model.Chat, error) {
	builderInsert := sq.Insert(tableName).
		PlaceholderFormat(sq.Dollar).
		Columns(usernamesColumn).
		Values(pq.Array(chat.Usernames)).
		Suffix(fmt.Sprintf("RETURNING %s, %s, %s, %s", idColumn, usernamesColumn, createdAtColumn, updatedAtColumn))

	query, args, err := builderInsert.ToSql()
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to build query: %v", err)
	}

	var repoChat modelRepo.Chat
	err = r.db.QueryRow(ctx, query, args...).Scan(&repoChat.ID, &repoChat.Usernames, &repoChat.CreatedAt, &repoChat.UpdatedAt)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to insert chat: %v", err)
	}

	return (*model.Chat)(converterRepo.FromRepoToChat(&repoChat)), nil
}

func (r *repo) Delete(ctx context.Context, chat *model.Chat) error {
	builderSelect := sq.Select("COUNT(*)").
		From(tableName).
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{idColumn: chat.ID})

	selectQuery, args, err := builderSelect.ToSql()
	if err != nil {
		return status.Errorf(codes.Internal, "Failed to build select query: %v", err)
	}

	var count int
	err = r.db.QueryRow(ctx, selectQuery, args...).Scan(&count)
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

	_, err = r.db.Exec(ctx, query, args...)
	if err != nil {
		return status.Errorf(codes.Internal, "Failed to delete chat: %v", err)
	}

	log.Printf("Chat with id %d deleted", chat.ID)

	return nil
}
