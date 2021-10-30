package db

import (
	"context"
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/sakhaei-wd/banker/util"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestCreateEntry(t *testing.T) {
	createRandomEntry(t, createRandomAccount(t).ID)
}

func TestGetEntry(t *testing.T) {
	entry1 := createRandomEntry(t, createRandomAccount(t).ID)
	entry2, err := testQueries.GetEntry(context.Background(), entry1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, entry2)

	require.Equal(t, entry1.ID, entry2.ID)
	require.Equal(t, entry1.AccountID, entry2.AccountID)
	require.Equal(t, entry1.Amount, entry2.Amount)
	require.WithinDuration(t, entry1.CreatedAt, entry1.CreatedAt, time.Second)

}

func TestListEntries(t *testing.T) {
	accountId := createRandomAccount(t).ID
	for i := 0; i < 6; i++ {
		createRandomEntry(t, accountId)
	}

	arg := ListEntriesParams{
		AccountID: accountId,
		Limit:     3, //take 3
		Offset:    3, //skip 3
	}

	listEntries, err := testQueries.ListEntries(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, listEntries)
	require.Len(t, listEntries, 3)

	for _, entry := range listEntries {
		require.NotEmpty(t, entry)
	}
}

func createRandomEntry(t *testing.T, accountId int64) Entry {
	arg := CreateEntryParams{
		AccountID: accountId,
		Amount:    util.RandomMoney(),
	}

	entry, err := testQueries.CreateEntry(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, entry)

	require.Equal(t, arg.AccountID, entry.AccountID)
	require.Equal(t, arg.Amount, entry.Amount)

	require.NotZero(t, entry.Amount)
	require.NotZero(t, entry.ID)
	require.NotZero(t, entry.AccountID)
	//require.NotZero(t, entry.CreatedAt)

	return entry
}
