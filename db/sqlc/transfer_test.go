package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/gmalheiro/golang-course/util"
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

	transfer, err := testQueries.CreateTransfer(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, transfer)

	require.Equal(t, transfer.FromAccountID, arg.FromAccountID)
	require.Equal(t, transfer.ToAccountID, arg.ToAccountID)
	require.Equal(t, transfer.Amount, arg.Amount)

	require.NotZero(t, transfer.ID)
	require.NotZero(t, transfer.CreatedAt)
	return transfer
}

func TestCreateTransfer(t *testing.T) {
	createRandomTransfer(t)
}

func TestGetTransfer(t *testing.T) {
	transfer := createRandomTransfer(t)
	transferInDb, err := testQueries.GetTransfer(context.Background(), transfer.ID)

	require.NoError(t, err)
	require.NotEmpty(t, transferInDb)
	require.NotZero(t, transferInDb.ID)

	require.Equal(t, transferInDb.ID, transfer.ID)
	require.Equal(t, transferInDb.FromAccountID, transfer.FromAccountID)
	require.Equal(t, transferInDb.ToAccountID, transfer.ToAccountID)
	require.Equal(t, transferInDb.Amount, transfer.Amount)
}

func TestUpdateTransfer(t *testing.T) {
	transfer := createRandomTransfer(t)

	arg := UpdateTransferParams{
		ID:     transfer.ID,
		Amount: util.RandomMoney(),
	}

	updatedTransfer, err := testQueries.UpdateTransfer(context.Background(), arg)

	require.NoError(t, err)
	require.NotZero(t, updatedTransfer)

	require.Equal(t, updatedTransfer.ID, transfer.ID)
	require.Equal(t, updatedTransfer.Amount, arg.Amount)
}

func TestDeleteTransfer(t *testing.T) {
	transfer := createRandomTransfer(t)

	err := testQueries.DeleteTransfer(context.Background(), transfer.ID)

	require.NoError(t, err)

	transferInDb, err := testQueries.GetTransfer(context.Background(), transfer.ID)
	require.Error(t, err, sql.ErrNoRows.Error())

	require.Empty(t, transferInDb)
}

func TestGetAllTransfers(t *testing.T) {
	for i := 0; i < 5; i++ {
		createRandomTransfer(t)
	}

	arg := ListTransferParams{
		Offset: 5,
		Limit:  5,
	}

	transfers, err := testQueries.ListTransfer(context.Background(), arg)

	require.NoError(t, err)

	require.Len(t, transfers, 5)

	for _, transfer := range transfers {
		require.NotEmpty(t, transfer)
	}
}
