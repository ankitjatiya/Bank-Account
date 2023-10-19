package db

import (
	"Bank-Account/util"
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func createRandomAccount(t *testing.T) Account {
	user := createRandomUser(t)
	arg := CreateAccountParams{
		user.Username,
		util.RandomMoney(),
		util.RandomCurrency(),
	}
	res, err := testQueries.CreateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, res)
	require.Equal(t, arg.Owner, res.Owner)
	require.Equal(t, arg.Balance, res.Balance)
	require.Equal(t, arg.Currency, res.Currency)
	require.NotEmpty(t, res.ID)
	require.NotEmpty(t, res.CreatedAt)
	return res
}
func TestCreateAccount(t *testing.T) {
	createRandomAccount(t)
}

func TestGetAccount(t *testing.T) {
	account := createRandomAccount(t)
	res, err := testQueries.GetAccount(context.Background(), account.ID)
	require.NoError(t, err)
	require.NotEmpty(t, res)
	require.Equal(t, account.Owner, res.Owner)
	require.Equal(t, account.Balance, res.Balance)
	require.Equal(t, account.Currency, res.Currency)
	require.NotEmpty(t, res.ID)
	require.NotEmpty(t, res.CreatedAt)
}

func TestListAccounts(t *testing.T) {
	var lastAccount Account
	for i := 0; i < 10; i++ {
		lastAccount = createRandomAccount(t)
	}
	arg := ListAccountsParams{
		Owner:  lastAccount.Owner,
		Limit:  5,
		Offset: 0,
	}
	res, err := testQueries.ListAccounts(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, res)

	for _, account := range res {
		require.NotEmpty(t, account)
		require.Equal(t, lastAccount.Owner, account.Owner)
	}
}

func TestUpdateAccount(t *testing.T) {
	account1 := createRandomAccount(t)
	arg := UpdateAccountParams{
		ID:      account1.ID,
		Balance: util.RandomMoney(),
	}
	account2, err := testQueries.UpdateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account2)
	require.Equal(t, account1.Owner, account2.Owner)
	require.Equal(t, account1.Currency, account2.Currency)
	require.Equal(t, arg.Balance, account2.Balance)
	require.NotEmpty(t, account2.ID)
	require.WithinDuration(t, account1.CreatedAt, account2.CreatedAt, time.Second)
}

func TestDeleteAccount(t *testing.T) {
	account := createRandomAccount(t)
	testQueries.DeleteAccount(context.Background(), account.ID)
	res, err := testQueries.GetAccount(context.Background(), account.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, res)
}
