package models

import (
	"time"

	"github.com/Shepherd-Go/Back-Nlj-Internal.git/dtos"
	"github.com/google/uuid"
)

type Employees []Employee

type Employee struct {
	ID              uuid.UUID
	FirstName       string
	LastName        string
	Username        string
	Email           string
	Phone           string
	Password        []byte
	Permissions     string
	Confirmed_Email *bool
	Status          *bool
	Deleted         *bool
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       *time.Time
}

func (e *Employee) BuildCreateEmployeeModel(empl dtos.RegisterEmployee) {

	isFalse, isTrue := false, true

	e.ID = empl.ID
	e.FirstName = empl.FirstName
	e.LastName = empl.LastName
	e.Username = e.FirstName[:3] + e.LastName[:3] + "-" + e.ID.String()[:5]
	e.Email = empl.Email
	e.Phone = empl.Phone
	e.Password = []byte(empl.Password)
	e.Permissions = empl.Permissions
	e.Confirmed_Email = &isFalse
	e.Status = &isTrue
	e.Deleted = &isFalse
}

func (e *Employee) ToDomainDTO() dtos.EmployeeResponse {
	return dtos.EmployeeResponse{
		ID:              e.ID,
		FirstName:       e.FirstName,
		LastName:        e.LastName,
		Username:        e.Username,
		Email:           e.Email,
		Phone:           e.Phone,
		Password:        e.Password,
		Permissions:     parsePermissions(e.Permissions),
		Confirmed_Email: e.Confirmed_Email,
		Status:          e.Status,
		Created_At:      e.CreatedAt,
		Updated_At:      e.UpdatedAt,
	}
}

func (e *Employees) ToDomainDTO() dtos.Employees {
	var employee dtos.Employees

	for _, v := range *e {
		employee.Add(v.ToDomainDTO())
	}

	return employee
}

func (e *Employee) BuildUpdatedEmployeeModel(empl dtos.UpdateEmployee) {
	e.FirstName = empl.FirstName
	e.LastName = empl.LastName
	e.Username = e.FirstName[:3] + e.LastName[:3] + "-" + empl.ID.String()[:5]
	e.Email = empl.Email
	e.Phone = empl.Phone
	e.Permissions = empl.Permissions
	e.Status = empl.Status
}

func parsePermissions(permissions string) string {
	switch permissions {
	case "1":
		return "administrador"
	case "2":
		return "vendedor"
	default:
		return ""
	}
}
