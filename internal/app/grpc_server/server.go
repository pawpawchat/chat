package grpcserver

import (
	"context"

	"github.com/pawpawchat/chat/api/pb"
	"github.com/pawpawchat/chat/internal/domain/model"
	"github.com/pawpawchat/chat/internal/domain/service"
)

type usecase interface {
	CreateChat(context.Context, *model.Member, *model.Chat) error
	GetChat(context.Context, uint64) (*model.Chat, error)

	AddMember(context.Context, *model.Member) error
	GetMembers(context.Context, uint64) (*[]model.Member, error)

	SendMessage(context.Context, *model.Message) error
	GetMessages(context.Context, uint64) (*[]model.Message, error)
}

var _ usecase = (*service.Service)(nil)

type ChatGRPCServer struct {
	pb.UnimplementedChatServiceServer
	usecase usecase
}

func New(usecase usecase) *ChatGRPCServer {
	return &ChatGRPCServer{usecase: usecase}
}
