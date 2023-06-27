package repositories

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/khilmi-aminudin/bank_api/utils"
)

func createRandomMerchant(t *testing.T) MMerchant {
	arg := CreateMerchantParams{
		Name:    utils.RandomString(8),
		Address: utils.RandomString(10),
		Website: utils.RandomString(15),
		Email:   utils.RandomEmail(),
	}
	merchant, err := testRepo.CreateMerchant(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, merchant)
	require.NotEmpty(t, merchant.ID)

	require.Equal(t, arg.Name, merchant.Name)
	require.Equal(t, arg.Address, merchant.Address)
	require.Equal(t, arg.Website, merchant.Website)
	require.Equal(t, arg.Email, merchant.Email)

	return merchant
}

func TestCreateMerchant(t *testing.T) {
	createRandomMerchant(t)
}

func TestUpdateMerchant(t *testing.T) {
	merchant1 := createRandomMerchant(t)

	arg := UpdateMerchantParams{
		ID:        merchant1.ID,
		Name:      utils.RandomString(10),
		Address:   utils.RandomString(15),
		Website:   utils.RandomString(10),
		UpdatedAt: time.Now(),
	}

	err := testRepo.UpdateMerchant(context.Background(), arg)
	require.NoError(t, err)

	merchant2, err := testRepo.GetMerchantById(context.Background(), arg.ID)
	require.NoError(t, err)
	require.NotEmpty(t, merchant2)

	require.Equal(t, arg.ID, merchant2.ID)
	require.Equal(t, arg.Name, merchant2.Name)
	require.Equal(t, arg.Address, merchant2.Address)
	require.Equal(t, arg.Website, merchant2.Website)
	require.Equal(t, arg.UpdatedAt.Second(), merchant2.UpdatedAt.Second())

	require.Equal(t, merchant1.Balance, merchant2.Balance)
	require.Equal(t, merchant1.Email, merchant2.Email)
}

func TestGetAllMerchant(t *testing.T) {
	n := 5

	for i := 0; i < n; i++ {
		createRandomMerchant(t)
	}

	merchants, err := testRepo.GetAllMerchants(context.Background())
	require.NoError(t, err)
	require.NotEmpty(t, merchants)

	for _, merchant := range merchants {
		require.NotEmpty(t, merchant)
		require.NotZero(t, merchant.ID)
	}
}

func TestGetMerchantById(t *testing.T) {
	merchant1 := createRandomMerchant(t)

	merchant2, err := testRepo.GetMerchantById(context.Background(), merchant1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, merchant2)

	require.Equal(t, merchant1.ID, merchant2.ID)
	require.Equal(t, merchant1.Name, merchant2.Name)
	require.Equal(t, merchant1.Address, merchant2.Address)
	require.Equal(t, merchant1.Website, merchant2.Website)
	require.Equal(t, merchant1.Balance, merchant2.Balance)
	require.Equal(t, merchant1.Email, merchant2.Email)
	require.Equal(t, merchant1.UpdatedAt, merchant2.UpdatedAt)
}

func TestGetMerchantByName(t *testing.T) {
	merchant1 := createRandomMerchant(t)

	merchant2, err := testRepo.GetMerchantByName(context.Background(), merchant1.Name)
	require.NoError(t, err)
	require.NotEmpty(t, merchant2)

	require.Equal(t, merchant1.ID, merchant2.ID)
	require.Equal(t, merchant1.Name, merchant2.Name)
	require.Equal(t, merchant1.Address, merchant2.Address)
	require.Equal(t, merchant1.Website, merchant2.Website)
	require.Equal(t, merchant1.Balance, merchant2.Balance)
	require.Equal(t, merchant1.Email, merchant2.Email)
	require.Equal(t, merchant1.UpdatedAt, merchant2.UpdatedAt)
}

func TestAddWalletBalance(t *testing.T) {
	merchant1 := createRandomMerchant(t)

	arg := AddMerchantBalanceParams{
		Balance: float64(utils.RandomInt(10000, 900000)),
		ID:      merchant1.ID,
	}

	merchant2, err := testRepo.AddMerchantBalance(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, merchant2)

	require.Equal(t, arg.ID, merchant2.ID)
	require.Equal(t, merchant1.Name, merchant2.Name)
	require.Equal(t, merchant1.Address, merchant2.Address)
	require.Equal(t, merchant1.Website, merchant2.Website)
	require.Equal(t, arg.Balance, merchant2.Balance-merchant1.Balance)
	require.Equal(t, merchant1.Email, merchant2.Email)
}
