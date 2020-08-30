package service

import (
	"context"

	"github.com/morphcloud/customer-service/internal/app/domain"
	"github.com/morphcloud/customer-service/internal/app/repository"
)

// CustomerService
type CustomerService interface {
	FindOne(ctx context.Context, customerID string) (domain.Customer, error)
}

type customerService struct {
	repository repository.CustomerRepository
}

// NewCustomerService
func NewCustomerService(repository repository.CustomerRepository) CustomerService {
	return &customerService{
		repository,
	}
}

func (s *customerService) FindOne(ctx context.Context, customerID string) (domain.Customer, error) {
	return s.repository.FindOne(ctx, customerID)
}
