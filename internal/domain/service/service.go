package service

import (
	"context"

	"github.com/pawpawchat/chat/internal/domain/model"
	"github.com/pawpawchat/chat/internal/infrastructure/repository"
)

type ChatRepository interface {
	CreateChat(context.Context, *model.Chat) error
	GetChat(context.Context, uint64) (*model.Chat, error)
}

type MemberRepository interface {
	AddMember(context.Context, *model.Member) error
	GetMember(context.Context, uint64, uint64) (*model.Member, error)
	GetMembers(context.Context, uint64) (*[]model.Member, error)
}

type MessageRepository interface {
	SendMessage(context.Context, *model.Message) error
	GetMessages(context.Context, uint64) (*[]model.Message, error)
}

var _ ChatRepository = (*repository.ChatRepository)(nil)
var _ MemberRepository = (*repository.MemberRepository)(nil)
var _ MessageRepository = (*repository.MessageRepository)(nil)

type Service struct {
	msgRepo    MessageRepository
	chatRepo   ChatRepository
	memberRepo MemberRepository
	msgChan    chan *model.Message
}

func New(me MessageRepository, ch ChatRepository, mb MemberRepository, mch chan *model.Message) *Service {
	return &Service{msgRepo: me, chatRepo: ch, memberRepo: mb, msgChan: mch}
}
