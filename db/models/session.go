package models

import (
	"github.com/Shepherd-Go/Back-Nlj-Internal.git/dtos"
	"github.com/google/uuid"
)

type Session struct {
	ID             uuid.UUID
	FirstName      string
	LastName       string
	Email          string
	Password       []byte
	Permissions    string
	ConfirmedEmail *bool
	Status         *bool
}

func (s *Session) ToDomainDTO() dtos.Session {
	return dtos.Session{
		ID:              s.ID,
		First_Name:      s.FirstName,
		Last_Name:       s.LastName,
		Email:           s.Email,
		Password:        s.Password,
		Permissions:     parsePermissions(s.Permissions),
		Confirmed_Email: s.ConfirmedEmail,
		Status:          s.Status,
	}
}
