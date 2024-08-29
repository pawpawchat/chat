package adapter

import (
	"context"
	"time"

	"github.com/pawpawchat/chat/api/pb"
	"github.com/pawpawchat/chat/internal/domain/model"
)

type messageProvider interface {
	SendMessage(context.Context, *model.Message) error
	GetMessages(context.Context, int64) (*[]model.Message, error)
}

func SendMessageAdapter(ctx context.Context, provider messageProvider, req *pb.SendMessageRequest) (*pb.SendMessageResponse, error) {
	var sentAt time.Time

	if req.SentAt != nil {
		sentAt = req.SentAt.AsTime()
	} else {
		sentAt = time.Now()
	}

	msg := &model.Message{
		ChatID:         req.ChatId,
		Body:           req.Body,
		SenderID:       req.SenderId,
		SenderUsername: req.SenderUsername,
		SentAt:         sentAt,
	}

	if err := provider.SendMessage(ctx, msg); err != nil {
		return nil, err
	}

	return &pb.SendMessageResponse{Message: msg.ToPb()}, nil
}

func GetMessagesAdapter(ctx context.Context, provider messageProvider, req *pb.GetMessagesRequest) (*pb.GetMessagesResponse, error) {
	chatID := req.ChatId

	messages, err := provider.GetMessages(ctx, chatID)
	if err != nil {
		return nil, err
	}

	messagespb := make([]*pb.Message, 0)

	for _, m := range *messages {
		messagespb = append(messagespb, m.ToPb())
	}

	return &pb.GetMessagesResponse{ChatId: chatID, Messages: messagespb}, nil
}
