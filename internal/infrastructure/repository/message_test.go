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

	mr := repository.NewMessageRepository(db)
	msg := testMessage()

	assert.NoError(t, mr.SendMessage(context.Background(), msg))
}

func TestMessageRepository_GetMessages(t *testing.T) {
	db := testingContext(t)
	defer db.Close()

	mr := repository.NewMessageRepository(db)
	msgs, err := mr.GetMessages(context.TODO(), 1)

	assert.NoError(t, err)
	assert.NotNil(t, msgs)
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
