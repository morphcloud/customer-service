package domain

import "time"

type Customer struct {
	ID         string    `json:"id"`
	FirstName  string    `json:"first_name"`
	MiddleName string    `json:"middle_name,omitempty"`
	LastName   string    `json:"last_name"`
	Email      string    `json:"email"`
	Password   string    `json:"-"`
	StatusID   int32     `json:"status_id"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
