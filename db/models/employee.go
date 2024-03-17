package models

import (
	"time"

	"github.com/BBCompanyca/Back-Nlj-Internal.git/dto"
	"github.com/google/uuid"
)

type Employee struct {
	ID              string
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
	Status          bool
	Created_By      string
	Updated_By      string
	Created_At      time.Time
	Updated_At      time.Time
}

func (e *Employee) BuildCreateEmployeeModel(empl dto.EmployeeRequest) {

	uuid := uuid.NewString()

	e.ID = uuid
	e.FirstName = empl.FirstName
	e.LastName = empl.LastName
	e.Username = empl.Username
	e.Email = empl.Email
	e.Phone = empl.Phone
	e.Password = []byte(empl.Password)
	e.Permissions = empl.Permissions
	e.Confirmed_Email = false
	e.Cod_Bank = empl.Cod_Bank
	e.Pay_Phone = empl.Pay_Phone
	e.Payment_Card = empl.Payment_Card
	e.Status = true
	e.Created_By = uuid
	e.Updated_By = uuid
	e.Created_At = time.Now()
	e.Updated_At = time.Now()

}
