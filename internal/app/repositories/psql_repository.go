package repositories

import (
	"context"
	"database/sql"

	"github.com/morphcloud/customer-service/internal/app/domain"
)

type psqlCustomerRepository struct {
	client *sql.DB
}

// NewPSQLOrderRepository
func NewPSQLCustomerRepository(client *sql.DB) CustomerRepository {
	return &psqlCustomerRepository{client}
}

// FindOne
func (r *psqlCustomerRepository) FindOne(ctx context.Context, customerID string) (domain.Customer, error) {
	// TODO
	return domain.Customer{}, nil
}
