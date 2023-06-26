package repositories

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/khilmi-aminudin/bank_api/utils"
)

func createRandomAccount(t *testing.T) MAccount {
	ctx := context.Background()

	customer := createRandomCustomer(t)

	arg := CreateAccountParams{
		CustomerID: customer.ID,
		Number:     utils.RandomNumber(8),
		Balance:    float64(utils.RandomInt(100000, 10000000)),
	}

	account, err := testStore.CreateAccount(ctx, arg)
	require.NoError(t, err)
	require.NotEmpty(t, account)
	require.NotEmpty(t, account.ID)

	require.Equal(t, arg.CustomerID, account.CustomerID)
	require.Equal(t, arg.Number, account.Number)
	require.Equal(t, arg.Balance, account.Balance)

	return account
}

func TestCreateAccount(t *testing.T) {
	createRandomAccount(t)
}

func TestAddAccountBalance(t *testing.T) {
	account1 := createRandomAccount(t)

	arg := AddAccountBalanceParams{
		Balance: float64(utils.RandomInt(100000, 10000000)),
		ID:      account1.ID,
	}

	account2, err := testStore.AddAccountBalance(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account2)

	require.Equal(t, account1.ID, account2.ID)
	require.Equal(t, account1.Number, account2.Number)
	require.Equal(t, account1.CustomerID, account2.CustomerID)
	require.Equal(t, account1.Balance+arg.Balance, account2.Balance)
	require.Equal(t, account1.CreatedAt, account2.CreatedAt)

}
