package repository

import (
	"context"

	"github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	"github.com/pawpawchat/chat/internal/domain/model"
)

type MemberRepository struct {
	db *sqlx.DB
}

func NewMemberRepository(db *sqlx.DB) *MemberRepository {
	return &MemberRepository{db}
}

func (s *MemberRepository) AddMember(ctx context.Context, member *model.Member) error {
	sql, args := squirrel.Insert("chat_members").
		Columns("chat_id", "member_id", "username", "role").
		Values(member.ChatID, member.MemberID, member.Username, member.Role).
		PlaceholderFormat(squirrel.Dollar).
		MustSql()

	_, err := s.db.ExecContext(ctx, sql, args...)
	return err
}

func (s *MemberRepository) GetMembers(ctx context.Context, chatID int64) (*[]model.Member, error) {
	sql, args := squirrel.Select("*").
		From("chat_members").
		Where(squirrel.Eq{"chat_id": chatID}).
		PlaceholderFormat(squirrel.Dollar).
		MustSql()

	members := new([]model.Member)
	return members, s.db.SelectContext(ctx, members, sql, args...)
}

func (s *MemberRepository) GetMember(ctx context.Context, chatID int64, memberID int64) (*model.Member, error) {
	sql, args := squirrel.Select("*").
		From("chat_members").
		Where(squirrel.Eq{"chat_id": chatID, "member_id": memberID}).
		PlaceholderFormat(squirrel.Dollar).
		MustSql()

	member := new(model.Member)
	return member, s.db.GetContext(ctx, member, sql, args...)
}
