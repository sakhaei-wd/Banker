package db

import (
	"context"
	"database/sql"
	"fmt"
)

// Store provides all functions to execute SQL queries and transactions
type Store struct {
	*Queries
	db *sql.DB
}
// TransferTxParams contains the input parameters of the transfer transaction
type TransferTxParams struct {
	FromAccountID int64 `json:"from_account_id"` //FromAccountID is the ID of the account where money will be sent from.
	ToAccountID   int64 `json:"to_account_id"`   //ToAccountID is the ID of the account where money will be sent to.
	Amount        int64 `json:"amount"`          //And the last field is the Amount of money to be sent.
}

// TransferTxResult is the result of the transfer transaction
type TransferTxResult struct {
	Transfer    Transfer `json:"transfer"`     //The created Transfer record.
	FromAccount Account  `json:"from_account"` //The FromAccount after its balance is subtracted.
	ToAccount   Account  `json:"to_account"`   //The ToAccount after its its balance is added.
	FromEntry   Entry    `json:"from_entry"`   //The FromEntry of the account which records that money is moving out of the FromAccount.
	ToEntry     Entry    `json:"to_entry"`     //And the ToEntry of the account which records that money is moving in to the ToAccount.
}

// NewStore creates a new store
func NewStore(db *sql.DB) *Store {
    return &Store{
        db:      db,
        Queries: New(db),
    }
}

// ExecTx executes a function within a database transaction
func (store *Store) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	q := New(tx)
	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
		}
		return err
	}

	return tx.Commit()
}


// TransferTx performs a money transfer from one account to the other.
// It creates the transfer, add account entries, and update accounts' balance within a database transaction
func (store *Store) TransferTx(ctx context.Context, arg TransferTxParams) (TransferTxResult, error) {
	var result TransferTxResult

	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		//First step is to create a transfer record
		result.Transfer, err = q.CreateTransfer(ctx, CreateTransferParams{
			FromAccountID: arg.FromAccountID,
			ToAccountID:   arg.ToAccountID,
			Amount:        arg.Amount,
		})
		if err != nil {
			return err
		}

		//Next step is to add 2 account entries: 1 for the FromAccount, and 1 for the ToAccount.
		result.FromEntry, err = q.CreateEntry(ctx, CreateEntryParams{
			AccountID: arg.FromAccountID,
			Amount:    -arg.Amount,
		})
		if err != nil {
			return err
		}

		result.ToEntry, err = q.CreateEntry(ctx, CreateEntryParams{
			AccountID: arg.ToAccountID,
			Amount:    arg.Amount,
		})
		if err != nil {
			return err
		}

		//TODO : we’re done with the account entries creation. 
		//The last step to update account balance will be more complicated because it involves 
		//locking and preventing potential deadlock.



		return nil
	})

	return result, err
}