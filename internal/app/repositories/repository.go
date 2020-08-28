package repositories

import (
	"context"

	"github.com/morphcloud/customer-service/internal/app/domain"
)

// CustomerRepository
type CustomerRepository interface {
	FindOne(ctx context.Context, customerID string) (domain.Customer, error)
}
