package repository

import (
	"context"
	"time"

	"github.com/Shepherd-Go/Back-Nlj-Internal.git/db/models"
	"github.com/Shepherd-Go/Back-Nlj-Internal.git/dtos"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Employee interface {
	CreateEmployee(ctx context.Context, empl models.Employee) error
	SearchEmployeeByID(ctx context.Context, id uuid.UUID) (dtos.EmployeeResponse, error)
	SearchEmployeeByEmail(ctx context.Context, email string) (dtos.EmployeeResponse, error)
	SearchEmployeeByEmailAndNotID(ctx context.Context, id uuid.UUID, email string) (dtos.EmployeeResponse, error)
	SearchAllEmployees(ctx context.Context) (dtos.Employees, error)
	UpdateEmployee(ctx context.Context, empl models.Employee, id uuid.UUID) error
	ActivateEmail(ctx context.Context, id, pass string) error
	DeleteEmployee(ctx context.Context, id uuid.UUID, idToken string) error
}

type employee struct {
	db *gorm.DB
}

func NewRepositoryEmployee(db *gorm.DB) Employee {
	return &employee{db}
}

func (e *employee) CreateEmployee(ctx context.Context, empl models.Employee) error {

	if err := e.db.WithContext(ctx).Table("employees").Create(&empl).Error; err != nil {
		return err
	}

	return nil
}

func (e *employee) SearchEmployeeByID(ctx context.Context, id uuid.UUID) (dtos.EmployeeResponse, error) {

	empl := &models.Employee{}

	if err := e.db.WithContext(ctx).Table("employees").
		Where("id=?", id).Not("deleted=?", true).
		Select("*").
		Scan(empl).Error; err != nil {
		return dtos.EmployeeResponse{}, err
	}

	return empl.ToDomainDTO(), nil
}

func (e *employee) SearchEmployeeByEmail(ctx context.Context, email string) (dtos.EmployeeResponse, error) {

	empl := models.Employee{}

	if err := e.db.WithContext(ctx).Table("employees").
		Where("email=?", email).Not("deleted=?", true).
		Select("*").
		Scan(&empl).Error; err != nil {
		return dtos.EmployeeResponse{}, err
	}

	return empl.ToDomainDTO(), nil
}

func (e *employee) SearchEmployeeByEmailAndNotID(ctx context.Context, id uuid.UUID, email string) (dtos.EmployeeResponse, error) {

	empl := models.Employee{}

	if err := e.db.WithContext(ctx).Table("employees").
		Where("email=?", email).Not("id=?", id).Not("deleted=?", true).
		Select("*").
		Scan(&empl).Error; err != nil {
		return dtos.EmployeeResponse{}, err
	}

	return empl.ToDomainDTO(), nil
}

func (e *employee) SearchAllEmployees(ctx context.Context) (dtos.Employees, error) {

	empl := models.Employees{}

	if err := e.db.WithContext(ctx).Table("employees").
		Where("deleted=?", false).
		Select("id, first_name, last_name, username, email, phone, permissions, confirmed_email, status, created_at, updated_at").
		Scan(&empl).Error; err != nil {
		return dtos.Employees{}, err
	}

	return empl.ToDomainDTO(), nil
}

func (e *employee) UpdateEmployee(ctx context.Context, empl models.Employee, id uuid.UUID) error {

	if err := e.db.WithContext(ctx).Where("id=?", id).Updates(&empl).Error; err != nil {
		return err
	}

	return nil
}

func (e *employee) ActivateEmail(ctx context.Context, id, pass string) error {
	if err := e.db.WithContext(ctx).Table("employees").Where("id=?", id).Updates(map[string]interface{}{"confirmed_email": true, "password": pass}).Error; err != nil {
		return err
	}

	return nil
}

func (e *employee) DeleteEmployee(ctx context.Context, id uuid.UUID, idToken string) error {

	if err := e.db.WithContext(ctx).Table("employees").Where("id=?", id).Updates(map[string]interface{}{"deleted": true, "deleted_at": time.Now()}).Error; err != nil {
		return err
	}

	return nil
}
