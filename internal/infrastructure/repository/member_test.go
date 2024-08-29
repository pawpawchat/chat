package repository_test

import (
	"context"
	"math/rand"
	"testing"

	"github.com/jmoiron/sqlx"
	"github.com/pawpawchat/chat/internal/domain/model"
	"github.com/pawpawchat/chat/internal/infrastructure/repository"
	"github.com/stretchr/testify/assert"
)

func TestMemberRepository_AddMember(t *testing.T) {
	db := testingContext(t)
	defer db.Close()

	_, member := addTestChat(db)

	testCases := []struct {
		desc   string
		member func() *model.Member
		valid  bool
	}{
		{
			"Chat exists, user has not been added",
			func() *model.Member { member.MemberID = int64(rand.Int31()); return member },
			true,
		},
		{
			"Chat exists, user has already been added",
			func() *model.Member { return member },
			false,
		},
		{
			"Chat doesn't exists",
			func() *model.Member { member.ChatID = 0; return member },
			false,
		},
	}

	mr := repository.NewMemberRepository(db)

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			// running query
			err := mr.AddMember(context.Background(), tc.member())
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

func TestMemberRepository_GetMembers(t *testing.T) {
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

	mr := repository.NewMemberRepository(db)

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			// running query
			members, err := mr.GetMembers(context.Background(), tc.chatID)
			// check expectations
			switch tc.valid {
			case true:
				assert.NoError(t, err)
				assert.NotNil(t, members)
			case false:
				assert.Empty(t, members)
			}
		})
	}
}

func TestMemberRepository_GetMember(t *testing.T) {
	db := testingContext(t)
	defer db.Close()

	// adding test data
	chat, member := addTestChat(db)

	testCases := []struct {
		desc     string
		chatID   int64
		memberID int64
		valid    bool
	}{
		{
			"Chat and member exist",
			chat.ChatID, member.MemberID,
			true,
		},
		{
			"Chat exists, member doesn't exists",
			chat.ChatID, 0,
			false,
		},
		{
			"Chat doens't exists, member exists in other chats",
			0, member.ChatID,
			false,
		},
	}

	mr := repository.NewMemberRepository(db)

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			// running query
			member, err := mr.GetMember(context.Background(), tc.chatID, tc.memberID)
			// check expectations
			switch tc.valid {
			case true:
				assert.NoError(t, err)
				assert.NotNil(t, member)
			case false:
				assert.Error(t, err)
				assert.Empty(t, member)
			}
		})
	}
}

func addTestChat(db *sqlx.DB) (*model.Chat, *model.Member) {
	chat := testChat()
	repository.NewChatRepository(db).CreateChat(context.Background(), chat)
	member := testMember()
	member.ChatID = chat.ChatID
	repository.NewMemberRepository(db).AddMember(context.Background(), member)

	return chat, member
}
