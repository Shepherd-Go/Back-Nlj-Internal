package repository

import (
	"context"

	"github.com/BBCompanyca/Back-Nlj-Internal.git/db/models"
	"github.com/BBCompanyca/Back-Nlj-Internal.git/dtos"
	"gorm.io/gorm"
)

type Employee interface {
	CreateEmployee(ctx context.Context, empl models.Employee) error
	SearchEmployeeByEmail(ctx context.Context, email string) (dtos.EmployeeResponse, error)
	SearchAllEmployees(ctx context.Context) (dtos.Employees, error)
}

type employee struct {
	db *gorm.DB
}

func NewRepositoryEmployee(db *gorm.DB) Employee {
	return &employee{db}
}

func (e *employee) CreateEmployee(ctx context.Context, empl models.Employee) error {

	if err := e.db.WithContext(ctx).Table("employees").Create(empl).Error; err != nil {
		return err
	}

	return nil
}

func (e *employee) SearchEmployeeByEmail(ctx context.Context, email string) (dtos.EmployeeResponse, error) {

	empl := models.Employee{}

	if err := e.db.WithContext(ctx).Table("employees").
		Select("id").Where("email=?", email).
		Scan(&empl).Error; err != nil {
		return dtos.EmployeeResponse{}, err
	}

	return empl.ToDomainDTO(), nil
}

func (e *employee) SearchAllEmployees(ctx context.Context) (dtos.Employees, error) {

	empl := models.Employees{}

	if err := e.db.WithContext(ctx).Table("employees").
		Select("id, first_name, last_name, username, email, phone, permissions, confirmed_email, cod_bank, pay_phone, payment_card, status, created_by, updated_by, created_at, updated_at").
		Where("deleted=?", false).
		Scan(&empl).Error; err != nil {
		return dtos.Employees{}, err
	}

	return empl.ToDomainDTO(), nil

}
