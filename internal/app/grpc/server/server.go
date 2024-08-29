package server

import (
	"context"

	"github.com/pawpawchat/chat/api/pb"
	"github.com/pawpawchat/chat/internal/app/grpc/adapter"
	"github.com/pawpawchat/chat/internal/domain/model"
	"github.com/pawpawchat/chat/internal/domain/service"
)

type profileService interface {
	CreateChat(context.Context, *model.Member, *model.Chat) error
	GetChat(context.Context, int64) (*model.Chat, error)

	AddMember(context.Context, *model.Member) error
	GetMembers(context.Context, int64) (*[]model.Member, error)

	SendMessage(context.Context, *model.Message) error
	GetMessages(context.Context, int64) (*[]model.Message, error)
}

var _ profileService = (*service.Service)(nil)

type ChatGRPCServer struct {
	pb.UnimplementedChatServiceServer
	service profileService
}

func NewGRPCServer(service profileService) *ChatGRPCServer {
	return &ChatGRPCServer{service: service}
}

func (s *ChatGRPCServer) CreateChat(ctx context.Context, req *pb.CreateChatRequest) (*pb.CreateChatResponse, error) {
	return adapter.CreateChatAdapter(ctx, s.service, req)
}

func (s *ChatGRPCServer) GetChat(ctx context.Context, req *pb.GetChatRequest) (*pb.GetChatResponse, error) {
	return adapter.GetChatAdapter(ctx, s.service, req)
}

func (s *ChatGRPCServer) AddMember(ctx context.Context, req *pb.AddMemberRequest) (*pb.AddMemberResponse, error) {
	return adapter.AddMemberAdapter(ctx, s.service, req)
}

func (s *ChatGRPCServer) GetMembers(ctx context.Context, req *pb.GetMembersRequest) (*pb.GetMembersResponse, error) {
	return adapter.GetMembersAdapater(ctx, s.service, req)
}

func (s *ChatGRPCServer) SendMessage(ctx context.Context, req *pb.SendMessageRequest) (*pb.SendMessageResponse, error) {
	return adapter.SendMessageAdapter(ctx, s.service, req)
}

func (s *ChatGRPCServer) GetMessages(ctx context.Context, req *pb.GetMessagesRequest) (*pb.GetMessagesResponse, error) {
	return adapter.GetMessagesAdapter(ctx, s.service, req)
}
