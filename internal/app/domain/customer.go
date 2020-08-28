package domain

import "time"

type Customer struct {
	ID         string    `json:"id"`
	FirstName  string    `json:"firstName,omitempty"`
	MiddleName string    `json:"middleName,omitempty"`
	LastName   string    `json:"lastName"`
	Email      string    `json:"email"`
	StatusID   int32     `json:"statusId"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
}
