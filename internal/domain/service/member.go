package service

import (
	"context"

	"github.com/pawpawchat/chat/internal/domain/model"
)

func (s *Service) AddMember(ctx context.Context, member *model.Member) error {
	return s.memberRepo.AddMember(ctx, member)
}

func (s *Service) GetMembers(ctx context.Context, chatID int64) (*[]model.Member, error) {
	return s.memberRepo.GetMembers(ctx, chatID)
}

func (s *Service) GetMember(ctx context.Context, chatID int64, memberId int64) (*model.Member, error) {
	return s.memberRepo.GetMember(ctx, chatID, memberId)
}
