package db

import (
	"context"
	"testing"
	"time"

	"github.com/luckyparakh/goBank/utils"
	"github.com/stretchr/testify/require"
)

func randomUserCreation(t *testing.T) User {
	hp, err := utils.HashedPassword(utils.GenerateRandomString(6))
	require.NoError(t, err)
	args := CreateUserParams{
		Username:       utils.GenerateRandomOwner(),
		HashedPassword: hp,
		FullName:       utils.GenerateRandomOwner(),
		Email:          utils.GenerateRandomEmail(),
	}
	user, err := testQueries.CreateUser(context.Background(), args)
	require.NotEmpty(t, user)
	require.NoError(t, err)
	require.Equal(t, user.Username, args.Username)
	require.Equal(t, user.FullName, args.FullName)
	require.Equal(t, user.Email, args.Email)
	require.Equal(t, user.HashedPassword, args.HashedPassword)
	require.NotZero(t, user.CreatedAt)
	require.True(t, user.PasswordChangedAt.IsZero())
	return user
}
func TestCreateUser(t *testing.T) {
	randomUserCreation(t)
}

func TestGetUser(t *testing.T) {
	user1 := randomUserCreation(t)
	user2, err := testQueries.GetUser(context.Background(), user1.Username)
	require.NotEmpty(t, user2)
	require.Empty(t, err)
	require.Equal(t, user1.Username, user2.Username)
	require.Equal(t, user1.Email, user2.Email)
	require.WithinDuration(t, user1.CreatedAt, user2.CreatedAt, time.Second)
	require.WithinDuration(t, user1.PasswordChangedAt, user2.PasswordChangedAt, time.Second)
}
