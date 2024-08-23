package repository_test

import (
	"context"
	"math/rand"
	"testing"

	"github.com/pawpawchat/chat/internal/infrastructure/repository"
	"github.com/stretchr/testify/assert"
)

func TestMemberRepository_AddMember(t *testing.T) {
	db := testingContext(t)
	defer db.Close()

	mr := repository.NewMemberRepository(db)

	member := testMember()
	member.MemberID = uint64(rand.Int63())

	assert.NoError(t, mr.AddMember(context.Background(), member))
}

func TestMemberRepository_GetMembers(t *testing.T) {
	db := testingContext(t)
	defer db.Close()

	mr := repository.NewMemberRepository(db)

	var chatID uint64 = 1
	members, err := mr.GetMembers(context.Background(), chatID)

	assert.NoError(t, err)
	assert.NotNil(t, members)
}
