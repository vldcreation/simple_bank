package db

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/vldcreation/simple_bank/util"
)

func createRandoomUser(t *testing.T) Users {
	hashedPassword, err := util.HashPassword(util.RandString(6))
	require.NoError(t, err)

	args := CreateUserParams{
		Username:       util.RandOwnersName(),
		HashedPassword: hashedPassword,
		FullName:       util.RandOwnersName(),
		Email:          util.RandEmail(),
	}

	user, err := testQueries.CreateUser(context.Background(), args)

	require.NoError(t, err)
	require.NotEmpty(t, user)

	// check if the created account is the same as the one we passed in
	require.Equal(t, args.Username, user.Username)
	require.Equal(t, args.HashedPassword, user.HashedPassword)
	require.Equal(t, args.FullName, user.FullName)
	require.Equal(t, args.Email, user.Email)

	// check if the created account has an ID
	require.NotZero(t, user.CreatedAt)
	require.True(t, user.PasswordChangedAt.IsZero())

	return user
}

func TestCreateUser(t *testing.T) {
	createRandoomUser(t)
}

func TestGetUser(t *testing.T) {
	user1 := createRandoomUser(t)

	user2, err := testQueries.GetUser(context.Background(), user1.Username)
	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, user1.Username, user2.Username)
	require.Equal(t, user1.HashedPassword, user2.HashedPassword)
	require.Equal(t, user1.FullName, user2.FullName)
	require.Equal(t, user1.Email, user2.Email)

	require.WithinDuration(t, user1.CreatedAt, user2.CreatedAt, time.Second)
	require.True(t, user2.PasswordChangedAt.IsZero())
	require.WithinDuration(t, user1.PasswordChangedAt, user2.PasswordChangedAt, time.Second)
}
