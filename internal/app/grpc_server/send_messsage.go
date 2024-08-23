package grpcserver

import (
	"context"
	"time"

	"github.com/pawpawchat/chat/api/pb"
	"github.com/pawpawchat/chat/internal/domain/model"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *ChatGRPCServer) SendMessage(ctx context.Context, req *pb.SendMessageRequest) (*pb.SendMessageResponse, error) {
	msg, err := parseSendMessageRequest(req)
	if err != nil {
		return nil, err
	}

	if err := s.service.SendMessage(ctx, msg); err != nil {
		return nil, err
	}

	pbmsg := parseMessage(*msg)
	return &pb.SendMessageResponse{Message: pbmsg}, nil
}

func parseMessage(m model.Message) *pb.Message {
	return &pb.Message{
		MessageId: m.MessageID,
		Body:      m.Body,
		IsDeleted: m.IsDeleted,
		SenderId:  m.SenderID,
		Username:  m.SenderUsername,
		ChatId:    m.ChatID,
		SentAt:    timestamppb.New(m.SentAt),
	}
}

func parseSendMessageRequest(pb *pb.SendMessageRequest) (*model.Message, error) {
	var sentAt time.Time

	switch pb.SentAt {
	case nil:
		sentAt = time.Now()
	default:
		sentAt = pb.SentAt.AsTime()
	}

	return &model.Message{
		ChatID:         pb.ChatId,
		Body:           pb.Body,
		SenderID:       pb.SenderId,
		SenderUsername: pb.SenderUsername,
		SentAt:         sentAt,
	}, nil
}
