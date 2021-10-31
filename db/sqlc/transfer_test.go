package db

import (
	"context"
	_ "github.com/lib/pq"
	"github.com/sakhaei-wd/banker/util"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestCreateTransfer(t *testing.T) {
	createRandomTransfer(t, createRandomAccount(t).ID, createRandomAccount(t).ID)
}

func createRandomTransfer(t *testing.T, fromAccountId int64, toAccountId int64) Transfer {
	arg := CreateTransferParams{
		FromAccountID: fromAccountId,
		ToAccountID:   toAccountId,
		Amount:        util.RandomMoney(),
	}

	transfer, err := testQueries.CreateTransfer(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, transfer)

	require.Equal(t, arg.FromAccountID, transfer.FromAccountID)
	require.Equal(t, arg.ToAccountID, transfer.ToAccountID)
	require.Equal(t, arg.Amount, transfer.Amount)

	require.NotZero(t, transfer.Amount)
	require.NotZero(t, transfer.ID)
	require.NotZero(t, transfer.FromAccountID)
	require.NotZero(t, transfer.ToAccountID)
	//require.NotZero(t, transfer.CreatedAt)

	return transfer
}

func TestGetTransfer(t *testing.T) {
	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)
	transfer1 := createRandomTransfer(t, account1.ID, account2.ID)

	transfer2, err := testQueries.GetTransfer(context.Background(), transfer1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, transfer2)

	require.Equal(t, transfer1.ID, transfer2.ID)
	require.Equal(t, transfer1.FromAccountID, transfer2.FromAccountID)
	require.Equal(t, transfer1.ToAccountID, transfer2.ToAccountID)
	require.Equal(t, transfer1.Amount, transfer2.Amount)
	require.WithinDuration(t, transfer1.CreatedAt, transfer2.CreatedAt, time.Second)
}

func TestListTransfer(t *testing.T) {

	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)

	for i := 0; i < 3; i++ {
		createRandomTransfer(t, account1.ID, account2.ID)
		createRandomTransfer(t, account2.ID, account1.ID)
	}

	arg := ListTransferParams{
		FromAccountID: account1.ID,
		ToAccountID:   account1.ID,
		Limit:         3,
		Offset:        3,
	}

	transfers, err := testQueries.ListTransfer(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, transfers, 3)

	for _, transfer := range transfers {
		require.NotEmpty(t, transfer)
		require.True(t, transfer.FromAccountID == account1.ID || transfer.ToAccountID == account1.ID)
	}
}
