package services

import (
	"context"

	"github.com/google/uuid"

	"github.com/khilmi-aminudin/bank_api/repositories"
)

type CustomerService interface {
	CreateCustomer(ctx context.Context, args repositories.CreateCustomerParams) (repositories.MCustomer, error)
	UpdateCustomer(ctx context.Context, args repositories.UpdateCustomerParams) (repositories.MCustomer, error)
	GetAllCustomers(ctx context.Context, args repositories.GetAllCustomersParams) ([]repositories.MCustomer, error)
	GetCustomerById(ctx context.Context, id uuid.UUID) (repositories.MCustomer, error)
	GetCustomerByEmail(ctx context.Context, email string) (repositories.GetCustomerByEmailRow, error)
	GetCustomerByUsername(ctx context.Context, username string) (repositories.GetCustomerByUsernameRow, error)
}

type customerService struct {
	repository repositories.Repository
}

func NewCustomerService(repository repositories.Repository) CustomerService {
	return &customerService{
		repository: repository,
	}
}

// UpdateCustomer implements CustomerService.
func (s *customerService) UpdateCustomer(ctx context.Context, args repositories.UpdateCustomerParams) (repositories.MCustomer, error) {
	if args.IDCardFile != "" && args.IDCardType != "" && args.IDCardNumber != "" {
		args.Status = repositories.CustomerEnumActive
	} else {
		args.Status = repositories.CustomerEnumInactive
	}
	return s.repository.UpdateCustomer(ctx, args)
}

// CreateCustomer implements CustomerService.
func (s *customerService) CreateCustomer(ctx context.Context, args repositories.CreateCustomerParams) (repositories.MCustomer, error) {
	customer, err := s.repository.CreateCustomer(ctx, args)
	return customer, err
}

// GetAllCustomers implements CustomerService.
func (s *customerService) GetAllCustomers(ctx context.Context, args repositories.GetAllCustomersParams) ([]repositories.MCustomer, error) {
	return s.repository.GetAllCustomers(ctx, args)
}

// GetCustomerByEmail implements CustomerService.
func (s *customerService) GetCustomerByEmail(ctx context.Context, email string) (repositories.GetCustomerByEmailRow, error) {
	return s.repository.GetCustomerByEmail(ctx, email)
}

// GetCustomerById implements CustomerService.
func (s *customerService) GetCustomerById(ctx context.Context, id uuid.UUID) (repositories.MCustomer, error) {
	return s.repository.GetCustomerById(ctx, id)
}

// GetCustomerByUsername implements CustomerService.
func (s *customerService) GetCustomerByUsername(ctx context.Context, username string) (repositories.GetCustomerByUsernameRow, error) {
	return s.repository.GetCustomerByUsername(ctx, username)
}
