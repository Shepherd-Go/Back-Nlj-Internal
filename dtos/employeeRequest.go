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
	ID          uuid.UUID `json:"id" mod:"trim" validate:"required,uuid4"`
	FirstName   string    `json:"first_name" mod:"trim,lcase" validate:"required,max=15"`
	LastName    string    `json:"last_name" mod:"trim,lcase" validate:"required,max=15"`
	Email       string    `json:"email" mod:"trim,lcase" validate:"required,email"`
	Phone       string    `json:"phone" mod:"trim" validate:"required,len=11"`
	Password    string    `json:"password"`
	Permissions string    `json:"permissions" mod:"trim" validate:"required"`
}

type UpdateEmployee struct {
	ID          uuid.UUID `json:"id"`
	FirstName   string    `json:"first_name" mod:"trim,lcase" validate:"required,max=15"`
	LastName    string    `json:"last_name" mod:"trim,lcase" validate:"required,max=15"`
	Email       string    `json:"email" mod:"trim,lcase" validate:"required,email"`
	Phone       string    `json:"phone" mod:"trim" validate:"required,len=11"`
	Permissions string    `json:"permissions" mod:"trim" validate:"required"`
	Status      *bool     `json:"status" mod:"trim" validate:"required,boolean"`
}

func (e *RegisterEmployee) Validate() error {
	_ = conform.Struct(context.Background(), e)
	return validate.Struct(e)
}

func (e *UpdateEmployee) Validate() error {
	_ = conform.Struct(context.Background(), e)
	return validate.Struct(e)
}
