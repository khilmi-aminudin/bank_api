package services

import (
	"context"
	"time"

	"github.com/google/uuid"

	"github.com/khilmi-aminudin/bank_api/repositories"
)

type MerchantService interface {
	CreateMerchant(ctx context.Context, args repositories.CreateMerchantParams) (repositories.MMerchant, error)
	UpdateMerchant(ctx context.Context, args repositories.UpdateMerchantParams) error
	GetAllMerchants(ctx context.Context) ([]repositories.MMerchant, error)
	GetMerchantByName(ctx context.Context, merchantName string) (repositories.MMerchant, error)
	GetMerchantById(ctx context.Context, id uuid.UUID) (repositories.MMerchant, error)
	AddMerchantBalance(ctx context.Context, args repositories.AddMerchantBalanceParams) (repositories.MMerchant, error)
}

type merchantService struct {
	repository repositories.Repository
}

func NewMerchantService(repository repositories.Repository) MerchantService {
	return &merchantService{
		repository: repository,
	}
}

// AddMerchantBalance implements MerchantService.
func (s *merchantService) AddMerchantBalance(ctx context.Context, args repositories.AddMerchantBalanceParams) (repositories.MMerchant, error) {
	return s.repository.AddMerchantBalance(ctx, args)
}

// CreateMerchant implements MerchantService.
func (s *merchantService) CreateMerchant(ctx context.Context, args repositories.CreateMerchantParams) (repositories.MMerchant, error) {
	return s.repository.CreateMerchant(ctx, args)
}

// GetAllMerchants implements MerchantService.
func (s *merchantService) GetAllMerchants(ctx context.Context) ([]repositories.MMerchant, error) {
	return s.repository.GetAllMerchants(ctx)
}

// GetMerchantById implements MerchantService.
func (s *merchantService) GetMerchantById(ctx context.Context, id uuid.UUID) (repositories.MMerchant, error) {
	return s.repository.GetMerchantById(ctx, id)
}

// GetMerchantByName implements MerchantService.
func (s *merchantService) GetMerchantByName(ctx context.Context, merchantName string) (repositories.MMerchant, error) {
	return s.repository.GetMerchantByName(ctx, merchantName)
}

// UpdateMerchant implements MerchantService.
func (s *merchantService) UpdateMerchant(ctx context.Context, args repositories.UpdateMerchantParams) error {
	args.UpdatedAt = time.Now()
	return s.repository.UpdateMerchant(ctx, args)
}
