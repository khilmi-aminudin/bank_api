package repositories

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/google/uuid"
)

type Repository interface {
	Querier
	TransferTx(ctx context.Context, arg TranferTxParams) (TransferTxResult, error)
	PaymentTx(ctx context.Context, arg PaymentTxParams) (PaymentTxResult, error)
	WithdrawalTx(ctx context.Context, args WithdrawalParams) (WithdrawalResult, error)
	TopupTx(ctx context.Context, args TopupParams) (TopupResult, error)
}

type repository struct {
	*Queries
	db *sql.DB
}

func NewRepo(db *sql.DB) Repository {
	return &repository{
		db:      db,
		Queries: New(db),
	}
}

func (repo *repository) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := repo.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	queries := New(tx)
	err = fn(queries)

	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("err : %v, rb err : %v", err, rbErr)
		}
		return err
	}

	return tx.Commit()
}

type TranferTxParams struct {
	FromAccountID uuid.UUID `json:"from_account_id"`
	ToAccountID   uuid.UUID `json:"to_account_id"`
	Amount        float64   `json:"amount"`
	Description   string    `json:"description"`
}

type TransferTxResult struct {
	TransactionID uuid.UUID `json:"transaction_id"`
	FromAccount   MAccount  `json:"from_account"`
	ToAccount     MAccount  `json:"to_account"`
}

// TransferTx implements Repositories.
func (repo *repository) TransferTx(ctx context.Context, arg TranferTxParams) (TransferTxResult, error) {
	var result TransferTxResult
	err := repo.execTx(ctx, func(q *Queries) error {
		var err error

		result.FromAccount, err = q.AddAccountBalance(ctx, AddAccountBalanceParams{
			Balance: -arg.Amount,
			ID:      arg.FromAccountID,
		})

		if err != nil {
			return err
		}

		result.ToAccount, err = q.AddAccountBalance(ctx, AddAccountBalanceParams{
			Balance: arg.Amount,
			ID:      arg.ToAccountID,
		})

		if err != nil {
			return err
		}

		trxHistory, err := q.CreateTransactionHistory(ctx, CreateTransactionHistoryParams{
			TransactionType: TransactionTypeTransfer,
			FromAccountID:   arg.FromAccountID,
			ToAccountID: uuid.NullUUID{
				UUID:  arg.ToAccountID,
				Valid: true,
			},
			Amount:      arg.Amount,
			Description: arg.Description,
		})

		if err != nil {
			return err
		}

		result.TransactionID = trxHistory.ID

		return err
	})

	return result, err
}

type PaymentTxParams struct {
	FromAccountID uuid.UUID `json:"from_account_id"`
	ToMerchantID  uuid.UUID `json:"to_merchant_id"`
	Amount        float64   `json:"amount"`
	Description   string    `json:"description"`
}

type PaymentTxResult struct {
	TransactionID uuid.UUID `json:"transaction_id"`
	FromAccount   MAccount  `json:"from_account"`
	ToMerchant    MMerchant `json:"to_merchant"`
}

// PaymentTx implements Repositories.
func (repo *repository) PaymentTx(ctx context.Context, arg PaymentTxParams) (PaymentTxResult, error) {
	var result PaymentTxResult

	err := repo.execTx(ctx, func(q *Queries) error {
		var err error

		result.FromAccount, err = q.AddAccountBalance(ctx, AddAccountBalanceParams{
			Balance: -arg.Amount,
			ID:      arg.FromAccountID,
		})

		if err != nil {
			return err
		}

		result.ToMerchant, err = q.AddMerchantBalance(ctx, AddMerchantBalanceParams{
			Balance: arg.Amount,
			ID:      arg.ToMerchantID,
		})

		if err != nil {
			return err
		}

		trxHistory, err := q.CreateTransactionHistory(ctx, CreateTransactionHistoryParams{
			TransactionType: TransactionTypePayment,
			FromAccountID:   arg.FromAccountID,
			ToMerchantID: uuid.NullUUID{
				UUID:  arg.ToMerchantID,
				Valid: true,
			},
			Amount:      arg.Amount,
			Description: arg.Description,
		})

		if err != nil {
			return err
		}

		result.TransactionID = trxHistory.ID

		return err
	})

	return result, err
}

type TopupParams struct {
	ToAccountId uuid.UUID `json:"to_account_id"`
	Amount      float64   `json:"amount"`
	Description string    `json:"description"`
}

type TopupResult struct {
	TransactionID uuid.UUID `json:"to_account_id"`
	Account       MAccount  `json:"account"`
}

// TopupTx implements Repositories.
func (repo *repository) TopupTx(ctx context.Context, args TopupParams) (TopupResult, error) {
	var result TopupResult
	err := repo.execTx(ctx, func(q *Queries) error {
		var err error
		result.Account, err = q.AddAccountBalance(ctx, AddAccountBalanceParams{
			Balance: args.Amount,
			ID:      args.ToAccountId,
		})

		if err != nil {
			return err
		}

		trx, err := q.CreateTransactionHistory(ctx, CreateTransactionHistoryParams{
			TransactionType: TransactionTypeTopup,
			ToAccountID: uuid.NullUUID{
				UUID:  args.ToAccountId,
				Valid: true,
			},
			Amount:      args.Amount,
			Description: args.Description,
		})

		if err != nil {
			return err
		}

		result.TransactionID = trx.ID
		return err
	})
	return result, err
}

type WithdrawalParams struct {
	FromAccountID uuid.UUID `json:"from_account_id"`
	Amount        float64   `json:"amount"`
	Description   string    `json:"description"`
}

type WithdrawalResult struct {
	TransactionID uuid.UUID `json:"to_account_id"`
	Account       MAccount  `json:"account"`
}

// WithdrawalTx implements Repositories.
func (repo *repository) WithdrawalTx(ctx context.Context, args WithdrawalParams) (WithdrawalResult, error) {
	var result WithdrawalResult
	err := repo.execTx(ctx, func(q *Queries) error {
		var err error

		result.Account, err = q.AddAccountBalance(ctx, AddAccountBalanceParams{
			Balance: args.Amount,
			ID:      args.FromAccountID,
		})

		if err != nil {
			return err
		}

		trx, err := q.CreateTransactionHistory(ctx, CreateTransactionHistoryParams{
			TransactionType: TransactionTypeWithdrawal,
			FromAccountID:   args.FromAccountID,
			Amount:          args.Amount,
			Description:     args.Description,
		})

		result.TransactionID = trx.ID
		return err
	})
	return result, err
}
