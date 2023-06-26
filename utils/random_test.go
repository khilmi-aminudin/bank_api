package utils

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRandomInt(t *testing.T) {
	randomInt := RandomInt(3, 100)
	require.GreaterOrEqual(t, randomInt, int64(3))
	require.LessOrEqual(t, randomInt, int64(100))
}

func TestRandomUsername(t *testing.T) {
	owner := RandomUsername()
	require.Equal(t, 5, len(owner))
}

func TestRandomEmail(t *testing.T) {
	email := RandomEmail()
	require.Contains(t, email, "@email.com")
}

func TestRandomTransactionType(t *testing.T) {
	var trxTypes = []string{"TRANSFER", "TOPUP", "WITHDRAWAL", "PAYMENT"}
	trxType := RandomTransactionTypes()
	require.NotEmpty(t, trxType)
	require.Contains(t, trxTypes, trxType)
}
