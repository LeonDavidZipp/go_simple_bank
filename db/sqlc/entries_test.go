package db

import (
	"context"
	"testing"
	"time"
	"database/sql"
	"github.com/stretchr/testify/require"
	"github.com/LeonDavidZipp/go_simple_bank/util"
)

func CreateRandomEntry(t *testing.T) Entry {
	arg := CreateEntryParams{
		AccountID := RandomID()
		Amount := RandomAmount()
	}

	account, err := testQueries.CreateEntry(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, arg.AccountID, account.AccountID)
	require.Equal(t, arg.Amount, account,Amount)

	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)
}
