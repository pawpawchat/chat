package repository

import (
	"context"

	"github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	"github.com/pawpawchat/chat/internal/domain/model"
)

var (
	ROLE_OWNER  = "owner"
	ROLE_MEMBER = "member"
)

type ChatRepository struct {
	db *sqlx.DB
}

func NewChatRepository(db *sqlx.DB) *ChatRepository {
	return &ChatRepository{db}
}

func (s *ChatRepository) CreateChat(ctx context.Context, chat *model.Chat) error {
	sql, args := squirrel.Insert("chats").
		Columns("title", "created_at").
		Values(chat.Title, chat.CreatedAt).
		Suffix("RETURNING chat_id").
		PlaceholderFormat(squirrel.Dollar).
		MustSql()

	return s.db.QueryRowContext(ctx, sql, args...).Scan(&chat.ChatID)
}

func (s *ChatRepository) GetChat(ctx context.Context, chatID int64) (*model.Chat, error) {
	sql, args := squirrel.Select("*").
		From("chats").
		Where(squirrel.Eq{"chat_id": chatID}).
		PlaceholderFormat(squirrel.Dollar).
		MustSql()

	chat := new(model.Chat)
	return chat, s.db.Get(chat, sql, args...)
}
