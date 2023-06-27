package repositories

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/khilmi-aminudin/bank_api/utils"
)

func createRandomTransactionHistory(t *testing.T, trxType string) TransactionHistory {
	account1 := createRandomAccount(t)
	arg := CreateTransactionHistoryParams{
		TransactionType: TransactionType(trxType),
		FromAccountID:   account1.ID,
		Amount:          float64(utils.RandomInt(50000, 100000)),
		Description:     utils.RandomString(100),
	}

	if trxType == string(TransactionTypePayment) {
		merchant := createRandomMerchant(t)
		arg.ToMerchantID.UUID = merchant.ID
		arg.ToMerchantID.Valid = true

	}

	if trxType == string(TransactionTypeTransfer) {
		account2 := createRandomAccount(t)
		arg.ToAccountID.UUID = account2.ID
		arg.ToAccountID.Valid = true
	}

	trx, err := testRepo.CreateTransactionHistory(context.Background(), arg)

	require.NoError(t, err)

	return trx
}

func TestCreateTransactionHistoryTransfer(t *testing.T) {
	createRandomTransactionHistory(t, string(TransactionTypeTransfer))
}

func TestCreateTransactionHistoryPayment(t *testing.T) {
	createRandomTransactionHistory(t, string(TransactionTypePayment))
}

func TestGetTransactionHistory(t *testing.T) {
	trx1 := createRandomTransactionHistory(t, string(TransactionTypeTransfer))

	transactions, err := testRepo.GetTransactionHistory(context.Background(), trx1.FromAccountID)
	require.NoError(t, err)
	require.NotZero(t, len(transactions))

	for _, trx := range transactions {
		require.NotEmpty(t, trx)
		require.NotEmpty(t, trx.ID)
	}
}

func TestGetTransactionHistoryByType(t *testing.T) {
	_ = createRandomTransactionHistory(t, string(TransactionTypeTransfer))

	transactions, err := testRepo.GetTransactionHistoryByType(context.Background(), TransactionTypeTransfer)
	require.NoError(t, err)
	require.NotZero(t, len(transactions))

	for _, trx := range transactions {
		require.NotEmpty(t, trx)
		require.NotEmpty(t, trx.ID)
	}
}
