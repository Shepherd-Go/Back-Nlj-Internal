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
