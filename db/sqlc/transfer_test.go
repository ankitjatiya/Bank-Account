package db

import (
	"Bank-Account/db/util"
	"context"
	"database/sql"
	"testing"

	"github.com/stretchr/testify/require"
)

func createRandomTransfer(t *testing.T) Transfer {
	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)
	arg := CreateTransferParams{
		FromAccountID: account1.ID,
		ToAccountID:   account2.ID,
		Amount:        util.RandomMoney(),
	}
	res, err := testQueries.CreateTransfer(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, res)
	require.Equal(t, arg.FromAccountID, res.FromAccountID)
	require.Equal(t, arg.ToAccountID, res.ToAccountID)
	require.Equal(t, arg.Amount, res.Amount)
	require.NotEmpty(t, res.ID)
	return res
}
func TestCreateTransfer(t *testing.T) {
	createRandomEntry(t)
}

func TestGetTransfer(t *testing.T) {
	transfer := createRandomTransfer(t)
	res, err := testQueries.GetTransfer(context.Background(), transfer.ID)
	require.NoError(t, err)
	require.NotEmpty(t, res)
	require.Equal(t, res.FromAccountID, transfer.FromAccountID)
	require.Equal(t, res.ToAccountID, transfer.ToAccountID)
	require.Equal(t, res.Amount, transfer.Amount)
}

func TestListTransfer(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomTransfer(t)
	}
	arg := ListTransfersParams{
		5,
		5,
	}
	res, err := testQueries.ListTransfers(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, res, 5)
	for _, transfer := range res {
		require.NotEmpty(t, transfer)
	}
}

func TestUpdateTransfer(t *testing.T) {
	transfer1 := createRandomTransfer(t)
	arg := UpdateTransferParams{
		ID:     transfer1.ID,
		Amount: util.RandomMoney(),
	}
	transfer2, err := testQueries.UpdateTransfer(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, transfer2)
	require.Equal(t, arg.Amount, transfer2.Amount)
	require.NotEmpty(t, transfer2.ID)
}

func TestDeleteTransfer(t *testing.T) {
	transfer := createRandomEntry(t)
	testQueries.DeleteTransfer(context.Background(), transfer.ID)
	res, err := testQueries.GetTransfer(context.Background(), transfer.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, res)
}
