package chat

import (
	"context"
	"fmt"

	"github.com/marinaaaniram/go-common-platform/pkg/db"

	"github.com/marinaaaniram/go-chat-server/internal/errors"
)

// Create chat in repository layer
func (r *repo) Create(ctx context.Context) (int64, error) {
	// builderInsert :=
	// 	sq.Insert(tableName).
	// 		Columns(). // Нет необходимости указывать колонки, если они будут автогенерированы
	// 		Values().  // Нет значений для автоинкрементного id
	// 		PlaceholderFormat(sq.Dollar).
	// 		Suffix("RETURNING id")

	query := fmt.Sprintf("INSERT INTO %s DEFAULT VALUES RETURNING id", tableName)

	// query, args, err := builderInsert.ToSql()
	// if err != nil {
	// 	return 0, errors.ErrFailedToBuildQuery(err)
	// }

	q := db.Query{
		Name:     "chat_repository.Create",
		QueryRaw: query,
	}

	var chatId int64
	err := r.db.DB().QueryRowContext(ctx, q).Scan(&chatId)
	if err != nil {
		return 0, errors.ErrFailedToInsertQuery(err)
	}

	return chatId, nil
}
