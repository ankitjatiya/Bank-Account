package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateAccount(t *testing.T) {
	arg := CreateAccountParams{
		"AJ", 1234, "INR",
	}
	res, err := testQueries.CreateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, res)
	require.Equal(t, arg.Owner, res.Owner)
	require.Equal(t, arg.Balance, res.Balance)
	require.Equal(t, arg.Currency, res.Currency)
	require.NotEmpty(t, res.ID)
	require.NotEmpty(t, res.CreatedAt)
}
