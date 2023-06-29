package services

import (
	"context"

	"github.com/khilmi-aminudin/bank_api/repositories"
	"github.com/khilmi-aminudin/bank_api/utils"
)

type AccoountService interface {
	CreateAccount(ctx context.Context, args repositories.CreateAccountParams) (repositories.MAccount, error)
	GetAccountByNumber(ctx context.Context, accountnumber string) (repositories.MAccount, error)
	AddAccountBalance(ctx context.Context, args AddAccountBalanceParams) (repositories.MAccount, error)
}

type accountService struct {
	repository repositories.Repository
}

func NewAccountService(repository repositories.Repository) AccoountService {
	return &accountService{
		repository: repository,
	}
}

type AddAccountBalanceParams struct {
	AccountNumber string  `json:"account_number"`
	Balance       float64 `json:"balance"`
}

// AddAccountBalance implements AccoountService.
func (s *accountService) AddAccountBalance(ctx context.Context, args AddAccountBalanceParams) (repositories.MAccount, error) {
	account, err := s.repository.GetAccountByNumber(ctx, args.AccountNumber)
	if err != nil {
		return repositories.MAccount{}, err
	}
	return s.repository.AddAccountBalance(ctx, repositories.AddAccountBalanceParams{
		Balance: args.Balance,
		ID:      account.ID,
	})
}

// GetAccountByNumber implements AccoountService.
func (s *accountService) GetAccountByNumber(ctx context.Context, accountnumber string) (repositories.MAccount, error) {
	return s.repository.GetAccountByNumber(ctx, accountnumber)
}

// CreateAccount implements AccoountService.
func (s *accountService) CreateAccount(ctx context.Context, args repositories.CreateAccountParams) (repositories.MAccount, error) {
	args.Number = utils.RandomNumber(10)
	return s.repository.CreateAccount(ctx, args)
}
