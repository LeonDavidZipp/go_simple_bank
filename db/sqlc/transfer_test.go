package db

import (
	"context"
	"testing"
	"time"
	"github.com/stretchr/testify/require"
	"github.com/LeonDavidZipp/go_simple_bank/util"
)

func CreateRandomTransfer(t *testing.T, sender Account, receiver Account) Transfer {
	arg := CreateTransferParams{
		FromAccountID : sender.ID,
		ToAccountID : receiver.ID,
		Amount : util.RandomAmount(),
	}

	transfer, err := testQueries.CreateTransfer(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, transfer)

	require.Equal(t, arg.FromAccountID, transfer.FromAccountID)
	require.Equal(t, arg.ToAccountID, transfer.ToAccountID)
	require.Equal(t, arg.Amount, transfer.Amount)

	require.NotZero(t, transfer.ID)
	require.NotZero(t, transfer.CreatedAt)

	return transfer
}

func TestCreateTransfer(t *testing.T) {
	sender := CreateRandomAccount(t)
	receiver := CreateRandomAccount(t)
	CreateRandomTransfer(t, sender, receiver)
}

func TestGetTransfer(t *testing.T) {
	transfer1 := CreateRandomTransfer(t)
	transfer2, err := testQueries.GetTransfer(context.Background(), transfer1.ID)

	require.NoError(t, err)
	require.NotEmpty(t, transfer2)

	require.Equal(t, transfer1.ID, transfer2.ID)
	require.Equal(t, transfer1.FromAccountID, transfer2.FromAccountID)
	require.Equal(t, transfer1.ToAccountID, transfer2.ToAccountID)
	require.Equal(t, transfer1.Amount, transfer2.Amount)

	require.WithinDuration(t, transfer1.CreatedAt, transfer2.CreatedAt, time.Second)

}

func TestListTransfers(t *testing.T) {
	sender := CreateRandomAccount(t)
	for i := 0; i < 5; i++ {
		receiver := CreateRandomAccount(t)
		CreateRandomTransfer(t, sender, receiver)

		arg := ListTransfersParams{
			FromAccountID : sender.ID,
			ToAccountID : receiver.ID,
			Limit : 5,
			Offset : 0,
		}
	}
}
