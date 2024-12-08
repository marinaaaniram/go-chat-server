package chat

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgconn"
	"github.com/marinaaaniram/go-common-platform/pkg/db"

	"github.com/marinaaaniram/go-chat-server/internal/errors"
	"github.com/marinaaaniram/go-chat-server/internal/model"
	modelRepo "github.com/marinaaaniram/go-chat-server/internal/repository/message/model"
)

// Send message in repository layer
func (r *repo) Send(ctx context.Context, message *model.Message) error {
	if message == nil {
		return errors.ErrPointerIsNil("message")
	}

	builderInsert := sq.Insert(tableName).
		PlaceholderFormat(sq.Dollar).
		Columns(chatId, sentBy, text).
		Values(message.ChatId, message.Username, message.Text).
		Suffix("RETURNING id")

	query, args, err := builderInsert.ToSql()
	if err != nil {
		return errors.ErrFailedToBuildQuery(err)
	}

	q := db.Query{
		Name:     "message_repository.Send",
		QueryRaw: query,
	}

	var messageRepo modelRepo.Message
	err = r.db.DB().QueryRowContext(ctx, q, args...).Scan(&messageRepo.ID)
	if err != nil {
		if pgErr, ok := err.(*pgconn.PgError); ok && pgErr.Code == "23503" {
			return errors.ErrObjectNotFount("chat", message.ChatId)
		}
		return errors.ErrFailedToInsertQuery(err)
	}

	return nil
}
