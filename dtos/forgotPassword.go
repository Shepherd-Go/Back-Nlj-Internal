package dtos

import "context"

type ForgotPassword struct {
	Email string `json:"email" mod:"trim,lcase" validate:"required,email"`
}

func (f *ForgotPassword) Validate() error {
	_ = conform.Struct(context.Background(), f)
	return validate.Struct(f)
}
