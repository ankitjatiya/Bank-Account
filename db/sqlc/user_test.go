package db

import (
	"Bank-Account/util"
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func createRandomUser(t *testing.T) User {
	hashedPassword, err := util.HashPassword(util.RandomString(6))
	require.NoError(t, err)

	arg := CreateUserParams{
		Username:       util.RandomOwner(),
		HashedPassword: hashedPassword,
		Email:          util.RandomEmail(),
		FullName:       util.RandomOwner(),
	}
	res, err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, res)
	require.Equal(t, arg.Username, res.Username)
	require.Equal(t, arg.HashedPassword, res.HashedPassword)
	require.Equal(t, arg.Email, res.Email)
	require.Equal(t, arg.FullName, res.FullName)
	require.NotEmpty(t, res.Username)
	require.NotEmpty(t, res.CreatedAt)
	return res
}
func TestCreateUser(t *testing.T) {
	createRandomUser(t)
}

func TestGetUser(t *testing.T) {
	user := createRandomUser(t)
	res, err := testQueries.GetUser(context.Background(), user.Username)
	require.NoError(t, err)
	require.NotEmpty(t, res)
	require.Equal(t, user.Username, res.Username)
	require.Equal(t, user.Email, res.Email)
	require.Equal(t, user.FullName, res.FullName)
	require.Equal(t, user.HashedPassword, res.HashedPassword)
	require.NotEmpty(t, res.Username)
	require.NotEmpty(t, res.CreatedAt)
}
