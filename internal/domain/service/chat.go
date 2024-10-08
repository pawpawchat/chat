package service

import (
	"context"

	"github.com/pawpawchat/chat/internal/domain/model"
)

func (s *Service) CreateChat(ctx context.Context, owner *model.Member, chat *model.Chat) error {
	if err := s.chatRepo.CreateChat(ctx, chat); err != nil {
		return nil
	}

	owner.ChatID = chat.ChatID
	return s.memberRepo.AddMember(ctx, owner)
}

func (s *Service) GetChat(ctx context.Context, chatID int64) (*model.Chat, error) {
	return s.chatRepo.GetChat(ctx, chatID)
}

func (s *Service) GetAllChats(ctx context.Context, profileID int64) ([]*model.Chat, error) {
	return s.chatRepo.GetAllChats(ctx, profileID)
}
