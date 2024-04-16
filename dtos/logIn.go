package dtos

import "context"

type LogIn struct {
	Identifier string `json:"identifier" mod:"trim" validate:"required"`
	Password   string `json:"password" mod:"trim" validate:"required"`
}

func (e *LogIn) Validate() error {
	_ = conform.Struct(context.Background(), e)
	return validate.Struct(e)
}
