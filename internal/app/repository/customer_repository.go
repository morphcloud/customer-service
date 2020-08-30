package repository

import (
	"context"

	"github.com/morphcloud/customer-service/internal/app/domain"
)

// CustomerRepository
type CustomerRepository interface {
	FindOne(ctx context.Context, customerID string) (domain.Customer, error)
	Register(ctx context.Context, customer domain.Customer) (domain.Customer, error)
}
