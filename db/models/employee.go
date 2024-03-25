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
	Confirmed_Email bool
	Cod_Bank        string
	Pay_Phone       string
	Payment_Card    string
	Status          string
	Deleted         bool
	Created_By      string
	Updated_By      string
	Created_At      time.Time
	Updated_At      time.Time
}

func (e *Employee) BuildCreateEmployeeModel(empl dtos.RegisterEmployee) {
	e.ID = uuid.New()
	e.FirstName = empl.FirstName
	e.LastName = empl.LastName
	e.Username = e.FirstName[:3] + e.LastName[:3] + "-" + e.ID.String()[:5]
	e.Email = empl.Email
	e.Phone = empl.Phone
	e.Password = []byte(empl.Password)
	e.Permissions = empl.Permissions
	e.Confirmed_Email = false
	e.Cod_Bank = empl.Cod_Bank
	e.Pay_Phone = empl.Pay_Phone
	e.Payment_Card = empl.Payment_Card
	e.Status = "true"
	e.Deleted = false
	e.Created_By = e.ID.String()
	e.Updated_By = e.ID.String()
	e.Created_At = time.Now()
	e.Updated_At = time.Now()
}

func (e *Employee) ToDomainDTO() dtos.EmployeeResponse {
	return dtos.EmployeeResponse{
		ID:              e.ID,
		FirstName:       e.FirstName,
		LastName:        e.LastName,
		Username:        e.Username,
		Email:           e.Email,
		Phone:           e.Phone,
		Permissions:     parsePermissions(e.Permissions),
		Confirmed_Email: e.Confirmed_Email,
		Cod_Bank:        e.Cod_Bank,
		Pay_Phone:       e.Pay_Phone,
		Payment_Card:    e.Payment_Card,
		Status:          e.Status,
		Created_By:      e.Created_By,
		Updated_By:      e.Updated_By,
		Created_At:      e.Created_At,
		Updated_At:      e.Updated_At,
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
	e.Username = e.FirstName[:3] + e.LastName[:3] + "-" + id.String()[:5]
	e.Email = empl.Email
	e.Phone = empl.Phone
	e.Password = []byte(empl.Password)
	e.Permissions = empl.Permissions
	e.Cod_Bank = empl.Cod_Bank
	e.Pay_Phone = empl.Pay_Phone
	e.Payment_Card = empl.Payment_Card
	e.Status = empl.Status
	e.Updated_By = id.String()
	e.Updated_At = time.Now()
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
