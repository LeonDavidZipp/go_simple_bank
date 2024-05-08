package db

import (
	"context"
	"database/sql"
	"fmt"
)

type Store interface {
	Querier
	TransferTx(ctx context.Context, arg TransferTxParams) (TransferTxResult, error)
}

// SQLStore provides all functions to execute SQL queries and transactions
type SQLStore struct {
	*Queries
	db *sql.DB
}

// NewStore creates a new Store struct
func NewStore(db *sql.DB) *SQLStore {
	return &SQLStore{
		db : db,
		Queries: New(db),
	}
}

// execTx executes a function within a database transaction
func (store *SQLStore) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	q := New(tx)
	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx err: %v; rb err: %v", err, rbErr)
		}
		return err
	}
	return tx.Commit()
}

// TransferTxParams defines the input parameters for the TransferTx function
type TransferTxParams struct {
	FromAccountID int64 `json:"from_account_id"`
	ToAccountID   int64 `json:"to_account_id"`
	Amount        int64 `json:"amount"`
}

// TransferTxResult defines the output result for the TransferTx function
type TransferTxResult struct {
	Transfer    Transfer `json:"transfer"`
	FromAccount Account `json:"from_account"`
	ToAccount   Account `json:"to_account"`
	FromEntry   Entry `json:"from_entry"`
	ToEntry     Entry `json:"to_entry"`
}

// TransferTx performs a money transfer from one account to the other
// It updates the balance of the 2 accounts in the accounts table, ...
// creates a transfer int the transfer table and ...
// creates 2 entries in the entries table
func (store *SQLStore) TransferTx(ctx context.Context, arg TransferTxParams) (TransferTxResult, error) {
	var result TransferTxResult

	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		result.Transfer, err = q.CreateTransfer(ctx, CreateTransferParams{
			FromAccountID : arg.FromAccountID,
			ToAccountID :   arg.ToAccountID,
			Amount :        arg.Amount,
		})
		if err != nil {
			return err
		}

		result.FromEntry, err = q.CreateEntry(ctx, CreateEntryParams{
			AccountID : arg.FromAccountID,
			Amount :    arg.Amount * -1,
		})
		if err != nil {
			return err
		}

		result.ToEntry, err = q.CreateEntry(ctx, CreateEntryParams{
			AccountID : arg.ToAccountID,
			Amount :    arg.Amount,
		})
		if err != nil {
			return err
		}

		// update account balances
		if arg.FromAccountID < arg.ToAccountID {
			result.FromAccount, result.ToAccount, err = addMoney(
				ctx, q, arg.FromAccountID, arg.Amount * -1, arg.ToAccountID, arg.Amount)
		} else {
			result.FromAccount, result.ToAccount, err = addMoney(
				ctx, q, arg.ToAccountID, arg.Amount, arg.FromAccountID, arg.Amount * -1)
		}
		if err != nil {
			return err
		}

		return nil
	})

	return result, err
}

func addMoney(
	ctx context.Context,
	q *Queries,
	accountID1 int64,
	amount1 int64,
	accountID2 int64,
	amount2 int64,
) (Account, Account, error) {
	account1, err := q.AddAccountBalance(ctx, AddAccountBalanceParams{
		ID:     accountID1,
		Amount: amount1,
	})
	if err != nil {
		return Account{}, Account{}, err
	}
	account2, err := q.AddAccountBalance(ctx, AddAccountBalanceParams{
		ID:     accountID2,
		Amount: amount2,
	})
	if err != nil {
		return Account{}, Account{}, err
	}
	return account1, account2, nil
}
