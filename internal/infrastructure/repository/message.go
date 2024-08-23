package repository

import (
	"context"

	"github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	"github.com/pawpawchat/chat/internal/domain/model"
)

type MessageRepository struct {
	db *sqlx.DB
}

func NewMessageRepository(db *sqlx.DB) *MessageRepository {
	return &MessageRepository{db}
}

func (r *MessageRepository) SendMessage(ctx context.Context, msg *model.Message) error {
	sql, args := squirrel.Insert("messages").
		Columns("chat_id", "sender_id", "sender_username", "body", "sent_at").
		Values(msg.ChatID, msg.SenderID, msg.SenderUsername, msg.Body, msg.SentAt).
		PlaceholderFormat(squirrel.Dollar).
		MustSql()

	_, err := r.db.ExecContext(ctx, sql, args...)
	return err
}

func (r *MessageRepository) GetMessages(ctx context.Context, chatID uint64) (*[]model.Message, error) {
	sql, args := squirrel.Select("*").
		From("messages").Where(squirrel.Eq{"chat_id": chatID}).
		PlaceholderFormat(squirrel.Dollar).
		MustSql()

	msgs := new([]model.Message)
	return msgs, r.db.SelectContext(ctx, msgs, sql, args...)
}
