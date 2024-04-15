package dtos

import "context"

type ActivateEmail struct {
	Password string `json:"password" mod:"trim" validate:"required,min=8,max=16"`
}

func (a *ActivateEmail) Validate() error {
	_ = conform.Struct(context.Background(), a)
	return validate.Struct(a)
}
