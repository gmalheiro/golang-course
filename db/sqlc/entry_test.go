package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/gmalheiro/golang-course/util"
	"github.com/stretchr/testify/require"
)

func createRandomEntry(t *testing.T) Entry {
	account := createRandomAccount(t)
	arg := CreateEntryParams{
		AccountID: account.ID,
		Amount:    util.RandomMoney(),
	}

	entry, err := testQueries.CreateEntry(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, entry)

	require.Equal(t, entry.AccountID, arg.AccountID)
	require.Equal(t, entry.Amount, arg.Amount)

	require.NotZero(t, entry.ID)
	require.NotZero(t, entry.CreatedAt)

	return entry
}

func TestCreateEntry(t *testing.T) {
	createRandomEntry(t)
}

func TestGetEntry(t *testing.T) {
	entry := createRandomEntry(t)

	entryInDb, err := testQueries.GetEntry(context.Background(), entry.ID)
	require.NoError(t, err)
	require.NotEmpty(t, entryInDb)
	require.NotZero(t, entryInDb.ID)

	require.Equal(t, entryInDb.Amount, entry.Amount)
	require.Equal(t, entryInDb.AccountID, entry.AccountID)
	require.Equal(t, entryInDb.ID, entry.ID)
}

func TestUpdateEntry(t *testing.T) {
	entry := createRandomEntry(t)

	arg := UpdateEntryParams{
		ID:     entry.ID,
		Amount: util.RandomMoney(),
	}

	entryInDb, err := testQueries.UpdateEntry(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, entryInDb)
	require.NotZero(t, entryInDb.CreatedAt)
	require.NotZero(t, entryInDb.ID)

	require.Equal(t, entryInDb.ID, entry.ID)
	require.Equal(t, entryInDb.AccountID, entry.AccountID)
	require.Equal(t, entryInDb.Amount, arg.Amount)
}

func TestDeleteEntry(t *testing.T) {
	entry := createRandomEntry(t)

	err := testQueries.DeleteEntry(context.Background(), entry.ID)
	require.NoError(t, err)

	entryInDb, err := testQueries.GetEntry(context.Background(), entry.ID)

	require.Error(t, err, sql.ErrNoRows.Error())
	require.Empty(t, entryInDb)
}

func TestListEntry(t *testing.T) {
	for i := 0; i < 5; i++ {
		createRandomEntry(t)
	}

	arg := ListEntriesParams{
		Offset: 5,
		Limit:  5,
	}

	entries, err := testQueries.ListEntries(context.Background(), arg)

	require.NoError(t, err)
	require.Len(t, entries, 5)

	for _, entry := range entries {
		require.NotEmpty(t, entry)
	}

}
