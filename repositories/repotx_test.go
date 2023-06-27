package repositories

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/khilmi-aminudin/bank_api/utils"
)

func TestPaymentTx(t *testing.T) {
	account := createRandomAccount(t)
	merchant := createRandomMerchant(t)

	arg := PaymentTxParams{
		FromAccountID: account.ID,
		ToMerchantID:  merchant.ID,
		Amount:        float64(utils.RandomInt(10000, 900000)),
		Description:   "random payment for test" + utils.RandomString(10),
	}

	result, err := testRepo.PaymentTx(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, result)
	require.NotZero(t, result.TransactionID)

	trsHistories, err := testRepo.GetTransactionHistory(context.Background(), result.FromAccount.ID)
	require.NoError(t, err)
	require.NotEmpty(t, trsHistories)

	var created bool = false
	var createdTx TransactionHistory
loop:
	for _, trx := range trsHistories {
		require.NotEmpty(t, trx)
		if trx.ID == result.TransactionID {
			createdTx = trx
			created = true
			break loop
		}
	}

	require.Equal(t, true, created)

	require.Equal(t, arg.Amount, createdTx.Amount)
	require.Equal(t, arg.Description, createdTx.Description)
	require.Equal(t, arg.FromAccountID, createdTx.FromAccountID)
	require.Equal(t, arg.ToMerchantID, createdTx.ToMerchantID.UUID)
}

func TestTransfeTx(t *testing.T) {
	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)

	arg := TranferTxParams{
		FromAccountID: account1.ID,
		ToAccountID:   account2.ID,
		Amount:        float64(utils.RandomInt(10000, 900000)),
		Description:   "random transfer for test" + utils.RandomString(10),
	}

	result, err := testRepo.TransferTx(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, result)
	require.NotZero(t, result.TransactionID)

	trsHistories, err := testRepo.GetTransactionHistory(context.Background(), result.FromAccount.ID)
	require.NoError(t, err)
	require.NotEmpty(t, trsHistories)

	var created bool = false
	var createdTx TransactionHistory
loop:
	for _, trx := range trsHistories {
		require.NotEmpty(t, trx)
		if trx.ID == result.TransactionID {
			createdTx = trx
			created = true
			break loop
		}
	}

	require.Equal(t, true, created)

	require.Equal(t, arg.Amount, createdTx.Amount)
	require.Equal(t, arg.Description, createdTx.Description)
	require.Equal(t, arg.FromAccountID, createdTx.FromAccountID)
	require.Equal(t, arg.ToAccountID, createdTx.ToAccountID.UUID)
}
