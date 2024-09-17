package service

import (
	"context"

	"github.com/pawpawchat/chat/internal/domain/model"
)

func (s *Service) SendMessage(ctx context.Context, msg *model.Message) error {
	if err := s.msgRepo.SendMessage(ctx, msg); err != nil {
		return err
	}

	s.msgChan <- msg
	return nil
}

func (s *Service) GetMessages(ctx context.Context, chatID int64) ([]*model.Message, error) {
	return s.msgRepo.GetMessages(ctx, chatID)
}
