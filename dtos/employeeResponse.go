package dtos

import (
	"time"

	"github.com/google/uuid"
)

type Employees []EmployeeResponse

type EmployeeResponse struct {
	ID              uuid.UUID `json:"id"`
	FirstName       string    `json:"first_name"`
	LastName        string    `json:"last_name"`
	Username        string    `json:"username"`
	Email           string    `json:"email"`
	Password        []byte    `json:"password,omitempty"`
	Phone           string    `json:"phone"`
	Permissions     string    `json:"permissions"`
	Confirmed_Email *bool     `json:"confirmed_email"`
	Status          *bool     `json:"status"`
	Created_At      time.Time `json:"created_at"`
	Updated_At      time.Time `json:"updated_at"`
}

func (e *Employees) Add(employee EmployeeResponse) {
	*e = append(*e, employee)
}
