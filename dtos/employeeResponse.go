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
	Phone           string    `json:"phone"`
	Permissions     string    `json:"permissios"`
	Confirmed_Email bool      `json:"confirmed_email"`
	Cod_Bank        string    `json:"cod_bank"`
	Pay_Phone       string    `json:"pay_phone"`
	Payment_Card    string    `json:"payment_card"`
	Status          string    `json:"status"`
	Created_By      string    `json:"created_by"`
	Updated_By      string    `json:"updated_by"`
	Created_At      time.Time `json:"created_at"`
	Updated_At      time.Time `json:"updated_at"`
}

func (e *Employees) Add(employee EmployeeResponse) {
	*e = append(*e, employee)
}
