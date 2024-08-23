package service

import (
	"context"

	"github.com/pawpawchat/chat/internal/domain/model"
)

func (s *Service) SendMessage(ctx context.Context, msg *model.Message) error {
	return s.msgRepo.SendMessage(ctx, msg)
}

func (s *Service) GetMessages(ctx context.Context, chatID uint64) (*[]model.Message, error) {
	return s.msgRepo.GetMessages(ctx, chatID)
}
