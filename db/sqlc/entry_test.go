package db

import (
	"bank_project/db/util"
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func createRandomEntry(t *testing.T) Entry {
	account := createRandomAccount(t)
	arg := CreateEntryParams{
		AccountID: account.ID,
		Amount:    util.RandomMoney(),
	}
	res, err := testQueries.CreateEntry(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, res)
	require.Equal(t, arg.AccountID, res.AccountID)
	require.Equal(t, arg.Amount, res.Amount)
	require.NotEmpty(t, res.ID)
	require.NotEmpty(t, res.CreatedAt)
	return res
}
func TestCreateEntry(t *testing.T) {
	createRandomEntry(t)
}

func TestGetEntry(t *testing.T) {
	entry := createRandomEntry(t)
	res, err := testQueries.GetEntry(context.Background(), entry.ID)
	require.NoError(t, err)
	require.NotEmpty(t, res)
	require.NotEmpty(t, res.ID)
	require.NotEmpty(t, res.CreatedAt)
}

func TestListEntries(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomEntry(t)
	}
	arg := ListEntriesParams{
		5,
		5,
	}
	res, err := testQueries.ListEntries(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, res, 5)
	for _, entry := range res {
		require.NotEmpty(t, entry)
	}
}

func TestUpdateEntry(t *testing.T) {
	entry1 := createRandomEntry(t)
	arg := UpdateEntryParams{
		ID:     entry1.ID,
		Amount: util.RandomMoney(),
	}
	entry2, err := testQueries.UpdateEntry(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, entry2)
	require.Equal(t, arg.Amount, entry2.Amount)
	require.NotEmpty(t, entry2.ID)
	require.WithinDuration(t, entry1.CreatedAt, entry2.CreatedAt, time.Second)
}

func TestDeleteEntry(t *testing.T) {
	entry := createRandomEntry(t)
	testQueries.DeleteEntry(context.Background(), entry.ID)
	res, err := testQueries.GetEntry(context.Background(), entry.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, res)
}
