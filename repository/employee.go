package repository

import (
	"context"

	"github.com/BBCompanyca/Back-Nlj-Internal.git/db/models"
	"gorm.io/gorm"
)

type Employee interface {
	CreateEmployee(ctx context.Context, empl models.Employee) error
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
