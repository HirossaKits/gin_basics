package db

import (
	"context"
	"database/sql"
	"gin_basics/util"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func createRandomEntry(t *testing.T) Entry {
	account := CreateRandomAccount(t)
	arg := CreateEntryParams{
		AccountID: account.ID,
		Amount:    util.RandomMoney(),
	}

	entry, err := testQueries.CreateEntry(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, entry)

	require.Equal(t, arg.AccountID, entry.AccountID)
	require.Equal(t, arg.Amount, entry.Amount)

	require.NotZero(t, entry.ID)
	require.NotZero(t, entry.CreatedAt)

	return entry
}

func TestCreateEntry(t *testing.T) {
	createRandomEntry(t)
}

func TestGetEntry(t *testing.T) {
	entry1 := createRandomEntry(t)
	entry2, err := testQueries.GetEntry(context.Background(), entry1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, entry2)

	require.Equal(t, entry1.ID, entry2.ID)
	require.Equal(t, entry1.AccountID, entry2.AccountID)
	require.Equal(t, entry1.Amount, entry2.Amount)
	require.Equal(t, entry1.CreatedAt, entry2.CreatedAt)
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

	require.Equal(t, entry1.ID, entry2.ID)
	require.Equal(t, entry1.AccountID, entry2.AccountID)
	require.Equal(t, arg.Amount, entry2.Amount)
	require.WithinDuration(t, entry1.CreatedAt, entry2.CreatedAt, time.Second)
}

func TestDeleteEntry(t *testing.T) {
	entry1 := createRandomEntry(t)
	err := testQueries.DeleteEntry(context.Background(), entry1.ID)
	require.NoError(t, err)

	entry2, err := testQueries.GetEntry(context.Background(), entry1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, entry2)
}

func TestListEntries(t *testing.T) {
	account := CreateRandomAccount(t)

	for i := 0; i < 10; i++ {
		arg := CreateEntryParams{
			AccountID: account.ID,
			Amount:    util.RandomMoney(),
		}
		testQueries.CreateEntry(context.Background(), arg)
	}

	entries, err := testQueries.ListEntries(context.Background(), account.ID)
	require.NoError(t, err)
	require.Len(t, entries, 10)

	for _, entry := range entries {
		require.NotEmpty(t, entry.ID)
		require.Equal(t, account.ID, entry.AccountID)
		require.NotEmpty(t, entry.Amount)
		require.NotEmpty(t, entry.CreatedAt)
	}
}
