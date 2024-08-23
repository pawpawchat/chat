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

func testingContext(t *testing.T) *sqlx.DB {
	flag.Set("env", "testing")
	flag.Parse()

	config, err := config.LoadDefaultConfig()
	assert.NoError(t, err)

	db, err := sqlx.Open("pgx", config.Env().DB_URL)
	assert.NoError(t, err)

	return db
}

func TestChatRepository_GetChat(t *testing.T) {
	db := testingContext(t)
	defer db.Close()

	cr := repository.NewChatRepository(db)

	chat, err := cr.GetChat(context.Background(), 1)

	assert.NoError(t, err)
	assert.NotNil(t, chat)
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
