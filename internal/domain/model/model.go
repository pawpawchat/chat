package model

import (
	"time"
)

type Chat struct {
	ChatID        uint64    `db:"chat_id"`           // идентификатор чата
	Title         string    `db:"title"`             // название чата
	NumberMembers int32     `db:"number_of_members"` // количество участников
	CreatedAt     time.Time `db:"created_at"`        // время создания
}

type Message struct {
	MessageID      uint64    `db:"message_id"`      // идентификатор сообщения
	ChatID         uint64    `db:"chat_id"`         // идентификатор чата
	SenderID       uint64    `db:"sender_id"`       // идентификатор отправителя
	SenderUsername string    `db:"sender_username"` // юзернейм отправителя
	Body           string    `db:"body"`            // тело сообщения
	SentAt         time.Time `db:"sent_at"`         // время отправки сообщения
	IsDeleted      bool      `db:"is_deleted"`      // флаг удаления
}

type Sender struct {
	SenderID uint64 `db:"sender_id"`
	Username string `db:"sender_username"`
}

type Member struct {
	MemberID uint64 `db:"member_id"` // идентификатор участника
	Username string `db:"username"`  // имя пользователя
	ChatID   uint64 `db:"chat_id"`   // идентификатор чата
	Role     string `db:"role"`      // роль в чате
}

type MessageReaders struct {
	MessageID uint64   `db:"message_id"` // идентификатор сообщения
	ReaderIDs []uint64 `db:"reader_id"`  // список идентификаторов прочитавших
}
