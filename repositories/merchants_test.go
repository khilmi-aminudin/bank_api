package repositories

import (
	"context"
	"testing"

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
	merchant, err := testStore.CreateMerchant(context.Background(), arg)

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
