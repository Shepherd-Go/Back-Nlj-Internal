package dtos

import (
	"context"

	"github.com/google/uuid"
)

type PaymentMobile struct {
	ID           uuid.UUID `json:"id" mod:"trim" validate:"required,uuid4"`
	Bank         string    `json:"bank" mod:"trim,lcase" validate:"required"`
	Phone        string    `json:"phone" mod:"trim,lcase" validate:"required,len=11"`
	Payment_Card string    `json:"payment_card" mod:"trim,lcase" validate:"required,min=7,max=8"`
	EmployeeID   uuid.UUID `json:"employee_id"`
}

func (p *PaymentMobile) Validate() error {
	_ = conform.Struct(context.Background(), p)
	return validate.Struct(p)
}
