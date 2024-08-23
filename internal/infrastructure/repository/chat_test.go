package repository_test

import (
	"context"
	"flag"
	"testing"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/pawpawchat/chat/config"
	"github.com/pawpawchat/chat/internal/domain/model"
	"github.com/pawpawchat/chat/internal/infrastructure/repository"
	"github.com/stretchr/testify/assert"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func TestChatRepository_CreateChat(t *testing.T) {
	db := testingContext(t)
	defer db.Close()

	cr := repository.NewChatRepository(db)
	template := func() *model.Chat {
		chat := testChat()
		chat.ChatID = 0

		return chat
	}

	testCases := []struct {
		desc  string
		req   func() *model.Chat
		valid bool
	}{
		{
			"Succesfully adding",
			func() *model.Chat {
				return template()
			},
			true,
		},
	}

	chat := testCases[0].req()
	err := cr.CreateChat(context.Background(), chat)

	assert.NoError(t, err)
	assert.NotZero(t, chat.ChatID)
}

func TestChatRepository_GetChat(t *testing.T) {
	db := testingContext(t)
	defer db.Close()

	// test data
	chat := testChat()
	repository.NewChatRepository(db).CreateChat(context.Background(), chat)

	testCases := []struct {
		desc   string
		chatID uint64
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

	cr := repository.NewChatRepository(db)

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			// running query
			chat, err := cr.GetChat(context.Background(), tc.chatID)
			// check expectations
			switch tc.valid {
			case true:
				assert.NoError(t, err)
				assert.NotNil(t, chat)
			case false:
				assert.Error(t, err)
			}
		})
	}

}

func testingContext(t *testing.T) *sqlx.DB {
	flag.Set("env", "testing")
	flag.Parse()

	config, err := config.LoadDefaultConfig()
	assert.NoError(t, err)

	db, err := sqlx.Open("pgx", config.Env().DbUrl)
	assert.NoError(t, err)

	return db
}

func testChat() *model.Chat {
	return &model.Chat{
		ChatID:        1,
		Title:         "title",
		NumberMembers: 1,
		CreatedAt:     time.Now(),
	}
}

func testMember() *model.Member {
	return &model.Member{
		MemberID: 1,
		Username: "username",
		ChatID:   1,
	}
}
