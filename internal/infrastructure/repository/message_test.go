package repository_test

import (
	"context"
	"testing"
	"time"

	"github.com/pawpawchat/chat/internal/domain/model"
	"github.com/pawpawchat/chat/internal/infrastructure/repository"
	"github.com/stretchr/testify/assert"
)

func TestMessageRepository_SendMessage(t *testing.T) {
	db := testingContext(t)
	defer db.Close()

	chat, _ := addTestChat(db)

	testCases := []struct {
		desc  string
		msg   func() *model.Message
		valid bool
	}{
		{
			"Chat exists",
			func() *model.Message {
				msg := testMessage()
				msg.ChatID = chat.ChatID
				return msg
			},
			true,
		},
		{
			"Chat doesn't exists",
			func() *model.Message {
				msg := testMessage()
				msg.ChatID = 0
				return msg
			},
			false,
		},
	}

	mr := repository.NewMessageRepository(db)

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			// running query
			err := mr.SendMessage(context.Background(), tc.msg())
			// check expectations
			switch tc.valid {
			case true:
				assert.NoError(t, err)
			case false:
				assert.Error(t, err)
			}
		})
	}
}

func TestMessageRepository_GetMessages(t *testing.T) {
	db := testingContext(t)
	defer db.Close()

	chat, _ := addTestChat(db)

	testCases := []struct {
		desc   string
		chatID int64
		valid  bool
	}{
		{
			"Chat exists",
			chat.ChatID,
			true,
		},
		{
			"Chat doesn't exist",
			0,
			false,
		},
	}

	mr := repository.NewMessageRepository(db)

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			// running query
			msgs, err := mr.GetMessages(context.Background(), tc.chatID)
			// check expectations
			switch tc.valid {
			case true:
				assert.NoError(t, err)
				assert.NotNil(t, msgs)
			case false:
				assert.Empty(t, msgs)
			}
		})
	}

}

func testMessage() *model.Message {
	return &model.Message{
		ChatID:         1,
		SenderID:       1,
		SenderUsername: "username",
		Body:           "message body",
		SentAt:         time.Now(),
		IsDeleted:      false,
	}
}
