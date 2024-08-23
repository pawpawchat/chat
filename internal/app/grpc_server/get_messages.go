package grpcserver

import (
	"context"

	"github.com/pawpawchat/chat/api/pb"
	"github.com/pawpawchat/chat/internal/domain/model"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *ChatGRPCServer) GetMessages(ctx context.Context, req *pb.GetMessagesRequest) (*pb.GetMessagesResponse, error) {
	chatID := req.ChatId

	messages, err := s.usecase.GetMessages(ctx, chatID)
	if err != nil {
		return nil, err
	}

	pbmsgs := make([]*pb.Message, 0)

	parseMessage := func(m model.Message) *pb.Message {
		return &pb.Message{
			MessageId: m.MessageID,
			Body:      m.Body,
			SenderId:  m.SenderID,
			Username:  m.SenderUsername,
			SentAt:    timestamppb.New(m.SentAt),
		}
	}

	for _, m := range *messages {
		pbmsgs = append(pbmsgs, parseMessage(m))
	}

	return &pb.GetMessagesResponse{ChatId: chatID, Messages: pbmsgs}, nil
}
