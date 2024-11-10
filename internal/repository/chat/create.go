package chat

import (
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/lib/pq"

	"github.com/marinaaaniram/go-chat-server/internal/errors"
	"github.com/marinaaaniram/go-chat-server/internal/model"
	"github.com/marinaaaniram/go-common-platform/pkg/db"
)

// Create chat in repository layer
func (r *repo) Create(ctx context.Context, chat *model.Chat) (int64, error) {
	if chat == nil {
		return 0, errors.ErrPointerIsNil("chat")
	}

	builderInsert := sq.Insert(tableName).
		PlaceholderFormat(sq.Dollar).
		Columns(usernamesColumn).
		Values(pq.Array(chat.Usernames)).
		Suffix(fmt.Sprintf("RETURNING %s", idColumn))

	query, args, err := builderInsert.ToSql()
	if err != nil {
		return 0, errors.ErrFailedToBuildQuery(err)
	}

	q := db.Query{
		Name:     "chat_repository.Create",
		QueryRaw: query,
	}

	var chatId int64
	err = r.db.DB().QueryRowContext(ctx, q, args...).Scan(&chatId)
	if err != nil {
		return 0, errors.ErrFailedToInsertQuery(err)
	}

	return chatId, nil
}
