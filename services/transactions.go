package services

import (
	"context"

	"github.com/google/uuid"

	"github.com/khilmi-aminudin/bank_api/repositories"
)

type TransactionsService interface {
	TransferTx(ctx context.Context, args repositories.TranferTxParams) (repositories.TransferTxResult, error)
	PaymentTx(ctx context.Context, args repositories.PaymentTxParams) (repositories.PaymentTxResult, error)
	WithdrawalTx(ctx context.Context, args repositories.WithdrawalParams) (repositories.WithdrawalResult, error)
	TopupTx(ctx context.Context, args repositories.TopupParams) (repositories.TopupResult, error)
	GetTransactionHistory(ctx context.Context, fromAccountId uuid.UUID) ([]repositories.TransactionHistory, error)
	GetTransactionHistoryByType(ctx context.Context, args repositories.GetTransactionHistoryByTypeParams) ([]repositories.TransactionHistory, error)
}

type transactionsService struct {
	repository repositories.Repository
}

func NewTransactionsService(repository repositories.Repository) TransactionsService {
	return &transactionsService{
		repository: repository,
	}
}

// GetTransactionHistory implements TransactionsService.
func (s *transactionsService) GetTransactionHistory(ctx context.Context, fromAccountId uuid.UUID) ([]repositories.TransactionHistory, error) {
	return s.repository.GetTransactionHistory(ctx, fromAccountId)
}

// GetTransactionHistoryByType implements TransactionsService.
func (s *transactionsService) GetTransactionHistoryByType(ctx context.Context, args repositories.GetTransactionHistoryByTypeParams) ([]repositories.TransactionHistory, error) {
	return s.repository.GetTransactionHistoryByType(ctx, args)
}

// PaymentTx implements TransactionsService.
func (s *transactionsService) PaymentTx(ctx context.Context, args repositories.PaymentTxParams) (repositories.PaymentTxResult, error) {
	return s.repository.PaymentTx(ctx, args)
}

// TopupTx implements TransactionsService.
func (s *transactionsService) TopupTx(ctx context.Context, args repositories.TopupParams) (repositories.TopupResult, error) {
	return s.repository.TopupTx(ctx, args)
}

// TransferTx implements TransactionsService.
func (s *transactionsService) TransferTx(ctx context.Context, args repositories.TranferTxParams) (repositories.TransferTxResult, error) {
	return s.repository.TransferTx(ctx, args)
}

// WithdrawalTx implements TransactionsService.
func (s *transactionsService) WithdrawalTx(ctx context.Context, args repositories.WithdrawalParams) (repositories.WithdrawalResult, error) {
	return s.repository.WithdrawalTx(ctx, args)
}
