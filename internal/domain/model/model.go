package model

import (
	"time"

	"github.com/pawpawchat/chat/api/pb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Chat struct {
	ChatID        int64     `db:"chat_id"`           // идентификатор чата
	Title         string    `db:"title"`             // название чата
	NumberMembers int32     `db:"number_of_members"` // количество участников
	CreatedAt     time.Time `db:"created_at"`        // время создания
}

func (c Chat) ToPb() *pb.Chat {
	return &pb.Chat{
		ChatId:        c.ChatID,
		Title:         c.Title,
		NumberMembers: c.NumberMembers,
		CreatedAt:     timestamppb.New(c.CreatedAt),
	}
}

type Message struct {
	MessageID      int64     `db:"message_id"`      // идентификатор сообщения
	ChatID         int64     `db:"chat_id"`         // идентификатор чата
	SenderID       int64     `db:"sender_id"`       // идентификатор отправителя
	SenderUsername string    `db:"sender_username"` // юзернейм отправителя
	Body           string    `db:"body"`            // тело сообщения
	SentAt         time.Time `db:"sent_at"`         // время отправки сообщения
	IsDeleted      bool      `db:"is_deleted"`      // флаг удаления
}

func (m Message) ToPb() *pb.Message {
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

func (m *Member) FromPb(member *pb.Member) *Member {
	m.MemberID = member.MemberId
	m.Username = member.Username
	m.ChatID = member.ChatId
	m.Role = member.Role
	return m
}

type Sender struct {
	SenderID int64  `db:"sender_id"`
	Username string `db:"sender_username"`
}

type Member struct {
	MemberID int64  `db:"member_id"` // идентификатор участника
	Username string `db:"username"`  // имя пользователя
	ChatID   int64  `db:"chat_id"`   // идентификатор чата
	Role     string `db:"role"`      // роль в чате
}

func (m *Member) ToPb() *pb.Member {
	return &pb.Member{
		MemberId: m.MemberID,
		Username: m.Username,
		ChatId:   m.ChatID,
		Role:     m.Role,
	}
}

type MessageReaders struct {
	MessageID int64   `db:"message_id"` // идентификатор сообщения
	ReaderIDs []int64 `db:"reader_id"`  // список идентификаторов прочитавших
}
