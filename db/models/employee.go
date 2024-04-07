package models

import (
	"time"

	"github.com/BBCompanyca/Back-Nlj-Internal.git/dtos"
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
	Code_Bank       string
	Pay_Phone       string
	Payment_Card    string
	Status          *bool
	Deleted         *bool
	Created_By      string
	Updated_By      string
	Deleted_By      *string
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       *time.Time
}

func (e *Employee) BuildCreateEmployeeModel(empl dtos.RegisterEmployee) {

	isFalse, isTrue := false, true

	e.ID = uuid.New()
	e.FirstName = empl.FirstName
	e.LastName = empl.LastName
	e.Username = e.FirstName[:3] + e.LastName[:3] + "-" + e.ID.String()[:5]
	e.Email = empl.Email
	e.Phone = empl.Phone
	e.Password = []byte(empl.Password)
	e.Permissions = empl.Permissions
	e.Confirmed_Email = &isFalse
	e.Code_Bank = empl.Code_Bank
	e.Pay_Phone = empl.Pay_Phone
	e.Payment_Card = empl.Payment_Card
	e.Status = &isTrue
	e.Deleted = &isFalse
	e.Created_By = e.ID.String()
	e.Updated_By = e.ID.String()
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
		Code_Bank:       e.Code_Bank,
		Pay_Phone:       e.Pay_Phone,
		Payment_Card:    e.Payment_Card,
		Status:          e.Status,
		Created_By:      e.Created_By,
		Updated_By:      e.Updated_By,
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

func (e *Employee) BuildUpdatedEmployeeModel(empl dtos.UpdateEmployee, id uuid.UUID) {
	e.FirstName = empl.FirstName
	e.LastName = empl.LastName
	e.Username = e.FirstName[:3] + e.LastName[:3] + "-" + empl.ID.String()[:5]
	e.Email = empl.Email
	e.Phone = empl.Phone
	e.Permissions = empl.Permissions
	e.Code_Bank = empl.Code_Bank
	e.Pay_Phone = empl.Pay_Phone
	e.Payment_Card = empl.Payment_Card
	e.Status = &empl.Status
	e.Updated_By = id.String()
}

func parsePermissions(permissions string) string {
	switch permissions {
	case "1":
		return "administrator"
	case "2":
		return "seller"
	default:
		return ""
	}
}
