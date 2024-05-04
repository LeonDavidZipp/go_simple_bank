package db

import (
	"context"
	"testing"
	"time"
	"github.com/stretchr/testify/require"
	"github.com/LeonDavidZipp/go_simple_bank/util"

)


func CreateRandomAccount(t *testing.T) Account {
	arg := CreateAccountParams{
		Owner: util.RandomOwner(),
		Balance: util.RandomBalance(),
		Currency: util.RandomCurrency(),
	};

	account, err := testQueries.CreateAccount(context.Background(), arg);

	require.NoError(t, err);
	require.NotEmpty(t, account);

	require.Equal(t, arg.Owner, account.Owner);
	require.Equal(t, arg.Balance, account.Balance);
	require.Equal(t, arg.Currency, account.Currency);

	require.NotZero(t, account.ID);
	require.NotZero(t, account.CreatedAt);

	return account
}

func TestCreateAccount(t *testing.T) {
	CreateRandomAccount(t)
}

func TestGetAccount(t *testing.T) {
	user1 := CreateRandomAccount(t)
	user2, err := testQueries.GetAccount(context.Background(), user1.ID)

	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, user1.ID, user2.ID)
	require.Equal(t, user1.Owner, user2.Owner)
	require.Equal(t, user1.Balance, user2.Balance)
	require.Equal(t, user1.Currency, user2.Currency)

	require.WithinDuration(t, user1.CreatedAt, user2.CreatedAt, time.Second)
}
