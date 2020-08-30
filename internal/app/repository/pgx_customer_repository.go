package repository

import (
	"context"

	"github.com/jackc/pgx/v4"

	"github.com/morphcloud/customer-service/internal/app/domain"
)

type pgxCustomerRepository struct {
	conn *pgx.Conn
}

// NewPGXCustomerRepository
func NewPGXCustomerRepository(conn *pgx.Conn) CustomerRepository {
	return &pgxCustomerRepository{conn}
}

// FindOne
func (r *pgxCustomerRepository) FindOne(ctx context.Context, customerID string) (domain.Customer, error) {
	// TODO
	return domain.Customer{}, nil
}

// Register
func (r *pgxCustomerRepository) Register(ctx context.Context, customer domain.Customer) (domain.Customer, error) {
	return domain.Customer{}, nil
}
