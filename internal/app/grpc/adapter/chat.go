package adapter

import (
	"context"
	"time"

	"github.com/pawpawchat/chat/api/pb"
	"github.com/pawpawchat/chat/internal/domain/model"
)

type chatProvider interface {
	CreateChat(context.Context, *model.Member, *model.Chat) error
	GetChat(context.Context, int64) (*model.Chat, error)
}

func CreateChatAdapter(ctx context.Context, provider chatProvider, req *pb.CreateChatRequest) (*pb.CreateChatResponse, error) {
	var createdAt time.Time
	switch req.CreatedAt {
	case nil:
		createdAt = req.CreatedAt.AsTime()
	default:
		createdAt = time.Now()
	}

	chat := &model.Chat{
		Title:         req.Title,
		CreatedAt:     createdAt,
		NumberMembers: 1,
	}

	owner := &model.Member{
		MemberID: req.OwnerId,
		Username: req.OwnerUsername,
		Role:     "owner",
	}

	if err := provider.CreateChat(ctx, owner, chat); err != nil {
		return nil, err
	}
	return &pb.CreateChatResponse{Chat: chat.ToPb()}, nil
}

func GetChatAdapter(ctx context.Context, provider chatProvider, req *pb.GetChatRequest) (*pb.GetChatResponse, error) {
	chatID, _ := req.ChatId, req.MemberId

	chat, err := provider.GetChat(ctx, chatID)
	if err != nil {
		return nil, err
	}

	chatMessages, _ := provider.(messageProvider)

	messages, err := chatMessages.GetMessages(ctx, chatID)
	if err != nil {
		return nil, err
	}

	var messagespb []*pb.Message = make([]*pb.Message, 0)
	for _, m := range *messages {
		messagespb = append(messagespb, m.ToPb())
	}

	return &pb.GetChatResponse{Chat: chat.ToPb(), Messages: messagespb}, nil
}
