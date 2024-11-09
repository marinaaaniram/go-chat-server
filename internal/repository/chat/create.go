package chat

import (
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/lib/pq"

	"github.com/marinaaaniram/go-chat-server/internal/client/db"
	"github.com/marinaaaniram/go-chat-server/internal/errors"
	"github.com/marinaaaniram/go-chat-server/internal/model"
	converterRepo "github.com/marinaaaniram/go-chat-server/internal/repository/chat/converter"
	modelRepo "github.com/marinaaaniram/go-chat-server/internal/repository/chat/model"
)

// Create chat in repository layer
func (r *repo) Create(ctx context.Context, chat *model.Chat) (*model.Chat, error) {
	if chat == nil {
		return nil, errors.ErrPointerIsNil("chat")
	}

	builderInsert := sq.Insert(tableName).
		PlaceholderFormat(sq.Dollar).
		Columns(usernamesColumn).
		Values(pq.Array(chat.Usernames)).
		Suffix(fmt.Sprintf("RETURNING %s, %s, %s, %s", idColumn, usernamesColumn, createdAtColumn, updatedAtColumn))

	query, args, err := builderInsert.ToSql()
	if err != nil {
		return nil, errors.ErrFailedToBuildQuery(err)
	}

	q := db.Query{
		Name:     "chat_repository.Create",
		QueryRaw: query,
	}

	var repoChat modelRepo.Chat
	err = r.db.DB().QueryRowContext(ctx, q, args...).Scan(&repoChat.ID, &repoChat.Usernames, &repoChat.CreatedAt, &repoChat.UpdatedAt)
	if err != nil {
		return nil, errors.ErrFailedToInsertQuery(err)
	}

	return (*model.Chat)(converterRepo.FromRepoToChat(&repoChat)), nil
}
