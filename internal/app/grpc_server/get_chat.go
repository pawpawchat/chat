package grpcserver

import (
	"context"

	"github.com/pawpawchat/chat/api/pb"
	"github.com/pawpawchat/chat/internal/domain/model"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *ChatGRPCServer) GetChat(ctx context.Context, req *pb.GetChatRequest) (*pb.GetChatResponse, error) {
	chatID, _ := req.ChatId, req.MemberId

	chat, err := s.service.GetChat(ctx, chatID)
	if err != nil {
		return nil, err
	}

	pbchat, err := parseGetChatResponse(chat)
	if err != nil {
		return nil, err
	}

	return &pb.GetChatResponse{Chat: pbchat}, nil
}

func parseGetChatResponse(chat *model.Chat) (*pb.Chat, error) {
	return &pb.Chat{
		ChatId:        chat.ChatID,
		Title:         chat.Title,
		NumberMembers: chat.NumberMembers,
		CreatedAt:     timestamppb.New(chat.CreatedAt),
	}, nil
}
