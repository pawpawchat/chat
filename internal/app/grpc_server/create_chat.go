package grpcserver

import (
	"context"
	"time"

	"github.com/pawpawchat/chat/api/pb"
	"github.com/pawpawchat/chat/internal/domain/model"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *ChatGRPCServer) CreateChat(ctx context.Context, req *pb.CreateChatRequest) (*pb.CreateChatResponse, error) {
	chat, owner, err := parseCreateChatRequest(req)
	if err != nil {
		return nil, err
	}

	if err := s.usecase.CreateChat(ctx, owner, chat); err != nil {
		return nil, err
	}

	pbchat, err := parseCreateChatResponse(chat)
	if err != nil {
		return nil, err
	}

	return &pb.CreateChatResponse{Chat: pbchat}, nil
}

func parseCreateChatResponse(chat *model.Chat) (*pb.Chat, error) {
	return &pb.Chat{
		ChatId:        chat.ChatID,
		Title:         chat.Title,
		NumberMembers: chat.NumberMembers,
		CreatedAt:     timestamppb.New(chat.CreatedAt),
	}, nil
}

func parseCreateChatRequest(r *pb.CreateChatRequest) (*model.Chat, *model.Member, error) {
	var createdAt time.Time
	switch r.CreatedAt {
	case nil:
		createdAt = r.CreatedAt.AsTime()
	default:
		createdAt = time.Now()
	}

	chat := &model.Chat{
		Title:         r.Title,
		CreatedAt:     createdAt,
		NumberMembers: 1,
	}

	owner := &model.Member{
		MemberID: r.OwnerId,
		Username: r.OwnerUsername,
		Role:     "owner",
	}

	return chat, owner, nil
}
