package repositories

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/khilmi-aminudin/bank_api/utils"
)

func createRandomCustomer(t *testing.T) MCustomer {
	ctx := context.Background()

	arg := CreateCustomerParams{
		IDCardType:   IDCardTypeKTP,
		IDCardNumber: utils.RandomNumber(10),
		IDCardFile:   utils.RandomString(10),
		FirstName:    utils.RandomString(5),
		LastName:     utils.RandomString(8),
		PhoneNumber:  utils.RandomNumber(12),
		Email:        utils.RandomEmail(),
		Username:     utils.RandomUsername(),
		Password:     utils.RandomString(8),
	}

	customer, err := testRepo.CreateCustomer(ctx, arg)
	require.NoError(t, err)
	require.NotEmpty(t, customer)
	require.NotEmpty(t, customer.ID)
	require.Equal(t, RoleUser, customer.Role)

	require.Equal(t, arg.IDCardType, customer.IDCardType)
	require.Equal(t, arg.IDCardNumber, customer.IDCardNumber)
	require.Equal(t, arg.IDCardFile, customer.IDCardFile)
	require.Equal(t, arg.FirstName, customer.FirstName)
	require.Equal(t, arg.LastName, customer.LastName)
	require.Equal(t, arg.PhoneNumber, customer.PhoneNumber)
	require.Equal(t, arg.Email, customer.Email)
	require.Equal(t, arg.Username, customer.Username)
	require.Equal(t, arg.Password, customer.Password)
	require.Equal(t, CustomerEnumPending, customer.Status)

	return customer
}

func TestCreateCustomer(t *testing.T) {
	createRandomCustomer(t)
}

func TestGetCustomerById(t *testing.T) {
	customer1 := createRandomCustomer(t)
	ctx := context.Background()

	customer2, err := testRepo.GetCustomerById(ctx, customer1.ID)

	require.NoError(t, err)
	require.NotEmpty(t, customer2)

	require.Equal(t, customer1.Role, customer2.Role)
	require.Equal(t, customer1.Email, customer2.Email)
	require.Equal(t, customer1.Username, customer2.Username)
	require.Equal(t, customer1.Password, customer2.Password)
}

func TestGetCustomerByEmail(t *testing.T) {
	customer1 := createRandomCustomer(t)
	ctx := context.Background()

	customer2, err := testRepo.GetCustomerByEmail(ctx, customer1.Email)

	require.NoError(t, err)
	require.NotEmpty(t, customer2)

	require.Equal(t, customer1.Role, customer2.Role)
	require.Equal(t, customer1.Email, customer2.Email)
	require.Equal(t, customer1.Username, customer2.Username)
	require.Equal(t, customer1.Password, customer2.Password)
}

func TestGetCustomerByUsername(t *testing.T) {
	customer1 := createRandomCustomer(t)
	ctx := context.Background()

	customer2, err := testRepo.GetCustomerByUsername(ctx, customer1.Username)

	require.NoError(t, err)
	require.NotEmpty(t, customer2)

	require.Equal(t, customer1.Role, customer2.Role)
	require.Equal(t, customer1.Email, customer2.Email)
	require.Equal(t, customer1.Username, customer2.Username)
	require.Equal(t, customer1.Password, customer2.Password)
}

func TestGetAllCustomers(t *testing.T) {
	var n int32 = 10

	for i := int32(0); i < n; i++ {
		createRandomCustomer(t)
	}

	ctx := context.Background()

	customers, err := testRepo.GetAllCustomers(ctx, GetAllCustomersParams{
		Limit:  n / 2,
		Offset: n / 2,
	})
	require.NoError(t, err)
	require.NotEmpty(t, customers)
	require.Equal(t, int(n/2), len(customers))
}

func TestUpdateCustomer(t *testing.T) {
	customer1 := createRandomCustomer(t)

	arg := UpdateCustomerParams{
		ID:           customer1.ID,
		IDCardType:   IDCardTypeKTP,
		IDCardNumber: utils.RandomNumber(12),
		IDCardFile:   utils.RandomString(10),
		Status:       CustomerEnumActive,
	}

	customer2, err := testRepo.UpdateCustomer(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, customer2)
	require.NotEmpty(t, customer2.ID)
	require.Equal(t, customer1.Role, customer2.Role)

	require.Equal(t, customer1.IDCardType, customer2.IDCardType)
	require.Equal(t, arg.IDCardNumber, customer2.IDCardNumber)
	require.Equal(t, arg.IDCardFile, customer2.IDCardFile)
	require.Equal(t, customer1.FirstName, customer2.FirstName)
	require.Equal(t, customer1.LastName, customer2.LastName)
	require.Equal(t, customer1.PhoneNumber, customer2.PhoneNumber)
	require.Equal(t, customer1.Email, customer2.Email)
	require.Equal(t, customer1.Username, customer2.Username)
	require.Equal(t, customer1.Password, customer2.Password)
	require.Equal(t, arg.Status, customer2.Status)
}
