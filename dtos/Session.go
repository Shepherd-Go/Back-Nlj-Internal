package dtos

import "github.com/google/uuid"

type Session struct {
	ID              uuid.UUID `json:"id"`
	First_Name      string    `json:"first_name"`
	Last_Name       string    `json:"last_name"`
	Email           string    `json:"email"`
	Password        []byte    `json:"password,omitempty"`
	Permissions     string    `json:"permissions"`
	Confirmed_Email *bool     `json:"confirmed_email"`
	Status          *bool     `json:"status,omitempty"`
}
