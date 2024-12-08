package chat

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/marinaaaniram/go-common-platform/pkg/db"

	"github.com/marinaaaniram/go-chat-server/internal/errors"
)

// Delete chat in repository layer
func (r *repo) Delete(ctx context.Context, chatId int64) error {
	builder := sq.Delete(tableName).
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{idColumn: chatId})

	query, args, err := builder.ToSql()
	if err != nil {
		return errors.ErrFailedToBuildQuery(err)
	}

	q := db.Query{
		Name:     "chat_repository.Delete",
		QueryRaw: query,
	}

	_, err = r.db.DB().ExecContext(ctx, q, args...)
	if err != nil {
		return errors.ErrFailedToDeleteQuery(err)
	}

	return nil
}
