package dtos

import (
	"context"

	"github.com/go-playground/mold/modifiers"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

var (
	validate = validator.New()
	conform  = modifiers.New()
)

type RegisterEmployee struct {
	FirstName    string `json:"first_name" mod:"trim,lcase" validate:"required,max=15"`
	LastName     string `json:"last_name" mod:"trim,lcase" validate:"required,max=15"`
	Email        string `json:"email" mod:"trim,lcase" validate:"required,email"`
	Phone        string `json:"phone" mod:"trim" validate:"required,len=11"`
	Password     string `json:"password"`
	Permissions  string `json:"permissions" mod:"trim" validate:"required"`
	Code_Bank    string `json:"code_bank" mod:"trim" validate:"required"`
	Pay_Phone    string `json:"pay_phone" mod:"trim" validate:"required,len=11"`
	Payment_Card string `json:"payment_card" mod:"trim" validate:"required"`
	Created_by   string `json:"created_by_by"`
	Updated_by   string `json:"updated_by"`
}

type UpdateEmployee struct {
	ID           uuid.UUID `json:"id"`
	FirstName    string    `json:"first_name" mod:"trim,lcase" validate:"required,max=15"`
	LastName     string    `json:"last_name" mod:"trim,lcase" validate:"required,max=15"`
	Email        string    `json:"email" mod:"trim,lcase" validate:"required,email"`
	Phone        string    `json:"phone" mod:"trim" validate:"required,len=11"`
	Permissions  string    `json:"permissions" mod:"trim" validate:"required"`
	Code_Bank    string    `json:"code_bank" mod:"trim" validate:"required"`
	Pay_Phone    string    `json:"pay_phone" mod:"trim" validate:"required,len=11"`
	Payment_Card string    `json:"payment_card" mod:"trim" validate:"required"`
	Status       *bool     `json:"status" mod:"trim" validate:"required,boolean"`
	Updated_by   string    `json:"updated_by"`
}

func (e *RegisterEmployee) Validate() error {
	_ = conform.Struct(context.Background(), e)
	return validate.Struct(e)
}

func (e *UpdateEmployee) Validate() error {
	_ = conform.Struct(context.Background(), e)
	return validate.Struct(e)
}
